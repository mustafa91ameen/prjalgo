# Crystal Indigo Theme System (Light Mode)

This document defines the **Crystal Indigo** design system for the Project Management Dashboard. This theme prioritizes legibility, professional aesthetics ("SaaS Look"), and Arabic RTL support.

## 1. Core Palette & Variables

| Role | Color Name | Hex Value | CSS Variable | Usage |
|------|------------|-----------|--------------|-------|
| **Primary** | Electric Indigo | `#4f46e5` | `--color-indigo-600` | Main actions, active states, links |
| **Secondary** | Soft Sky | `#e0e7ff` | `--color-indigo-100` | Background accents, halos, active backgrounds |
| **Background (App)** | Soft Cloud | `#f8fafc` | `--color-slate-50` | Main application background (behind cards) |
| **Background (Surface)** | Pure Crystal | `#ffffff` | `--background-primary` | Cards, Sidebar, Modals |
| **Text (Heading)** | Midnight | `#0f172a` | `--color-slate-900` | Page titles, Section headers |
| **Text (Body)** | Deep Slate | `#334155` | `--color-slate-700` | Standard text, Descriptions |
| **Border** | Mist | `#e2e8f0` | `--color-slate-200` | Dividers, Card borders |

---

## 2. Layout & Structure

### Sidebar Navigation
The sidebar should rely on whitespace and subtle indicators rather than heavy backgrounds.

- **Background**: `#ffffff` (White)
- **Border**: Right border `#e2e8f0` (1px solid)
- **Navigation Item (Default)**:
  - **Text**: `#64748b` (Slate 500)
  - **Icon**: `#94a3b8` (Slate 400)
- **Navigation Item (Active)**:
  - **Background**: `linear-gradient(90deg, #eff6ff 0%, #ffffff 100%)`
  - **Text**: `#4f46e5` (Indigo 600, Bold)
  - **Icon**: `#4f46e5` (Indigo 600)
  - **Indicator**: 4px solid `#4f46e5` on the Right edge

### Page Headers
- **Background**: Transparent (or Glass blur `rgba(255,255,255,0.8)`)
- **Title Text**: `#0f172a` (Size: 1.875rem/30px, FontWeight: 700)
- **Breadcrumbs**: `#64748b` (Size: 0.875rem/14px)

---

## 3. Data Presentation

### Tables
Designed for high density data without visual clutter.

- **Container Card**: White, Rounded corners (16px), Shadow `sm`
- **Table Header (Thead)**:
  - **Background**: `#f8fafc` (Slate 50)
  - **Text**: `#475569` (Slate 600), Uppercase, Bold, Size 0.75rem
  - **Border Bottom**: `#e2e8f0` (2px solid)
- **Table Row (Tr)**:
  - **Background**: `#ffffff`
  - **Border Bottom**: `#f1f5f9` (1px solid)
  - **Hover Effect**: `#eff6ff` (Blue 50) - *Crucial for interactivity*
- **Text**: `#334155` (Slate 700)

### Status Chips (Badges)
Use **Soft/Pastel** backgrounds with **Dark/Vivid** text. Never use solid dark blocks.

- **Completed**: Bg `#dcfce7` (Green 100) / Text `#166534` (Green 800)
- **In Progress**: Bg `#dbeafe` (Blue 100) / Text `#1e40af` (Blue 800)
- **Pending**: Bg `#fef3c7` (Amber 100) / Text `#92400e` (Amber 800)
- **Cancelled**: Bg `#fee2e2` (Red 100) / Text `#991b1b` (Red 800)

---

## 4. Forms & Input

### Input Fields
- **Background**: `#ffffff`
- **Border (Default)**: `#cbd5e1` (Slate 300)
- **Text**: `#1e293b` (Slate 800)
- **Placeholder**: `#94a3b8` (Slate 400)
- **Focus State**:
  - **Border**: `#6366f1` (Indigo 500)
  - **Ring/Glow**: `0 0 0 4px #e0e7ff` (Indigo 100, 4px spread) - *The "Premium" Halo*

### Selection / Dropdowns
- **Menu Background**: `#ffffff`
- **Shadow**: `0 10px 15px -3px rgba(0, 0, 0, 0.1)` (Large Shadow)
- **Active Option**: `#eff6ff` (Indigo 50)

