package main

// Print struct template for a given json string.
// If there are multiple definitions for the same object, print the first one only.

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	j := string(inBytes)

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		panic(err)
	}

	printStructForMap("MyStruct", m)

	for _, v := range structz {
		fmt.Printf("%s\n", v)
	}
}

var structz = make(map[string]string)

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
			case []interface{}:
				b.WriteString(fmt.Sprintf("  %s []string `json:\"%s,omitempty\"`\n", capitalize(k), k))
				for _, item := range t {
					switch item.(type) {
					case map[string]interface{}:
						printStructForMap(capitalize(k)+"Item", item.(map[string]interface{}))
					}
				}
			case float64:
				vv := v.(float64)
				if vv == float64(int64(vv)) {
					b.WriteString(fmt.Sprintf("  %s int `json\"%s,omitempty\"`\n", capitalize(k), k))
				} else {
					b.WriteString(fmt.Sprintf("  %s float `json\"%s,omitempty\"`\n", capitalize(k), k))
				}
			default:
				b.WriteString(fmt.Sprintf("  %s %T `json:\"%s,omitempty\"`\n", capitalize(k), t, k))
			}
		}
	}
	b.WriteString(fmt.Sprintf("}\n"))
	structz[name] = b.String()
}

func capitalize(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
