package templates

import "embed"

//go:embed go/** java/** ruby/**
var FS embed.FS
