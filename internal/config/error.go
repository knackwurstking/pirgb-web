package config

import (
	"fmt"
	"net/http"
)

type RespError struct {
	Resp *http.Response
}

func (err *RespError) Error() string {
	return fmt.Sprintf("[%s] %s", err.Resp.Status, err.Resp.Request.URL)
}
