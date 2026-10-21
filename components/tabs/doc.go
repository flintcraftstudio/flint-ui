// Package tabs provides tabbed interfaces in two interaction modes.
//
// ModeAlpine (default): every panel is server-rendered up front; Alpine's
// x-show toggles which one is visible based on the active tab. Switching
// is instant — no roundtrip. Pick this when each tab's content is static
// or cheap to render with whatever the page already has in scope.
//
// ModeServer: the tablist and the active panel are server-rendered; each
// Tab carries hx-* attrs (caller-provided via Attrs) that swap a fresh
// server-rendered tabs section into the same hx-target. Pick this when
// content depends on server state that other tabs' actions can mutate
// (a form in tab A changes what tab B should show) — the server stays
// the source of truth, every tab switch reflects the latest state.
//
// Both modes follow the WAI-ARIA Tabs Pattern: role="tablist" containing
// role="tab"s, panels with role="tabpanel" linked via aria-controls and
// aria-labelledby, only-active tab is tabindex=0 (roving tabindex).
//
// Keyboard navigation in ModeAlpine: ArrowLeft/ArrowRight rotate focus
// AND activation, Home/End jump to first/last. Disabled tabs are skipped.
// ModeServer relies on native Tab-key focus between buttons; arrow-key
// roving with htmx activation is left for a future version.
//
// Sub-components (List, Tab, Panel) carry Mode and (where relevant)
// Active that must match the parent Tabs — same convention as
// modal.SubProps.Alert.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/tabs.md in the flint-ui module.
package tabs
