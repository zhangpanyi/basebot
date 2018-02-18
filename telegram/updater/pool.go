package updater

import (
	"log"

	"github.com/Jeffail/tunny"
)

// Pool 工作队列
type Pool struct {
	pool *tunny.Pool
}

// NewPool 创建工作队列
func NewPool(numWorkers int) *Pool {
	pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Pool goroutine recoverd, %v\n", err)
			}
		}()

		callback, ok := payload.(func())
		if ok {
			callback()
		}
		return nil
	})
	return &Pool{pool: pool}
}

// Async 添加异步任务
func (pool *Pool) Async(callback func()) {
	go pool.pool.Process(callback)
}
