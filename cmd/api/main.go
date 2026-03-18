package main

import (
	"fmt"
	"net/http"

	apphttp "github.com/ajama475/trust-api/internal/http"
)

func main() {
	//create new router
	router := apphttp.NewRouter()
	fmt.Println("Trust API is starting...")
	
	//start in port 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}