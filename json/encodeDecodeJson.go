package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type gResult struct {
	IsOk      bool `json:"is_ok"`
	TestField int  `json:"test_field"`
}

func decodeJsonApi() {
	url := ""

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()

	var r gResult

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%#v\n", r)
	fmt.Println(r.IsOk)
}

func decodeJsonStr() {
	jsonStr := `{
		"is_ok": false,
		"data": {"name": "marvin"}
	}`

	var r interface{}

	err := json.Unmarshal([]byte(jsonStr), &r)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%#v\n", r)
	//fmt.Println(r.IsOk)

	var r1 map[string]interface{}

	err = json.Unmarshal([]byte(jsonStr), &r1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%#v\n", r1)
	fmt.Println(r1["is_ok"], r1["data"].(map[string]interface{})["name"])
}

func encodeJson() {
	var m = make(map[string]interface{})

	m["code"] = 200
	m["data"] = map[string]interface{}{
		"name": "marvin",
	}
	data, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	data, err = json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	array := [...]int{1, 2, 3, 4}
	data, err = json.Marshal(array)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func main() {
	//decodeJsonApi()
	//decodeJsonStr()
	encodeJson()
}
