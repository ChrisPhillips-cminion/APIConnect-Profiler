package utils
/*
      Licensed Materials - Property of IBM
      © IBM Corp. 2019
*/
import (
	"encoding/json"
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
)

func JSONify(data model.TopLevel) []byte {
	TraceEnter("JSONify")
	br()
	json, _ := json.MarshalIndent(data, "  ", " ")
	fmt.Printf("%v \n", string(json))
	// fmt.Printf("%v \n",err)
	TraceExit("JSONify")
	return json
}
