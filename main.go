package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/montanaflynn/stats"
)

func main() {
	// Verify and parse arguments
	if err := run(os.Args[1], os.Args[2], 100); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out string, N int) error {
	// Create output file
	outfile, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("Cannot create file: %w", err)
	}

	// Close output file after function returns
	defer outfile.Close()

	// Define statistics that we will compute
	statlist := []string{"count", "mean", "std", "min", "25%", "50%", "75%", "max"}

	// Run the function N times
	for i := 0; i < N; i++ {
		// Open file for reading
		infile, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("Cannot open file: %w", err)
		}

		headers, datamap, err := csv2float(infile)
		if err != nil {
			return err
		}

		// Close input file
		infile.Close()

		// Write header row to output file
		fmt.Fprint(outfile, strings.Join(headers, "\t")+"\n")

		// Compute and write stats for each column
		for _, rstat := range statlist {
			fmt.Fprint(outfile, rstat)
			for _, field := range headers {
				switch rstat {
				case "count":
					fmt.Fprintf(outfile, "\t%f", float64(len(datamap[field])))
				case "mean":
					m, _ := stats.Mean(datamap[field])
					fmt.Fprintf(outfile, "\t%f", m)
				case "std":
					stdev, _ := stats.StandardDeviation(datamap[field])
					fmt.Fprintf(outfile, "\t%f", stdev)
				case "min":
					mn, _ := stats.Min(datamap[field])
					fmt.Fprintf(outfile, "\t%f", mn)
				case "25%":
					tfp, _ := stats.Quartile(datamap[field])
					fmt.Fprintf(outfile, "\t%f", tfp.Q1)
				case "50%":
					fp, _ := stats.Quartile(datamap[field])
					fmt.Fprintf(outfile, "\t%f", fp.Q2)
				case "75%":
					sfp, _ := stats.Quartile(datamap[field])
					fmt.Fprintf(outfile, "\t%f", sfp.Q3)
				case "max":
					mx, _ := stats.Max(datamap[field])
					fmt.Fprintf(outfile, "\t%f", mx)
				}
			}
			// Write new line
			fmt.Fprintln(outfile)
		}

	}
	fmt.Println("Summary statistics have been written to", out)
	return nil

}
