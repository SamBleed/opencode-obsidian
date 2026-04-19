# React Patterns (2026)

**Status**: Core Standard
**Source**: [[Context7]]
**Parent**: [[_index]]

## React 19+ Evolution
The **React Compiler** is now the default, meaning manual `useMemo` and `useCallback` are largely unnecessary in optimized environments.

### Key Hooks
- **`useActionState`**: For handling form actions and transitions.
- **`useFormStatus`**: To access form pending state in child components.

### Performance
- The compiler optimizes JSX and handler stability automatically.
- Focus on **Server Components** for data fetching and **Client Components** for interactivity.

## Coding Style
- Prefer functional components.
- Use atomic design for component structure.
- State management: Use native hooks first, then Context, then external libraries if needed.

---
**Last Updated**: 2026-04-19
