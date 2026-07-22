user = ("tanishk",22)

#unpack tupule
name, age  = user
print(name)
print(age)

#star unpacking

nums = [1,2,3,4,5,6]

first , *middle , last = nums
print(first)
print(*middle)
print(last)

#upacking dictionareis
user = {
    "name" : "tanishk",
    "age":22
}
extra = {
    "city": "bhilai"
}
#merge
result = {
    **user,
    **extra
}
print(result)

#functions

def greet(name="tanishk"):#default paraemeter (fallbck if not arg pass)
    print(f"hey ${name}")


greet(name="suresh")#keyword argument(better reading)\

#. *agrs
#sometimes we dont know how many arguments we will recieve 

def add(*nums):
    print(nums)#return a tupule
    print(sum(nums))


add(1,2,3,4,5,5)


def dicdata(**data):
    print(data)
dicdata(name="tanishk" , age=22 , hobbie=["playing cricket" , "coding.."])#this iwill give out dictionary

# ⭐ One Important Rule
# Syntax	Type
# *args	tuple
# **kwargs	dict
# *rest (unpacking)	list

