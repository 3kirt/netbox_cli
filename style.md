# Go Style Guide Summary

Practical summary of the style rules applied in this project. Authoritative sources:

- [Effective Go](https://go.dev/doc/effective_go)
- [Google Go Style Guide](https://google.github.io/styleguide/go/guide)
- [Google Go Best Practices](https://google.github.io/styleguide/go/best-practices)
- [Google Go Decisions](https://google.github.io/styleguide/go/decisions)

Compliance is enforced mechanically by `make lint` (`golangci-lint`, config in `.golangci.yml`).

---

## Formatting

- **`gofmt`** is mandatory. Tabs for indentation, not spaces.
- Run `gofmt -w .` to reformat; `gofmt -l .` to check (no output = clean).
- `.editorconfig` has a `[*.go]` section that matches `gofmt` settings.

---

## Naming

- **MixedCaps** throughout — no underscores in identifiers (`MaxRetries`, not `max_retries`).
- **Acronyms** are all-caps or all-lower consistently: `URL`, `HTTP`, `ID`, `urlStr`, `httpClient`, `userID`. Never `Url`, `Http`, `Id`.
- **Package names** are lowercase, single words, no underscores (`cmdutil`, not `cmd_util`).
- **Receiver names** are short (1–2 chars) and consistent across a type's methods (`c *Config`, not `config *Config`).
- **Local variable names** are short; longer names are used only when needed for clarity.
- **Unexported** context key types are empty structs (`type contextKey struct{}`), not strings, to prevent collisions.

---

## Comments

- All exported symbols and packages must have **godoc comments**.
- Package comments: `// Package foo does X.`
- Function/type comments: `// FunctionName does X.` — start with the symbol name.
- Comments are complete sentences ending with a **period**.
- No blank line between a declaration and its doc comment.

---

## Errors

- **Error strings are lowercase** with no trailing punctuation: `"no token found"`, not `"No token found."`.
- Use **`errors.New`** for static strings (no format verbs). Use `fmt.Errorf` only when formatting is needed.
- Use **`%w`** (not `%v` or `%s`) when wrapping an error so callers can use `errors.Is` / `errors.As`.
- Use **`errors.Is`** (not `==` or `os.IsNotExist`) for error comparisons.
- Never silently discard errors. If discarding is unavoidable, add an explanatory comment and use `_ =`.

---

## Control flow — indent error flow

Handle errors and special cases first; keep the happy path at the lowest indent level.

```go
// correct
data, err := os.ReadFile(path)
if errors.Is(err, os.ErrNotExist) {
    return &Config{}, nil
}
if err != nil {
    return nil, fmt.Errorf("reading %s: %w", path, err)
}
// happy path continues here unindented

// wrong — nested ifs push happy path right
if err != nil {
    if os.IsNotExist(err) { // deprecated; use errors.Is
        return &Config{}, nil
    }
    return nil, err
}
```

---

## Imports

Three groups separated by blank lines: **stdlib**, **third-party**, **internal**.

```go
import (
    "context"
    "fmt"

    netbox "github.com/netbox-community/go-netbox/v4"
    "github.com/spf13/cobra"

    "github.com/kirtis/netbox-cli/internal/cmdutil"
)
```

---

## Context

- `context.Context` is always the **first parameter**, named `ctx`.
- Never store a context in a struct field.
- Pass the context through; don't use `context.Background()` inside library or command code.

---

## Testing

- Use **table-driven tests** when multiple cases share the same structure.
- Use `t.Run` for subtests, including inside table tests.
- Test helpers must call `t.Helper()` as their first statement.
- Use `t.Setenv` (not `os.Setenv`) for environment variable isolation.
- Use **`t.TempDir()`** for temporary files — automatically cleaned up.
- Test failure messages follow the `"got X, want Y"` convention.
- Test packages use **black-box** naming (`package foo_test`) to test the public API only.

---

## Linting

`make lint` runs `golangci-lint` with the following linters enabled beyond the standard set:

| Linter | Purpose |
|---|---|
| `bodyclose` | HTTP response body must be closed. |
| `dupword` | Duplicate words in comments or strings. |
| `errorlint` | Correct use of Go 1.13 error wrapping (`errors.Is`, `%w`). |
| `fatcontext` | Context not stored in loops or closures. |
| `gocritic` | Broad set of style and correctness checks. |
| `godot` | Comments must end with a period. |
| `gosec` | Common security issues. |
| `misspell` | Spelling in comments and strings. |
| `revive` | Idiomatic Go; replaces `golint`. |

Use `//nolint:<linter> // reason` to suppress a specific false positive with an explanation.

---

## Unused parameters

Cobra `RunE` callbacks require `args []string` even when unused. Name it `_`:

```go
RunE: func(cmd *cobra.Command, _ []string) error {
```
