package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// fungsi ParseBody() bertujuan untuk melakukan debugging HTTP response
// dengan cara membaca seluruh body dengan melakukan parsing body tersebut menjadi byte

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
