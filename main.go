package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func pick(data interface{}, path string) (interface{}, bool) {
	parts := strings.Split(path, ".")
	current := data
	for _, part := range parts {
		if part == "" {
			continue
		}
		m, ok := current.(map[string]interface{})
		if !ok {
			return nil, false
		}
		current, ok = m[part]
		if !ok {
			return nil, false
		}
	}
	return current, true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: jsonpick <field.path> [field2...]\nReads JSON from stdin, prints selected fields.\n")
		os.Exit(1)
	}
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		os.Exit(1)
	}
	var arr []interface{}
	if err := json.Unmarshal(input, &arr); err == nil {
		for _, item := range arr {
			printPicks(item, os.Args[1:])
		}
		return
	}
	var obj interface{}
	if err := json.Unmarshal(input, &obj); err != nil {
		fmt.Fprintf(os.Stderr, "invalid JSON: %v\n", err)
		os.Exit(1)
	}
	printPicks(obj, os.Args[1:])
}

func printPicks(data interface{}, paths []string) {
	if len(paths) == 1 {
		val, ok := pick(data, paths[0])
		if !ok {
			fmt.Println("null")
			return
		}
		printValue(val)
		return
	}
	result := make(map[string]interface{})
	for _, p := range paths {
		val, ok := pick(data, p)
		if ok {
			result[p] = val
		} else {
			result[p] = nil
		}
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}

func printValue(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Println(val)
	case nil:
		fmt.Println("null")
	default:
		out, _ := json.Marshal(val)
		fmt.Println(string(out))
	}
}
