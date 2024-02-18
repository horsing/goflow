package main

import (
	"fmt"
	"github.com/horsing/goflow/samples/condition"
	"github.com/horsing/goflow/samples/loop"
	"github.com/horsing/goflow/samples/myflow"
	"github.com/horsing/goflow/samples/parallel"
	"github.com/horsing/goflow/samples/serial"
	"github.com/horsing/goflow/samples/single"

	goflow "github.com/horsing/goflow/v1"
)

func main() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		RedisPassword:     "redis",
		OpenTraceUrl:      "localhost:5775",
		WorkerConcurrency: 5,
		EnableMonitoring:  true,
		DebugEnabled:      true,
	}
	fs.Register("single", single.DefineWorkflow)
	fs.Register("serial", serial.DefineWorkflow)
	fs.Register("parallel", parallel.DefineWorkflow)
	fs.Register("condition", condition.DefineWorkflow)
	fs.Register("loop", loop.DefineWorkflow)
	fs.Register("myflow", myflow.DefineWorkflow)
	fmt.Println(fs.Start())
}
