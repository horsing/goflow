package runtime

import (
	"github.com/horsing/goflow/core/sdk/executor"
)

type Runtime interface {
	Init() error
	CreateExecutor(*Request) (executor.Executor, error)
}
