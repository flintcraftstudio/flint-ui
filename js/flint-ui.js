// flint-ui's Alpine components, registered with Alpine.data so they run
// under both the regular Alpine build and the CSP build (@alpinejs/csp).
// The CSP build parses only expressions, never function bodies, so a
// component's state and methods must live here, not inline in x-data.
//
// Load this before Alpine: Alpine fires alpine:init as it starts, and the
// components register in that event. Serve it from the Go module with
// flintui.Script(), so its version always matches the components.
document.addEventListener("alpine:init", function () {
	"use strict";

	// Slide-over: see components/slideover. x-data="flintSlideover", with
	// its name in data-flint-name. Opened and closed by the window events
	// flint:slideover-open and flint:slideover-close.
	Alpine.data("flintSlideover", function () {
		return {
			open: false,

			// trapped drives the panel's x-trap. It follows open, but turns
			// true only once the panel has a box: x-show shows it on a later
			// animation frame, and a trap started before then (the Focus
			// plugin waits a fixed 15ms) finds nothing to focus and never
			// retries.
			trapped: false,

			// eventName is the name an open or close event addresses. $dispatch
			// sends the bare name; an htmx HX-Trigger header wraps a scalar as
			// {value: name}. Anything else is no name.
			_eventName: function (e) {
				var d = e.detail;
				if (d && typeof d === "object" && "value" in d) d = d.value;
				return typeof d === "string" ? d : "";
			},

			onOpen: function (e) {
				if (this._eventName(e) === this.$el.dataset.flintName) this.open = true;
			},

			// No name (none, true, an object) closes every slide-over; a name
			// closes only that one. htmx 4 can't send a null detail.
			onClose: function (e) {
				var n = this._eventName(e);
				if (!n || n === this.$el.dataset.flintName) this.open = false;
			},

			close: function () {
				this.open = false;
			},

			// With data-flint-history, the panel owns one history entry while
			// open, marked in history.state. Back closes it; any other close
			// calls history.back() to remove the entry. back() is
			// asynchronous, and in Chromium a pushState made before it lands
			// doesn't stack on top: the back() still lands on the entry before
			// the panel's. So a panel reopened in that gap waits for the
			// back() and then pushes, and _backing (true only while a back()
			// of ours is in flight) tells that popstate apart from Back.
			init: function () {
				var self = this;
				this.$watch("open", function (open) {
					if (!open) {
						self.trapped = false;
						return;
					}
					var panel = self.$refs.panel;
					(function shown() {
						if (!self.open) return;
						if (panel.getClientRects().length) self.trapped = true;
						else requestAnimationFrame(shown);
					})();
				});
				if (this.$el.dataset.flintHistory === undefined) return;
				var name = this.$el.dataset.flintName;
				var ours = function () { return !!history.state && history.state.flintSlideover === name; };
				// Keep whatever the page or htmx stored in the entry's state.
				var push = function () { history.pushState(Object.assign({}, history.state, { flintSlideover: name }), ""); };
				// Loaded onto an entry still marked (a reload while open): unmark
				// it, or Back from a reopened panel would land on it and stay open.
				if (ours()) {
					var unmarked = Object.assign({}, history.state);
					delete unmarked.flintSlideover;
					history.replaceState(unmarked, "");
				}
				this._backing = false;
				this.$watch("open", function (open) {
					if (open) {
						if (!self._backing) push();
					} else if (ours() && !self._backing) {
						self._backing = true;
						history.back();
					}
				});
				this._onPopState = function () {
					if (self._backing) {
						self._backing = false;
						if (self.open) push();
					} else if (self.open && !ours()) {
						self.open = false;
					}
				};
				window.addEventListener("popstate", this._onPopState);
			},

			destroy: function () {
				if (this._onPopState) window.removeEventListener("popstate", this._onPopState);
			}
		};
	});

	// Toast container: see components/toast. x-data="flintToast".
	Alpine.data("flintToast", function () {
		return {
			toasts: [],
			nextId: 0,
			leaveDuration: 200,

			add: function (detail) {
				if (!detail) return;
				var allowed = ["info", "success", "warning", "danger"];
				var t = {
					id: ++this.nextId,
					variant: allowed.indexOf(detail.variant) >= 0 ? detail.variant : "info",
					title: String(detail.title || ""),
					body: String(detail.body || ""),
					duration: typeof detail.duration === "number" ? detail.duration : 5000,
					paused: false,
					leaving: false,
					_timeout: null
				};
				this.toasts.push(t);
				if (t.duration > 0) this._schedule(t);
			},

			_schedule: function (t) {
				var self = this;
				t._timeout = setTimeout(function () { self.dismiss(t.id); }, t.duration);
			},

			pause: function (t) {
				if (t.leaving) return;
				if (t._timeout) {
					clearTimeout(t._timeout);
					t._timeout = null;
					t.paused = true;
				}
			},

			resume: function (t) {
				if (t.leaving) return;
				if (t.paused && t.duration > 0) {
					t.paused = false;
					this._schedule(t);
				}
			},

			// dismiss runs the leave animation by hand (t.leaving switches on
			// the flint-toast-out-* class), then removes the toast.
			dismiss: function (id) {
				var t = null;
				for (var i = 0; i < this.toasts.length; i++) {
					if (this.toasts[i].id === id) t = this.toasts[i];
				}
				if (!t || t.leaving) return;
				if (t._timeout) {
					clearTimeout(t._timeout);
					t._timeout = null;
				}
				t.leaving = true;
				var self = this;
				setTimeout(function () {
					self.toasts = self.toasts.filter(function (x) { return x.id !== id; });
				}, this.leaveDuration);
			},

			// Escape dismisses only the most recent toast.
			dismissLatest: function () {
				if (this.toasts.length) this.dismiss(this.toasts[this.toasts.length - 1].id);
			},

			progressStyle: function (t) {
				return {
					animation: "flint-toast-progress " + t.duration + "ms linear forwards",
					animationPlayState: t.paused ? "paused" : "running"
				};
			}
		};
	});
});
