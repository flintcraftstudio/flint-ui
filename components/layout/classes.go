package layout

// Non-color class strings preserved verbatim from catalyst-ui-kit/
// typescript/sidebar-layout.tsx, stacked-layout.tsx, and auth-layout.tsx.
// white → surface, zinc-100 → canvas, zinc-950/5 → foreground/5, black/30
// → scrim. The dark: variants collapse into the token contract; canvas
// exists because Catalyst's dark page (zinc-950) is darker than the panel
// while its hovers are lighter, which one muted token can't express.

const sidebarRootClasses = "relative isolate flex min-h-svh w-full bg-surface max-lg:flex-col lg:bg-canvas"

const stackedRootClasses = "relative isolate flex min-h-svh w-full flex-col bg-surface lg:bg-canvas"

const sidebarMainClasses = "flex flex-1 flex-col pb-2 lg:min-w-0 lg:pt-2 lg:pr-2 lg:pl-64"

// contentPanelClasses is the raised white card the page content sits in
// at lg+. Below lg it's flat so content runs edge to edge.
const contentPanelClasses = "grow p-6 lg:rounded-lg lg:bg-surface lg:p-10 lg:shadow-xs lg:ring-1 lg:ring-foreground/5"

const authPanelClasses = "flex grow items-center justify-center p-6 lg:rounded-lg lg:bg-surface lg:p-10 lg:shadow-xs lg:ring-1 lg:ring-foreground/5"

const drawerPanelClasses = "flex h-full flex-col rounded-lg bg-surface shadow-xs ring-1 ring-foreground/5"

// alpineData holds the drawer state. closeOnDesktop releases the drawer
// (and x-trap's scroll lock) if the viewport grows past lg while it's
// open — at lg+ the drawer is display:none but the trap would linger.
const alpineData = "{ mobileNav: false, closeOnDesktop() { if (window.matchMedia('(min-width: 64rem)').matches) this.mobileNav = false } }"

func join(base, extra string) string {
	if extra == "" {
		return base
	}
	return base + " " + extra
}
