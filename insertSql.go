package rest

import (
	"strings"
)

type InsertMiddleware struct {
	f func(string, string)
}

// MiddlewareFunc makes InsertMiddleware implement the Middleware interface.
func (im *InsertMiddleware) MiddlewareFunc(h HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, r *Request) {
		// call the handler
		h(w, r)

		if im.f != nil {
			ip := r.Header.Get("X-Real-IP")
			if ip == "" {
				ip = strings.Split(r.RemoteAddr, ":")[0]
			}
			go im.f(r.RequestURI, ip)
		}
	}
}

func (im *InsertMiddleware) SetHandler(f func(string, string)) {
	im.f = f
}
