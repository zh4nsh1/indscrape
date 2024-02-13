package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func ReadPage(url string){
	
	//Get our URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)	
	}
	
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)


}


func main(){
	
	//Take args
	args := os.Args

	//No args? Quit and print a message telling the user to enter a URL
	if len(args) < 2 {
		fmt.Println("Usage: indscrape [URL]")
		return
	}


}
