package main
import (
"log"
"net/http"
"io/ioutil"
"context"
_ "errors"
"time"
)

func Downloader(ctx context.Context, desc interface{}) (func(interface{}) (interface{}, error)) {
	log.Println("Initializing Downloader")
	url := desc.(string)
	
	// use togen and set cancel
	var localCtx context.Context
	var cancel func()
	localCtx = context.WithValue(ctx, "id",byte('A'))

	return func(rtdesc interface{}) (interface{}, error) {

		localCtx, cancel = context.WithTimeout(localCtx,time.Duration(rtdesc.(int)) * time.Second)

		response, err := http.Get(url);

		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()
		defer cancel()

		// Dummy implementation reading all in memory. TODO Streaming
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		} else {
			cancel() // Calling cancel now. no need to wait in select
		}

		localCtx = context.WithValue(localCtx, "responseData", responseData)

		for {	
			select {
				// Timeout or cancel
				case <- localCtx.Done(): log.Println("context Done"); goto end
			}
		}
end:
		return localCtx, nil
		}
}
