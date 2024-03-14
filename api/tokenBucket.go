package api

import "time"

// TokenBucket 令牌桶算法
type TokenBucket struct {
	capacity int           // 桶的容量
	rate     int           // 令牌生成速率，每秒生成的令牌数量
	tokens   chan struct{} // 令牌通道
}

func NewTokenBucket(capacity, rate int) *TokenBucket {
	tb := &TokenBucket{
		capacity: capacity,
		rate:     rate,
		tokens:   make(chan struct{}, capacity),
	}
	// 生成令牌
	go func() {
		for {
			time.Sleep(time.Second / time.Duration(tb.rate))
			select {
			case tb.tokens <- struct{}{}:
			default:
			}
		}
	}()
	return tb
}

func (tb *TokenBucket) AllowRequest() bool {
	select {
	case <-tb.tokens:
		return true
	default:
		return false
	}
}
