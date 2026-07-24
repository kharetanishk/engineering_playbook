from dataclasses import dataclass
from enum import Enum


@dataclass
class User:
    name: str
    age: int


user = User("tanishk", 22)
print(user.name)
print(user.age)

# enums


class Color(Enum):
    RED = 1
    GREEN = 2
    BLUE = 3


if color := Color.RED:
    print(color.name)
    print(color.value)


class Role(Enum):
    ADMIN = "admin"
    USER = "user"


print(role := Role.ADMIN)

# exceptions created by own


def banking(amount: int, balance: int):
    if amount > balance:
        raise ValueError("insufficient funds")
    b = balance - amount
    return f"remainig balance {b}"


print(balance := banking(10, 80))


def mypytest(name : str)-> str:
    return name

person: int = mypytest("tanishk")
print(person)

def greet(name: str) -> str:
    return f"Hello {name}"

age: int = "22"