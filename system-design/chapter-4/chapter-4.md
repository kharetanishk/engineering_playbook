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
