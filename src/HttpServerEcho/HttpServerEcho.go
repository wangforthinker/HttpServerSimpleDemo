package main

import "net/http"
import "fmt"

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r * http.Request){
		s := r.FormValue("data") //form data
		b1 := [] byte(s)
		w.Write(b1)
	})
	
	err := http.ListenAndServe(":8888", nil )
	if err != nil {
		fmt.Println("http server create error: ",err.Error())
		return
	}
}
