package executor

import(
	"testing"
	"context"
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

	// Set arguments for worker. Can be used to Initialize values.
	//rootContext, cancel :=  context.WithTimeout(context.Background(), 5*1000)
	funWork := worker(context.Background(), "http://example.com")

	// Execute work closure. Shares scope with worker. Go thread test
	//defer cancel()
	//go funWork()
	
	// Storing the result to same variable in Context. Taking runtime parameter timetout.
	value, err := funWork(int(3))

	ctx := value.(context.Context)

	// Convert results to string
	// Id is the unique id of the worker
	if err != nil {
		t.Log(err)
	} else if ctx.Value("id") == nil {
		t.Log("Fail id is nill")
	} else {
		t.Log(string(ctx.Value("id").(byte)))
	}

	t.Log(string(ctx.Value("responseData").([]byte)))
}
