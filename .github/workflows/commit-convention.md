# 🧩 Commit Message Convention

This repository follows a structured commit convention to keep history clean and consistent.

---

## ✏️ Format
`<type>`(scope): `<short summary>` (Fixes|Refs #issue, task:n)

### Examples
feat(auth): implement JWT validation (#12, task:2)
feat(auth): complete auth flow (Fixes #12)
fix(order): correct price rounding issue (Fixes #33)
refactor(db): simplify query builder (#21)
test(user): add test for registration API (#10)
docs: update setup instruction (#7)
chore(ci): add commit convention and GitHub Actions for linting (#45)

---

## 🧱 Allowed Types

| Type | Description |
|------|--------------|
| feat | Add a new feature |
| fix | Bug fix |
| refactor | Code refactor (no new features or bug fixes) |
| test | Add or update tests |
| docs | Documentation only changes |
| chore | Maintenance or CI/CD changes |
| style | Code style changes (formatting, linting) |

---

## ⚙️ Notes

- Use **`Fixes #X`** only when the issue is fully completed → it will auto-close.
- Use **`(#X, task:N)`** for partial progress to reference specific checklist items.
- Scope (the part in parentheses after the type) is optional but recommended, e.g. `(auth)` or `(db)`.
- Keep the summary under 80 characters.
- Commit messages must be in **English** to stay consistent with GitHub automation tools.