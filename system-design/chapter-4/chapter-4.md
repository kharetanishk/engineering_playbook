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
