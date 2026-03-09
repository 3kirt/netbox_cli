# Ship

Complete the implement → verify → commit → merge → push cycle for the current feature branch.

## Steps

1. Run `gofmt -w .` and fix any formatting issues.
2. Run `go build ./...` and fix any compilation errors.
3. Run `make lint` and fix all reported issues.
4. Stage all relevant changes (avoid committing secrets or generated binaries).
5. Commit with a concise conventional commit message.
6. Merge the current feature branch into `main`: `git checkout main && git merge <branch>`.
7. Push to origin: `git push`.
8. Report what was shipped.
