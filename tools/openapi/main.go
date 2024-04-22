package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/nonchan7720/user-flex-feature/tools/openapi/internal/merge"
)

func readOpenAPI(file string, first bool) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	mp := make(map[string]interface{})
	if err := yaml.NewDecoder(f).Decode(&mp); err != nil {
		return nil, err
	}
	if !first {
		cpMp := map[string]interface{}{}
		fields := []string{"paths", "components"}
		for _, field := range fields {
			if v, ok := mp[field]; ok {
				cpMp[field] = v
			}
		}
		mp = cpMp
	}
	return yaml.Marshal(&mp)
}

func main() {
	files := os.Args[1:]
	var apis [][]byte
	for idx, file := range files {
		api, err := readOpenAPI(file, idx == 0)
		if err != nil {
			panic(err)
		}
		apis = append(apis, api)
	}
	buf, err := merge.YAML(apis...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf.String())
}
