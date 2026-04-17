# FlintCraft UI: Catalyst to Templ Conversion Guide

## Objective

Convert Tailwind Catalyst UI Kit components into idiomatic templ components for use across FlintCraft Studio client projects. The resulting package (`github.com/flintcraft/flint-ui`) provides production-ready, server-rendered UI components for service business dashboards built on Go + templ + htmx + Alpine.js + Tailwind CSS.

## Context

FlintCraft Studio builds integrated platforms for service businesses (contractors, cleaning services, guides, trades). Client projects include marketing websites, lead analytics dashboards, scheduling systems, and invoicing tools. All projects share the same stack:

- **Backend**: Go (standard library HTTP, sqlc, SQLite/PostgreSQL)
- **Templating**: templ (type-safe HTML templates)
- **Interactivity**: htmx for server-driven UI, Alpine.js for client-side state
- **Styling**: Tailwind CSS (no custom CSS frameworks)

Catalyst was chosen because it's Tailwind-native, designed for business dashboards, and maintained by the Tailwind team. The conversion preserves Catalyst's visual design while adapting to server-rendered patterns.

## Architectural Principles

Follow these principles for every component conversion. They are non-negotiable.

### 1. Server-Rendered First
Components render complete HTML on the server. No client-side rendering, no React state, no component lifecycle. If a component needs to update, it either uses htmx to re-fetch from the server or Alpine.js for local UI state.

### 2. htmx-Compatible by Default
Every interactive component must pass through htmx attributes cleanly. Accept `Attrs templ.Attributes` or equivalent so callers can add `hx-get`, `hx-post`, `hx-target`, `hx-swap` as needed.

### 3. Alpine.js for Local UI State Only
Use Alpine.js (`x-data`, `x-show`, `x-on:click`) only for UI state that doesn't need server persistence: dropdown open/closed, tab selection, modal visibility, form field focus states. Never use Alpine for business data.

### 4. Preserve Catalyst Visual Design (Structure, Not Colors)
Keep Catalyst's spacing, sizing, typography, borders, shadows, transitions, and layout patterns exactly as defined. Do not "improve" the design or modernize patterns. **The one exception is color**: replace Catalyst's hardcoded color classes (`bg-zinc-900`, `text-white`, `bg-red-600`) with semantic token classes (`bg-primary`, `text-primary-foreground`, `bg-danger`) per the Theming section below. Structure is fixed; color is themed.

### 5. Minimal Component API
Components should have small, focused prop structs. If a component needs 15 props, it's probably two components. Favor composition over configuration.

### 6. Accessibility Preserved
Catalyst includes ARIA attributes, keyboard handling, and focus management. Preserve all of it. If Catalyst's original uses `aria-*` attributes or focus trapping, the templ version must too.

### 7. Semantic Color Tokens Only
Components never use raw Tailwind color classes (`bg-zinc-900`, `text-red-600`, `border-gray-200`). They use semantic tokens (`bg-primary`, `text-danger`, `border-border`) defined in the Theming section. This is what makes per-client theming possible without component changes. Violating this principle defeats the entire theming system.

## Go Module Structure

```
github.com/flintcraft/flint-ui/
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── components/
│   ├── button/
│   │   ├── button.templ
│   │   ├── button_templ.go       // generated
│   │   └── button_test.go
│   ├── table/
│   │   ├── table.templ
│   │   ├── table_templ.go
│   │   └── table_test.go
│   ├── form/
│   │   ├── input.templ
│   │   ├── select.templ
│   │   ├── textarea.templ
│   │   └── checkbox.templ
│   ├── modal/
│   │   └── modal.templ
│   ├── dropdown/
│   │   └── dropdown.templ
│   └── shared/
│       ├── types.go              // shared prop types
│       └── classes.go            // class utilities (cn, etc.)
├── styles/
│   ├── flint.css                 // any base styles not covered by Tailwind
│   └── tailwind.config.example.js // config clients need to extend
├── examples/
│   ├── showcase/                 // running Go app showing all components
│   │   ├── main.go
│   │   └── templates/
│   └── README.md
└── docs/
    ├── installation.md
    ├── theming.md
    └── components/
        ├── button.md
        ├── table.md
        └── ...
```

