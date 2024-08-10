package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:generate go run .
const amisDir = "components"

const temp = `package components

func %s() SchemaNode {
	%sOption
}

type %sOption struct {
	
}



`

func main() {
	err := filepath.WalkDir(amisDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		fmt.Println(d.Name())
		name := strings.TrimSuffix(d.Name(), ".md")
		name = toPascalCase(name)

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func toPascalCase(s string) string {
	arr := strings.Split(s, "-")
	for i, v := range arr {
		arr[i] = strings.Title(v)
	}
	return strings.Join(arr, "")
}
