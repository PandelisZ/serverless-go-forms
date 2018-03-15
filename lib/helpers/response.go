package helpers

import (
	"encoding/json"
)

type Response struct {
	Status string
	Error  interface{}
}

func (r *Response) ToString() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func ResponseFail(err string) string {
	r := Response{
		Status: "fail",
		Error:  err,
	}
	return r.ToString()
}

func ResponseSuccess() string {
	r := Response{
		Status: "ok",
	}
	return r.ToString()
}
