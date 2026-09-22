# Chapter 1 — Scale From Zero To Millions Of Users

> Source: *System Design Interview* (Alex Xu), Chapter 1. Notes are my own summaries.
>
> Goal of the chapter: start with one server, then add one building block at a time
> until the system can serve millions of users.

---

## 1. Basic System Architecture

Everything starts on **one server**: web app, database, cache — all on the same machine.

| Piece | Role |
|---|---|
| User | Browser or mobile app that sends requests |
| DNS | Converts a domain name (`api.mysite.com`) into an IP address |
| Server | Machine that runs the application and returns responses |
| HTTP/HTTPS | Protocol used to talk to the server (HTTPS = HTTP + TLS encryption) |

### Single-server request flow

```
 User (browser / mobile)
   │  1. "What is the IP of api.mysite.com?"
   ▼
  DNS  ──── 2. returns 15.125.23.214
   │
   │  3. HTTP(S) request to 15.125.23.214
   ▼
 Web Server (app + DB on same machine)
   │  4. HTML page / JSON response
   ▼
 User
```

- **Web app** usually returns HTML pages.
- **Mobile app** usually talks to an API and gets **JSON** back.

Problem: one machine = one point of failure and a hard limit on capacity.

---

## 2. DNS

**DNS (Domain Name System)** = the internet's phone book.

- Humans remember names: `google.com`
- Computers route to IPs: `142.250.x.x`
- DNS maps **domain → IP**.

> DNS is a **naming / lookup system**, not a networking layer.
> It only tells you *where* to go. Actually getting there is done by TCP/IP.

```
 Browser ──"mysite.com?"──▶ DNS resolver ──▶ (root → .com → mysite's nameserver)
    ▲                                                     │
    └──────────────── "15.125.23.214" ◀───────────────────┘

 Browser ──── HTTP request ────▶ 15.125.23.214 (server)
```

- DNS is usually a **3rd-party managed service** (Route 53, Cloudflare), not something we host.
- Results are **cached** (browser, OS, resolver) based on TTL, so lookups are rare.

---

## 3. HTTP

**Mental model:** client asks, server answers. One request → one response.

```
 Client ──── Request  (GET /users/42) ────▶ Server
 Client ◀─── Response (200 OK + JSON) ───── Server
```

| Request has | Response has |
|---|---|
| Method (`GET`, `POST`, `PUT`, `DELETE`) | Status code (`200`, `404`, `500`) |
| URL / path | Headers |
| Headers (auth, content-type) | Body (HTML / JSON) |
| Body (for `POST` / `PUT`) | |

- HTTP is **stateless**: every request carries what the server needs (e.g. cookie / token).
- HTTPS = same thing, encrypted with TLS.

---

## 4. Web Tier

| Term | What it does | Examples |
|---|---|---|
| Web server | Accepts HTTP connections, serves static files, forwards requests | Nginx, Apache |
| Application server | Runs business logic, talks to DB | Node.js/Express, Django, Spring |

In small setups both are often the same process (e.g. an Express app).

### Step 1 of scaling: separate web tier and data tier

```
 Before:                         After:

 ┌──────────────────┐           ┌────────────┐      ┌────────────┐
 │ Server           │           │ Web / App  │ ───▶ │ Database   │
 │  app + database  │           │ server     │      │ server     │
 └──────────────────┘           └────────────┘      └────────────┘
                                  (web tier)          (data tier)
```

Why split?
- Web tier handles **traffic**, data tier handles **storage** → scale each **independently**.
- A crash or CPU spike in the app doesn't take the database down with it.

---

## 5. Load Balancer

### Why

With one web server:
- server dies → site is down
- traffic spike → slow responses / timeouts

Fix: run **many** web servers and put a **load balancer (LB)** in front.

### Flow

```
                    User
                     │
                     ▼
                    DNS  → returns LB's public IP
                     │
                     ▼
        ┌──────────────────────────┐
        │  Load Balancer           │   public IP  (e.g. 88.88.88.1)
        └──────────────────────────┘
             │                │        private network
             ▼                ▼
      ┌────────────┐   ┌────────────┐
      │ Server 1   │   │ Server 2   │   private IPs (10.0.0.1, 10.0.0.2)
      └────────────┘   └────────────┘
```

| Concept | Meaning |
|---|---|
| Public IP (LB) | The only address users see / DNS returns |
| Private IPs (servers) | Reachable only inside the network → servers are not directly exposed |
| Request distribution | LB picks a server: round robin, least connections, IP hash, … |
| Health checks | LB pings servers (`/health`); unhealthy ones stop receiving traffic |

What we gain:
- **Failover:** Server 1 down → all traffic goes to Server 2.
- **Horizontal scaling:** traffic grows → add Server 3, register it with LB.

### Vertical vs horizontal scaling

| | Vertical (scale up) | Horizontal (scale out) |
|---|---|---|
| How | Bigger machine (more CPU/RAM) | More machines |
| Limit | Hardware ceiling | Practically unlimited |
| Failure | Still a single point of failure | Others keep serving |
| Complexity | Simple | Needs LB, stateless servers, etc. |

