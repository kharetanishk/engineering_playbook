# Chapter 4 — Rate Limiter

> Source: *System Design Interview* (Alex Xu), Chapter 4. Notes are my own summaries.

## 1. What Is a Rate Limiter?

A **rate limiter** caps how many requests a client can make in a given time window.
Once the limit is hit, extra requests are rejected (usually `HTTP 429 Too Many Requests`).

**Purpose:**

- Protect servers from being overwhelmed (abuse, bugs, traffic spikes)
- Control cost (e.g. limit calls to a paid third-party API)
- Prevent one client from starving others (fair usage)

> **Memory:** a rate limiter is a bouncer at the door — past the headcount, no more entries.

---

## 2. Rate Limiter Requirements

- **Accurate limiting** — should not let more requests through than the configured limit.
- **Low latency** — the check must not slow down the request path.
- **Low memory usage** — the algorithm's state should be cheap to store per client.
- **Distributed rate limiting** — must work correctly across many servers, not just one.
- **Exception handling** — clients that are rate-limited should get a clear error, not a silent failure.
- **High fault tolerance** — if the rate limiter itself fails, it should fail in a safe way (not take the whole system down).

> **Memory:** accurate, fast, cheap, shared, clear, resilient.

---

## 3. Client-Side vs Server-Side Rate Limiting

- **Client-side** — the client throttles its own requests. Easy to bypass (a malicious or
  buggy client just ignores it), so it can't be trusted for protection.
- **Server-side** — the server (or a gateway in front of it) enforces the limit. This is
  the trustworthy place to enforce limits, since the client doesn't control it.

> **Memory:** never trust the client to rate-limit itself — enforce it server-side.

---

## 4. API Gateway

**Definition:** a single entry point that sits in front of backend services and handles
cross-cutting concerns before a request reaches business logic.

**Responsibilities:**

- Rate limiting
- Authentication
- TLS termination
- Routing requests to the right service
- Logging / metrics

**Simple request flow:**

```
Client → API Gateway → Backend Service
           |
           +-- rate limit check
           +-- auth check
           +-- TLS termination
```

> **Memory:** the API Gateway is the front desk — it checks your ID and your appointment
> slot before letting you into the building.

---

## 5. TLS / SSL

- **SSL** is the older protocol; **TLS** is its modern successor. In practice "SSL" is
  often used loosely to mean TLS.
- **HTTPS = HTTP + TLS** — same HTTP semantics, encrypted in transit.
- **TLS termination** — the point (often the API Gateway or load balancer) where
  encrypted traffic is decrypted. Traffic can then travel as plain HTTP internally.
- **TLS re-encryption** — after termination, the gateway opens a *new* TLS connection to
  the backend, so traffic is encrypted again for the next hop.

```
Client --TLS--> Gateway --(plain or re-encrypted TLS)--> Backend
              (termination)
```

> **Memory:** TLS termination = take off the envelope at the door; re-encryption = put it
> in a new envelope for the next leg.

---

## 6. TLS Certificates

- A **certificate** proves the server's identity (like an ID card) — it says "this server
  is really who it claims to be."
- A certificate is **not encryption itself** — it enables trust and key exchange during
  the handshake, but the actual traffic is protected by session keys negotiated afterward.
- **CA (Certificate Authority)** — issues and signs the certificate.
- **Certificate manager** — automates requesting, renewing, and installing certificates.
- **Server** — presents the certificate to clients; it doesn't issue or sign it.

> **Memory:** CA = the passport office, certificate manager = the assistant who renews
> your passport, server = the person showing the passport.

---

## 7. TLS Handshake

The handshake happens once, when a connection starts:

1. **Server authentication** — client checks the server's certificate against a trusted CA.
2. **Key establishment** — client and server agree on session keys used to encrypt the
   rest of the conversation.

```
Client                      Server
   |--- Hello ---------------->|
   |<-- Certificate -----------|
   | (verify cert against CA)  |
   |--- agree on session keys->|
   |<==== encrypted traffic ==>|
```

> **Memory:** the handshake proves *who* you're talking to, then agrees on a *secret*
> for the rest of the chat.

---

## 8. Certificate Managers

Tools that automate issuing, renewing, and installing TLS certificates:

- **cert-manager** (Kubernetes)
- **AWS Certificate Manager**
- **Google Cloud Certificate Manager**
- **HashiCorp Vault**
- **Azure Key Vault**
- **Venafi**

---

## 9. API Gateway Examples

- **Kong**
- **AWS API Gateway**
- **Apigee**
- **NGINX**
- **Traefik**
- **Envoy Gateway**
- **Tyk**
- **Azure API Management**

---

## 10. API Gateway vs Service Mesh

- **API Gateway** — sits at the edge, handles **north-south** traffic (outside clients
  calling into the system).
- **Service Mesh** — sits between internal services, handles **east-west** traffic
  (service-to-service calls), e.g. retries, internal routing, mTLS.

> **Memory:** gateway = front door for outsiders; mesh = hallways between rooms inside.

---

## 11. When to Use or Build a Custom API Gateway

