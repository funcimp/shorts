package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
	+----------------------------+
	| learn things for the lulz. |
	+----------------------------+

`))
	})

	http.ListenAndServe(":1337", nil)
}
