package mvc

import (
	"github.com/ReSc/fmt"
	"net/http"
	"time"
)

type logHandler struct {
	Inner http.Handler
}

func (lh *logHandler) Name() string                          { return "log" }
func (lh *logHandler) Build(inner http.Handler) http.Handler { return &logHandler{inner} }
func (lh *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rw := NewResponseWriter(w)
	lh.Inner.ServeHTTP(rw, r)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printline("[%d]\t%v\t%v", rw.Status(), r.URL.Path, duration)
}
