package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var priorityTypes = []reflect.Kind{
	reflect.Bool,
	reflect.Uint8,
	reflect.Int8,
	reflect.Uint16,
	reflect.Int16,
	reflect.Uint32,
	reflect.Int32,
	reflect.Uint64,
	reflect.Int64,
	reflect.String}

var priorityTypeConverters = map[reflect.Kind]func(string) (bool, any){
	reflect.Bool: func(val string) (bool, any) {
		if convertedVal, err := strconv.ParseBool(val); err == nil {
			return true, convertedVal
		} else {
			return false, ""
		}
	},
	reflect.Uint8: func(val string) (bool, any) {
		if convertedVal, err := strconv.ParseUint(val, 10, 8); err == nil {
			return true, convertedVal
		} else {
			return false, ""
		}
	},
	reflect.Int8: func(val string) (bool, any) {
		if convertedval, err := strconv.ParseInt(val, 10, 8); err == nil {
			return true, convertedval
		} else {
			return false, ""
		}
	},
	reflect.Uint16: func(val string) (bool, any) {
		if convertedVal, err := strconv.ParseUint(val, 10, 16); err == nil {
			return true, convertedVal
		} else {
			return false, ""
		}
	},
	reflect.Int16: func(val string) (bool, any) {
		if convertedval, err := strconv.ParseInt(val, 10, 16); err == nil {
			return true, convertedval
		} else {
			return false, ""
		}
	},
	reflect.Uint32: func(val string) (bool, any) {
		if convertedVal, err := strconv.ParseUint(val, 10, 32); err == nil {
			return true, convertedVal
		} else {
			return false, ""
		}
	},
	reflect.Int32: func(val string) (bool, any) {
		if convertedval, err := strconv.ParseInt(val, 10, 32); err == nil {
			return true, convertedval
		} else {
			return false, ""
		}
	},
	reflect.Uint64: func(val string) (bool, any) {
		if convertedVal, err := strconv.ParseUint(val, 10, 64); err == nil {
			return true, convertedVal
		} else {
			return false, ""
		}
	},
	reflect.Int64: func(val string) (bool, any) {
		if convertedval, err := strconv.ParseInt(val, 10, 64); err == nil {
			return true, convertedval
		} else {
			return false, ""
		}
	},
	reflect.String: func(val string) (bool, any) {
		return true, val
	},
}

func getPriorityType(val string) reflect.Kind {
	sanitizedVal := strings.Trim(val, " ")
	for _, kind := range priorityTypes {
		convFunc := priorityTypeConverters[kind]
		if ok, _ := convFunc(sanitizedVal); ok {
			return kind
		}
	}
	return reflect.String
}

func mergeSchemas(current map[int]reflect.Kind, other map[int]reflect.Kind) map[int]reflect.Kind {
	if len(other) == 0 {
		return current
	}
	findPriority := func(kind reflect.Kind) int {
		for i, val := range priorityTypes {
			if val == kind {
				return i
			}
		}
		return -1
	}
	schema := make(map[int]reflect.Kind)
	for key, currentVal := range current {
		if otherVal, ok := other[key]; ok {
			if findPriority(currentVal) > findPriority(otherVal) {
				schema[key] = currentVal
			} else {
				schema[key] = otherVal
			}
		}
	}
	return schema
}

func inferTypes(frame [][]string) map[int]reflect.Kind {
	var finalSchema map[int]reflect.Kind
	for i, row := range frame {
		rowSchema := make(map[int]reflect.Kind)
		for j := range row {
			inferredType := getPriorityType(frame[i][j])
			rowSchema[j] = inferredType
		}
		finalSchema = mergeSchemas(rowSchema, finalSchema)
	}
	return finalSchema
}

func InferSchema(frames <-chan [][]string) map[int]reflect.Kind {
	var finalSchema map[int]reflect.Kind
	frameCount := 0
	for frame := range frames {
		finalSchema = mergeSchemas(inferTypes(frame), finalSchema)
		fmt.Printf("Frame %d: Schema => %v\n", frameCount, finalSchema)
		frameCount++
	}
	fmt.Printf("Total frames: %d\n\n", frameCount)
	return finalSchema
}
