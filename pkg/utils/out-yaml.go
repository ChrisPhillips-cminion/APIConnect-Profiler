package utils

import (
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
	"gopkg.in/yaml.v2"
)

func Yamlify(data model.TopLevel) {
	TraceEnter("Yamlify")
	br()
	yaml, _ := yaml.Marshal(data)
	fmt.Printf("%v \n", string(yaml))
	// fmt.Printf("%v \n",err)
	TraceExit("Yamlify")
}
