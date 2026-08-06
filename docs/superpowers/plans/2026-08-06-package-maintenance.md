# Package Maintenance Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Update every managed dependency in the Go, Python, TypeScript, and repository hook configurations to the latest compatible/current release and verify the repository.

**Architecture:** Preserve each starter's existing structure and behavior. Use each ecosystem's native resolver to update direct and transitive dependencies, update declared direct-version floors where applicable, and only change source code if an upstream compatibility break requires it.

**Tech Stack:** Go modules, uv, npm, prek, Go test, pytest, Node.js test runner, TypeScript.

## Global Constraints

- The Go implementation remains the behavioral reference.
- Shared environment keys and runtime behavior remain unchanged.
- Dependency and lock files must agree after resolution.
- Existing user changes must be preserved.

---

### Task 1: Go dependencies

**Files:**

- Modify: `golang/go.mod`
- Modify: `golang/go.sum`

- [x] Run `go get -u ./...` and `go mod tidy` in `golang/`.
- [x] Run `go test ./... -race -cover` and confirm zero failures.

### Task 2: Python dependencies

**Files:**

- Modify: `python/pyproject.toml`
- Modify: `python/uv.lock`

- [x] Resolve the newest direct releases and update the lower bounds in `pyproject.toml`.
- [x] Run `uv lock --upgrade` and `uv sync` in `python/`.
- [x] Run `uv run pytest` and confirm zero failures.

### Task 3: TypeScript dependencies

**Files:**

- Modify: `ts/package.json`
- Modify: `ts/package-lock.json`

- [x] Query npm for the newest direct releases and update `package.json`.
- [x] Run `npm install` in `ts/` to refresh the lockfile.
- [x] Run `npm test`, `npm run typecheck`, and `npm run build` and confirm zero failures.

### Task 4: Repository hooks and final verification

**Files:**

- Modify: `prek.toml`

- [x] Run `prek auto-update` to update hook repositories.
- [x] Run `prek run --all-files` and address only update-related failures.
- [x] Run the root test target and dependency-manager outdated checks.
- [x] Review `git diff --check`, `git status --short`, and the final diff for unintended changes.
