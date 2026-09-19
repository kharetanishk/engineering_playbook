Generators

- `yield` pauses a function and hands back a value; state is kept between calls.
- Lazy: values are produced on demand, so memory is O(1) and infinite sequences work.
- Generator expression `(x for x in xs)` vs list comprehension `[x for x in xs]`.
- `yield from` delegates to another iterable.
- A generator is exhausted after one pass.

Code: code_python/generators.py
