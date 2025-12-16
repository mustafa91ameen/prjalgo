# Frontend Development Rules

Rules for maintaining code consistency. Follow these strictly.

---

## Structure Overview

```
src/
├── styles/          # Global design tokens (DO NOT duplicate)
├── components/      # Feature-based folders with co-located styles
├── pages/           # Page components with styles in pages/styles/
├── composables/     # Reusable Vue logic
├── stores/          # Pinia stores
└── utils/           # Helper functions
```

---

## Rule 1: Use CSS Variables

**NEVER hardcode colors, fonts, spacing, or shadows.**

```css
/* WRONG */
color: #4338ca;
font-size: 14px;
padding: 16px;
box-shadow: 0 4px 6px rgba(0,0,0,0.1);

/* CORRECT */
color: var(--color-primary);
font-size: var(--font-size-sm);
padding: var(--space-4);
box-shadow: var(--shadow-base);
```

### Available Variables

| Category | File | Examples |
|----------|------|----------|
| Colors | `colors.css` | `--color-primary`, `--color-success`, `--color-indigo-500` |
| Typography | `typography.css` | `--font-size-sm`, `--font-weight-bold`, `--line-height-normal` |
| Spacing | `spacing.css` | `--space-4`, `--radius-lg`, `--shadow-card` |
| Effects | `visual-effects.css` | Classes: `.hover-lift`, `.glass-effect`, `.smooth-transition` |

---

## Rule 2: Component Organization

### New Component in Existing Feature

1. Create `ComponentName.vue` in the feature folder
2. Add export to `index.js`
3. Use existing `styles/[feature].css` - add styles there

```
components/expenses/
├── ExpenseStats.vue
├── ExpenseTable.vue
├── NewComponent.vue      ← Add here
├── index.js              ← Export here
└── styles/
    └── expenses.css      ← Styles here
```

### New Feature Module

1. Create folder: `components/[feature-name]/`
2. Create `index.js` with barrel exports
3. Create `styles/[feature-name].css`
4. Import styles in components using `@import`

```javascript
// index.js
export { default as FeatureStats } from './FeatureStats.vue'
export { default as FeatureTable } from './FeatureTable.vue'
```

---

## Rule 3: Page Styles & Overrides

Page-specific styles go in `pages/styles/[page-name].css`. 
- **Inheritance**: Pages automatically inherit global variables from `colors.css`.
- **Overrides**: To customize a specific page (e.g., a "Gold Theme" page), simply re-declare the variable inside the page's `.css` file using the specific class scope.

```css
/* pages/styles/golden-page.css */
.golden-page-wrapper {
  /* This overrides the global primary color ONLY for this page */
  --color-primary: #FFD700; 
}
```

```vue
<!-- pages/golden-page.vue -->
<style scoped>
@import './styles/golden-page.css';
</style>
```

---

## Rule 4: Adding New Design Tokens

### New Color

Add to `styles/colors.css` under `:root`:

```css
--color-[name]-[shade]: #hexvalue;
```

Follow pattern: `50, 100, 200... 900` for shades.

### New Font Size

Add to `styles/typography.css`:

```css
--font-size-[name]: [value]rem;
```

### New Spacing/Shadow

Add to `styles/spacing.css`:

```css
--space-[name]: [value];
--shadow-[name]: [shadow-value];
```

---

## Rule 5: Utility Classes

Use existing utility classes from `visual-effects.css`:

- `.hover-lift` - Lift on hover
- `.card-glow` - Glow effect on cards
- `.smooth-transition` - Standard transition
- `.glass-effect` - Glassmorphism
- `.icon-glow` - Icon glow effect

**DO NOT** create one-off transition/animation styles. Add to `visual-effects.css` if reusable.

---

## Rule 6: RTL Support

This is an Arabic RTL application.

- Use logical properties: `margin-inline-start` not `margin-left`
- Text alignment: `text-align: end` not `text-align: right`
- RTL overrides go in `styles/rtl.css`

---

## Rule 7: Vuetify Components

- Use Vuetify's built-in props for spacing (`pa-4`, `mb-6`)
- Theme overrides go in `styles/indigo-theme.css`
- Table header styles are in `styles/table-headers.css`
- Dialog/form styles are in `styles/form-dialog.css`

---

## Rule 8: Logic Separation (State Management)

- **Stores (`src/stores/`)**: Contain all business logic, API calls, and state.
- **Components**: Should trigger actions in stores and display state. **Avoid** complex logic or direct API calls in `.vue` files.
- **API Layer**: All HTTP requests go in `src/api/`.

---

## Rule 9: Naming Conventions

- **Components**: `PascalCase` (e.g., `ExpenseCard.vue`, `ProjectList.vue`)
- **Files/Folders**: `kebab-case` (e.g., `pages/project-management.vue`, `styles/visual-effects.css`)
- **Stores**: `use[Feature]Store` (e.g., `useProjectsStore`, `useAuthStore`)

---

## Rule 10: Golden Rule of Architecture

**"Smart Pages, Dumb Components"**
- **Pages**: Connect to the store, fetch data, and pass it down.
- **Components**: Receive data via `props`, emit events for actions. They should generally NOT access the store directly unless necessary.

---

## Quick Reference

| Task | Location |
|------|----------|
| Change primary color | `styles/colors.css` → `--color-primary` |
| Change font | `styles/typography.css` → `--font-primary` |
| Change button shadows | `styles/spacing.css` → `--shadow-*` |
| Add hover effect | `styles/visual-effects.css` |
| Table header style | `styles/table-headers.css` |
| Dialog/form style | `styles/form-dialog.css` |
| RTL fixes | `styles/rtl.css` |
| Component styles | `components/[feature]/styles/` |
| Page styles | `pages/styles/` |

---

## Checklist Before Commit

- [ ] No hardcoded colors/fonts/spacing
- [ ] Used existing CSS variables
- [ ] Component in correct feature folder
- [ ] Exported in `index.js`
- [ ] Styles in correct location
- [ ] RTL compatible