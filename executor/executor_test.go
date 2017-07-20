package executor

import(
"testing"
"context"
_ "os"
_ "io"
"encoding/hex"
)

type mockEngine struct { }	

func TestInitialize(t *testing.T) {
	m := mockEngine{}
	_, err := NewEngine(m)
	if err != nil {
		t.Error("error", err)
	}
}


func TestWorker(t *testing.T) {

	// Load Downloader plugin
	worker, err := NewWorker("Downloader")

	if err != nil {
		t.Error("error", err)
	}

	funWork := worker(context.Background(), "http://example.com")

	value, err := funWork(int(3))

	ctx := value.(context.Context)

	if err != nil {
		t.Log("Got error")
		t.Log(err)
	} else if ctx.Value("id") == nil {
		t.Log("Fail id is nill")
	} else {
		t.Log("Output id")
		t.Log(hex.Dump(ctx.Value("id").([]byte)[:]))
	}

//	t.Log(string(ctx.Value("responseData").([]byte)))
//	io.Copy(os.Stdout, ctx.Value("ResponseBody").(io.ReadCloser))
}
