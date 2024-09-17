package gocsvmapper

import (
	"strings"
	"testing"
)

func TestProcessCSV(t *testing.T) {
	csvData := `학생,나이
홍길동,20
김철수,22`

	columnMap := map[string]string{
		"학생": "student",
		"나이": "age",
	}

	processor := NewCSVColumnProcessor(columnMap)
	records, err := processor.MapCSVColumns(strings.NewReader(csvData))
	if err != nil {
		t.Fatalf("Error processing CSV: %v", err)
	}

	if records[0][0] != "student" || records[0][1] != "age" {
		t.Errorf("Expected headers to be changed, got %v", records[0])
	}

	csvString := CSVToString(records)
	expected := "student,age\n홍길동,20\n김철수,22\n"
	if csvString != expected {
		t.Errorf("Expected CSV string: %v, got %v", expected, csvString)
	}
}
