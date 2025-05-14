package server

import (
	"fmt"
	"net/http"
)

func initRoutes(){
	http.HandleFunc("/",index)

	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			initCompanies(w,r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
	})

	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			getCompanies(w,r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
	})
}