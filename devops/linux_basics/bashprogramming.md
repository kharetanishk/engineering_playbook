# Bash Programming Notes

# 1. What is Bash?

**Bash (Bourne Again Shell)** is:

* A shell (command interpreter)
* A scripting language
* An interface between the user and the Linux kernel

Flow:

```text
User
   │
   ▼
Terminal
   │
   ▼
Bash Shell
   │
   ▼
Programs (ls, cat, mkdir...)
   │
   ▼
Linux Kernel
   │
   ▼
Hardware
```

---

# 2. Shebang

Every Bash script should start with:

```bash
#!/bin/bash
```

This tells Linux to execute the script using Bash.

---

# 3. Running a Script

Give execute permission:

```bash
chmod +x script.sh
```

Run:

```bash
./script.sh
```

Or execute directly with Bash:

```bash
bash script.sh
```

---

# 4. Variables

Assign a value:

```bash
name="Tanishk"
age=22
```

Print:

```bash
echo "$name"
echo "$age"
```

Empty variable:

```bash
name=""
```

or

```bash
name=
```

Delete a variable:

```bash
unset name
```

---

# 5. Variable Rules

Correct:

```bash
firstName="Tanishk"
_age=22
```

Wrong:

```bash
2name="abc"
my name="abc"
```

Variable names can contain:

* Letters
* Numbers
* Underscore (_)

Cannot start with a number.

---

# 6. Comments

Single-line:

```bash
# This is a comment
```

There is no true multi-line comment syntax.

---

# 7. Printing Output

```bash
echo "Hello"

echo "$name"

echo "Age = $age"
```

---

# 8. Quotes

## Double Quotes

Variables expand.

```bash
name="Tanishk"

echo "Hello $name"
```

Output:

```
Hello Tanishk
```

---

## Single Quotes

Nothing expands.

```bash
echo '$name'
```

Output:

```
$name
```

---

# 9. Command Substitution

Run a command and store its output.

```bash
files=$(ls)

currentUser=$(whoami)

today=$(date)
```

---

# 10. User Input

```bash
read name
```

Prompt:

```bash
read -p "Enter your name: " name
```

Password:

```bash
read -sp "Password: " password
echo
```

---

# 11. Positional Parameters

Suppose:

```bash
./script.sh tanishk devops 22
```

Then:

```
$1 -> tanishk
$2 -> devops
$3 -> 22
```

Useful variables:

```
$1      First argument
$2      Second argument
$#      Number of arguments
$@      All arguments (recommended)
$*      All arguments as one string
$?      Exit status
$$      Current Process ID
```

---

# 12. Arithmetic

Arithmetic expansion:

```bash
sum=$((10+20))
```

Using variables:

```bash
a=10
b=20

echo $((a+b))
```

Operations:

```
+
-
*
/
%
**
```

Increment:

```bash
((count++))
```

Decrement:

```bash
((count--))
```

---

# 13. if Statement

```bash
if [ condition ]; then
    commands
fi
```

Example:

```bash
if [ "$name" = "Tanishk" ]; then
    echo "Found"
fi
```

---

# 14. if-else

```bash
if [ condition ]; then
    commands
else
    commands
fi
```

---

# 15. if-elif-else

```bash
if [ condition ]; then
    ...
elif [ condition ]; then
    ...
else
    ...
fi
```

---

# 16. File Tests

```
-f  Regular file

-d  Directory

-e  Exists

-r  Readable

-w  Writable

-x  Executable

-s  File not empty
```

Example:

```bash
if [ -f "config.yaml" ]; then
    cat config.yaml
fi
```

---

# 17. String Operators

Equal

```
=
```

Bash also allows

```
==
```

Not Equal

```
!=
```

Empty string

```
-z
```

Not empty

```
-n
```

Examples:

```bash
if [ -z "$name" ]; then
```

```bash
if [ -n "$name" ]; then
```

---

# 18. Numeric Operators

```
-eq   equal

-ne   not equal

-gt   greater than

-lt   less than

-ge   greater than or equal

-le   less than or equal
```

Example:

```bash
if [ "$age" -ge 18 ]; then
```

Modern Bash:

