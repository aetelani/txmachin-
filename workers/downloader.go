package main
import (
"log"
"net/http"
"context"
"io"
idgen "txmachinae/tokengenerator"
"time"
"bytes"
"fmt"
)

func Downloader(ctx context.Context, desc interface{}) (func(interface{}) (interface{}, error)) {

	const (
		name string = "downloader"
		version int = 0
	)

	var (
		localCtx context.Context
		client *http.Client
		url string = desc.(string)
	)

	uid := idgen.NewTokenGenerator().New()

	localCtx = context.WithValue(ctx, "id", uid)
	localCtx = context.WithValue(localCtx, "Name", name)
	localCtx = context.WithValue(localCtx, "Version",version)
	localCtx = context.WithValue(localCtx, "State", "init")
	
	tr := http.DefaultTransport

	client = &http.Client{Transport: tr}
	
	return func(rtdesc interface{}) (interface{}, error) {

		req, err := http.NewRequest("GET", url, nil)
		
		var response *http.Response

		req = req.WithContext(localCtx)

		failedCtx, failed := context.WithTimeout(localCtx,time.Duration(rtdesc.(int)) * time.Second)
		
		successCtx, success := context.WithCancel(localCtx)

		respChan := make(chan []byte)
		
		go func() {

			response, err = client.Do(req);
			
			var readSize int64 = 0
			for {
				// Needs buffer to separate requests or results get garbaged. Adjust as you will for perf
				buf := make([]byte, 8)
				if length, err := response.Body.Read(buf); length > 0 {
					readSize += int64(length)
					respChan <- buf[:length]
				} else if err == io.EOF && readSize > 0 {
					close(respChan)
					success()
					return
				} else {
					failed()
					log.Fatal(err)
					return
				}
			}
		}()
		
		var buffer bytes.Buffer

		for {
			select {
				case <-successCtx.Done():
					log.Println("Got Done from context")
					
					// PipeReader would proably be better or just channel for tight control 
					localCtx = context.WithValue(localCtx, "ResponseBody", buffer)
					localCtx = context.WithValue(localCtx, "State", "success")
					goto ready

				case <- failedCtx.Done():
					localCtx = context.WithValue(localCtx, "State", "failed")
					return localCtx, fmt.Errorf("Failed to download")

				case d := <-respChan:
					buffer.Write(d)
					break
			}
		}
ready:
		log.Println(buffer.String())

		return localCtx, nil
		}
}
