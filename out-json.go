package main

import (
	"encoding/json"
	"fmt"
)

func JSONify(data topLevel) []byte {
	TraceEnter("JSONify")
	br()
	json, _ := json.MarshalIndent(data, "  ", " ")
	fmt.Printf("%v \n", string(json))
	// fmt.Printf("%v \n",err)
	TraceExit("JSONify")
	return json
}
