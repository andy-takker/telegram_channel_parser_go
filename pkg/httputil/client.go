package httputil

import (
	"net"
	"net/http"
	"time"
)

func NewDefaultClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
		TLSHandshakeTimeout: 10 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
	}
	return &http.Client{
		Timeout:   20 * time.Second,
		Transport: tr,
	}
}
