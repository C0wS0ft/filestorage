package main

type (
	Name   = string
	URL    = string
	Volume struct {
		URL  URL    `json:"url"`
		Size uint64 `json:"size"`
		Seq  uint64 `json:"seq"`
	}

	CatalogItem struct {
		HashSum []byte   `json:"hashSum"`
		Size    uint64   `json:"size"`
		Volumes []Volume `json:"volumes"`
	}
)

var (
	catalog map[Name]*CatalogItem
	volumes map[URL]struct{}
)
