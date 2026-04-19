# Error Handling & Observability

**Status**: Core Standard
**Type**: Engineering Pattern
**Parent**: [[_index]]

## 1. Errors as Values (Go Style)
Even in JavaScript/TypeScript, prefer returning errors or using explicit try-catch blocks that wrap errors with context.
- **Good**: `throw new Error("Failed to fetch user: " + err.message)`
- **Better**: Use a custom error class with codes.

## 2. Logging
- Use structured logging (JSON).
- Levels: `DEBUG` (dev only), `INFO` (milestones), `WARN` (recoverable), `ERROR` (action required).
- Include `trace_id` or `request_id` in logs for distributed tracing.

## 3. Silent Failures
- **NEVER** let an error pass silently.
- At minimum, log the error and provide a fallback or user-friendly message.

## 4. Monitoring
- Define key metrics (SRE Golden Signals): Latency, Traffic, Errors, Saturation.

---
**Last Updated**: 2026-04-19
