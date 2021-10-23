package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":		
		 http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		number := r.FormValue("number")
		res := happyNumber(number)
		fmt.Fprintf(w, res)
	default:
		fmt.Fprintf(w, "Sorry")
	}
}

func happyNumber(s string) string{
	//s := req.FormValue("n")
	var a [1000]int
	n, er := strconv.Atoi(s)
	er = er
	for i := 0; i < 1000; i++ {
		sum := 0
		x := n
		for x > 0 {
			t := x % 10
			sum += t * t
			x = x / 10
		}
		n = sum
		if findInMas(a, n) {
			break
		}
		a[i] = n
	}
	
	if n == 1 {
		return "true"
	} else {
		return "false"
	}
}

func findInMas(arr [1000]int, x int) bool {
	for i := 0; i < 1000; i++ {
		if x == arr[i] {
			return true
		}
	}
	return false
}
