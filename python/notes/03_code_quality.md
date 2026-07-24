Topic 1 — Ruff (5 minutes)

Think of Ruff as combining the roles of ESLint + Prettier for Python.

Suppose you write:

import os
import sys

x = 10

print("Hello")

Problems:

os isn't used.
sys isn't used.
Variable names might not follow conventions.
Formatting might be inconsistent.

Running Ruff:

ruff check .

It reports issues like:

F401 'os' imported but unused
F401 'sys' imported but unused

To automatically fix many issues:

ruff check . --fix

To format code:

ruff format .
Why Every Company Uses It

Before code is merged:

Developer
      │
      ▼
ruff check
      │
      ▼
No issues?
      │
      ▼
Git Commit

This keeps the entire codebase consistent.