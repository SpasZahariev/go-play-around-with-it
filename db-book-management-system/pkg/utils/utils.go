package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, absorber interface{}) {

	// attemts to read the json body and put it in a variable called "body"
	if body, err := ioutil.ReadAll(request.Body); err != nil {
		// try writing the json into a go interface/struct
		if err := json.Unmarshal([]byte(body), absorber); err != nil {
			return
		}
	}
}