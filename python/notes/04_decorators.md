Decorator Pattern


def decorator(func):

    def wrapper(*args, **kwargs):
        # Before
        result = func(*args, **kwargs)
        # After
        return result

    return wrapper

This is the template you'll see everywhere.

usage - @ - shorthand to use 

e.g - 
def logger(func):
    def wrapper():
        print("Started")
        func()
        print("Ended")

    return wrapper

usage - 
@logger
def hello():
    print("Hello")


now hello() - outputs-  Started
Hello
Ended

here - @ is just shorthand for  - hello = logger(hello)


🚀 Topic: Context Managers (with)

Instead of

file = open("data.txt")

content = file.read()

file.close()

What if an exception occurs before close()?

The file remains open.

Python solution:

with open("data.txt") as file:
    content = file.read()

When execution leaves the with block—even if an exception occurs—Python automatically closes the file.

Used everywhere:

with open(...)
with lock:
with Session(engine) as session:
with transaction():