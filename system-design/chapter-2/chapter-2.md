# Chapter 2 — Back-of-the-Envelope Estimation

> Source: *System Design Interview* (Alex Xu), Chapter 2. Notes are my own summaries.

**Back-of-the-envelope estimation** = using reasonable assumptions and simple arithmetic to
estimate a system's capacity and performance **before** designing the actual infrastructure.

The goal is **not** perfect accuracy. The goal is the approximate scale of:

- QPS (average and peak)
- Storage
- Bandwidth
- Latency
- Availability / allowed downtime
- Infrastructure needed (servers, DB size, cache size, CDN)

### Mental model

```
 Requirements
     ↓
 Assumptions
     ↓
 Simple calculations
     ↓
 QPS / Storage / Bandwidth / Latency
     ↓
 Infrastructure decisions
```

---

## 1. What Is Back-Of-The-Envelope Estimation?

Rough math done "on the back of an envelope" — no spreadsheets, no profiler, just a few
assumptions and multiplication/division you can do in your head.

**Why designers use it**
- Decides the shape of the architecture: 10 QPS and 1 GB needs one server; 100k QPS and 50 PB does not.
- Tells you which resource is the bottleneck (CPU? storage? bandwidth?).
- Catches impossible designs early, before anything is built.

**Why exact numbers are unnecessary**
- The inputs are guesses anyway (future users, growth, payload size).
- A 20% error doesn't change the design; a 100× error does.
- Real traffic will differ from any estimate, so you design with headroom.

**Why reasonable assumptions matter**
- Every number downstream comes from them — state them out loud so they can be checked.
- If an assumption is wrong, you can redo the math in seconds instead of redoing the design.

**How estimates drive architecture**

| Estimate says | Design reacts with |
|---|---|
| High QPS | Load balancer, more servers, caching |
| Read-heavy | Cache + read replicas |
| Write-heavy | Sharding, queues, batching |
| Huge storage | Sharding, object storage, retention limits |
| High bandwidth | CDN, compression, smaller payloads |

### Simple example

```
 100M users
    ↓ × 10 requests/day
 1B requests/day
    ↓ ÷ 86,400 seconds
 ~11,600 average QPS
    ↓ × 2 (peak assumption)
 ~23,000 peak QPS
    ↓
 → needs a load balancer, many app servers, caching; one server is nowhere near enough
```

### Interview mindset

> "The goal is to get the correct order of magnitude, not an exact number."

Say your assumptions out loud, round aggressively (300M ÷ 86,400 ≈ 3,500, not 3,472.2),
and keep the arithmetic simple.
