import asyncio


async def fetch(name: str, delay: float) -> str:
    await asyncio.sleep(delay)  # stands in for network I/O
    return f"{name} done"


async def main():
    # sequential: ~2s
    # await fetch("a", 1); await fetch("b", 1)

    # concurrent: ~1s (same idea as Promise.all in JS)
    results = await asyncio.gather(fetch("a", 1), fetch("b", 1))
    print(results)

    # timeout (like Promise.race with a timer)
    try:
        await asyncio.wait_for(fetch("slow", 5), timeout=0.1)
    except asyncio.TimeoutError:
        print("timed out")


asyncio.run(main())
