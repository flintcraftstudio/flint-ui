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
