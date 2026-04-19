# Go Patterns (2026)

**Status**: Core Standard
**Source**: [[Context7]]
**Parent**: [[_index]]

## Project Structure
Modern Go 1.24+ projects should follow the standard layout:
- `/cmd`: Entry points.
- `/internal`: Private code.
- `/pkg`: Public library code.
- `go.mod`: Dependency management (v1.24+).

## Error Handling
Errors are values. Use the `%w` verb for wrapping:
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

## Concurrency
- Use **Goroutines** for lightweight concurrency (`go func()`).
- Use **Channels** for communication and synchronization.
- Always close channels in the sender goroutine.

## Performance
- Prefer `fmt.Printf` and `fmt.Println` for basic I/O.
- Use `os.Getenv` for environment-driven configuration.

---
**Last Updated**: 2026-04-19
