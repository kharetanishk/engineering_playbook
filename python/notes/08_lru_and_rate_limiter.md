LRU cache and rate limiter (write-from-scratch questions)

- LRU: OrderedDict gives O(1) get/put; `move_to_end` on use, `popitem(last=False)` to evict.
- Rate limiter: sliding window with a deque of timestamps; drop the old ones, then check the count.
- Use `time.monotonic()`, not `time.time()`, so clock changes don't break it.
- Production: `functools.lru_cache` for caching; token bucket or Redis for distributed limits.

Code: code_python/lru_and_rate_limiter.py