- Use an **existing gateway** for standard needs: rate limiting, auth, routing, TLS.
- **Build/customize** only when there's a specialized requirement an existing gateway
  can't meet (unusual protocol, very specific business logic at the edge, etc.).

## 12. Principle

> Use an existing API Gateway by default. Only customize or build your own when you have
> a specialized requirement that off-the-shelf gateways don't cover.

---

## 13. Rate Limiting Algorithms

### 13.1 Token Bucket

- Bucket has a fixed **max capacity**.
- Tokens are added at a fixed **refill rate**.
- Each request consumes **one token**.
- Token available → request allowed. No token → request rejected.
- **Bucket size controls burst capacity**; **refill rate controls sustained rate**.

```
[ capacity: 10 ]
[ ●●●●●○○○○○ ]  <- tokens, refilled over time
     |
  request consumes 1 token
```

**Why memory efficient:** it only stores a token count and last-refill timestamp per
bucket — not a log of every request. Multiple logical buckets (per user, per IP, per
endpoint) can exist, but each one is tiny.

**Burst example:** bucket capacity = 10, refill = 1 token/sec. If no requests come in for
10 seconds, the bucket fills to 10 — the client can then fire 10 requests instantly (a
burst), then must wait for refills.

> **Memory:** Bucket size = burst capacity | Refill rate = sustained request rate

### 13.2 Leaky Bucket

- Requests are placed into a **FIFO queue**.
- If the queue is full, new requests are **rejected**.
- Requests **leave the queue at a fixed rate** (processed one at a time, steadily).
- **Bucket size = queue capacity**; **outflow rate = requests processed per unit time**.

```
requests -> [ queue: FIFO, fixed size ] -> leaks out at fixed rate -> processed
              (full? reject new ones)
```

- **Benefit:** smooth, stable outflow — downstream never sees a burst.
- **Drawback:** a burst can fill the queue, so newer requests wait or get rejected even if
  they'd otherwise be within a longer-term average.
- **Memory efficient:** the queue has a fixed max size, so worst-case memory is bounded.

> **Memory:** Leaky Bucket = queue requests → process them at a fixed rate.

### 13.3 Fixed Window Counter

- Divide time into **fixed windows** (e.g. every 1-minute clock boundary).
- Maintain a **counter** for each window.
- Each request **increments the counter**.
- Once the limit is reached, further requests are **rejected until the next window**.
- At the next window, the **counter resets to 0**.

- **Pros:** simple, memory efficient (one counter per window).
- **Main problem: boundary burst.**

**Boundary example:** limit = 5 requests/minute.

```
window 1 [00:00 - 00:59]         window 2 [01:00 - 01:59]
                    5 requests |5 requests
                          ^ 00:59            ^ 01:00
```

5 requests arrive right before `00:59` (end of window 1) and 5 more arrive right after
`01:00` (start of window 2). Each window individually stays within its 5/minute limit —
but the rolling one-minute period spanning `00:30`–`01:30` actually saw **10 requests**.
A rolling window can straddle two fixed windows, letting bursts slip through.

> **Memory:** Fixed Window = fixed time box + counter + reset.

### 13.4 Sliding Window Counter

- A **hybrid** of Fixed Window Counter and Sliding Window Log.
- **Rolling window** = "look backward a fixed amount of time from NOW."
- A rolling window can **overlap parts of two fixed windows** (the tail of the previous
  one and the head of the current one).
- Estimates the request count as:

```
estimated count = current window requests + (previous window requests × overlap %)
```

**Example:** limit = 7 requests/minute, previous window = 5 requests, current window = 3
requests, current position in window = 30% elapsed → previous window overlap = 70%.

```
estimated count = 3 + (5 × 0.7) = 6.5  → round down to 6
6 < 7  →  request allowed
```

- The 70% overlap exists because we're only 30% into the current window, so 70% of "now
  looking back one minute" still falls inside the previous window.
- Assumes requests in the previous window were **evenly distributed** — it's an
  **approximation**, not an exact count.

- **Pros:** memory efficient (still just two counters), smoother than fixed window.
- **Cons:** not perfectly accurate for strict look-back windows.

> **Memory:** Sliding Window Counter = previous counter × overlap + current counter.

### 13.5 Comparison

| Algorithm              | Main idea                          | Burst handling                          | Memory              | Main drawback                          |
|-------------------------|-------------------------------------|------------------------------------------|----------------------|------------------------------------------|
| Token Bucket            | Tokens refill, request consumes one | Allows bursts up to bucket size          | Very low (count + timestamp) | Burst can still hit downstream at once |
| Leaky Bucket            | FIFO queue drained at fixed rate    | Smooths bursts into steady outflow       | Low (bounded queue)  | Bursts fill queue, delay/reject newer requests |
| Fixed Window Counter    | Counter per fixed time window       | Poor — boundary burst (2x limit possible) | Very low (one counter) | Boundary burst                          |
| Sliding Window Counter  | Weighted blend of two window counters | Good approximation, smooths boundary    | Low (two counters)   | Approximation, not exact for strict windows |

---
