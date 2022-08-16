package router

import "net/http"

func CopyHeader(dst, src http.Header) {
	for n, v := range src {
		for _, vv := range v {
			dst.Add(n, vv)
		}
	}
}
