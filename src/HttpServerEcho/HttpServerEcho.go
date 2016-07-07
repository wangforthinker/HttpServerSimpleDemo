package main

import (
	"net/http"
	"io/ioutil"
)
import "fmt"

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r * http.Request){
	/*	s := r.FormValue("data") //form data
		b1 := [] byte(s)
		w.Write(b1)
		*/
		body,err := ioutil.ReadAll(r.Body)
		if err != nil {
			//error
		} else {
			w.Write(body)
		}
	})
	
	err := http.ListenAndServe(":8888", nil )
	if err != nil {
		fmt.Println("http server create error: ",err.Error())
		return
	}
}
