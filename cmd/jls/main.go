package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
		kilobyte = 1024
		megabyte = 1024 * kilobyte
		gigabyte = 1024 * megabyte
		terabyte = 1024 * gigabyte
	)

	var unit string
	value := float32(size)

	switch {
	case size >= terabyte:
		unit = "T"
		value = value / terabyte
	case size >= gigabyte:
		unit = "G"
		value = value / gigabyte
	case size >= megabyte:
		unit = "M"
		value = value / megabyte
	case size >= kilobyte:
		unit = "K"
		value = value / kilobyte
	}

	stringValue := fmt.Sprintf("%3.2f", value)
	stringValue = strings.TrimSuffix(stringValue, ".0")
	return fmt.Sprintf("%s%s", stringValue, unit)
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

func main() {
	fileInfos, _ := ioutil.ReadDir(".") // TODO: Read other directories

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(NewFileInfos(fileInfos))
}
