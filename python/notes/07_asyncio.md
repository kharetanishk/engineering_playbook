asyncio

- `async def` makes a coroutine; `await` yields control while waiting on I/O.
- `asyncio.run(main())` is the entry point.
- `asyncio.gather(...)` runs coroutines concurrently (like `Promise.all`).
- `asyncio.wait_for(coro, timeout)` adds a timeout.
- Best for I/O-bound work. CPU-bound work: use multiprocessing (the GIL blocks threads).

Code: code_python/async_basics.py
