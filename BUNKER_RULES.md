# 🛡️ Bunker Rules (v1.3.1)
*Autonomous governance and technical standards for the Bunker OS.*

## 🏛️ General Protocol
1. **Context First**: Always read `wiki/hot.md` at the start of a session.
2. **Deep Capacity**: Do not write "seed" notes. Aim for 100+ lines for core concepts.
3. **Traceability**: All architectural decisions must be documented as [[ADR-Lifecycle|AgDRs]].
4. **Integrity**: Every new note must be linked from the [[index]] or a parent concept.

## 📐 Technical Standards (2026)

### Go Development
- **Architecture**: Hexagonal / Clean Architecture only.
- **Concurrency**: Use `errgroup` for structured concurrency.
- **Logging**: Use `slog` with JSON output. No `fmt.Print`.
- **DB**: Use `pgx/v5` and handle zero-downtime migrations.

### React Development
- **Version**: React 19 / Next.js 15+.
- **Patterns**: Server Components by default, Actions for mutations.
- **Styling**: Tailwind CSS v4 (CSS-first config). OKLCH colors.
- **Quality**: Maintain [[react-health-score]] > 80.

### Infrastructure & Security
- **IaC**: OpenTofu with encrypted state. Three-layer module architecture.
- **Zero-Trust**: Mandatory [[Trivy]] scans and SLSA Level 1 attestations.
- **Secrets**: Zero keys in code. Use ephemeral OIDC tokens in CI.

## 🔄 Ingestion Workflow
When a new resource is provided:
1. Create source summary in `wiki/sources/`.
2. Create/Expand entity and concept pages.
3. Include real code examples.
4. Update [[index]], [[log]], and [[hot]].
5. Add to [[BUNKER-OS.canvas]].

## 💸 FinOps Guardrails
- **Logic Density**: Optimize for "Utility per Token."
- **Model Routing**: Use Tier 1 models for routine tasks, Tier 2 for planning.
- **Budgeting**: Hard loop limits (max 15 steps) per agent task.

---
*Last Updated: 2026-07-02 — Bunker OS v1.3.1.*
