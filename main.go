package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	// メインルーチンが終わってしまったら、サブルーチン(getLuckyNum())が
	// 終わる前に処理が終了してしまう

	// サブルーチンが終わる前に、メインルーチンで待ちたい場合はsync.WaitGroupを使う
	var wg sync.WaitGroup

	thread := 100

	// 内部カウンタにthreadをセットする
	wg.Add(thread)

	for i := 0; i < thread; i++ {

		fmt.Printf("スレッドNo. = %d\n", i)

		go func(i int) {
			// 内部カウンタを-1する
			defer wg.Done()
			timeWaite(i)
		}(i)
	}

	// 内部カウンタが0になるまでまつ
	wg.Wait()
	fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
}

func timeWaite(i int) {
	fmt.Printf("--- time wait --- スレッドNo. = %d\n", i)
	time.Sleep(time.Second * 10)
	fmt.Printf("--- finish --- スレッドNo. = %d\n", i)
}
