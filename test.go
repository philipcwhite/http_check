package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")


func web_check(url string) {
	resp, err := http.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
}

func main() {
	k.Load(file.Provider("test2.yaml"), yaml.Parser())

	mymap := k.StringMap("receivers.http_check.urls")


	fmt.Println("Duration:", k.String("receivers.http_check.duration"))
	for key, val := range mymap {
        //fmt.Println(key, val)
		fmt.Printf("Name: %s, URL: %s\n", key, val)
		web_check(val)

    }
	
}
