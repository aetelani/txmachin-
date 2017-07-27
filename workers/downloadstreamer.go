package main

import (
	"context"
	"log"
	"net/http"
	"time"
	idgen "txmachinae/tokengenerator"
)

func DownloadStreamer(ctx context.Context, desc interface{}) func(interface{}) (context.Context, interface{}, error) {

	const (
		name    string = "DownloadStreamer"
		version int    = 0
	)

	log.Println("Initializing " + name + "." + string(version))

	var (
		localCtx context.Context
		client   *http.Client
		response *http.Response
		url      string = desc.(string)
	)

	uid := idgen.NewTokenGenerator().New()

	localCtx = context.WithValue(ctx, "InboundId", uid)

	tr := http.DefaultTransport

	client = &http.Client{Transport: tr}

	return func(rt interface{}) (context.Context, interface{}, error) {

		req, err := http.NewRequest("GET", url, nil)

		localCtx = context.WithValue(localCtx, "InboundInitiated", time.Now().UTC())

		req = req.WithContext(localCtx)

		response, err = client.Do(req)

		return localCtx, &response.Body, err
	}
}
