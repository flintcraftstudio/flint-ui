// Package toggle provides an accessible on/off switch backed by a
// native <input type="checkbox" role="switch">. Converted from
// catalyst-ui-kit/typescript/switch.tsx.
//
// Package name: `toggle` instead of `switch` because `switch` is a Go
// reserved keyword. The exported type is still called Switch so the
// API reads naturally:
//
//	@toggle.Switch(toggle.Props{Name: "notifications", Checked: true})
//
// Structure mirrors Checkbox and Radio:
//
//   - Switch      single styled switch
//   - SwitchGroup vertical stack wrapper
//   - SwitchField grid layout — label left, switch right (reversed
//     from Checkbox/Radio, which put the control first)
//
// Under the hood: sr-only <input type="checkbox" role="switch"> so
// the browser handles form submission and keyboard (Tab focuses,
// Space toggles). The label renders the track and thumb; has-checked
// and has-focus-visible drive state changes on the label from the
// input's state.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/switch.md in the flint-ui module.
package toggle
