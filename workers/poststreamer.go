package main
import (
	"log"
	"net/http"
	"context"
	idgen "txmachinae/tokengenerator"
	"time"
	"io"
)

func PostStreamer(ctx context.Context, desc interface{}) (func(interface{}) (context.Context, interface{}, error)) {

	const (
		name string = "PostStreamer"
		version int = 0
		OutboundIdKey = "OutboundId"
		OutboundInitiatedKey = "OutboundInitiated"
		OutboundHTTPMethod = "POST"
	)
	
	log.Println("Initializing " + name + "." + string(version))

	var (
		localCtx context.Context
		client *http.Client
		response *http.Response
		url string = desc.(string)
	)

	uid := idgen.NewTokenGenerator().New()

	localCtx = context.WithValue(ctx, OutboundIdKey, uid)

	tr := http.DefaultTransport

	client = &http.Client{Transport: tr}
	
	return func(rt interface{}) (context.Context, interface{}, error) {

		req, err := http.NewRequest(OutboundHTTPMethod, url, rt.(io.ReadCloser))

		localCtx = context.WithValue(localCtx, OutboundInitiatedKey, time.Now().UTC())

		req = req.WithContext(localCtx)

		response, err = client.Do(req);
		
		return localCtx, nil, err
		}
}

