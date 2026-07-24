#step 1
initializing the project

uv init app_name

this project will generate  - pyproject.toml file

step 2 - now install dependencies (for e,g)

uv add --dev mypy
(uv run mypy .)
this will create a .venv folder - equivalent to node_modules in node

Write Code
      │
      ▼
ruff check
      │
      ▼
mypy
      │
      ▼
pytest
      │
      ▼
Git Push