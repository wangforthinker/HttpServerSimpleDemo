package main

import (
	"io/ioutil"
	"bytes"
	"net/http"
)
import "encoding/json"
import "os"
import "fmt"
import "strconv"

type AttachInfo struct {
	Age int  	`json:"age"`
	Id_card string  `json:"id_card"`
	Label []string   `json:"label"`
}

type UserInfo struct {
	ID int 
	Name string
	Message []string
	Info AttachInfo
}

func main () {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("args error")
		return
	}
	
	id,err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("ID imput error")
		return
	}
	
	
	msg_len := len(args) - 4
	Msg := make([]string,msg_len)
	
	for i:= 4; i < msg_len + 4; i++ {
		Msg[i - 4] = args[i]
	}
	
	attach_info := AttachInfo {
		Age : 30,
		Id_card : "511xxxxxxxxx0001",
		Label : []string {"high","strength"},
	}
	
	var userinfo UserInfo = UserInfo{
		ID : id,
		Name : args[3],
		Message : Msg,
		Info : attach_info,
	}
	
	
	b,err := json.Marshal(userinfo)
	if err != nil {
		fmt.Println("Json Marshal error ")
		return
	}
	
	os.Stdout.Write(b)
	
	req, err := http.NewRequest("POST", args[1], bytes.NewReader(b))
	if err != nil {
		//error
		fmt.Println("http request error")
		return
	}
	
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	resp,err := client.Do(req)
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error")
		return
	}
	
	fmt.Println(string(body))
	
	var users UserInfo
	err = json.Unmarshal(body, &users) // json decode to the detail object 
	
	
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	
	fmt.Println("Output:")
	fmt.Println("%+v",users)
	fmt.Println("%+v",users.Info)
	
	//json decode to map[string] {}
	var map_info map[string] interface{}
	err = json.Unmarshal(body, &map_info)
	
	if err != nil {
		fmt.Println("json decode map_info error")
		return
	}
	
	for i,v := range map_info {
		fmt.Println(i,":",v)
	}
	
	
//	attch_info_map := map_info["Info"]
/*	
	var attach_info_maps map[string] interface {}
	err = json.Unmarshal(attch_info_map, &attach_info_maps)
	
	if err != nil {
		fmt.Println("json decode attch_info_map error")
		return
	}
	
	for i,v := range attach_info_maps {
		fmt.Println(i,":",v)
	}
*/	
	
}