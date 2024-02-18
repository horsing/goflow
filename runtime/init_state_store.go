package runtime

import (
	redisStateStore "github.com/horsing/goflow/core/redis-statestore"
	"github.com/horsing/goflow/core/sdk"
)

func initStateStore(redisURI string, password string) (stateStore sdk.StateStore, err error) {
	stateStore, err = redisStateStore.GetRedisStateStore(redisURI, password)
	return stateStore, err
}
