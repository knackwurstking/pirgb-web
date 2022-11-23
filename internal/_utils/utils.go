package utils

import "net/http"

func CopyHeaders(dstHeader, srcHeader http.Header) {
	for key, values := range srcHeader {
		for _, value := range values {
			dstHeader.Add(key, value)
		}
	}
}
