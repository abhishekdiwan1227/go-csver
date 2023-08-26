package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var helpText string = `go-csver
Command Usage:
gocsver -i(--input) [inputFile] 
gocsver -i(--input) [inputFile] -s(--size) [sampleSize]
inputFile : path for the input csv file
samepleSize : size of one sample (default is 10000)`

func main() {
	_, args := os.Args[0], os.Args[1:]
	if len(args) == 0 || len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		fmt.Println(helpText)
	} else if len(args) == 4 && (args[0] == "-i" || args[0] == "--input") && args[1] != "" && (args[2] == "-s" || args[2] == "--size") && args[3] != "" {
		inputPath := strings.Trim(args[1], " ")
		sampleSize, err := strconv.ParseInt(strings.Trim(args[3], " "), 0, 32)
		if err != nil {
			panic(err)
		}
		frames := ReadCsv(inputPath, int(sampleSize))
		schema := InferSchema(frames)
		fmt.Println(schema)
	} else if len(args) == 2 && (args[0] == "-i" || args[0] == "--input") && args[1] != "" {
		inputPath := strings.Trim(args[1], " ")
		frames := ReadCsv(inputPath, 10000)
		schema := InferSchema(frames)
		fmt.Println(schema)
	} else {
		fmt.Println("Invalid arguments.")
		fmt.Println(helpText)
	}
}
