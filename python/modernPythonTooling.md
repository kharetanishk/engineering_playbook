What is pip?
Python ka default package installer. pip install fastapi → package current environment mein aa jaata hai. Bas itna hi kaam. Python ke saath bundled aata hai.

What is uv?
Astral (ruff wale) ka banaya ek Rust-based package + project manager. Ek hi binary jo pip, virtualenv, pip-tools, pipx aur pyenv — sabka kaam karta hai.

Difference

               pip	            uv
Kaam	   irf install	install + venv + Python versions + lockfile
Speed	   slow	        10–100x fast (Rust, parallel, global cache)
Lockfile   nahi (pip freeze sirf snapshot)	uv.lock — cross-platform + hashes
Env	active venv mein daal deta hai	project ka .venv khud dhoondh/bana leta hai

Kaun better?
uv — for new projects. Reason: one tool instead of four, reproducible builds via real lockfile, aur installs jo turant khatam ho jaate hain. pip tab better jab environment restricted ho (old CI images, corporate setups) kyunki wo har jagah pehle se maujood hai.

Modern practice
pyproject.toml + uv add + uv.lock committed to git + uv sync in CI + uv run to execute (venv activate karne ki zarurat nahi). requirements.txt ab legacy hai — chahiye toh uv pip compile se generate kar lo.

Think of it like this:

Node	        Python
package.json	pyproject.toml
pnpm-lock.yaml	uv.lock
node_modules	.venv
pnpm add	    uv add
pnpm install	uv sync
pnpm exec	    uv run

What is uv?

Think of it as

pnpm + nvm + virtualenv + poetry

in one executable.

It can

✅ Install Python

✅ Create virtual environments

✅ Install packages

✅ Lock dependencies

✅ Run scripts

✅ Manage versions


TOML stands for Tom’s Obvious, Minimal Language and is a configuration file format designed to be easy for humans to read and write. 

Virtual Environment (The .venv Folder)

This is one concept every Python developer must understand.

Suppose you have:

Project A

FastAPI 0.115

and

Project B

FastAPI 0.90

If both install packages globally, they'll conflict.

Instead, every project gets its own isolated environment:

todo-api/

.venv/

Think of .venv as a private installation of Python packages for that project.

Node vs Python

Node:

node_modules/

contains packages.

Python:

.venv/

contains the interpreter and installed packages.

That's a key difference.

.venv is a project-local folder containing its own Python interpreter and its own site-packages (where packages get installed) — isolation happens because that interpreter's sys.path only points to its own site-packages, not the system's or any other project's, so each project can have different package versions without conflicting — same reason node_modules exists in JS.

Running Python

Instead of

python app.py

with uv, you typically use

uv run app.py

Why?

Because uv run automatically uses the correct Python version and the project's virtual environment, so you don't have to activate .venv manually.

Installing Packages

Node

pnpm add fastapi

Python

uv add fastapi

This updates:

pyproject.toml
uv.lock

and installs the package into the project's .venv.

uv.lock is a file uv auto-generates that records the exact resolved dependency tree for your project — every package (direct + transitive), its precise version, hash, and which other packages it needs.

It's called a graph, not a list, because packages depend on other packages, which depend on more packages — so the structure isn't flat, it branches (fastapi → starlette → anyio, etc.), forming a network of dependencies rather than a simple line. uv.lock is that network, frozen to exact versions so every machine installs the identical tree.

small python fundamental - 
Variables never own objects.
They only refer to objects.

rule - Never use 'is' to compare numbers or strings.

Mutable Objects (can be modified)
list
dict
set
bytearray

Immutable Objects (cannot be modified)
int
float
bool
str
tuple
frozenset