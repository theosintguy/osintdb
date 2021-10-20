package model

//go:generate enumer -type=ResourceType -json -text

type ResourceType int

const (
	App ResourceType = iota
	Link
	BrowserExtension
	None
)