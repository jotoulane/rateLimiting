package api

import "sync"

// LeakyBucket 漏桶算法
type LeakyBucket struct {
	capacity int        // 桶的容量
	rate     int        // 漏桶速率，每秒流出的数量
	tokens   int        // 当前桶中的令牌数量
	mutex    sync.Mutex // 互斥锁
}

func NewLeakyBucket(capacity, rate int) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		tokens:   0,
	}
}

func (lb *LeakyBucket) AllowRequest() bool {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	// 生成令牌
	lb.tokens += lb.rate
	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	// 检查是否有令牌
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	return false
}
