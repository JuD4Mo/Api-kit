package templates

import "embed"

//go:embed go/** java/**
var FS embed.FS
