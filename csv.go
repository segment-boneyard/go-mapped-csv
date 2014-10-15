//
// Like stdlib's encoding/csv except that it writes maps based on the columns passed.
//
package csv

import "encoding/csv"
import "io"

// CSV event writer
type Writer struct {
	*csv.Writer
	cols []string
}

// New CSV writter which maps columns on .Write().
func New(w io.Writer, cols []string) *Writer {
	return &Writer{
		Writer: csv.NewWriter(w),
		cols:   cols,
	}
}

// Write a map.
func (w *Writer) Write(m map[string]string) error {
	row := make([]string, len(w.cols))

	for i, col := range w.cols {
		row[i] = m[col]
	}

	return w.Writer.Write(row)
}
