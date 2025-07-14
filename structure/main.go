package main

import (
	"fmt"
	"net/http"
	"strconv"

)

func bmiCaculatorHandler(w http.ResponseWriter, r *http.Request){
	heightStr := r.URL.Query().Get("height")
	weightStr := r.URL.Query().Get("weight")

	height, err1 := strconv.ParseFloat(heightStr, 64)
	weight, err2 := strconv.ParseFloat(weightStr, 64)

	if err1 != nil || err2 != nil || height == 0{
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	bmi := weight / ((height / 100) * (height/ 100))
	fmt.Fprintf(w, "You BMI index is :%.2f\n", bmi)
}

func main(){
	http.HandleFunc("/bmi", bmiCaculatorHandler)

	fmt.Println("Server is running on :8080.....")
	http.ListenAndServe(":8080", nil)
}