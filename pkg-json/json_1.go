package main

import (
	"encoding/json"
	"fmt"
)

type spec struct {
	Version string `json:"cdiVersion"`
	Kind    string `json:"kind"`
}

func main() {
	// Example JSON data
	jsonData := `{"cdiVersion": "1.0", "kind": "example"}`

	// Unmarshal JSON into struct
	var s spec
	err := json.Unmarshal([]byte(jsonData), &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the struct
	fmt.Printf("%+v\n", s)

	// Marshal struct to JSON
	jsonOutput, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON output
	fmt.Println(string(jsonOutput))
}
