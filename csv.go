package main

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCsv(path string, sampleSize int) <-chan [][]string {
	frames := make(chan [][]string)
	go func(returnChan chan<- [][]string) {
		if file, err := os.Open(path); err == nil {
			defer file.Close()
			defer close(returnChan)
			reader := csv.NewReader(file)
			frame := [][]string{}
			for {
				record, err := reader.Read()
				if err == io.EOF {
					returnChan <- frame
					break
				} else if err != nil {
					panic(err)
				} else {
					if len(frame) == sampleSize {
						returnChan <- frame
						frame = [][]string{}
					}
					frame = append(frame, record)
				}
			}
		} else {
			panic(err)
		}
	}(frames)
	return frames
}