### Module Conventions

- **One package per component family**: `button`, `table`, `form`, `modal` each get their own package. This keeps imports explicit and allows independent versioning of component groups.
- **Shared types in `components/shared`**: Common types (Size, Variant, Color) live here to avoid circular imports.
- **Generated files committed**: Commit `*_templ.go` files so consumers don't need to run `templ generate`.
- **Semantic versioning**: Start at `v0.1.0`. Breaking API changes bump minor until `v1.0.0`.

## Component Conversion Pattern

Every component follows this pattern. Use it as the template for all conversions.

### Step 1: Identify the Catalyst Component

Pull the component source from Catalyst. Note:
- All prop variants (size, color, disabled states)
- All Tailwind classes used, including conditional classes
- Any JavaScript behavior (dropdown toggles, focus management)
- ARIA attributes and accessibility patterns

### Step 2: Define the Props Struct

Create a Go struct that captures only the props the component actually needs. Example for a Button:

```go
// components/button/button.templ

package button

type Variant string

const (
    VariantPrimary   Variant = "primary"
    VariantSecondary Variant = "secondary"
    VariantDanger    Variant = "danger"
    VariantGhost     Variant = "ghost"
)

type Size string

const (
    SizeSmall  Size = "sm"
    SizeMedium Size = "md"
    SizeLarge  Size = "lg"
)

type Props struct {
    Variant  Variant
    Size     Size
    Disabled bool
    Type     string           // "button", "submit", "reset"
    Attrs    templ.Attributes // htmx attrs, data attrs, etc.
}
```

**Rules**:
- Use typed constants for variants, not raw strings. Prevents typos.
- Provide sensible zero-value defaults (empty `Variant` = primary).
- Accept `templ.Attributes` for htmx pass-through.
- Never include content/children in props — use templ's children syntax.

### Step 3: Write the templ Component

```go
templ Button(props Props) {
    <button
        type={ defaultType(props.Type) }
        class={ buttonClasses(props) }
        disabled?={ props.Disabled }
        { props.Attrs... }
    >
        { children... }
    </button>
}
```

**Rules**:
- Use class helper functions for complex conditional logic (not inline ternaries).
- Spread `Attrs` at the end so callers can override defaults if needed.
- Use `disabled?={ }` conditional attribute syntax.

### Step 4: Extract Class Logic

Keep class decisions in plain Go functions, not template logic:

```go
// components/button/classes.go

package button

func buttonClasses(props Props) string {
    base := "inline-flex items-center justify-center rounded-md font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2"

    variant := variantClasses(props.Variant)
    size := sizeClasses(props.Size)
    state := stateClasses(props.Disabled)

    return strings.Join([]string{base, variant, size, state}, " ")
}

func variantClasses(v Variant) string {
    switch v {
    case VariantSecondary:
        return "bg-surface text-foreground ring-1 ring-inset ring-border hover:bg-muted focus:ring-ring"
    case VariantDanger:
        return "bg-danger text-danger-foreground hover:bg-danger/90 focus:ring-danger"
    case VariantGhost:
        return "text-foreground hover:bg-muted focus:ring-ring"
    default: // Primary
        return "bg-primary text-primary-foreground hover:bg-primary/90 focus:ring-ring"
    }
}
```

**Rules**:
- Use semantic color tokens (`bg-primary`, `bg-danger`), never raw Tailwind colors (`bg-zinc-900`, `bg-red-600`). See the Theming section for the full token list.
- Preserve Catalyst's non-color classes exactly — spacing, sizing, borders, transitions.
- One function per decision dimension (variant, size, state).
- Use `switch` with default for sensible fallbacks.

### Step 5: Add Alpine.js Only If Needed

