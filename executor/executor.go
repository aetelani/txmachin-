package executor

import(
	"log"
//togen "txmachinae/tokengenerator"
"context"
_ "container/list"
"plugin"
"strings"
)

type DoWorkFun (func() (interface{}, error))
type Worker func (ctx context.Context, desc interface{}) DoWorkFun

//type D
type W func(context.Context, interface{}) (DoWorkFun)

type Process struct {
	
	// Id for prcess
	id []byte
	
	// Root context
	rootContext context.Context
}

func NewProcess() (*Process) {
	var process *Process = &Process {
		//id:togen.NewTokenGenerator().New(),
		rootContext:context.Background()}
	return process
}

func NewWorker(moniker string) ((func(context.Context, interface{}) (func(interface{}) (interface{}, error))), error) {
	const (
		pluginPath string = "../workers/"
		postfix string = ".so"
	)
	p, e := plugin.Open(pluginPath + strings.ToLower(moniker) + postfix)
	if e != nil {
		log.Fatal(e)
	}
	fun, e := p.Lookup(moniker)
	if e != nil {
		log.Fatal(e)
	}
	// Fix this with type alias when supported.
	funWorker := fun.(func(context.Context, interface{}) (func(interface{}) (interface{}, error)))
	return funWorker, nil
}

// Engine
type Engine interface {
}

type EngineImpl struct {
}

func NewEngine(engineImpl Engine) (Engine, error)  {
	log.Println("Initialized")
	return engineImpl, nil
}
