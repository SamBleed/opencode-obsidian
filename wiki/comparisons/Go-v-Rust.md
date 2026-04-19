# Go vs Rust para CLI Tools

> ¿Cuándo elegir uno u otro para herramientas de línea de comando?

## Comparativa Rápida

| Criterio | Go | Rust |
|----------|-----|------|
| **Tiempo de dev** | ⭐⭐⭐⭐⭐ | ⭐⭐ |
| **Performance** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Binary size** | ~10MB | ~3MB |
| **Cross-compile** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Error handling** | Explicit (result) | ? operator |
| **Learning curve** | Bajo | Alto |
| **Ecosistema** | Maduro | Creciendo |

## Cuándo Go

- **CLI tools simples** — wrappers, scripts enhancidos
- **Tiempo a producción** — necesito algo ayer
- **Equipo diverso** — Go es más accesible
- **Microservices** — el standard de la industria
- **DevOps tools** — Terraform, Docker, Kubernetes

Ejemplos: `rclone`, `teleport`, `docker`

## Cuándo Rust

- **Performance crítico** — cada ciclo cuenta
- **Memory safety** — sin GC, sin overflow exploits
- **Tools complejos** — editors, compilers
- **Embedded** — sistemas terbatas
- **Cuando me importa el binary size**

Ejemplos: `ripgrep`, `starship`, `bat`

## Regla de Thumb

```
¿Necesitás algo rápido? → Go
¿Necesitá ser perfecto? → Rust
```

## En Este Proyecto

- **Go** → APIs, CLI tools, microservices
- **Rust** → solo si el benchmark lo justifica

> 💡 Para la mayoría de las CLI tools, **Go es suficiente** y 10x más rápido de desarrollar.