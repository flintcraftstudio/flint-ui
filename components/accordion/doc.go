// Package accordion provides disclosure panels — click a trigger to reveal
// its paired content. From-scratch design (Catalyst has no accordion.tsx).
//
// Two interaction types:
//
//   - TypeMultiple (default): each panel toggles independently. Natural
//     for settings groups or reference material where the reader may want
//     several panels open at once.
//
//   - TypeSingle: opening a panel closes whichever was previously open.
//     Natural for FAQs or mobile navigation where focus should stay on
//     one item.
//
// State lives in Alpine, centralized on the Accordion root. The root owns
// `active` (an array of open Item names) and a `toggle(n)` method; child
// components reference both via the inherited x-data scope. Centralizing
// the state is what lets TypeSingle close the previously-open item when
// a new one opens.
//
// Accordion is always Alpine-driven — no ModeServer variant. Which panels
// are expanded is local, ephemeral UI state; if you need server state to
// dictate which section is visible, reach for Tabs (ModeServer) instead.
//
// Collapse animation (smooth height transitions) requires the Alpine
// Collapse plugin: load @alpinejs/collapse before alpinejs core. The
// Panel emits `x-collapse`; without the plugin Alpine silently ignores
// the directive and panels snap open/closed via plain x-show. Same
// progressive-enhancement pattern as @alpinejs/focus on Modal/Dropdown.
//
// A11y: each Trigger is a <button> with aria-expanded + aria-controls;
// each Panel is a <div role="region"> with id + aria-labelledby; the
// button is wrapped in an <h3> by default (override via Item.HeadingLevel
// 1-6) so the document outline stays correct.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/accordion.md in the flint-ui module.
package accordion
