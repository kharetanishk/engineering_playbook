#old way

num1 = [1,2,3,4]
twice_num1 = []
for n in num1:
    twice_num1.append(n * 2)
print(twice_num1)

#pythonic way

twice_num1_new = [n * 2 for n in num1 ]
print(twice_num1_new)

#adding a condition
even = [n for n in num1 if n % 2 == 0]
print(even)

#strings
names =["tanishk","shreya","omi"]
caps = [n.upper() for n in names]
print(caps)

#dictionary comprehension

#old
squares = {}
for x in range(5):
    squares[x] = x*x
print(squares)

#pythonic way 
squares_1 = {
    x:x*x
    for x in range(5)
}
print(squares_1)

#questions
# nums = [1,2,3,4,5] if these are nums 
# write a list comprehension to produce [2,4,6,8,10]
nums_2 = [1,2,3,4,5]
list_1 = [x*2 for x in nums_2]
print(list_1)

#list comprehension to produce this [1,9,25]
list_2 = [x*x for x in nums_2 if x % 2 != 0]
print(list_2)
# Q3

# Given

# names = ["tanu","john","alice"]

# Produce

# ["TANU","JOHN","ALICE"]

# using a comprehension.

names = ["tanu","john","alice"]
capitalize = [c.upper() for c in names]
print(capitalize)

# Given

# scores = {
#     "math":90,
#     "science":95
# }

# Write a dictionary comprehension that adds 10 to every score.
scores = {
     "math":90,
     "science":95
 }

add_10 = {
    x: scores[x]+10
    for x in scores
}
print(add_10)

add_10_2 = {
    sub : score + 10
    for sub , score in scores.items()
}
print(add_10_2)

logs_of_million = (f"Log entry {i}" for i in range(1000000))

print(logs_of_million)

#generatrors and yield

#generators are used to save memory and time and generate and give values
#when asked , they are used to read logs of big files and also in case infinite sequences


#wheneve yield is called generator function is paused and the value is returned to the caller and when next is called the function resumes from where it left off
def numbers():
    for i in range(10):
        yield i#Key point: yield pauses the function. The next next() resumes from exactly where it stopped.

gen = numbers()
print(next(gen))
print(next(gen))\
# ✅ Generator Complete

# Remember:

# return → function ends.
# yield → function pauses.

#decorators

def logger(func):
    def wrapper():
        print("function execution started")
        func()
        print("function execution stopped")

    return wrapper



#usage
@logger
def greet():
    print("hello world")

greet()

#type hint
list_of_integers : list[int] = [1,2,3,4,]



dict_with_type : dict[str,float] = {
    "marks" : 90.0,
}

integer : int | None = None


def greet(name: str)-> str:
    return f"Hello {name}"