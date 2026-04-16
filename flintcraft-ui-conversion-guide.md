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

### 4. Preserve Catalyst Visual Design
Keep Tailwind classes exactly as Catalyst defines them. Do not "improve" the design, adjust spacing, change colors, or modernize patterns. The design system has already been solved — our job is translating it to templ.

### 5. Minimal Component API
Components should have small, focused prop structs. If a component needs 15 props, it's probably two components. Favor composition over configuration.

### 6. Accessibility Preserved
Catalyst includes ARIA attributes, keyboard handling, and focus management. Preserve all of it. If Catalyst's original uses `aria-*` attributes or focus trapping, the templ version must too.

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
        return "bg-white text-zinc-900 ring-1 ring-inset ring-zinc-300 hover:bg-zinc-50 focus:ring-zinc-500"
    case VariantDanger:
        return "bg-red-600 text-white hover:bg-red-700 focus:ring-red-500"
    case VariantGhost:
        return "text-zinc-700 hover:bg-zinc-100 focus:ring-zinc-500"
    default: // Primary
        return "bg-zinc-900 text-white hover:bg-zinc-800 focus:ring-zinc-500"
    }
}
```

**Rules**:
- Preserve Catalyst's exact class strings. Do not shorten or reorganize.
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

Build these only when a client project requires them.

15. **DatePicker** — for scheduling tools.
16. **Combobox** — searchable select for long lists.
17. **Pagination** — for long tables.
18. **Breadcrumbs** — for nested navigation.

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
- [ ] Preserve Catalyst's Tailwind classes exactly
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
- **Do not** modify Catalyst's visual design or Tailwind classes
- **Do not** create "smart" components that fetch their own data
- **Do not** add dependencies beyond templ, htmx, Alpine.js, Tailwind
- **Do not** use CSS-in-JS, styled-components, or similar patterns
- **Do not** build client-side routing or state management

## Theming and Client Customization

Catalyst uses Tailwind's default gray (zinc) scale as its neutral. For FlintCraft client projects, each client site extends Tailwind's config to add their brand accent color.

Example client `tailwind.config.js`:

```js
module.exports = {
    content: [
        "./templates/**/*.templ",
        "./node_modules/github.com/flintcraft/flint-ui/**/*.templ",
    ],
    theme: {
        extend: {
            colors: {
                accent: {
                    50:  '#fef3c7',
                    500: '#d97706', // A-Team Gutters red, or Lo Mo teal, etc.
                    900: '#78350f',
                },
            },
        },
    },
}
```

Components use semantic color names (`bg-zinc-900` for primary buttons) by default. Clients can override via their Tailwind config or by passing custom classes through `Attrs`.

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