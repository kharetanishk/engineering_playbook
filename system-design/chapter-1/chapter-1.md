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
