//go:build !ratatoskr_embed

package data

import "embed"

// Data is a zero-valued embed.FS when Ratatoskr is built without the
// `ratatoskr_embed` build tag. In that configuration NewEmbeddedPython returns
// an error directing the user to depend on a release tag or rebuild with the
// tag set. The real Data variable is declared in the generated per-platform
// embed_<goos>_<goarch>.go files, which are produced by the release workflow.
var Data embed.FS