### Text Selection (Highlight)
These are global styles for when a user highlights text.
- **Selection Background**: `#c7d2fe` (Indigo 200)
- **Selection Text Color**: `#1e1b4b` (Indigo 950)

---

## 5. System UI & Overlays

### Scrollbars (Custom)
- **Track**: `#f1f5f9` (Slate 100)
- **Thumb**: `#cbd5e1` (Slate 300)
- **Thumb (Hover)**: `#94a3b8` (Slate 400)
- **Width**: `8px` (Slim)

### Modals & Dialogs
- **Overlay (Backdrop)**: `rgba(15, 23, 42, 0.4)` (Slate 900 at 40%) with `backdrop-filter: blur(4px)`
- **Dialog Container**: `#ffffff` (White)
- **Dialog Shadow**: `0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)` (Shadow 2XL)

---

## 6. Visual Effects & Shadows

### Shadows
We use "Colored Shadows" for a modern feel.

- **Card Shadow (Default)**: `0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.1)`
- **Card Shadow (Hover)**: `0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -4px rgba(0, 0, 0, 0.1)`
- **Primary Button Shadow**: `0 4px 6px -1px rgba(79, 70, 229, 0.3)` (Indigo glow)

### Typography
- **Font Family (Arabic)**: `Cairo` (Headings), `Tajawal` (UI), `Noto Sans Arabic` (Body)
- **Weights**:
  - Regular: 400
  - Medium: 500
  - Bold: 700

---

## 6. Implementation CSS Snippet

Copy these overrides into `colors.css` or your specific page CSS to enforce the theme.

```css
:root {
  /* THEME OVERRIDES */
  --background-primary: #ffffff;
  --background-secondary: #f8fafc; /* Slate 50 */
  --background-tertiary: #f1f5f9;  /* Slate 100 */
  
  --color-primary: #4f46e5;        /* Indigo 600 */
  --color-primary-light: #6366f1;
  --color-primary-dark: #3730a3;
  
  --text-primary: #0f172a;         /* Slate 900 */
  --text-secondary: #334155;       /* Slate 700 */
  --text-tertiary: #64748b;        /* Slate 500 */
  
  --border-light: #e2e8f0;
  --border-primary: #e0e7ff;
  
  /* Shadows */
  --shadow-card: 0 1px 3px rgba(0,0,0,0.1);
  --shadow-card-hover: 0 10px 15px -3px rgba(0,0,0,0.1);
  --shadow-primary-glow: 0 4px 14px 0 rgba(79, 70, 229, 0.4);
}

/* Sidebar Reset */
.sidebar-bg {
  background: #ffffff !important;
  border-left: 1px solid var(--border-light) !important;
}

/* Table Clean Up */
.table-header {
  background-color: var(--background-secondary) !important;
  color: var(--text-tertiary) !important;
  text-transform: uppercase !important;
  font-size: 0.75rem !important;
  letter-spacing: 0.05em !important;
  font-weight: 700 !important;
}

/* Input Focus Ring */
.form-field-focused {
  border-color: var(--color-primary) !important;
  box-shadow: 0 0 0 4px var(--color-indigo-100) !important;
}

/* Global Selection */
::selection {
  background-color: #c7d2fe !important;
  color: #1e1b4b !important;
}

/* Custom Scrollbar */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-track {
  background: #f1f5f9;
}
::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 4px;
}
::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
```

## 7. Components

### Statistic Cards
Cards used for high-level metrics (Project Management & Dashboard).
- **Isolation**: Dashboard metrics use `.dashboard-stat-*` classes; Project specific metrics use `.project-stat-*`. This prevents style leakage.
- **Card Padding**: Generous spacing (`2.5rem` or `var(--space-6)`) to ensure content breathes.
- **Icon Styling**:
  - **Size**: strictly `42px`.
  - **Alignment**: Flex-centered (`display: flex; justify-content: center; align-items: center`).
  - **Layering**: Icons must have `z-index: 5` and `position: relative` to sit above background effects.
- **Icon Wrapper**:
  - **Size**: `64px` x `64px`.
  - **Shape**: Circle (`border-radius: 50%`).
  - **Background**: Theme-specific soft gradient (e.g., Blue 50 for Primary).

