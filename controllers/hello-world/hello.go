package new_hello

import "net/http"

func NewIndexHello() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Dwi"))
	}
}
