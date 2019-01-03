package vgwebsrvr

import (
	"fmt"
	"net/http"
)

type Flex struct{}
func (f Flex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>")
	fmt.Fprint(w, "<img src=flex.png>")
	fmt.Fprint(w, "<h1>Flex Appliance Console</h1>")
	fmt.Fprint(w, "</body></html>")
}

func main() {
	f := Flex{}
	fmt.Println("Hello, world!")
	http.ListenAndServe(":4000", f)
}
