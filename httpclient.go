package main

import "fmt"
import "io/ioutil"
import "net/http"

func main() {
	resp, err := http.Get("http://localhost:9090/api/query")
	if err != nil {
		// handle error
		fmt.Printf("error:%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}