For components with local UI state (dropdowns, modals, tabs), use Alpine.js inline:

```go
templ Dropdown(props Props) {
    <div x-data="{ open: false }" class="relative">
        <button
            x-on:click="open = !open"
            x-bind:aria-expanded="open"
            class={ triggerClasses(props) }
        >
            { children... }
        </button>
        <div
            x-show="open"
            x-on:click.outside="open = false"
            x-transition
            class="absolute mt-2 w-56 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5"
        >
            <!-- dropdown content -->
        </div>
    </div>
}
```

**Rules**:
- Keep Alpine scope minimal (`x-data` with just what's needed).
- Always include `x-on:click.outside` for dismissable elements.
- Use `x-transition` for smooth open/close animations.
- Set `aria-expanded` / `aria-hidden` correctly.

### Step 6: Write the Usage Example

Every component needs a working example in `examples/showcase/`. This serves as both documentation and regression test.

```go
// examples/showcase/templates/buttons.templ

templ ButtonShowcase() {
    <div class="space-y-4">
        <h2>Buttons</h2>

        <div class="flex gap-4">
            @button.Button(button.Props{Variant: button.VariantPrimary}) {
                Primary
            }
            @button.Button(button.Props{Variant: button.VariantSecondary}) {
                Secondary
            }
            @button.Button(button.Props{Variant: button.VariantDanger}) {
                Delete
            }
        </div>

        <div class="flex gap-4">
            @button.Button(button.Props{
                Variant: button.VariantPrimary,
                Attrs: templ.Attributes{
                    "hx-post":   "/api/save",
                    "hx-target": "#result",
                },
            }) {
                Save via htmx
            }
        </div>
    </div>
}
```

## Priority Component List

Convert in this order. Each phase delivers value before moving to the next.

### Phase 1: Foundation (convert first)

These are used in almost every dashboard view. Get these right first.

1. **Button** — primary, secondary, danger, ghost variants. sm/md/lg sizes.
2. **Input** — text, email, number, tel, password. With label and error state support.
3. **Select** — native select with consistent styling.
4. **Textarea** — for notes, messages, longer form content.
5. **Checkbox** — including checkbox with label pattern.
6. **Badge** — for status indicators (active, pending, completed).

### Phase 2: Layout (convert second)

Dashboard structure components.

7. **Table** — sortable columns, row actions, empty state. htmx-friendly.
8. **Card** — content container with optional header/footer.
9. **Alert** — info, success, warning, error variants.
10. **Heading** — consistent h1/h2/h3 styles with subtitle support.

### Phase 3: Interactive (convert third)

These require Alpine.js for local state.

11. **Modal** — dialog with backdrop, ESC-to-close, focus trap.
12. **Dropdown** — menu with positioning, keyboard nav.
13. **Tabs** — with server-side and client-side variants.
14. **Toast/Notification** — timed dismissible alerts.

### Phase 4: Specialized (convert as needed)

Build these only when a client project requires them. Order reflects usefulness-per-complexity — earlier items are cheap and widely applicable, later items are heavy and narrow.

**Foundational gaps** (surfaced by Pines catalog review; Catalyst doesn't ship them):

15. **Tooltip** — hover/focus hint over any trigger. Small Alpine x-data, `aria-describedby`, no portal needed for most cases. Universal utility.
16. **Accordion** — disclosure panel for FAQs, settings groups, mobile nav. Trivial Alpine state, `aria-expanded` + `aria-controls`. High value on brochure sites.
17. **Slide-over** — side drawer for filters, detail panes, mobile nav. Reuses Modal's teleport + `x-trap.inert.noscroll` + event-driven open pattern; only the enter/leave transition changes (translate-x vs scale).
18. **Copy-to-Clipboard** — trigger + clipboard write + transient "copied" state. Tiny Alpine primitive; common for share links, API keys, wallet addresses.
19. **Popover** — click-anchored floating panel with arbitrary content (not a menu). Shares anchor/positioning logic with Dropdown — extract a shared `components/floating/` helper (anchor enum, transition classes per anchor, `x-on:click.outside`) before building Popover so Dropdown/Tooltip/Popover/HoverCard don't each reinvent it.

**From the original roadmap:**

20. **Pagination** — table companion; server-driven links with current-page styling.
21. **Breadcrumbs** — nested-navigation trail, last item unlinked.
22. **Combobox** — searchable select for long lists. Non-trivial keyboard nav; uses the same Alpine Focus plugin pattern as Dropdown.
23. **Command Palette** — cmd+K search across app actions. High-impact for admin dashboards; complex (keyboard nav, fuzzy match, section grouping). Build only when an admin tool warrants it.
24. **DatePicker** — for scheduling tools. Heaviest component in the set; consider whether a native `<input type="date">` wrapped with flint-ui styling covers 80% of cases before building a full calendar grid.

**Shared infrastructure to build alongside:**

- `components/floating/` — anchor positioning + transition classes, consumed by Dropdown (refactor), Tooltip, Popover, HoverCard.
- Anchor-aware Dropdown transitions — current Dropdown uses uniform opacity+scale; a `bottom-end` menu should slide down, a `top-start` menu should slide up. Pick this up when extracting `floating/`.

**Explicitly out of scope:**

- Marketing section blocks (heros, pricing tables, footers). These belong in per-client site scaffolds, not in a shared component library — every client's hero needs to look different on purpose.
- Copy-paste HTML snippet model (Pines-style). flint-ui's templ-package model is the intentional alternative; it's what makes consistent rebranding via semantic tokens possible.

## Implementation Rules for the Coding Agent

When converting each component, follow this checklist:

### Before Writing Code
- [ ] Read the Catalyst source for the component in full
- [ ] Identify every variant, size, and state the component supports
- [ ] Note any JavaScript behavior that needs Alpine.js equivalent
- [ ] Check Catalyst's accessibility patterns (ARIA, keyboard, focus)

### Writing the Component
- [ ] Create the package directory under `components/`
- [ ] Define props struct with typed constants for variants
- [ ] Write the `.templ` file following the conversion pattern
- [ ] Extract class logic to a separate `classes.go` file
- [ ] Preserve Catalyst's non-color Tailwind classes exactly (spacing, sizing, borders, transitions)
- [ ] Translate Catalyst color classes to semantic tokens (`bg-primary`, not `bg-zinc-900`) per the Theming section
- [ ] Verify no raw Tailwind color classes remain: `grep -E "(zinc|gray|slate|red|green|blue|yellow)-[0-9]" components/{name}/`
- [ ] Add `Attrs templ.Attributes` for htmx pass-through
- [ ] Add Alpine.js only if local UI state is needed
- [ ] Preserve all ARIA attributes from Catalyst

### After Writing the Component
- [ ] Add example usage to `examples/showcase/`
- [ ] Write a short markdown doc at `docs/components/{name}.md`
- [ ] Run `templ generate` and commit the generated file
- [ ] Test that the showcase renders correctly in the browser

### What NOT to Do
- **Do not** introduce React patterns (useState, useEffect equivalents)
- **Do not** add JavaScript beyond Alpine.js for local UI state
- **Do not** use raw Tailwind color classes (`bg-zinc-900`, `text-red-600`, `border-gray-200`) in components — use semantic tokens (`bg-primary`, `text-danger`, `border-border`)
- **Do not** modify Catalyst's non-color design decisions (spacing, sizing, borders, transitions stay exactly as Catalyst defines them)
- **Do not** create "smart" components that fetch their own data
- **Do not** add dependencies beyond templ, htmx, Alpine.js, Tailwind
- **Do not** use CSS-in-JS, styled-components, or similar patterns
- **Do not** build client-side routing or state management
- **Do not** introduce new semantic tokens without updating the Theming contract (breaking change)

## Theming and Client Customization

FlintCraft clients have meaningfully different brand palettes (A-Team Gutters: black/silver/red; Lo Mo: earthy outfitter tones; Rockabilly: warm amber with rockabilly red; Western Skies: forest pine with gold). Components must adapt to each client's brand without component code changes. The mechanism is **semantic color tokens backed by CSS variables**.

### How It Works

1. **flint-ui components reference semantic tokens** (`bg-primary`, `text-danger-foreground`) instead of raw Tailwind colors (`bg-zinc-900`, `text-white`).
2. **Each client site defines CSS variables** for those semantic tokens in their base stylesheet, using their brand palette.
3. **Each client's `tailwind.config.js` maps the semantic tokens to those CSS variables**, so Tailwind's JIT generates the right utility classes.
4. **Result**: the same `bg-primary` class renders A-Team's red on A-Team's site and Lo Mo's teal on Lo Mo's site, with zero component changes.

### The Semantic Token Contract

Every flint-ui component is built against exactly this set of tokens. Clients must define all of them. Components may not introduce new tokens without updating this contract (a breaking change that requires a version bump).

| Token | Purpose | Typical Use |
|---|---|---|
| `background` | Page background | `bg-background` on `<body>` |
| `foreground` | Default text color | `text-foreground` on body copy |
| `surface` | Card/panel backgrounds | `bg-surface` on cards, modals, dropdowns |
| `surface-foreground` | Text on surface | `text-surface-foreground` |
| `muted` | Subtle backgrounds | `bg-muted` on hover states, table stripes |
| `muted-foreground` | Secondary text | `text-muted-foreground` for captions, placeholders |
| `primary` | Primary action color | `bg-primary` on primary buttons |
| `primary-foreground` | Text on primary | `text-primary-foreground` |
| `accent` | Brand accent / secondary action | `bg-accent` for client brand highlights |
| `accent-foreground` | Text on accent | `text-accent-foreground` |
| `danger` | Destructive actions, errors | `bg-danger` on delete buttons, error alerts |
| `danger-foreground` | Text on danger | `text-danger-foreground` |
| `success` | Positive states | `bg-success` on success alerts, badges |
| `success-foreground` | Text on success | `text-success-foreground` |
| `warning` | Caution states | `bg-warning` on warning alerts |
| `warning-foreground` | Text on warning | `text-warning-foreground` |
| `border` | Default borders | `border-border` on cards, inputs, dividers |
| `input` | Input backgrounds | `bg-input` on form fields |
| `ring` | Focus rings | `focus:ring-ring` on focusable elements |

### Client Site Setup

Each client site needs three things to use flint-ui: CSS variable definitions, Tailwind config mapping, and the content globs pointing at flint-ui.

**1. CSS variables — `styles/theme.css`**

Define all tokens as space-separated RGB triplets. The `<alpha-value>` placeholder in Tailwind lets opacity modifiers work (`bg-primary/50`).

```css
@layer base {
  :root {
    /* Neutrals */
    --color-background: 255 255 255;
    --color-foreground: 24 24 27;
    --color-surface: 255 255 255;
    --color-surface-foreground: 24 24 27;
    --color-muted: 244 244 245;
    --color-muted-foreground: 113 113 122;

    /* Brand — CLIENT CUSTOMIZES THESE */
    --color-primary: 24 24 27;              /* zinc-900 default; override per client */
    --color-primary-foreground: 255 255 255;
    --color-accent: 217 119 6;              /* client brand accent */
    --color-accent-foreground: 255 255 255;

    /* Semantic states */
    --color-danger: 220 38 38;
    --color-danger-foreground: 255 255 255;
    --color-success: 22 163 74;
    --color-success-foreground: 255 255 255;
    --color-warning: 234 179 8;
    --color-warning-foreground: 24 24 27;

    /* Structural */
    --color-border: 228 228 231;
    --color-input: 255 255 255;
    --color-ring: 24 24 27;
  }
}
```

**2. Tailwind config — `tailwind.config.js`**

Map semantic token names to the CSS variables. The `rgb(var(--...) / <alpha-value>)` pattern is what enables `bg-primary/50` to work.

```js
module.exports = {
  content: [
    "./templates/**/*.templ",
    "./internal/**/*.templ",
    // Include flint-ui so Tailwind scans its classes
    "./vendor/github.com/flintcraft/flint-ui/**/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        background: 'rgb(var(--color-background) / <alpha-value>)',
        foreground: 'rgb(var(--color-foreground) / <alpha-value>)',
        surface: {
          DEFAULT: 'rgb(var(--color-surface) / <alpha-value>)',
          foreground: 'rgb(var(--color-surface-foreground) / <alpha-value>)',
        },
        muted: {
          DEFAULT: 'rgb(var(--color-muted) / <alpha-value>)',
          foreground: 'rgb(var(--color-muted-foreground) / <alpha-value>)',
        },
        primary: {
          DEFAULT: 'rgb(var(--color-primary) / <alpha-value>)',
          foreground: 'rgb(var(--color-primary-foreground) / <alpha-value>)',
        },
        accent: {
          DEFAULT: 'rgb(var(--color-accent) / <alpha-value>)',
          foreground: 'rgb(var(--color-accent-foreground) / <alpha-value>)',
        },
        danger: {
          DEFAULT: 'rgb(var(--color-danger) / <alpha-value>)',
          foreground: 'rgb(var(--color-danger-foreground) / <alpha-value>)',
        },
        success: {
          DEFAULT: 'rgb(var(--color-success) / <alpha-value>)',
          foreground: 'rgb(var(--color-success-foreground) / <alpha-value>)',
        },
        warning: {
          DEFAULT: 'rgb(var(--color-warning) / <alpha-value>)',
          foreground: 'rgb(var(--color-warning-foreground) / <alpha-value>)',
        },
        border: 'rgb(var(--color-border) / <alpha-value>)',
        input: 'rgb(var(--color-input) / <alpha-value>)',
        ring: 'rgb(var(--color-ring) / <alpha-value>)',
      },
    },
  },
}
```

**3. flint-ui ships a default theme**

flint-ui itself includes `styles/flint.css` with a neutral default theme (essentially Catalyst's zinc-based palette). This lets the showcase app render without additional setup, and gives clients a working baseline they can override selectively — a client only overriding `--color-primary` and `--color-accent` still gets sensible defaults for everything else.

### Example: Per-Client Overrides

Once the contract is in place, each client's site only needs to override the tokens that differ from the default. Most clients only change primary, accent, and maybe border.

**A-Team Gutters** (`/sites/ateamgutters/styles/theme.css`):
```css
@layer base {
  :root {
    --color-primary: 0 0 0;          /* A-Team black */
    --color-primary-foreground: 255 255 255;
    --color-accent: 220 38 38;       /* A-Team red */
    --color-accent-foreground: 255 255 255;
  }
}
```

**Lo Mo Outfitting** (`/sites/lomo/styles/theme.css`):
```css
@layer base {
  :root {
    --color-primary: 45 55 72;       /* Slate */
    --color-primary-foreground: 255 255 255;
    --color-accent: 13 148 136;      /* Teal & Bone accent */
    --color-accent-foreground: 255 255 255;
  }
}
```

**Western Skies Contracting** (`/sites/westernskies/styles/theme.css`):
```css
@layer base {
  :root {
    --color-primary: 20 83 45;       /* Forest pine */
    --color-primary-foreground: 255 255 255;
    --color-accent: 202 138 4;       /* Gold */
    --color-accent-foreground: 24 24 27;
  }
}
```

Each site imports its `theme.css` after flint-ui's base stylesheet. Cascade handles the rest.

### Dark Mode (Future)

The CSS variable pattern makes dark mode trivial to add later without touching component code: define a `.dark` selector (or `@media (prefers-color-scheme: dark)`) that overrides the tokens with dark values. Components stay identical. Not required for v0.x but worth designing the tokens with this in mind — e.g., `background` and `surface` separate so dark mode can invert them independently.

### What This Means for Component Authors

When converting a Catalyst component, translate raw color classes to semantic tokens using this mapping:

| Catalyst class | flint-ui class | Notes |
|---|---|---|
| `bg-zinc-900`, `bg-gray-900` | `bg-primary` | Primary button, dark elements |
| `text-white` (on dark bg) | `text-primary-foreground` | When paired with `bg-primary` |
| `bg-white` | `bg-surface` | Cards, modals, dropdowns |
| `text-zinc-900` (on light bg) | `text-foreground` | Default body text |
| `bg-zinc-50`, `bg-gray-100` | `bg-muted` | Hover states, subtle backgrounds |
| `text-zinc-500`, `text-gray-500` | `text-muted-foreground` | Captions, placeholders |
| `border-zinc-200`, `border-gray-300` | `border-border` | Dividers, input borders |
| `bg-red-600` | `bg-danger` | Delete buttons, error states |
| `bg-green-600` | `bg-success` | Success badges |
| `bg-yellow-500` | `bg-warning` | Caution states |
| `ring-zinc-500`, `ring-blue-500` | `ring-ring` | Focus rings |

Non-color classes (`rounded-md`, `px-4`, `py-2`, `text-sm`, `font-medium`, `transition-colors`, etc.) stay exactly as Catalyst defines them. Only colors get the semantic treatment.

## Versioning Strategy

- **v0.x.x**: Pre-1.0, expect breaking changes between minors. Use during active development.
- **v1.0.0**: Stable API. Breaking changes require a major version bump.
- **Tag releases on GitHub**: `git tag v0.1.0 && git push --tags`
- **Pin in client projects**: `go get github.com/flintcraft/flint-ui@v0.1.0`

When updating flint-ui across multiple client projects, bump the version in each client's `go.mod` deliberately. Do not auto-update.

## Testing Approach

### Visual Testing via Showcase App
The `examples/showcase/` app is the primary testing mechanism. It should:
- Render every component in every variant
- Include interactive examples (buttons that actually post, tables that sort)
- Run locally with `go run examples/showcase/main.go`

### Unit Testing Class Logic
Class generation functions are pure Go and should have tests:

```go
// components/button/classes_test.go

func TestVariantClasses(t *testing.T) {
    tests := []struct {
        variant  Variant
        contains string
    }{
        {VariantPrimary, "bg-zinc-900"},
        {VariantDanger, "bg-red-600"},
        {VariantSecondary, "ring-1"},
    }

    for _, tt := range tests {
        got := variantClasses(tt.variant)
        if !strings.Contains(got, tt.contains) {
            t.Errorf("variantClasses(%q) = %q, should contain %q",
                tt.variant, got, tt.contains)
        }
    }
}
```

### Integration Testing in Client Projects
When a component is used in a real client project (e.g., Lo Mo booking dashboard), that usage acts as an integration test. Bugs found in client work get fixed in flint-ui, not worked around in the client project.

## Success Criteria

The conversion project is successful when:

1. **First client dashboard built entirely with flint-ui** — no custom components, no one-off Tailwind patterns.
2. **Second client project reuses the same components** — proving the abstraction works across clients.
3. **Component additions take less time than building from scratch** — net time saved on client work.
4. **Design consistency across all FlintCraft client tools** — when Jaden shows a prospect the scheduling tool and the invoicing tool, they feel like the same product.

## Getting Started

For the coding agent starting this project:

1. Initialize the Go module: `go mod init github.com/flintcraft/flint-ui`
2. Install templ: `go install github.com/a-h/templ/cmd/templ@latest`
3. Create the directory structure from the "Go Module Structure" section above
4. Start with `components/button/` following the conversion pattern
5. Set up the showcase app in `examples/showcase/` with one working button example
6. Once Button is done and rendering in the showcase, move to Input, then Select, then Table

Convert one component completely (props, template, classes, example, docs) before moving to the next. Incremental completion beats parallel half-finished work.
