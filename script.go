package flintui

import _ "embed"

//go:embed js/flint-ui.js
var script []byte

// ScriptName is the file name to serve Script under (before hashing).
const ScriptName = "flint-ui.js"

// Script returns flint-ui's Alpine components (js/flint-ui.js). Serve it as
// a file and load it before Alpine, so components that use x-data="flint…"
// are registered when Alpine starts. It works with the regular Alpine build
// and the CSP build alike.
func Script() []byte {
	return script
}
