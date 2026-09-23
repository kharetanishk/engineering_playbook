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

---

## 2. Power Of Two

Data volumes are based on powers of two.

```
 2^10 = 1,024        ≈ 1 thousand  (1K)
 2^20 = 1,048,576    ≈ 1 million   (1M)
 2^30 ≈ 1,073 million ≈ 1 billion  (1B)
 2^40 ≈ 1 trillion   (1T)
 2^50 ≈ 1 quadrillion (1P)
```

| Power | Approx value | Unit (bytes) |
|---|---|---|
| 2^10 | 1 thousand | 1 KB |
| 2^20 | 1 million | 1 MB |
| 2^30 | 1 billion | 1 GB |
| 2^40 | 1 trillion | 1 TB |
| 2^50 | 1 quadrillion | 1 PB |

For estimation, treat each step as **×1000** (1 GB ≈ 1000 MB). The 2.4% error per step
does not matter at this level.

### Bytes vs bits

```
 1 byte = 8 bits
```

| Measured in | Usually written as | Example |
|---|---|---|
| **Storage** | **bytes** (KB, MB, GB, TB, PB) | "30 TB of media per day" |
| **Network bandwidth** | **bits per second** (Kbps, Mbps, Gbps) | "a 1 Gbps link" |

> ⚠️ `MB` = megabyte, `Mb`/`Mbps` = megabit. They differ by 8×.
> Mixing them up is the most common estimation mistake.

### Quick conversions

```
 1 Gbps = 1,000 Mbps
 1 Gbps ÷ 8 = 125 MB/s
 1 Mbps ÷ 8 = 125 KB/s

 100 MB/s × 8 = 800 Mbps
```

### Mental math tips

- Round to one significant digit: 3,472 → ~3,500 → "about 3.5k".
- Work in powers of 10: 300M = 3 × 10^8, 86,400 ≈ 10^5 → 3 × 10^8 / 10^5 = 3 × 10^3 = 3,000 (close to the real 3,472).
- 1 million MB = 1 TB, 1 million GB = 1 PB.
- To go bytes → bits, multiply by 8; bits → bytes, divide by 8.
