package playground

import (
	"encoding/json"
	"os"
)

func ToJson(i interface{}) string {

	bytes, err := json.Marshal(i)
	if err != nil {
		_, err = os.Stderr.WriteString("error json marshal")
		if err != nil {
			panic(err)
		}
		return "error"
	}

	return string(bytes)

}
