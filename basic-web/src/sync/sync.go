package main

import (
	"bytes"
	"sync"
)

// sync.Mutex 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区

// 假设 info 是一个需要上锁的放在共享内存中的变量
type Info struct {
	mu  sync.Mutex
	Str string
}

// 在 sync 包中还有一个 RWMutex 锁  读写锁
// 他能通过 RLock() 来允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {

}

// 加锁的方式修改
func Update(info *Info) {
	info.mu.Lock()
	info.Str = "qxw"
	info.mu.Unlock()
}
