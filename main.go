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

	var obj map[string]interface{}
	json.Unmarshal(input.Bytes(), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	if len(string(s)) < 10 {
		fmt.Println(input.Text())
		return
	}

	fmt.Println(string(s))
}
