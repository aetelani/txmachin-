package main
import (
"log"
"net/http"
"io/ioutil"
"context"
)

func Downloader(ctx context.Context, url interface{}) {
	log.Println("downloader")
	response, err := http.Get(url.(string))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	log.Println(responseString)
}
