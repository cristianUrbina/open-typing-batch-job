package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func jsonFormatter(s []byte) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, s, "", " ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %v", err)
	}
	return prettyJSON.String()
}
