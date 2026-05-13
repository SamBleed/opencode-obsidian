---
type: concept
title: "React"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [frontend, react, library, ui]
status: mature
---

# React

React is the dominant library for building composable user interfaces. With the release of **React 19**, the paradigm has shifted significantly towards **Server Components**, **Actions**, and **Automatic Memoization** via the React Compiler.

## 🚀 React 19 Core Patterns

### Actions & Forms
React 19 simplifies form handling by introducing first-class support for async actions.

```tsx
function UpdateName() {
  const [error, submitAction, isPending] = useActionState(async (previousState, formData) => {
    const error = await updateName(formData.get("name"));
    if (error) return error;
    return null;
  }, null);

  return (
    <form action={submitAction}>
      <input type="text" name="name" />
      <button type="submit" disabled={isPending}>Update</button>
      {error && <p>{error}</p>}
    </form>
  );
}
```

### useOptimistic
Update the UI immediately while a server action is in progress.

```tsx
const [optimisticMessages, addOptimisticMessage] = useOptimistic(
  messages,
  (state, newMessage) => [...state, { text: newMessage, sending: true }]
);
```

### The React Compiler (Forget)
Manual memoization with `useMemo` and `useCallback` is largely becoming obsolete. The compiler automatically optimizes re-renders by analyzing component dependencies.

## 🎣 Hooks Deep Dive

### 1. useTransition
Used to mark updates as non-urgent. This allows React to prioritize urgent updates (like typing) over heavy re-renders (like filtering a long list).

```tsx
const [isPending, startTransition] = useTransition();
const handleFilter = (e) => {
  startTransition(() => {
    setFilter(e.target.value);
  });
};
```

### 2. useId
Generates unique IDs that are stable across server and client. Perfect for accessibility (`aria-labelledby`).

### 3. useDeferredValue
Similar to debouncing but managed by React's scheduler. It defers updating a part of the UI until more urgent parts have rendered.

## 🧪 Testing Best Practices (Testing Library)

1. **Avoid Testing Implementation Details**: Don't test internal state or private methods. Test what the user sees.
2. **User Event**: Use `@testing-library/user-event` instead of `fireEvent` for more realistic interactions.
3. **Screen Queries**: Prefer `getByRole`, `getByLabelText`, and `getByText` in that order for accessibility-first testing.

```tsx
test('submits the form with user input', async () => {
  render(<MyForm />);
  const input = screen.getByLabelText(/name/i);
  const button = screen.getByRole('button', { name: /submit/i });
  
  await userEvent.type(input, 'John Doe');
  await userEvent.click(button);
  
  expect(await screen.findByText(/success/i)).toBeInTheDocument();
});
```

## 🏗️ Architecture: Atomic Design + Container/Presentational
We combine these patterns for maximum reusability.

1. **Atoms**: Basic UI elements (Button, Input).
2. **Molecules**: Groups of atoms (SearchField).
3. **Organisms**: Complex sections (Navbar, ProductGrid).
4. **Templates**: Page layouts.
5. **Pages/Containers**: Where state lives and data is fetched.

## ⚡ Performance Optimization

- **Selective Hydration**: Use `Suspense` to load parts of the UI independently.
- **Server Components (RSC)**: Fetch data on the server to reduce client-side bundle size and waterfall requests.
- **Vite Integration**: Use `vite-plugin-react-swc` for ultra-fast HMR (Hot Module Replacement) during development.

## 🛡️ Security Best Practices

- **Sanitization**: Never use `dangerouslySetInnerHTML` with unsanitized user input.
- **XSS Prevention**: React escapes strings by default.
- **Secret Leaks**: Use `NEXT_PUBLIC_` prefix only for public variables.

## ⚠️ Common Pitfalls (Caught by [[react-doctor]])
- **Fetch in useEffect**: Causes race conditions. Use **Actions** or TanStack Query.
- **Derived State in useEffect**: Synchronizing state via effects leads to cascading renders. Use `useMemo`.
- **Prop Drilling**: Overusing props for deep trees. Use **Context** or **Composition**.

## Related
- [[react-doctor]]
- [[react-health-score]]
- [[DESIGN-TEMPLATE]]
- [[coding-agent-integration]]
