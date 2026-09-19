import time
from collections import OrderedDict, deque


class LRUCache:
    """O(1) get/put. OrderedDict keeps insertion order; move_to_end = 'recently used'."""

    def __init__(self, capacity: int) -> None:
        self.capacity = capacity
        self.data: OrderedDict[int, int] = OrderedDict()

    def get(self, key: int) -> int:
        if key not in self.data:
            return -1
        self.data.move_to_end(key)
        return self.data[key]

    def put(self, key: int, value: int) -> None:
        self.data[key] = value
        self.data.move_to_end(key)
        if len(self.data) > self.capacity:
            self.data.popitem(last=False)  # evict least recently used


class RateLimiter:
    """Sliding window: allow at most `limit` calls per `window` seconds."""

    def __init__(self, limit: int, window: float) -> None:
        self.limit = limit
        self.window = window
        self.calls: deque[float] = deque()

    def allow(self) -> bool:
        now = time.monotonic()
        while self.calls and now - self.calls[0] >= self.window:
            self.calls.popleft()
        if len(self.calls) < self.limit:
            self.calls.append(now)
            return True
        return False


if __name__ == "__main__":
    c = LRUCache(2)
    c.put(1, 1)
    c.put(2, 2)
    c.get(1)
    c.put(3, 3)  # evicts 2
    assert c.get(2) == -1 and c.get(1) == 1

    r = RateLimiter(2, 1.0)
    assert [r.allow() for _ in range(3)] == [True, True, False]
    print("ok")
