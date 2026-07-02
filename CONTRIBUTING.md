# Contributing to Bunker OS

PRs welcome. Please read this before opening one.

---

## Workflow

1. Fork the repo and create a branch from `main`
2. Make your changes
3. Run `make test` — all suites must pass
4. If you changed skills, update corresponding tests
5. Commit with descriptive message (conventional commits)
6. Open a PR to `main`

## Commit Conventions

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add /autoresearch skill
fix: fix Dead Letter Queue connection
docs: update README with comparison table
test: add n8n workflow tests
chore: update docker-compose
```

## Self-Review Checklist

Before requesting review:

- [ ] `make test` passes (431 tests, 0 failures)
- [ ] If you changed skills: did you update SKILL.md?
- [ ] If you changed scripts: did you update tests?
- [ ] If you changed structure: did you update README / PROJECT / WIKI?
- [ ] No hardcoded secrets
- [ ] Changes are backwards-compatible or documented

## Tests

All PRs must pass `make test`. CI verifies this automatically.

```bash
make test
```

## Reporting Issues

Use GitHub Issues for:
- Bugs in skills or scripts
- n8n workflow failures
- Missing or incorrect documentation
- Feature suggestions

Include:
- Bunker OS version (see `PROJECT.md`)
- OpenCode version
- Steps to reproduce
- Relevant error logs
