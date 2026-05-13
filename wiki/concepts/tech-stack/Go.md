---
type: concept
title: "Go (Golang)"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [backend, go, language, performance]
status: mature
---

# Go (Golang)

Go is an open-source, statically typed, compiled programming language designed at Google for simplicity, reliability, and efficiency. As of 2026 (Go 1.26+), it remains the industry standard for high-performance cloud infrastructure and microservices.

## 🏛️ Senior Architecture Patterns

### Hexagonal / Clean Architecture
In Go, we favor the separation of concerns by using interfaces to define boundaries.

```go
// internal/domain/user.go
type User struct {
    ID    string
    Email string
}

type UserRepository interface {
    Save(ctx context.Context, u *User) error
    GetByID(ctx context.Context, id string) (*User, error)
}

// internal/service/user_service.go
type UserService struct {
    repo UserRepository
}

func (s *UserService) Register(ctx context.Context, email string) error {
    // Business logic here
    return s.repo.Save(ctx, &User{Email: email})
}
```

### Structured Concurrency
Avoid "naked" goroutines. Use `errgroup.Group` from `golang.org/x/sync/errgroup` to manage life-cycles and error propagation.

```go
func processItems(ctx context.Context, items []Item) error {
    g, ctx := errgroup.WithContext(ctx)
    
    for _, item := range items {
        item := item // shadow for closure
        g.Go(func() error {
            return process(ctx, item)
        })
    }
    
    return g.Wait()
}
```

## 🛠️ Modern Features (2024-2026)

### Range-over-func (Iterators)
Go 1.23+ introduced a standard way to implement custom iterators that work with `for...range`.

```go
func AllUsers(db *sql.DB) func(yield func(u User) bool) {
    return func(yield func(u User) bool) {
        rows, _ := db.Query("SELECT id, email FROM users")
        defer rows.Close()
        for rows.Next() {
            var u User
            rows.Scan(&u.ID, &u.Email)
            if !yield(u) {
                return
            }
        }
    }
}

// Usage:
for user := range AllUsers(db) {
    fmt.Println(user.Email)
}
```

### Structured Logging (slog)
Ditch external dependencies for logging. Use the standard `log/slog`.

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("user_login",
    slog.String("user_id", "u123"),
    slog.Group("metadata",
        slog.String("ip", "1.1.1.1"),
        slog.Int("attempts", 1),
    ),
)
```

## 🧪 Testing Best Practices

1. **Table-Driven Tests**: The idiomatic way to test multiple scenarios.
2. **Golden Files**: For testing complex outputs (JSON, HTML).
3. **Fuzzing**: Use `testing.F` to find edge cases in parsers or logic.
4. **Mocking**: Use manual mocks or tools like `moq` instead of reflection-heavy frameworks.

## ⚠️ Anti-patterns (Avoid these!)
- **Global State**: Using `init()` to setup DB connections.
- **Context Misuse**: Passing context in structs (it should only be a function parameter).
- **Pointer Obsession**: Using pointers for everything. Copying small structs is faster and safer.
- **Ignored Errors**: Never use `_ = doSomething()`. Wrap and return: `fmt.Errorf("doing thing: %w", err)`.

## Related
- [[Go-v-Rust]] (Comparison)
- [[GO-API-BLUEPRINT]]
- [[go-rate-limiting]]
- [[Trivy]]
- [[agentmemory]]
