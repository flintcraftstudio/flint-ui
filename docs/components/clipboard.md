# Copy-to-Clipboard

Tiny Alpine primitive for copying a value to the system clipboard — API keys, share links, wallet addresses, license codes. The component is a wrapper that exposes state, not an opinionated button.

## Import

```go
import "github.com/flintcraftstudio/flint-ui/components/clipboard"
```

## Quick example

```go
@clipboard.Copy(clipboard.Props{Value: "wlg-live-abc123xyz"}) {
    @button.Button(button.Props{
        Plain: true,
        Attrs: templ.Attributes{
            "x-on:click": "copy()",
            "aria-label": "Copy API key",
        },
    }) {
        <span x-show="!copied">Copy</span>
        <span x-show="copied" class="text-success">Copied</span>
    }
}
```

Clicking the button writes `wlg-live-abc123xyz` to the clipboard and flips `copied` to `true` for two seconds, then back to `false`.

## How it works

The wrapper owns the Alpine scope:

```js
{
    copied: false,
    copy() {
        const v = this.$el.dataset.value;
        navigator.clipboard.writeText(v).then(() => {
            this.copied = true;
            setTimeout(() => { this.copied = false }, 2000);
        }).catch(() => {
            this.$dispatch('clipboard-error', v);
        });
    }
}
```

- The value rides on a `data-value` attribute. `copy()` reads it via `this.$el.dataset.value` instead of embedding it in the JS source. This means templ's HTML-attribute escaping handles any quotes, backslashes, or Unicode in the value — no JS-string escaping concerns.
- The `x-data` expression is a constant (it doesn't change per-instance), so no per-call codegen.
- A failed write (insecure context, permission denied) dispatches a `clipboard-error` event with the attempted value as `event.detail`.

## Three common patterns

### Icon swap

A silent icon-only button that flips copy → check:

```go
@clipboard.Copy(clipboard.Props{Value: value}) {
    @button.Button(button.Props{
        Plain: true,
        Attrs: templ.Attributes{
            "x-on:click": "copy()",
            "aria-label": "Copy",
        },
    }) {
        <span x-show="!copied">@copyIcon()</span>
        <span x-show="copied" class="text-success">@checkIcon()</span>
    }
}
```

### Text swap

A labeled button that changes its text during the feedback window:

```go
@clipboard.Copy(clipboard.Props{Value: shareURL}) {
    @button.Button(button.Props{
        Outline: true,
        Attrs:   templ.Attributes{"x-on:click": "copy()"},
    }) {
        <span x-show="!copied">Copy share link</span>
        <span x-show="copied">Copied</span>
    }
}
```

### Inline with displayed value

API-key-display style — the value is visible, the button sits at the trailing edge:

```go
<div class="flex items-center gap-2 rounded-md border border-border bg-muted/40 px-3 py-2">
    <code class="flex-1 truncate font-mono text-sm">flint_live_sk_...</code>
    @clipboard.Copy(clipboard.Props{Value: apiKey}) {
        @button.Button(button.Props{
            Plain: true,
            Attrs: templ.Attributes{"x-on:click": "copy()", "aria-label": "Copy API key"},
        }) {
            <span x-show="!copied">@copyIcon()</span>
            <span x-show="copied" class="text-success">@checkIcon()</span>
        }
    }
</div>
```

## Accessibility

The component doesn't render a trigger of its own, so there's no prescribed a11y pattern. Two things to include in your trigger markup:

1. **`aria-label="Copy X"`** on the button when the visible text doesn't describe the action (icon-only buttons especially).
2. **`aria-live="polite"`** on a visible element that reflects `copied`, so screen-reader users get confirmation:

```go
@clipboard.Copy(clipboard.Props{Value: value}) {
    @button.Button(...) { ... }
    <span aria-live="polite" class="sr-only" x-text="copied ? 'Copied to clipboard' : ''"></span>
}
```

Without the aria-live element, the only feedback is visual.

## Error handling

```go
@clipboard.Copy(clipboard.Props{
    Value: value,
    Attrs: templ.Attributes{
        "x-on:clipboard-error": "$dispatch('flint:toast', { variant: 'danger', title: 'Copy failed' })",
    },
}) {
    ...
}
```

`navigator.clipboard.writeText` rejects in two common cases:

- **Insecure context** — the page isn't HTTPS or localhost.
- **Permission denied** — the browser blocked clipboard access (rare in 2026 browsers but possible).

The component catches the rejection and dispatches a `clipboard-error` event on the wrapper element with the attempted value as `event.detail`. Listen for it to show a toast, log, or retry with a fallback.

## Browser requirements

`navigator.clipboard.writeText` is supported in every current browser but requires a **secure context** — HTTPS or localhost. On plain HTTP the API is undefined and `copy()` throws synchronously (caught and dispatched as `clipboard-error`).

For legacy HTTP deployments where you still need clipboard support, swap `copy()` for a fallback using a hidden `<textarea>` + `document.execCommand('copy')`. Not shipped in v0.1 because the use case is extremely rare for FlintCraft's deployments.

## Required surrounding setup

**`x-data` on a common ancestor.** The wrapper's own `x-data` is what provides the scope, but Alpine initializes inside the tree so a surrounding `x-data` on `<body>` ensures events propagate correctly. Same baseline as Modal / Dropdown / Toast.

No plugins required. Clipboard is the lightest Phase 4 component — no `@alpinejs/focus`, no `@alpinejs/collapse`, no teleport.

## Props reference

| Field   | Type               | Default | Notes                                                       |
| ------- | ------------------ | ------- | ----------------------------------------------------------- |
| `Value` | `string`           | `""`    | Required. Text written to the clipboard when `copy()` fires. |
| `Class` | `string`           | `""`    | Appended to the wrapper `<span>`. Default: `inline-flex items-center`. |
| `Attrs` | `templ.Attributes` | `nil`   | Spread onto the wrapper. Use for `x-on:clipboard-error` or extra data attributes. |