```bash
if (( age >= 18 )); then
```

---

# 19. Logical Operators

AND

```bash
&&
```

OR

```bash
||
```

NOT

```bash
!
```

Example:

```bash
if [ "$age" -ge 18 ] && [ "$age" -le 30 ]; then
```

---

# 20. Loops

## for

```bash
for fruit in apple banana mango
do
    echo "$fruit"
done
```

---

C-style loop

```bash
for ((i=1;i<=5;i++))
do
    echo "$i"
done
```

---

Loop through arguments

```bash
for arg in "$@"
do
    echo "$arg"
done
```

---

# 21. while Loop

```bash
count=1

while [ $count -le 5 ]
do
    echo "$count"
    ((count++))
done
```

---

# 22. until Loop

Runs until the condition becomes true.

```bash
count=1

until [ $count -gt 5 ]
do
    echo "$count"
    ((count++))
done
```

---

# 23. break

```bash
break
```

Stops the loop.

---

# 24. continue

```bash
continue
```

Skips the current iteration.

---

# 25. case Statement

```bash
case $choice in

1)
echo "Start"
;;

2)
echo "Stop"
;;

*)
echo "Invalid"
;;

esac
```

---

# 26. Functions

```bash
greet() {

echo "Hello"

}
```

With parameters:

```bash
greet(){

echo "Hello $1"

}

greet Tanishk
```

---

# 27. Arrays

```bash
languages=("Bash" "Python" "Java")
```

Access:

```bash
echo "${languages[0]}"
```

Loop:

```bash
for lang in "${languages[@]}"
do
    echo "$lang"
done
```

---

# 28. Redirection

Overwrite

```bash
>
```

Append

```bash
>>
```

Input

```bash
<
```

Pipe

```bash
|
```

Examples:

```bash
echo "Hello" > file.txt
```

```bash
echo "Hello" >> file.txt
```

---

# 29. Exit Status

Every command returns an exit code.

Success

```
0
```

Failure

```
Non-zero
```

Check:

```bash
echo $?
```

---

# 30. chmod

Add execute

```bash
chmod u+x script.sh
```

Only owner gets rwx

```bash
chmod 700 script.sh
```

Numeric permissions

```
7 = rwx

6 = rw-

5 = r-x

4 = r--

3 = -wx

2 = -w-

1 = --x

0 = ---
```

---

# 31. chown

Owner only

```bash
sudo chown tanishk file.txt
```

Owner and group

```bash
sudo chown tanishk:devops file.txt
```

---

# 32. Common Bash Syntax

Variable

```bash
name="Tanishk"
```

Read

```bash
read name
```

Prompt

```bash
read -p "Name: " name
```

Print

```bash
echo "$name"
```

Arithmetic

```bash
$((a+b))
```

Condition

```bash
if [ condition ]
```

Loop

```bash
for item in list
```

Function

```bash
functionName(){

}
```

---

# 33. Best Practices

* Always start scripts with `#!/bin/bash`
* Quote variables: `"$name"`
* Use meaningful variable names.
* Indent code consistently (Bash ignores indentation, but humans don't).
* Prefer `$(command)` over backticks.
* Use `[[ ... ]]` for Bash-specific scripts and `[ ... ]` with `=` for POSIX compatibility.
* Check command exit codes when writing automation.
* Test scripts before running them with `sudo`.

---

# 34. Real-World Bash Use Cases

* Server provisioning
* User and group management
* Package installation
* Backups
* Log cleanup
* File management
* Deployment scripts
* Monitoring CPU, RAM, and disk usage
* Cron jobs
* Docker setup
* Git automation
* CI/CD helper scripts
* Environment configuration
* Health checks
* Service restarts

---

# 35. Learning Roadmap

1. Variables
2. Input (`read`)
3. Positional parameters
4. Operators
5. Conditions
6. File tests
7. Loops
8. Functions
9. Arrays
10. Case statements
11. Redirection
12. Exit codes
13. Real-world automation projects

Master these fundamentals before moving on to advanced topics like regular expressions, process management, signals, traps, `find`, `xargs`, `sed`, `awk`, and shell scripting best practices for DevOps.
