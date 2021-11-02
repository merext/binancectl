package utils

import (
	"encoding/json"
	"fmt"
)

func OPrint(obj interface{}) {
	a, _ := json.Marshal(obj)
	fmt.Println(string(a))
}
