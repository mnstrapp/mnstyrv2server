package logger

import (
	"log"
	"net/http"
	"strings"
)

type Logger struct {
	Next http.Handler
}

func NewLogger(h http.Handler) Logger {
	return Logger{Next: h}
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	url := r.RequestURI

	log.Printf("[%s] %s", method, url)
	l.Next.ServeHTTP(w, r)
}
