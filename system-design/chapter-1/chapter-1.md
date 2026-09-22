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

---

## 10. Sharding

**Sharding** = splitting one big database into smaller pieces (**shards**), each on its
own server. Also called **horizontal partitioning**.

- Every shard has the **same schema**.
- Each shard holds **different rows**.

```
 Vertical partitioning:   split COLUMNS  (users_profile | users_settings)
 Horizontal (sharding):   split ROWS     (users 1..1M on A | users 1M..2M on B)
```

### Shard key

The **shard key** decides which shard a row lives on. Example: `user_id`.

```
 shard = user_id % 4
```

| user_id | user_id % 4 | Shard |
|---|---|---|
| 0 | 0 | Shard 0 |
| 1 | 1 | Shard 1 |
| 6 | 2 | Shard 2 |
| 7 | 3 | Shard 3 |
| 8 | 0 | Shard 0 |

### Routing

```
 user_id ──▶ shard function (user_id % 4) ──▶ shard N ──▶ DB server N

                    ┌──────────────────┐
  user_id = 6  ───▶ │  6 % 4 = 2       │
                    └────────┬─────────┘
       ┌──────────┬──────────┼──────────┬──────────┐
       ▼          ▼          ▼          ▼
   Shard 0    Shard 1    Shard 2 ✅   Shard 3
   (DB-0)     (DB-1)     (DB-2)      (DB-3)
```

- **Write:** new row for `user_id = 6` → compute `6 % 4 = 2` → insert into Shard 2.
- **Read:** fetch `user_id = 6` → same function → go straight to Shard 2. No need to search all shards.

The routing logic lives in the app, a DB proxy, or the database itself (e.g. Vitess, Citus, MongoDB).

### Why sharding scales

- **Storage:** data spread across N machines → N× capacity.
- **Throughput:** reads/writes spread across N machines → each handles ~1/N of the load.
- **Smaller indexes** per shard → faster queries.

Good shard key = **evenly distributes data** and matches **how data is queried**.

---

## 11. Sharding Challenges

### 11.1 Resharding

Needed when:
- **Shard exhaustion:** a shard runs out of space / CPU due to fast growth.
- **Uneven distribution:** some shards fill up much faster than others.

Problem with `user_id % 4` → `% 5`: almost **every key maps to a new shard** → massive data movement.

**Consistent hashing** fixes most of this: when a shard is added/removed, only the keys
next to it on the hash ring move (≈ 1/N of the data), not everything.

#### Resharding flow (online, no downtime)

```
 OLD SHARD
    │
    ▼
 MIGRATION ─────────── copy existing data (snapshot / bulk copy)
    │
    ▼
 NEW SHARD
    │
    ▼
 SYNCHRONIZATION ───── writes keep arriving on the old shard during the copy
    │
    ▼
 WAL / CDC ─────────── capture ongoing changes
    │
    ▼
 Apply changes ─────── replay them on the new shard
    │
    ▼
 New shard catches up  (lag ≈ 0)
    │
    ▼
 Switch routing ────── shard map now points these keys to the new shard
    │
    ▼
 Verify ────────────── row counts / checksums / sample reads match
    │
    ▼
 Cleanup old data ──── delete moved rows from the old shard
```

| Term | Meaning |
|---|---|
| Migration | One-time copy of **existing** data |
| Synchronization | Keeping the copy up to date with **new** writes until cutover |
| **WAL** (Write-Ahead Log) | DB's internal append-only log; every change is written here first (for crash recovery) |
| **CDC** (Change Data Capture) | Reading those changes (often from the WAL) and streaming them elsewhere (e.g. Debezium → Kafka) |

### 11.2 Celebrity / Hotspot Problem

**Hot shard** = one shard gets far more traffic than the rest.

Example: a celebrity's profile lives on Shard 2. Millions of users read it at once →
Shard 2 is overloaded while Shards 0, 1, 3 are idle. Sharding by `user_id` evenly spread
the **data**, but not the **traffic**.

```
 Shard 0  ▁
 Shard 1  ▁
 Shard 2  █████████████  ← celebrity key
 Shard 3  ▁
```

Fixes:
- **Cache** the hot data (Redis / CDN) → most reads never reach the DB.
- **Read replicas** for the hot shard → spread reads over multiple copies.
- **Dedicated shard** for very hot keys.
- **Split the hot key** (e.g. `celebrity_id#1..#10`) to spread writes like counters/likes where appropriate.

> Caching mainly reduces **read traffic and latency** on the DB.
> It does **NOT** increase database **storage capacity**.

### 11.3 Cross-Shard Joins

Before sharding, `JOIN users ↔ orders` is one query on one DB.
After sharding, related rows may live on **different servers** → the DB can't join them.

Options:
- **Cross-shard query:** query many shards, merge results in the app (slow, complex).
- **Shard related data together:** shard `orders` by `user_id` too, so a user's orders sit with the user.
- **Denormalize:** copy the needed fields into the table you read, so no join is needed. (Most common.)

---

## 12. Normalization vs Denormalization

| | Normalization | Denormalization |
|---|---|---|
| Idea | Split data, store each fact **once** | **Duplicate** selected data where it's read |
| Duplication | Minimal | Intentional |
| Consistency | Easy (one place to update) | Harder (update many copies) |
| Reads | More joins | Fewer joins, faster/simpler |
| Writes | Simple | More work |
| Storage | Less | More |

### Example

Normalized:

```
 users:  { id: 1, name: "Asha" }
 posts:  { id: 99, user_id: 1, text: "hello" }

 Show post with author → JOIN posts + users
```

