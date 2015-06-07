package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JsonFileInfo struct {
	os.FileInfo
}

func (fileInfo JsonFileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string       // base name of the file
		Size    JsonFileSize // length in bytes for regular files; system-dependent for others
		Mode    JsonFileMode // file mode bits
		ModTime time.Time    // modification time
		IsDir   bool         // abbreviation for Mode().IsDir()
	}{
		Name:    fileInfo.Name(),
		Size:    JsonFileSize(fileInfo.Size()),
		Mode:    JsonFileMode(fileInfo.Mode()),
		ModTime: fileInfo.ModTime(),
		IsDir:   fileInfo.IsDir(),
	})
}

type JsonFileMode os.FileMode

func (m JsonFileMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Octal   string
		Numeric uint32
	}{
		Octal:   fmt.Sprintf("%#o", m),
		Numeric: uint32(m),
	})
}

func humanizeByteSize(size JsonFileSize) string {
	const (
		_        = iota
		kilobyte = 1 << (10 * iota)
		megabyte
		gigabyte
		terabyte
	)

	format := "%.f"
	value := float32(size)

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

type JsonFileSize int64

func (s JsonFileSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Human   string
		Numeric int64
	}{
		Human:   humanizeByteSize(s),
		Numeric: int64(s),
	})
}

func NewFileInfos(fileInfos []os.FileInfo) []JsonFileInfo {
	jsonFileInfos := make([]JsonFileInfo, 0, len(fileInfos))
	for _, v := range fileInfos {
		jsonFileInfos = append(jsonFileInfos, JsonFileInfo{v})
	}
	return jsonFileInfos
}
