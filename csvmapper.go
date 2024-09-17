package gocsvmapper

import (
	"encoding/csv"
	"io"
	"strings"
)

// CSVColumnProcessor is responsible for processing CSV column names based on a given mapping.
// The ColumnMap contains the mapping of column names to be replaced.
// For example, {"student": "stu", "age": "old"} would map the "student" column to "stu".
type CSVColumnProcessor struct {
	ColumnMap map[string]string
}

// NewCSVColumnProcessor creates a new CSVColumnProcessor instance.
// It takes a column mapping (ColumnMap) as an argument, where keys are the original column names
// and values are the corresponding names to replace them with.
// For example, calling this with map[string]string{"student": "stu"} would result in "student"
// being replaced by "stu" in the CSV data.
func NewCSVColumnProcessor(columnMap map[string]string) *CSVColumnProcessor {
	return &CSVColumnProcessor{ColumnMap: columnMap}
}

// CSVToString converts a 2D slice of CSV records into a CSV-formatted string.
// Each row is joined by commas, and rows are separated by newlines.
func CSVToString(records [][]string) string {
	var sb strings.Builder
	for _, row := range records {
		sb.WriteString(strings.Join(row, ","))
		sb.WriteString("\n")
	}
	return sb.String()
}

// MapCSVColumns reads a CSV file and renames its columns based on the provided column map.
func (c *CSVColumnProcessor) MapCSVColumns(file io.Reader) ([][]string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) > 0 {
		headers := records[0]
		for i, col := range headers {
			if newCol, ok := c.ColumnMap[col]; ok {
				headers[i] = newCol
			}
		}
		records[0] = headers
	}

	return records, nil
}
