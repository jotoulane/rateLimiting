package api

import (
	"sync"
	"time"
)

// SlidingWindow 滑动窗口算法
type SlidingWindow struct {
	window    []time.Time   // 时间窗口
	threshold int           // 时间窗口内的请求阈值
	duration  time.Duration // 时间窗口的大小
	mutex     sync.Mutex    // 互斥锁
}

func NewSlidingWindow(threshold int, duration time.Duration) *SlidingWindow {
	return &SlidingWindow{
		window:    make([]time.Time, 0),
		threshold: threshold,
		duration:  duration,
	}
}

func (sw *SlidingWindow) AllowRequest() bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	// 移除超出时间窗口的请求
	currentTime := time.Now()
	for len(sw.window) > 0 && currentTime.Sub(sw.window[0]) > sw.duration {
		sw.window = sw.window[1:]
	}

	// 检查请求数量是否超出阈值
	if len(sw.window) < sw.threshold {
		sw.window = append(sw.window, currentTime)
		return true
	}

	return false
}
