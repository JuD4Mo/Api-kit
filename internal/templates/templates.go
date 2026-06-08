package templates

import "embed"

//go:embed go/** java/** ruby/** ts/**
var FS embed.FS
