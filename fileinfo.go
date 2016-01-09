package junix

import (
	"fmt"
	"os"
	"time"
)

// FileInfoColumns are columns for use by file info utilities
var FileInfoColumns = Result{
	Columns: []Column{
		{"name", "base name of the file"},
		{"size", "length in bytes for regular files; system-dependent for others"},
		{"size_humanized", "human friendly representation of the filesize"},
		{"mode", "file mode bits as a string"},
		{"mode_int", "file mode bits as a number"},
		{"mod_time", "modification time"},
		{"is_dir", "true if path is a directory"},
	},
}

func humanizeByteSize(size int64) string {
	const (
		_        = iota
		kilobyte = 1 << (10 * iota)
		megabyte
		gigabyte
		terabyte
	)

	format := "%.f"
	value := float64(size)

	switch {
	case size >= terabyte:
		format = "%3.1f TB"
		value = value / terabyte
	case size >= gigabyte:
		format = "%3.1f GB"
		value = value / gigabyte
	case size >= megabyte:
		format = "%3.1f MB"
		value = value / megabyte
	case size >= kilobyte:
		format = "%3.1f KB"
		value = value / kilobyte
	}
	return fmt.Sprintf(format, value)
}

// FileInfo is a struct for (un)marshalling os.FileInfo
type FileInfo struct {
	// base name of the file
	Name string `json:"name"`
	// length in bytes for regular files; system-dependent for others
	Size int `json:"size"`
	// human friendly representation of the filesize
	SizeHumanized string `json:"size_humanized"`
	// file mode bits as a string
	Mode string `json:"mode"`
	// file mode bits as a number
	ModeInt int `json:"mode_int"`
	// modification time
	ModTime time.Time `json:"mod_time"`
	// abbreviation for Mode().IsDir()
	IsDir bool `json:"is_dir"`
}

func NewFileInfo(fi os.FileInfo) FileInfo {
	size := fi.Size()
	mode := fi.Mode()
	return FileInfo{
		Name:          fi.Name(),
		Size:          int(size),
		SizeHumanized: humanizeByteSize(size),
		Mode:          fmt.Sprintf("%#o", mode),
		ModeInt:       int(mode),
		ModTime:       fi.ModTime(),
		IsDir:         fi.IsDir(),
	}
}
