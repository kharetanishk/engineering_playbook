🚀 Next Topic — Dataclasses

Without dataclasses, suppose you want a User.

class User:
    def __init__(self, name: str, age: int):
        self.name = name
        self.age = age

Create one:

user = User("Tanu", 22)

This works, but you're writing a lot of boilerplate.

Dataclass
from dataclasses import dataclass

@dataclass
class User:
    name: str
    age: int

Now

user = User("Tanu", 22)

print(user)

Output

User(name='Tanu', age=22)

Notice something?

You never wrote an __init__() method.

The @dataclass decorator generated it for you.

What Does @dataclass Generate?

This:

@dataclass
class User:
    name: str
    age: int

is roughly equivalent to writing:

class User:
    def __init__(self, name: str, age: int):
        self.name = name
        self.age = age

    def __repr__(self):
        ...

    def __eq__(self):
        ...

Automatically.

Why Is It Useful?

Imagine a class with 15 fields.

Without dataclasses, you'd manually assign all 15 in __init__.

With @dataclass, you just declare the fields.

@dataclass automatically generates boilerplate methods like __init__(), __repr__(), and __eq__() based on the fields you define, allowing you to write much less code.