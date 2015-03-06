package main

// print struct template for a given json string

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

func main() {

	var j = `{"data": null, "id": 1234, "error": {"code":200, "message": "abc"}}`

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		panic(err)
	}

	printStructForMap("MyStruct", m)

	fmt.Println(strings.Join(structz, "\n"))

}

var structz []string

func printStructForMap(name string, m map[string]interface{}) {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("type %s struct {\n", capitalize(name)))
	for k, v := range m {
		if v == nil {
			b.WriteString(fmt.Sprintf("  %s string `json\"%s,omitempty\"`\n", capitalize(k), k))
		} else {
			switch t := v.(type) {
			case map[string]interface{}:
				b.WriteString(fmt.Sprintf("  %s %s `json:\"%s,omitempty\"`\n", capitalize(k), capitalize(k), k))
				printStructForMap(k, v.(map[string]interface{}))
			case float64:
				b.WriteString(fmt.Sprintf("  %s int `json\"%s,omitempty\"`\n", capitalize(k), k))
			default:
				b.WriteString(fmt.Sprintf("  %s %T `json:\"%s,omitempty\"`\n", capitalize(k), t, k))
			}
		}
	}
	b.WriteString(fmt.Sprintf("}\n"))
	structz = append(structz, b.String())
}

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
