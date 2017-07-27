package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
	idgen "txmachinae/tokengenerator"
)

const (
	Name                 string = "PostStreamer"
	Version              int    = 0
	OutboundIdKey               = "OutboundId"
	OutboundInitiatedKey        = "OutboundInitiated"
)

func PostStreamer(ctx context.Context, desc interface{}) func(interface{}) (context.Context, interface{}, error) {

	const (
		outboundHTTPMethod = "POST"
	)

	log.Println("Initializing " + Name + "." + string(Version))

	var (
		localCtx context.Context
		client   *http.Client
		response *http.Response
		url      string = desc.(string)
	)

	uid := idgen.NewTokenGenerator().New()

	localCtx = context.WithValue(ctx, OutboundIdKey, uid)

	tr := http.DefaultTransport

	client = &http.Client{Transport: tr}

	return func(rt interface{}) (context.Context, interface{}, error) {

		req, err := http.NewRequest(outboundHTTPMethod, url, rt.(io.ReadCloser))

		localCtx = context.WithValue(localCtx, OutboundInitiatedKey, time.Now().UTC())

		req = req.WithContext(localCtx)

		response, err = client.Do(req)

		return localCtx, nil, err
	}
}
