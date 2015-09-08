```
$ printf '{"id":123, "name":"abc", "details": {"desc": "Help", "data": [10,20]}}' | go run json_to_struct.go 

type Details struct {
  Desc string `json:"desc,omitempty"`
  Data []string `json:"data,omitempty"`
}

type MyStruct struct {
  Details Details `json:"details,omitempty"`
  Id int `json"id,omitempty"`
  Name string `json:"name,omitempty"`
}
```
