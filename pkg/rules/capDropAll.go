package rules

import (
  "bytes"
  "fmt"
  "github.com/thedevsaddam/gojsonq"
  "strings"
)

func CapDropAll(json []byte) int {
	spec := getSpecSelector(json)
	containers := 0

	capDrop := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".containers").
		Only("securityContext.capabilities.drop")

	if capDrop != nil && strings.Contains(fmt.Sprintf("%v", capDrop), "[map[drop:[ALL]") {
		containers++
	}

	capDropInit := gojsonq.New().Reader(bytes.NewReader(json)).
		From(spec + ".initContainers").
		Only("securityContext.capabilities.drop")

  if capDropInit != nil && strings.Contains(fmt.Sprintf("%v", capDropInit), "[map[drop:[ALL]") {
		containers++
	}

	return containers
}
