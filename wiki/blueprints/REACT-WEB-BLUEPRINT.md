# Blueprint: React Web (2026)

**Target**: Modern Web Apps
**Stack**: React 19, Tailwind 4, TypeScript

## Directory Structure
```text
/
├── src/
│   ├── components/     # Atomic Design (atoms, molecules, etc.)
│   ├── hooks/          # Custom hooks
│   ├── services/       # API calls
│   ├── store/          # Global state (if needed)
│   ├── utils/          # Pure functions
│   └── App.tsx
├── public/
├── tailwind.config.js
└── Dockerfile
```

## Mandatory Checks
- [ ] React Compiler enabled.
- [ ] No `useMemo`/`useCallback` unless strictly necessary.
- [ ] Error Boundaries implemented.
- [ ] Responsive design (Mobile first).

---
**Template Type**: Project Blueprint
