package button

import (
	"strings"
	"testing"
)

// TestClassesUnchanged pins, byte for byte, the classes these props got
// before Variant could tint outline and plain buttons. Outline and plain
// with no Variant must stay neutral, and solid buttons must not move.
func TestClassesUnchanged(t *testing.T) {
	for _, tc := range []struct {
		p    Props
		want string
	}{
		{Props{Outline: true},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-border text-foreground data-active:bg-muted data-hover:bg-muted [--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)]"},
		{Props{Plain: true},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent text-foreground data-active:bg-muted data-hover:bg-muted [--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)]"},
		{Props{Outline: true, Indicator: true, Class: "w-full"},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-border text-foreground data-active:bg-muted data-hover:bg-muted [--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)] htmx-request:pointer-events-none htmx-request:text-transparent htmx-request:opacity-100 htmx-request:*:not-data-[slot=spinner]:invisible htmx-request:*:data-[slot=spinner]:flex w-full"},
		{Props{Plain: true, Variant: "bogus"},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent text-foreground data-active:bg-muted data-hover:bg-muted [--btn-icon:var(--color-muted-foreground)] data-active:[--btn-icon:var(--color-foreground)] data-hover:[--btn-icon:var(--color-foreground)]"},
		{Props{},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent bg-(--btn-border) before:absolute before:inset-0 before:-z-10 before:rounded-[calc(var(--radius-lg)-1px)] before:bg-(--btn-bg) before:shadow-sm after:absolute after:inset-0 after:-z-10 after:rounded-[calc(var(--radius-lg)-1px)] after:shadow-[inset_0_1px_--theme(--color-white/15%)] data-active:after:bg-(--btn-hover-overlay) data-hover:after:bg-(--btn-hover-overlay) data-disabled:before:shadow-none data-disabled:after:shadow-none text-primary-foreground [--btn-bg:var(--color-primary)] [--btn-border:var(--color-primary)] [--btn-hover-overlay:var(--color-primary-foreground)]/10 [--btn-icon:var(--color-primary-foreground)]/70 data-active:[--btn-icon:var(--color-primary-foreground)] data-hover:[--btn-icon:var(--color-primary-foreground)]"},
		{Props{Variant: VariantDanger},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent bg-(--btn-border) before:absolute before:inset-0 before:-z-10 before:rounded-[calc(var(--radius-lg)-1px)] before:bg-(--btn-bg) before:shadow-sm after:absolute after:inset-0 after:-z-10 after:rounded-[calc(var(--radius-lg)-1px)] after:shadow-[inset_0_1px_--theme(--color-white/15%)] data-active:after:bg-(--btn-hover-overlay) data-hover:after:bg-(--btn-hover-overlay) data-disabled:before:shadow-none data-disabled:after:shadow-none text-danger-foreground [--btn-bg:var(--color-danger)] [--btn-border:var(--color-danger)] [--btn-hover-overlay:var(--color-danger-foreground)]/10 [--btn-icon:var(--color-danger-foreground)]/70 data-active:[--btn-icon:var(--color-danger-foreground)] data-hover:[--btn-icon:var(--color-danger-foreground)]"},
		{Props{Variant: "bogus"},
			"relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:not-data-focus:outline-hidden data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-ring data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent bg-(--btn-border) before:absolute before:inset-0 before:-z-10 before:rounded-[calc(var(--radius-lg)-1px)] before:bg-(--btn-bg) before:shadow-sm after:absolute after:inset-0 after:-z-10 after:rounded-[calc(var(--radius-lg)-1px)] after:shadow-[inset_0_1px_--theme(--color-white/15%)] data-active:after:bg-(--btn-hover-overlay) data-hover:after:bg-(--btn-hover-overlay) data-disabled:before:shadow-none data-disabled:after:shadow-none text-primary-foreground [--btn-bg:var(--color-primary)] [--btn-border:var(--color-primary)] [--btn-hover-overlay:var(--color-primary-foreground)]/10 [--btn-icon:var(--color-primary-foreground)]/70 data-active:[--btn-icon:var(--color-primary-foreground)] data-hover:[--btn-icon:var(--color-primary-foreground)]"},
	} {
		if got := Classes(tc.p); got != tc.want {
			t.Errorf("Classes(%+v)\n got %q\nwant %q", tc.p, got, tc.want)
		}
	}
}

// TestTintedOutlineAndPlain checks every variant tints both unfilled
// styles from tintClasses, on top of the style's own structure.
func TestTintedOutlineAndPlain(t *testing.T) {
	for v, tint := range tintClasses {
		for _, p := range []Props{{Outline: true, Variant: v}, {Plain: true, Variant: v}} {
			got := Classes(p)
			structure := plainClasses
			if p.Outline {
				structure = outlineClasses
			}
			if want := baseClasses + " " + structure + " " + tint; got != want {
				t.Errorf("Classes(%+v)\n got %q\nwant %q", p, got, want)
			}
			if strings.Contains(got, "text-foreground") || strings.Contains(got, "bg-muted") {
				t.Errorf("Classes(%+v) still has neutral color classes: %q", p, got)
			}
		}
	}
	if len(tintClasses) != len(variantClasses) {
		t.Errorf("tintClasses has %d variants, variantClasses has %d", len(tintClasses), len(variantClasses))
	}
}
