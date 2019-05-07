package globalstore

import (
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	globalStoreInstance     GlobalStore
	initGlobalStoreInstance sync.Once
)

// Singleton returns the singleton global node instance.
func Singleton() GlobalStore {
	initGlobalStoreInstance.Do(func() {
		globalStoreInstance = NewGlobalStore(globaldb.GetGlobalDB())
	})
	return globalStoreInstance
}