Rule: vertical is fine early; large systems scale **horizontally**.

---

## 6. Database

The database is the **source of truth** — where data permanently lives.

### Relational databases (RDBMS)

- Data in **tables** (rows + columns) with a fixed **schema**.
- Tables connect via **keys** → queried with **SQL** using **joins**.
- Strong **transactions** (ACID): all steps succeed or none do.
- Examples: **PostgreSQL**, MySQL, Oracle, SQL Server.

> **PostgreSQL** = popular open-source RDBMS. Strong SQL support, transactions,
> indexes, JSONB columns, extensions. A safe default for most apps.

```
 users                          orders
 ┌────┬────────┐                ┌────┬─────────┬────────┐
 │ id │ name   │                │ id │ user_id │ amount │
 ├────┼────────┤   1 ─── many   ├────┼─────────┼────────┤
 │ 1  │ Asha   │ ◀───────────── │ 10 │ 1       │ 500    │
 │ 2  │ Ravi   │                │ 11 │ 1       │ 250    │
 └────┴────────┘                └────┴─────────┴────────┘
```

### NoSQL databases

Not table-first. Main families:

| Type | Stores | Examples |
|---|---|---|
| Key-value | `key → value` | Redis, DynamoDB |
| Document | JSON-like documents | MongoDB, Couchbase |
| Wide-column | Rows with flexible columns, built for huge scale | Cassandra, HBase |
| Graph | Nodes + edges | Neo4j |

When NoSQL is useful:
- Data is **flexible / semi-structured** (fields vary between records).
- You need very **low latency** for simple lookups by key.
- Massive data volume that must **scale out horizontally** across many machines.
- Few relationships between entities.

### One-line memory

```
SQL   → structured relationships + transactions
NoSQL → flexible schema + scale-out
```

---

## 7. SQL vs NoSQL

No absolute winner — pick based on **data shape** and **access patterns**.

| Need | SQL (RDBMS) | NoSQL |
|---|---|---|
| Structured data, known schema | ✅ Great fit | Works, but no enforcement |
| Relationships between entities | ✅ Foreign keys | Usually modeled by embedding / duplicating |
| Joins | ✅ Native | Limited or none → join in app code |
| Multi-row transactions (ACID) | ✅ Core feature | Varies (often limited to one item/document) |
| Flexible / changing schema | Needs migrations | ✅ Natural |
| Huge distributed workloads | Possible, harder (sharding by hand / extensions) | ✅ Often built-in |
| Horizontal scaling | Harder | ✅ Usually designed for it |

Interview framing:
- **Payments, orders, inventory** → SQL (correctness, transactions, relations).
- **Activity feeds, logs, sessions, IoT events, catalogs with varied fields** → NoSQL can fit well.
- Many real systems use **both** (Postgres for core data, Redis/Cassandra for specific workloads).

---

## 8. ORM

**ORM (Object-Relational Mapper)** = a library that lets you work with DB rows as
objects in your language.

> An ORM is **not a database**. It is a translation layer between code and a (usually SQL) database.

```
 Code (objects)                ORM                      Database (tables)
 ───────────────        ─────────────────         ──────────────────────
 prisma.user.findUnique  →  generates SQL  →   SELECT * FROM users WHERE id = 42;
 ({ where: { id: 42 } })                              │
        ▲                                             │
        └──── row mapped back to a User object ◀──────┘
```

Examples: **Prisma**, **Sequelize**, **TypeORM** (Node.js), SQLAlchemy (Python), Hibernate (Java).

What it gives you:
- **Object ↔ table mapping** (class/model ↔ table, field ↔ column).
- **SQL generation** — write code, ORM writes the query.
- **Migrations** — versioned schema changes (`prisma migrate`), tracked in git.
- Type safety, less boilerplate, some protection from SQL injection (parameterized queries).

Limitations:
- Generated SQL can be **inefficient** (e.g. N+1 queries).
- Complex queries / reports are often easier in **raw SQL**.
- Hides what's happening → you still need to understand SQL and indexes.

---

## 9. Serialization / Deserialization

| Term | Meaning |
|---|---|
| **Serialization** | In-memory object → format that can be sent/stored (JSON string, bytes) |
| **Deserialization** | That format → back into an in-memory object |

```js
const user = { id: 42, name: "Asha" };

const text = JSON.stringify(user);   // serialize   → '{"id":42,"name":"Asha"}'
const again = JSON.parse(text);      // deserialize → { id: 42, name: "Asha" }
```

### API flow

```
 Client object ─serialize─▶ JSON over HTTP ─deserialize─▶ Server object
 Server object ─serialize─▶ JSON over HTTP ─deserialize─▶ Client object
```

Other formats: Protocol Buffers, Avro, MessagePack (binary → smaller/faster than JSON).

> ⚠️ Serialization is **NOT** why NoSQL is fast.
> NoSQL performance comes from **access patterns** (fetch by key, data stored the way
> it's read), **indexing**, **no joins**, and **distribution** across many machines.
> A badly modeled NoSQL DB can be slower than Postgres.
