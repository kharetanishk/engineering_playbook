from itertools import islice


# a generator yields lazily - nothing runs until you ask for the next value
def countdown(n: int):
    while n > 0:
        yield n
        n -= 1


print(list(countdown(3)))  # [3, 2, 1]


# infinite sequence - only possible because it is lazy
def fib():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b


print(list(islice(fib(), 8)))  # [0, 1, 1, 2, 3, 5, 8, 13]

# generator expression: like a list comprehension but lazy, O(1) memory
total = sum(n * n for n in range(1_000_000))
print(total)


# pipeline: read -> filter -> transform, one item at a time
def read_lines(lines):
    yield from lines


def non_empty(lines):
    return (line for line in lines if line.strip())


print(list(non_empty(read_lines(["a", "", "b"]))))  # ['a', 'b']
