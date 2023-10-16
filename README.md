# Week 4 Assignment: Command-Line Applications

## Project Summary
This project builds a command-line application in Go that reads in a CSV file, computes summary statistics and writes the output to a file. The program is designed to mimick the output of Python's .describe() function and R's summary() function. 

[montanaflynn/stats](https://github.com/montanaflynn/stats) package is used to compute the statistical calculations.

The input data for this assignment are provided in a comma-delimited text file (housesInput.csv) derived from a study of California housing prices (Miller 2015). There are R (runHouses.R) and Python/Pandas (runHouses.py) scripts for reading these data and computing summary statistics for each of seven input variables. Output from running these scripts should be plain text files of summary statistics: housesOutputR.txt and housesOutputPy.txt show output from running the R and Python/Pandas scripts, respectively.

As part of this assignment, benchmarking is conducted to test the CPU processing times using console/terminal commands for all 3 programming languages. To help compare results, each program runs (and writes the output) 100 times.

## Files
### *csv.go*
This file defines the csv2float() function which has a single parameter of the source of CSV data. It reads in the CSV data and converts the column values into float64 values. The function returns a slice of header values and a map with the header values as keys and slice of column values as values. I used Gerardi's (Gerardi 2021) code for reading comma-delimited text files as a base for building this function, however, changed the function to return a map containing the results of all the columns rather than returning a single column.

### *main.go*
This file contains the run() function which takes in 3 parameters: input filename, output filename and an integer value (N) for how many times to run the function. Note that N is hard-coded to 100 in the main() function. The run() function calls the csv2float() function in csv.go with the input file. After receiving the returned slice and map, it computes the statistics for each column and writes the result to the output file. The main() function in this file validates input parameters and calls the run() function.

> ./msds431_mod4 housesInput.csv housesOutputGo.txt

### *errors.go*
This file contains custom error definitions.

## Results

### Comparison
The results from the Go application match (within rounding error) the results from the R and Python/Pandas scripts.

### Benchmarking
Processing times were computed using the time command.

> time ./msds431_mod4 housesInput.csv housesOutputGo.txt
```
7.15s user
0.35s system
6.230 total
```
> time python runHouses.py
```
2.53s user
0.39s system
2.422 total
```
> time Rscript runHouses.R
```
3.58s user
0.18s system
3.799 total
```


## References

Miller, Thomas W. 2015. Modeling Techniques in Predictive Analytics with Python and R: A Guide to Data Science. Upper Saddle River, NJ: Pearson Education. [ISBN-13: 978-0-13-389206-2].

Gerardi, Ricardo. 2021. Powerful Command-Line Applications in Go: Build Fast and Maintainable Tools. Raleigh, NC: The Pragmatic Bookshelf. [ISBN-13: 978-1-68050-696-9].