Denormalized:

```
 posts:  { id: 99, user_id: 1, author_name: "Asha", text: "hello" }

 Show post with author → one read, no join
 Cost: if Asha renames herself, every post copy must be updated
```

Rule of thumb: **normalize by default**, denormalize the specific **read-heavy** paths that need it.

---

## 13. Replication

**Database replication** = keeping **copies** of the same data on multiple DB servers.

> Sharding **splits** data (each server has different rows).
> Replication **copies** data (each server has the same rows).

### Primary / replica

```
                 writes
  App servers ────────────▶ ┌──────────┐
      │                     │ Primary  │
      │                     └────┬─────┘
      │                          │  WAL / change stream
      │            ┌─────────────┼─────────────┐
      │            ▼             ▼             ▼
      │       ┌─────────┐  ┌─────────┐  ┌─────────┐
      └──────▶│Replica 1│  │Replica 2│  │Replica 3│
       reads  └─────────┘  └─────────┘  └─────────┘
```

| Role | Handles |
|---|---|
| **Primary** | All **writes** (insert/update/delete); source of changes |
| **Replica (read replica)** | **Reads** only; receives changes from the primary |

Most apps read far more than they write → many replicas, one primary.

### How changes flow: WAL

1. Primary writes every change to its **WAL** first.
2. WAL records are shipped to replicas.
3. Replicas **replay** them → same data.

> WAL is the **log**, not replication itself. Replication *uses* the WAL as the change stream.

### Replication lag

Replication is often **asynchronous** → replicas are a little behind the primary.

Example: user updates their bio (write → primary), page reloads (read → replica) → sees the **old** bio.

Common fixes: read-your-own-writes from the primary for a short window, or synchronous replication for critical data (slower writes).

### Failures & failover

| What fails | What happens |
|---|---|
| A replica | Reads go to other replicas (or primary). A **new replica** is created to replace it. |
| The primary | A replica is **promoted** to new primary (**failover**); the rest follow it. |

Why is a failed replica **recreated from a healthy source** (not just restarted)?
- Its data may be **stale, incomplete, or corrupted**.
- Clean path: take a snapshot/backup of a healthy node → restore → replay WAL to catch up.

Promoting a replica has a catch: with async replication it may be **missing the last few
writes** → those must be recovered or accepted as lost.

### Terms that get mixed up

| Term | Meaning |
|---|---|
| Replication | Copying / synchronizing data to other servers |
| Failover | **Switching traffic** to another DB when one fails |
| High availability (HA) | Replication + **automatic** failover → system keeps working through failures |

Benefits: **performance** (parallel reads), **reliability** (data survives a server loss), **availability** (another copy can take over).

---

## 14. Multi-Data Center

**Why:** users are worldwide, and a whole data center (DC) can go down (power, network, region outage).

```
            User
              │
              ▼
           GeoDNS  ── routes by user location + DC health
              │
      ┌───────┴────────┐
      ▼                ▼
   DC: US-East      DC: EU-West        (nearest / healthy DC)
      │                │
      ▼                ▼
 Load Balancer    Load Balancer
      │                │
      ▼                ▼
 Web Servers      Web Servers
      │                │
      ▼                ▼
 DB + Cache  ◀─async replication─▶  DB + Cache
```

| Concept | Meaning |
|---|---|
| **GeoDNS** | DNS that returns a different IP based on where the user is (e.g. EU user → EU DC) |
| Lower latency | Requests travel a shorter distance |
| Failover | DC down → GeoDNS sends **all** traffic to the healthy DC |
| Cross-DC replication | Data copied between DCs so either can serve any user |
| Asynchronous replication | Cross-DC is usually async (sync over long distance is too slow) → some lag |

### Three main challenges

| # | Challenge | What it means |
|---|---|---|
| 1 | **Traffic redirection** | Send users to the right DC (GeoDNS), and reroute on failure |
| 2 | **Data synchronization** | On failover, users must still find their data → replicate across DCs, handle lag/conflicts |
| 3 | **Testing & deployment** | Same code, config, and versions in every DC → automated deploys, test in each location |

> **Memory:** Multi-DC = **Route traffic + Sync data + Deploy consistently**

---

## 15. Stateless Web Tier

**State** = data about a user's session (logged-in user, cart, preferences).

| | Stateful server | Stateless server |
|---|---|---|
| Session data kept | In the server's own memory | In a **shared store** (Redis, DB, NoSQL) |
| Next request must hit | The **same** server | **Any** server |
| Scaling / failure | Hard — server dies, sessions lost | Easy — add/remove servers freely |

### The problem with local state

```
 Request 1 (login)  ──▶ LB ──▶ Server A   (session saved in A's memory)
 Request 2          ──▶ LB ──▶ Server B   ❌ "who are you?" → user logged out
```

Workaround: **sticky sessions** (LB always sends user to the same server) — works, but
makes balancing uneven and breaks when that server dies.

### Stateless design

```
 Client  (cookie: session_id=abc123)
   │
   ▼
 Load Balancer
   │
   ▼
 Any Web Server ──── lookup session abc123 ───▶ Shared State Store (Redis / DB)
```

- Client carries a **session ID** (cookie) or a **token** (e.g. JWT).
- Any server can look up the session → LB can route anywhere.
- Enables **autoscaling**: servers come and go based on load.

> **Memory:** Stateless = the server doesn't depend on **locally stored** client state.
> The application still has state — it just lives in a shared store, not in the web server.
