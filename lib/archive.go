package lib

type ArchiveHeader struct {
	Version string `json:"version"`
	Backend string `json:"backend"`
}

const ArchiveVersion = "1.0.0"
