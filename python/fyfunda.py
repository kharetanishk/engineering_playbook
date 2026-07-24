# keyword only and positional only parameters


def greet(name, age):
    return f"Hello {name}, you are {age} years old"


greet(age=22, name="tanishk")
greet("tanishk", 22)


# Keyword-only Parameters
# def create_user(name, *, age):
#     print(name, age)

# Notice the *.

# Call

# create_user("Tanu", age=22)

# ✅ Works

# But

# create_user("Tanu",22)

# ❌ Error

# The * means:

# Everything after me must be passed using keywords.

# Positional-only Parameters

# Opposite.

# def divide(a,b,/):
#     return a/b

# Now

# divide(10,2)

# ✅

# But

# divide(a=10,b=2)

# ❌

# The / means

# Everything before me is positional only.


def addItems(item, items=None):
    if items is None:
        items = []
    items.append(item)
    return items


print(addItems(1))
print(addItems(2))
print(addItems(3))
