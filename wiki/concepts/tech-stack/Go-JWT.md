# JWT Implementation (Go 1.26 + Hexagonal)

**Status**: Core Standard (2026)
**Category**: Security / Authentication
**Library**: `github.com/golang-jwt/jwt/v5`
**Parent**: [[Go]]

## 🏗️ Architectural Pattern: Decoupling via Ports

To maintain [[opencode-obsidian-architecture]], auth logic must be treated as an infrastructure adapter.

### 1. The Port (Domain Layer)
```go
// internal/domain/ports/auth.go
type TokenService interface {
    Generate(userID string) (string, error)
    Validate(token string) (string, error)
}
```

### 2. The Adapter (Infrastructure Layer)
```go
// internal/repository/auth/jwt_adapter.go
import "github.com/golang-jwt/jwt/v5"

type JWTAdapter struct {
    secret []byte
}

func (a *JWTAdapter) Generate(userID string) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(a.secret)
}
```

## 🔒 Security Guardrails (2026)
- **Algorithm**: Prefer **RS256** (Asymmetric) for distributed systems.
- **Expiration**: Keep short lifetimes (5-15 min) and use Refresh Tokens.
- **Validation**: Always check `iss` (issuer) and `aud` (audience) claims.

---
**Last Updated**: 2026-04-19
