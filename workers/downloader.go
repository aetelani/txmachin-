package main
import (
"log"
"net/http"
"io/ioutil"
"context"
)

type foo struct { url string }

func Downloader(ctx context.Context, desc interface{}) (func() (interface{}, error)) {
	log.Println("Initializing Downloader")
	url := desc.(string)
	
	// use togen and set cancel
	localCtx := context.WithValue(ctx, "id",byte('A'))

	return func() (interface{}, error) { 

		response, err := http.Get(url);

		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		localCtx = context.WithValue(localCtx, "responseData", responseData)

		return localCtx, nil
		}
}
