package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// csv2float takes in the source of CSV data
func csv2float(r io.Reader) ([]string, map[string][]float64, error) {
	// Create the CSV Reader used to read in data from CSV files
	cr := csv.NewReader(r)

	// Grab header values from first line of file
	headers, err := cr.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot read data from file: %w", err)
	}

	// Read in all CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot read data from file: %w", err)
	}

	// Initialize a map to hold the data for each column.
	columns := make(map[string][]float64)
	for _, field := range headers {
		columns[field] = []float64{}
	}

	// Looping through all records
	for _, row := range allData {

		// Loop through columns in row and convert data read into a float
		for j, val := range row {
			v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
			}
			columns[headers[j]] = append(columns[headers[j]], v)
		}

	}

	// Return the map
	return headers, columns, nil
}
