# Interval Merge
## Abstract
A Go application to merge overlapping intervals of integers. The intervals must be of the form [1,10][22,23][6,19][2,7]...<br>
A file or direct input can be used to supply intervals. File input is used for big interval quantities.
## Usage
### interval
You can pass intervals as argument using the command 'interval' followed by :<br>
`./merge interval "[1,10][22,23][6,19][2,7]"`<br>
### file
You can create a file containing intervals in the same form. These intervals can be multilined for better readability. Use the command 'file' followed by the file path:<br>
`./merge file ./input.intervals`.<br>
An interval must not be ordered by lower bound to upper bound. In case of file-input, also whitespaces will be ignored during parsing.<br>
### verbose
If the -v flag (verbose) is used, the application outputs durations of execution and memory consumption.
### gendata
You can generate a randomized interval list file using the command 'gendata' followed by the number of intervals to be generated. For better readability every 20 intervals comes a newline. Merge was tested with 10 Mio intervals. The resulting file is 'gen.intervals'. To test duration with generated data type:<br>
`./merge -v file ./gen.intervals`
## Build
There is a Makefile with the following targets:
- run<br>builds and runs directly with some test data
- build<br>builds an executable 'merge' using the version.yaml file
- test<br>builds and runs all test with coverage<br>

`make build` will build an executable 'merge'
## Performance
On my machine MacBook Pro 2.3 GHz 8-Core i9 SSD, merging 1 Mio intervals took less than a second. If more than 100 input intervals are given, they are displayed as count. If more than 100 resulting intervals exit, only the count is displayed. 
## Development duration
It took about 12h to realize and test all.
