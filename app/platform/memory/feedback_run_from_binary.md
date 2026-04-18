---
name: Go services run from binary
description: Always run Go services from a compiled binary, not via go run
type: feedback
---

Always run Go services from the compiled binary, not `go run`. In Makefiles, the `run-*` target must have the corresponding `build-*` target as a prerequisite.

**Why:** Project convention — services are always executed from their binary.

**How to apply:** For any `run-X` Makefile target in this repo, add `build-X` as a prerequisite and invoke `./bin/X` instead of `go run ./...`.
