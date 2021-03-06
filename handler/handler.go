package handler


import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	"path"
	"log"

	

)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First time build Web on Golang"))
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}

	_, err := template.ParseFiles(path.Join("views","index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
	
	//err = tmpl.Execute(w,nil)
	//err = tmpl.Execute(w,r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
	
}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w,r)
		return
	} 

	fmt.Fprintf(w,"Product page: %d", idNumb)
}