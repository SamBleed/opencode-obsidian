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

## 5. Deep Verification (Inspired by [[DeepTutor]])
- **Logical Consistency Check**: Before concluding a task, run a dedicated "Critique" pass to ensure the solution matches the spec's intent, not just the syntax.
- **Self-Correction Loop**: If a verification fail occurs, document *why* the logic failed, not just the error message.
- **Cross-Source Validation**: When researching new tech, cross-reference at least 3 sources to avoid "hallucinated standards".

---
**Last Updated**: 2026-04-19
