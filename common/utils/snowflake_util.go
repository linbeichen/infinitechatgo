package utils

import (
	"log"
	"sync"

	"github.com/bwmarrin/snowflake"
)

const (
	WorkerID     int64 = 1 // 机器 ID (0-31)
	DataCenterID int64 = 1 // 数据中心 ID (0-31)
)

var (
	node *snowflake.Node
	once sync.Once
)

func initNode() {
	once.Do(func() {
		nodeID := (DataCenterID << 5) | WorkerID
		var err error
		node, err = snowflake.NewNode(nodeID)
		if err != nil {
			log.Fatalf("初始化 Snowflake 节点失败: %v", err)
		}
		log.Println("Snowflake 节点初始化完成！")
	})
}

func NextId() int64 {
	initNode() // 确保已初始化
	return node.Generate().Int64()
}

func NextIdStr() string {
	initNode() // 确保已初始化
	return node.Generate().String()
}
