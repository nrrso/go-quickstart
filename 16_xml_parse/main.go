package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Foaas Struct
type Foaas struct {
	Message  string `xml:"http://foaas.com/fuckoff message"`
	Subtitle string `xml:"http://foaas.com/fuckoff subtitle"`
}

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://foaas.com/off/Tom/Everyone", nil)
	req.Header.Add("Accept", "application/xml")
	resp, _ := client.Do(req)
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Foaas
	xml.Unmarshal(bytes, &s)
	// string_body := string(bytes)
	fmt.Println(s.Message)
	fmt.Println(s.Subtitle)
	resp.Body.Close()
}
