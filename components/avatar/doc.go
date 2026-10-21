// Package avatar provides user-identity avatars (image, initials, or
// both) and an interactive button/link wrapper. Converted from
// catalyst-ui-kit/typescript/avatar.tsx.
//
// Sizing is caller-controlled — avatars don't ship a Size prop. Pass
// size-8 / size-10 / size-12 via Class or let an ancestor's slot rule
// pick it (Dropdown items style `*:data-[slot=avatar]:size-6` on
// small viewports, for instance). This matches Catalyst's approach
// and keeps the component API tiny.
//
// Display semantics: an avatar with both Src and Initials stacks both
// in the same grid cell — the img sits above, initials below. If the
// img fails to load the initials show through; if the img is still
// loading the initials bridge the gap. Set only Initials for a static
// initials-only avatar.
//
// data-slot="avatar" is stamped on the rendered element so upstream
// components (Dropdown, eventually table rows, nav headers) can pick
// out avatars for consistent sizing without fragile selectors.
//
// Full reference (props, variants, accessibility, htmx wiring):
// docs/components/avatar.md in the flint-ui module.
package avatar
