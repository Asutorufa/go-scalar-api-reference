package scalar

import _ "embed"

// embeddedAPIReference is the @scalar/api-reference 1.64.0 browser bundle.
// It is used when Options.CDN is left at its default value.
//
//go:embed assets/api-reference.js
var embeddedAPIReference string
