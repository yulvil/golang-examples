package main

// print struct template for a given json string

import (
	"fmt"
	"unicode"
)
import "encoding/json"

// import "bytes"

func main() {

	var j = `{"data": null, "id": 1234, "error": {"code":200, "message": "abc"}}`

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		panic(err)
	}

	printStructForMap("MyStruct", m)

	//fmt.Printf("%+v\n", m)
}

// var structList[]

func printStructForMap(name string, m map[string]interface{}) {
	fmt.Printf("type %s struct {\n", capitalize(name))
	for k, v := range m {
		if v == nil {
			fmt.Printf("  %s string `json\"%s,omitempty\"`\n", capitalize(k), k)
		} else {
			switch t := v.(type) {
			case map[string]interface{}:
				fmt.Printf("  %s %s `json:\"%s,omitempty\"`\n", capitalize(k), capitalize(k), k)
				printStructForMap(k, v.(map[string]interface{}))
			case float64:
				fmt.Printf("  %s int `json\"%s,omitempty\"`\n", capitalize(k), k)
			default:
				fmt.Printf("  %s %s `json:\"%s,omitempty\"`\n", capitalize(k), t, k)
			}
		}
	}
	fmt.Printf("}\n")
}

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
