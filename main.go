package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	//
	//	str := `{
	//      "str": "foo",
	//      "num": 100,
	//      "bool": false,
	//      "null": null,
	//      "array": ["foo", "bar", "baz"],
	//      "obj": { "a": 1, "b": 2 }
	//    }`
	//
	//
	var obj map[string]interface{}
	json.Unmarshal(input.Bytes(), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, err := f.Marshal(obj)
	if err != nil {
		fmt.Printf("%v+\n", obj)
		return
	}

	fmt.Println(string(s))
}
