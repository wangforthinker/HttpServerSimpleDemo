package main

import "net/http"
import "os"
import "fmt"
//import "strings"
import "io/ioutil"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("args error")
		return
	}
	
	body := args[2];
	
	url := args[1]+"?data="+ body
	fmt.Println("url:",url)
	resp, err := http.Get(url) //url needs http://
	
	
	if err != nil {
		fmt.Println("http post error:",err)
	}
	
	defer resp.Body.Close()
	
	recv_data, err :=  ioutil.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println("recv data error")
		return
	}
	
	fmt.Println("data: ",string(recv_data))
}