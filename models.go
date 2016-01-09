package junix

// Result is the type for the first blob returned by a Junix tool.
// It describes the subsequent items.
type Result struct {
	Columns []Column    `json:"columns"`
	Errors  []string    `json:"errors"`
	Meta    interface{} `json:"meta"`
}

// Column describes one data attribute
type Column struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
