package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func Yamlify(data topLevel) {
	TraceEnter("Yamlify")
	br()
	yaml, _ := yaml.Marshal(data)
	fmt.Printf("%v \n", string(yaml))
	// fmt.Printf("%v \n",err)
	TraceExit("Yamlify")
}
