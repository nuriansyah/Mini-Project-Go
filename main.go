package main

import "fmt"

//"log"
//"net/http"
//"learn-go/handler"

func main() {
	/* mux := http.NewServeMux()



	mux.HandleFunc("/hello",handler.HelloHandler)
	mux.HandleFunc("/product",handler.ProductHandler)
	mux.HandleFunc("/",handler.HomeHandler)
	log.Println("Starting port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err) */

	scores := []int{8, 9, 9, 9, 10, 6, 5, 7, 3, 4, 5}

	var total int

	for _, score := range scores {
		total += score
	}

	fmt.Println(total)
}
