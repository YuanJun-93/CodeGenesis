package idgen

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// Init initializes the snowflake node with a machine/node ID.
// machineId must be between 0 and 1023.
func Init(machineId int64) {
	once.Do(func() {
		// Set custom epoch to 2024-01-01 (Best Practice)
		snowflake.Epoch = 1704067200000
		var err error
		node, err = snowflake.NewNode(machineId)
		if err != nil {
			// Fallback or panic? For an ID generator, failure to init is critical.
			// However, NewNode only errors if machineId is out of bounds.
			panic("failed to initialize snowflake node: " + err.Error())
		}
	})
}

// NextId generates a new snowflake ID.
func NextId() int64 {
	if node == nil {
		// Fallback for safety/tests if Init wasn't called.
		// In production, Init should always be called first.
		Init(1)
	}
	return node.Generate().Int64()
}
