package main
import (
"log"
"net/http"
"context"
idgen "txmachinae/tokengenerator"
"time"
)

func DownloadStreamer(ctx context.Context, desc interface{}) (func(interface{}) (context.Context, interface{}, error)) {

	const (
		name string = "DownloadStreamer"
		version int = 0
	)
	
	log.Println("Initializing " + name + "." + string(version))

	var (
		localCtx context.Context
		client *http.Client
		url string = desc.(string)
	)

	uid := idgen.NewTokenGenerator().New()

	localCtx = context.WithValue(ctx, "InboundId", uid)

	tr := http.DefaultTransport

	client = &http.Client{Transport: tr}
	
	return func(rt interface{}) (context.Context, interface{}, error) {

		req, err := http.NewRequest("GET", url, nil)
		
		var response *http.Response

		req = req.WithContext(localCtx)
		
		localCtx = context.WithValue(localCtx, "InboundInitiated", time.Now().UTC())

		response, err = client.Do(req);
		
		return localCtx, response.Body, err
		}
}
