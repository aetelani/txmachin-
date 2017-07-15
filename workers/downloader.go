package main
import (
"log"
"net/http"
"io/ioutil"
"context"
)

type foo struct { url string }

func Downloader(ctx context.Context, desc interface{}) (func() (interface{}, error)) {
	log.Println("downloader")
	url := desc.(string)
	response, err := http.Get(url);
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	return func() (interface{}, error) { return responseData, nil }
}
