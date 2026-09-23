# Chapter 2 — Back-of-the-Envelope Estimation

> Source: *System Design Interview* (Alex Xu), Chapter 2. Notes are my own summaries.

## Contents

1. [What Is Back-Of-The-Envelope Estimation?](#1-what-is-back-of-the-envelope-estimation)
2. [Power Of Two](#2-power-of-two)
3. [Time Units](#3-time-units)
4. [Latency Numbers](#4-latency-numbers)
5. [CPU Cache vs RAM](#5-cpu-cache-vs-ram)
6. [Availability](#6-availability)
7. [QPS](#7-qps)
8. [Storage Estimation](#8-storage-estimation)
9. [Bandwidth Estimation](#9-bandwidth-estimation)
10. [Read/Write Ratio](#10-readwrite-ratio)
11. [Worked Example — Twitter-Style Service](#11-worked-example--twitter-style-service)
12. [Estimation Workflow](#12-estimation-workflow)
13. [Conversion Cheat Sheet](#13-conversion-cheat-sheet)
14. [Common Mistakes](#14-common-mistakes)
15. [Latency vs Throughput](#15-latency-vs-throughput)

---

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

---

## 3. Time Units

```
 1 second = 1,000 ms
          = 1,000,000 μs
          = 1,000,000,000 ns

 1 ms = 1,000 μs
 1 μs = 1,000 ns
```

| Unit | Name | Fraction of a second |
|---|---|---|
| ns | nanosecond | 1 / 1,000,000,000 |
| μs | microsecond | 1 / 1,000,000 |
| ms | millisecond | 1 / 1,000 |
| s | second | 1 |

### Where each shows up

| Unit | Typical operation |
|---|---|
| ns | CPU cache read, one CPU instruction |
| μs | RAM access, lock/unlock, compressing a small buffer |
| ms | Disk read, network round trip, a full API request |
| s | Batch job, video upload, report generation |

Useful for estimation:

```
 1 minute =        60 seconds
 1 hour   =     3,600 seconds
 1 day    =    86,400 seconds
 1 month  ≈ 2,600,000 seconds  (~2.6M)
 1 year   ≈    31.5M  seconds
```

> **Memory:**
> ns → extremely small (CPU-level)
> μs → micro (memory-level)
> ms → milliseconds (disk and network level, what users feel)

---

## 4. Latency Numbers

Purpose: know **how expensive** each kind of operation is **relative to the others**, so you
can guess where time goes without measuring.

> ⚠️ The numbers below are the classic reference values (Jeff Dean's list, used in the book).
> They are **old** — modern CPUs, NVMe SSDs and networks differ, sometimes by a lot.
> Treat them as **orders of magnitude**, not as benchmarks of current hardware.

| Operation | Reference latency | What it actually means |
|---|---:|---|
| L1 cache reference | ~0.5 ns | CPU reads data from its smallest, closest cache |
| Branch misprediction | ~5 ns | CPU guessed the wrong `if` path and must discard work and restart |
| L2 cache reference | ~7 ns | CPU reads from the next cache level — bigger, slightly slower than L1 |
| Mutex lock/unlock | ~100 ns | Taking and releasing a lock so only one thread enters a critical section |
| Main memory (RAM) reference | ~100 ns | CPU reads data from main memory — a cache miss |
| Compress 1 KB (Zippy/Snappy) | ~10 μs | CPU work to shrink 1 KB of data before storing/sending it |
| Send 2 KB over 1 Gbps network | ~20 μs | Pushing a small payload onto the wire |
| Read 1 MB sequentially from memory | ~250 μs | Streaming a megabyte that's already in RAM |
| Round trip within same data center | ~500 μs | Request to another server in the same DC **and** its reply back |
| Disk seek (HDD) | ~10 ms | Mechanical head moves to the right track/position |
| Read 1 MB sequentially from network | ~10 ms | Pulling a megabyte across the network |
| Read 1 MB sequentially from disk | ~30 ms | Streaming a megabyte off spinning disk |
| Round trip across continents | ~150 ms | e.g. California ↔ Netherlands and back |

### What some of these mean

**L1 / L2 cache** — small, very fast memory built into the CPU. The CPU looks here first;
a hit costs nanoseconds. L1 is smallest and fastest, then L2, then L3.

**RAM (main memory)** — where running programs keep their data. Much bigger than cache,
but a read costs ~100× an L1 hit. Every cache miss ends up here.

**Disk seek** — on an **HDD**, the read/write head physically moves to the track holding
the data. Physical movement is why it costs milliseconds.

> **Seek ≠ read.**
> Seek = **locating / positioning** (finding where the data is).
> Read = **actually retrieving** the bytes.
> A random read = seek + read; a sequential read pays the seek once and then streams.
>
> **SSDs have no moving head**, so HDD seek numbers must not be used as modern SSD
> benchmarks. SSD random access is roughly microseconds-to-tens-of-microseconds, not ~10 ms.

**Same-data-center round trip**

```
 Server A
   │  request  (network hop)
   ▼
 Server B
   │  response (network hop)
   ▼
 Server A
```

**Round trip = request there + response back.** One extra service call adds a full round
trip, which is why chatty microservices get slow.

**Long-distance network** — a signal can't beat the speed of light, plus it passes through
routers, switches and undersea cables. Continents apart ≈ 150 ms round trip, no matter how
fast your servers are. This is why CDNs and multi-DC deployments exist.

### Rough hierarchy

```
 CPU cache          ns
    ↓
 RAM                ns → μs
    ↓
 Storage (disk)     μs (SSD) → ms (HDD)
    ↓
 Network (same DC)  μs → ms
    ↓
 Long-distance net  ~100 ms
```

> This is a **rule of thumb**, not an exact universal ordering — a fast NVMe read can beat
> a slow same-DC round trip, and workloads vary.

**Key lesson:**

> Local memory is extremely fast; storage and especially network operations can be
> orders of magnitude more expensive.

### Why it matters in design

| Idea | Because |
|---|---|
| **Caching** | Serving from memory avoids disk and network entirely |
| **Fewer network calls** | Each call costs at least one round trip; N+1 calls kill latency |
| **Data locality** | Keep data near the compute that reads it (same DC, same shard, same region) |
| **Batching** | 1 request for 100 items beats 100 requests for 1 item |
| **Avoid unnecessary disk access** | Especially **random** access on spinning disks |

---

## 5. CPU Cache vs RAM

```
 CPU
 ├── L1   (smallest, fastest, per core)
 ├── L2   (bigger, slightly slower)
 └── L3   (biggest cache, shared between cores)
        ↓
      RAM        (main memory — GBs)
        ↓
    Storage      (SSD / HDD — TBs, survives reboot)
```

| | CPU cache (L1/L2/L3) | RAM (main memory) |
|---|---|---|
| Size | KB → tens of MB | GB |
| Speed | ~1–20 ns | ~100 ns |
| Location | Inside the CPU | On the motherboard |
| Holds | The data being worked on right now | All running programs' data |
| Cost per byte | Very high | Lower |
| Persistent? | No | No (both lost on power off) |

The tradeoff is always the same: **smaller and closer = faster and more expensive**.

### Analogy

> **Cache** = the papers **on your desk** — tiny space, grab instantly.
> **RAM** = the **bookshelf** in the room — much more, a few seconds to walk over.
> **Disk** = the **storage room down the hall** — huge, but a trip each time.
> **Network** = a book in **another building** — you have to ask someone to send it.

A **cache hit** = what you need is already on the desk. A **miss** = get up and fetch it.
The same hit/miss idea repeats at every layer: CPU cache, Redis, CDN.

> **Memory:**
> Cache = information on your desk.
> RAM = information in your bookshelf.

---

## 6. Availability

**Availability** = the percentage of time a system is operational and able to serve requests.

```
 Availability = Uptime / Total time

 Downtime = Total time × (1 - Availability)
```

Example, over one year (365 days ≈ 8,760 hours):

```
 Availability = 99.9% = 0.999
 Downtime = 8,760 h × (1 - 0.999) = 8.76 hours/year
```

### The "nines"

| Availability | Name | Approx downtime/year | Per month | Per day |
|---|---|---:|---:|---:|
| 99% | two nines | 3.65 days | 7.2 hours | 14.4 min |
| 99.9% | three nines | 8.77 hours | 43.8 min | 1.44 min |
| 99.99% | four nines | 52.6 minutes | 4.38 min | 8.6 s |
| 99.999% | five nines | 5.26 minutes | 26 s | 0.86 s |
| 99.9999% | six nines | 31.56 seconds | 2.6 s | 0.086 s |

Each extra nine = **10× less allowed downtime** — and a big jump in cost and complexity
(redundancy everywhere, multi-DC, automated failover, no manual recovery steps).

Note that a **single** 10-minute outage already breaks a 99.99% yearly target.

### SLA

**SLA (Service Level Agreement)** = a contract between a provider and its customers stating
the service level that is promised — typically an availability target, plus penalties or
credits if it's missed.

Examples: cloud providers commonly publish SLAs like 99.9% or 99.99% for their services.

> An SLA is a **promise**, not a measurement.
> Actual measured availability can be better or worse than the SLA — the SLA just defines
> what counts as a breach and what compensation follows.

Related: **SLO** = the internal target a team aims for (usually stricter than the SLA);
**SLI** = the metric actually measured (e.g. % of successful requests).

> **Memory:**
> Availability = uptime percentage.
> More 9s = less downtime.

---

## 7. QPS

**QPS** = **Queries Per Second**. Often used interchangeably with **RPS** (requests per
second) — "query" for databases/search, "request" for APIs. Same math either way.

### Core conversion

```
 Average QPS = Requests per day ÷ 86,400

 because 24 × 60 × 60 = 86,400 seconds in a day
```

Example:

```
 864,000 requests/day ÷ 86,400 = 10 QPS
```

More:

| Requests/day | Average QPS |
|---|---:|
| 86,400 | 1 |
| 1,000,000 | ~12 |
| 100,000,000 | ~1,160 |
| 1,000,000,000 | ~11,600 |

### Average vs peak QPS

Traffic is **not** spread evenly over 24 hours:

```
 QPS
  │             ▁▃▅███▅▃▁
  │        ▁▃▅██        ██▅▃▁
  │   ▁▃▅██                  ██▅▃▁
  └──────────────────────────────────── time of day
   night      morning   evening peak   night
```

- Users sleep, work, and browse at the same times.
- Regional concentration, launches, notifications and events all create spikes.

Capacity must be planned for the **peak**, not the average — otherwise the system falls
over exactly when it's busiest.

```
 Peak QPS = Average QPS × peak multiplier
```

Example:

```
 Average QPS   = 3,500
 Peak multiplier = 2
 Peak QPS      ≈ 7,000
```

> ⚠️ The peak multiplier is an **assumption**, not a universal constant.
> 2× is a common interview default. Real systems vary — a global app may be flatter (~1.5×),
> a regional or event-driven one much spikier (5–10×). State the number you're assuming.

---

## 8. Storage Estimation

### The pattern

```
 Data generated per day
   = Active users per day
   × actions per user per day
   × data size per action

 Total storage
   = Storage per day
   × number of days retained
```

Example:

```
 10M DAU × 5 posts/day × 2 KB/post
   = 100M KB/day
   ≈ 100 GB/day

 × 365 days × 3 years  ≈ 110 TB
```

Do it **per data type** when sizes differ wildly — text and media are not in the same league:

| Data | Typical size | Effect |
|---|---|---|
| Text post / row | ~1 KB | Small; QPS matters more than bytes |
| Metadata (ids, timestamps) | ~100 B | Usually negligible |
| Image | ~0.5–2 MB | Dominates storage |
| Video | ~10 MB–GBs | Dominates everything |

> Media almost always decides the storage answer. Estimate it separately from text.

### Multipliers to remember

| Factor | Effect on raw storage |
|---|---|
| **Replication** | ×2, ×3 — copies for durability and availability |
| **Backups** | + snapshots/archives, often kept for months |
| **Metadata** | + a small % (ids, timestamps, ownership) |
| **Indexes** | + 10–30% of table size, sometimes more |
| **Compression** | ÷ 2–10 for text/logs; ~×1 for already-compressed media (JPEG, MP4) |
| **Retention** | Deleting after N days can cap total storage entirely |
| **Growth** | Users grow → next year's daily rate is higher than today's |

Example of how fast this adds up:

```
 Raw           30 TB/day
 × 3 replicas  90 TB/day
 + backups    ~110 TB/day
```

### In an interview

Start simple: `users × actions × size × days`. Get the order of magnitude first, then
say "with 3× replication and backups this is roughly 3–4× higher". Adding every factor
up front just buries the main number.

---

## 9. Bandwidth Estimation

**Bandwidth** = how much data is transferred per unit of time.

```
 Bandwidth = Requests per second × data per request
```

Example:

```
 1,000 requests/sec × 100 KB/request
   = 100,000 KB/sec
   ≈ 100 MB/sec
```

Convert to bits, because network links are sold in bits per second:

```
 100 MB/s × 8 = 800 Mbps
```

So this workload needs roughly a **1 Gbps** link — and that's the average, so size for peak.

### MB/s vs Mbps

| Written | Means | Note |
|---|---|---|
| **MB/s** | Mega**bytes** per second | Data volume — how storage is measured |
| **Mbps** / **Mb/s** | Mega**bits** per second | Network speed — how links are sold |

```
 MB/s → Mbps :  × 8
 Mbps → MB/s :  ÷ 8

 1 Gbps ≈ 125 MB/s
```

> Lowercase **b** = bits, uppercase **B** = bytes. Getting this wrong makes you 8× off.

### Estimate both directions

- **Egress / outbound** (server → user): usually the big one, and the one cloud providers bill.
- **Ingress / inbound** (user → server): uploads, writes.

Example: a video service streams far more than it receives, so its outbound bandwidth
dominates — which is exactly why a CDN is used to serve it.

> **Memory:** Data × requests = bandwidth.

---

## 10. Read/Write Ratio

Most systems are **read-heavy** — knowing by how much changes the whole design.

Example:

```
 1,000 total requests/sec
   80% reads  →   800 reads/sec
   20% writes →   200 writes/sec
```

Common ratios: social feeds and catalogs are often 10:1 to 100:1 reads-to-writes;
logging/analytics ingestion can be write-heavy instead.

### What the ratio changes

| Ratio | Design implication |
|---|---|
| **Read-heavy** | Cache aggressively; add **read replicas**; use a CDN; denormalize for fast reads |
| **Write-heavy** | Shard to spread writes; **queue** and batch them; keep indexes lean (each index slows writes) |
| Either | Sizes the DB, cache capacity and replica count |

```
 Read-heavy:   Client → Cache → (miss) → Read replicas → Primary
 Write-heavy:  Client → Queue → Workers → Primary → shards
```

Rule: **reads scale with copies** (cache, replicas, CDN); **writes scale with partitions**
(sharding, queues).

---

## 11. Worked Example — Twitter-Style Service

> ⚠️ These are **exercise assumptions for practice**, not real statistics about any company.

### Assumptions

| Assumption | Value |
|---|---|
| Monthly active users (MAU) | 300M |
| Share who are daily active | 50% |
| Tweets per user per day | 2 |
| Tweets containing media | 10% |
| Average media size | 1 MB |
| Retention | 5 years |
| Peak multiplier | 2× |

### Step 1 — Daily Active Users

**DAU (Daily Active Users)** = users who use the app on a given day.

```
 300M MAU × 50% = 150M DAU
```

### Step 2 — Tweets per day

```
 150M DAU × 2 tweets = 300M tweets/day
```

### Step 3 — Average QPS

```
 300M ÷ 86,400 ≈ 3,500 QPS
```

We divide by **86,400** because that's the number of seconds in a day
(24 h × 60 min × 60 s). Dividing a per-day number by it gives the per-second rate,
spread evenly across the whole day.

### Step 4 — Peak QPS

```
 3,500 × 2 ≈ 7,000 QPS
```

Traffic clusters in waking hours, so the peak is higher than the average. The **2× is an
assumption for this exercise** — a different multiplier is equally defensible if stated.

Note: this is **write** QPS (tweets posted). Read QPS (timelines viewed) would be far
higher — easily 100× — which is the real driver of the design.

### Step 5 — Media tweets per day

```
 300M tweets × 10% = 30M media tweets/day
```

### Step 6 — Media storage per day

```
 30M × 1 MB = 30M MB ≈ 30 TB/day
```

(1M MB = 1 TB.) Text is tiny by comparison: 300M × ~300 bytes ≈ 90 GB/day, ~0.3% of the media.

### Step 7 — Five-year storage

```
 30 TB × 365 × 5 ≈ 54,750 TB ≈ 55 PB
```

And that's **before** replication and backups — with 3 copies it's ~165 PB.

### Full flow

```
 300M MAU
   ↓ × 50%
 150M DAU
   ↓ × 2 tweets
 300M tweets/day
   ↓ ÷ 86,400
 ~3,500 QPS (average)
   ↓ × 2 (peak assumption)
 ~7,000 QPS (peak)
```

```
 300M tweets/day
   ↓ × 10% media
 30M media tweets/day
   ↓ × 1 MB
 30 TB/day
   ↓ × 365 × 5
 ~55 PB (5-year media storage)
```

### What the numbers tell the design

| Number | Implication |
|---|---|
| ~7,000 peak write QPS | Beyond one DB — shard writes, buffer with a queue |
| Read QPS far higher | Heavy caching + read replicas + fan-out on write |
| 30 TB/day media | Object storage (S3-style) + CDN, not the main database |
| ~55 PB over 5 years | Retention policy, tiered/cold storage, compression matter a lot |

---

## 12. Estimation Workflow

A reusable order to follow in an interview.

| Step | Do this | Example |
|---|---|---|
| 1 | **Clarify requirements** — what are we estimating and why? | "Write-heavy or read-heavy? Media included?" |
| 2 | **State assumptions** out loud | "300M MAU, 50% daily active" |
| 3 | **Estimate DAU** | 300M × 50% = 150M |
| 4 | **Actions per day** | 150M × 2 = 300M tweets/day |
| 5 | **Convert to QPS** | 300M ÷ 86,400 ≈ 3,500 |
| 6 | **Peak QPS** | 3,500 × 2 ≈ 7,000 |
| 7 | **Data generated per day** | 30M media × 1 MB = 30 TB/day |
| 8 | **Total storage** | 30 TB × 365 × 5 ≈ 55 PB |
| 9 | **Bandwidth** | QPS × payload size, then × 8 for bits |
| 10 | **Apply replication / backups / growth** | ×3 replicas → ~165 PB |
| 11 | **Use results to shape the architecture** | Shard the DB, cache reads, CDN for media |

```
 Requirements
     ↓
 Assumptions
     ↓
 DAU
     ↓
 Requests/day
     ↓
 QPS
     ↓
 Peak QPS
     ↓
 Storage + Bandwidth
     ↓
 Infrastructure
```

Tip: write the assumptions in a corner of the whiteboard and keep them visible. When the
interviewer changes one ("what if media is 20%?"), you only redo the arithmetic.

---

## 13. Conversion Cheat Sheet

**Time**

```
 1 hour  =  3,600 seconds
 1 day   = 86,400 seconds
 1 month ≈ 2.6M seconds
 1 year  ≈ 31.5M seconds

 1 s  = 1,000 ms
 1 ms = 1,000 μs
 1 μs = 1,000 ns
```

**Storage**

```
 1 byte = 8 bits

 2^10 ≈ 1K  → 1 KB
 2^20 ≈ 1M  → 1 MB
 2^30 ≈ 1G  → 1 GB
 2^40 ≈ 1T  → 1 TB
 2^50 ≈ 1P  → 1 PB

 1M KB = 1 GB     1M MB = 1 TB     1M GB = 1 PB
```

**Network**

```
 1 byte = 8 bits
 MB/s × 8 = Mbps
 Mbps ÷ 8 = MB/s
 1 Gbps ≈ 125 MB/s
```

**QPS**

```
 Average QPS = requests/day ÷ 86,400
 Peak QPS    = average QPS × peak multiplier (assumption, often 2)
```

**Common shortcuts**

```
 1M requests/day    ≈    12 QPS
 100M requests/day  ≈ 1,160 QPS
 1B requests/day    ≈ 11,600 QPS
```

---

## 14. Common Mistakes

| # | Mistake | Fix |
|---|---|---|
| 1 | Confusing **MB** with **Mb** | Uppercase B = bytes, lowercase b = bits; factor of 8 |
| 2 | Forgetting to convert **days → seconds** | Always ÷ 86,400 for QPS |
| 3 | Forgetting **peak** traffic | Size for peak, not average |
| 4 | Treating **assumptions as facts** | Say "assuming…" every time |
| 5 | **Overcomplicating** a simple estimate | 3 multiplications beat a spreadsheet |
| 6 | Chasing **exact numbers** | Round hard; order of magnitude is the goal |
| 7 | Forgetting the **retention period** | "Per day" is meaningless without "for how long" |
| 8 | Ignoring **media size** | Media usually dominates storage and bandwidth |
| 9 | Ignoring the **read/write ratio** | It decides caching, replicas and sharding |
| 10 | Forgetting **replication/backups** in real storage | Raw × 3 (or more) for the real number |
| 11 | Confusing **latency** with **throughput** | See below — they're independent |
| 12 | Treating old **latency reference numbers** as modern benchmarks | They show relative cost, not today's hardware |

---

## 15. Latency vs Throughput

| | Latency | Throughput |
|---|---|---|
| Question | How long does **one** operation take? | How much work per unit time? |
| Unit | ms, μs, ns | requests/sec, MB/s |
| Felt by | A single user waiting | The system under load |
| Improved by | Caching, fewer round trips, faster queries, closer servers | More servers/workers, parallelism, batching, bigger pipes |

```
 Latency:    one request takes 50 ms
 Throughput: the system handles 1,000 requests/sec
```

They are **not** the same number and don't move together:

- **Low latency, low throughput** — one fast server that can only handle 10 requests/sec.
- **High latency, high throughput** — a batch pipeline: each job takes minutes, but millions finish per hour.
- **Batching** often **raises** throughput while **raising** latency for the individual request.
- **Adding servers** raises throughput but does **nothing** for a single request's latency.

### Analogy

> A **highway**: latency = how long your car takes to drive the route.
> Throughput = how many cars pass per minute.
> Adding lanes moves more cars (throughput) without making any single car faster (latency).

### Related: percentiles

Average latency hides the bad cases — always ask for **p95 / p99**.

```
 p99 = 500 ms  →  1 in 100 requests is slower than 500 ms
```

On a page making 10 backend calls, a p99 of 500 ms means many page loads hit at least one
slow call — which is why tail latency matters more than the average.

> **Memory:**
> Latency = time per operation.
> Throughput = operations per time.
