// Package radio provides a native-input-backed radio button styled to
// match Catalyst's custom Headless UI control. Converted from
// catalyst-ui-kit/typescript/radio.tsx; the underlying element is a
// real <input type="radio"> so browser-native exclusive selection
// (via shared `name` attribute) and form submission work with zero JS.
//
// Structure mirrors Checkbox:
//
//   - Radio      single styled radio input
//   - RadioGroup vertical stack wrapper (matches CheckboxGroup)
//   - RadioField grid layout for Radio + Label + Description (matches
//     CheckboxField)
//
// For a set of exclusive choices, all Radios share a Name. The browser
// handles "only one checked at a time" without JS — same as any HTML
// form from 1995.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/radio.md in the flint-ui module.
package radio
