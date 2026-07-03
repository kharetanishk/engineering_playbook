# Bash Special Variables (Special Parameters)

Special variables are predefined by Bash. They provide information about the current script, command-line arguments, process, and command execution.

---

# 1. `$1`, `$2`, `$3` ...

Represents positional parameters.

Example:

```bash
#!/bin/bash

echo "First : $1"
echo "Second: $2"
echo "Third : $3"
```

Run:

```bash
./script.sh Tanishk DevOps Linux
```

Output:

```text
First : Tanishk
Second: DevOps
Third : Linux
```

---

# 2. `$#`

Returns the total number of positional parameters.

Example:

```bash
echo $#
```

Run:

```bash
./script.sh a b c d
```

Output:

```text
4
```

Very useful for validating input.

Example:

```bash
if [ $# -lt 2 ]; then
    echo "Usage: ./script.sh <username> <group>"
    exit 1
fi
```

---

# 3. `$*`

Represents **all positional parameters** as one string.

Example:

```bash
echo "$*"
```

Run:

```bash
./script.sh apple banana mango
```

Output:

```text
apple banana mango
```

---

# 4. `$@`

Represents **all positional parameters**, preserving each one as a separate argument.

Example:

```bash
for arg in "$@"
do
    echo "$arg"
done
```

Run:

```bash
./script.sh "apple pie" banana
```

Output:

```text
apple pie
banana
```

This is the preferred way to process all arguments.

---

# Difference: `$*` vs `$@`

Suppose:

```bash
./script.sh "apple pie" banana
```

Using:

```bash
echo "$*"
```

Produces one string:

```text
apple pie banana
```

Using:

```bash
for arg in "$@"
do
    echo "$arg"
done
```

Produces:

```text
apple pie
banana
```

**Recommendation:** Prefer `"$@"`.

---

# 5. `$?`

Returns the exit status of the last executed command.

Example:

```bash
mkdir demo

echo $?
```

Output:

```text
0
```

If the command fails:

```bash
mkdir demo
mkdir demo

echo $?
```

Output:

```text
1
```

Convention:

```text
0  -> Success
Non-zero -> Failure
```

Used for error checking:

```bash
cp file.txt backup/

if [ $? -eq 0 ]; then
    echo "Copied successfully"
else
    echo "Copy failed"
fi
```

---

# 6. `$$`

Returns the Process ID (PID) of the current script.

Example:

```bash
echo $$
```

Output:

```text
5421
```

Useful for:

* Temporary filenames
* Logging
* Debugging

Example:

```bash
touch temp_$$.txt
```

Possible output:

```text
temp_5421.txt
```

---

# 7. `$!`

Returns the Process ID of the most recently started background process.

Example:

```bash
sleep 100 &
echo $!
```

Output:

```text
6482
```

Useful when managing background jobs.

---

# 8. `$0`

Returns the name (or path) of the current script.

Example:

```bash
echo $0
```

Run:

```bash
./backup.sh
```

Output:

```text
./backup.sh
```

Useful for usage messages:

```bash
echo "Usage: $0 <filename>"
```

---

# 9. `$_`

Contains the last argument of the previous command.

Example:

```bash
mkdir project
echo $_
```

Output:

```text
project
```

Another example:

```bash
touch notes.txt
cat $_
```

Equivalent to:

```bash
cat notes.txt
```

---

# 10. `$-`

Shows the current shell options.

Example:

```bash
echo $-
```

Possible output:

```text
himBH
```

Useful for debugging shell behavior.

---

# 11. `$IFS`

Internal Field Separator.

Default value:

```text
space
tab
newline
```

Bash uses it to split words.

Normally, you don't change it unless you're doing advanced text processing.

---

# Summary Table

| Variable | Description                            |
| -------- | -------------------------------------- |
| `$0`     | Current script name                    |
| `$1`     | First argument                         |
| `$2`     | Second argument                        |
| `$3`     | Third argument                         |
| `$n`     | nth argument                           |
| `$#`     | Number of arguments                    |
| `$*`     | All arguments as one string            |
| `$@`     | All arguments separately (recommended) |
| `$?`     | Exit status of previous command        |
| `$$`     | Current script's Process ID            |
| `$!`     | Last background process ID             |
| `$_`     | Last argument of previous command      |
| `$-`     | Current shell options                  |
| `$IFS`   | Internal Field Separator               |

---

# Most Frequently Used in Real Projects

You'll use these constantly:

```text
$0
$1
$2
$#
$@
$?
$$
```

These seven cover the vast majority of Bash scripting in DevOps and automation.

---

# Example Script

```bash
#!/bin/bash

echo "Script Name : $0"
echo "Arguments   : $#"
echo "All Args    : $@"
echo "First Arg   : $1"

mkdir demo

echo "Exit Status : $?"
echo "PID         : $$"
```
