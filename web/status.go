package web

import (
	"fmt"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UP")
}
