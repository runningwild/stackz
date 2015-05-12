package stack

import (
	"net/http"
	"runtime"
)

func init() {
	http.HandleFunc("/stackz", func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 100000)
		n := runtime.Stack(buf, true)
		w.Write(buf[0:n])
	})
	go http.ListenAndServe(":8899", nil)
}
