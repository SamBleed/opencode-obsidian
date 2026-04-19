# Docker Patterns (2026)

**Status**: Core Standard
**Source**: [[Context7]]
**Parent**: [[_index]]

## Multi-Stage Builds (The Gold Standard)
Always separate the build environment from the runtime environment to minimize image size and attack surface.

### Example Node.js Multi-Stage:
```dockerfile
# Stage 1: Build
FROM node:24-alpine AS build
WORKDIR /app
COPY package* ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Runtime
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
```

## Security Best Practices
- **Non-Root User**: Always create and switch to a non-root user.
- **DHI (Docker Hardened Images)**: Use hardened base images like `alpine` or specialized hardened versions.
- **Cache Mounting**: Use `--mount=type=cache` in `RUN` commands to speed up builds.

## Optimization
- Use `.dockerignore` to avoid copying `node_modules` or `.git`.
- Order layers by frequency of change (Dependencies first, Source code last).

---
**Last Updated**: 2026-04-19
