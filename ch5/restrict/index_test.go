//package restrict

package restrict_test

import (
	"fmt"
	"gocamp/ch5/restrict"
	"sync"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	d := restrict.New(10)
	wg := sync.WaitGroup{}
	c := 10000
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			if !d.IsPass() {
				time.Sleep(1 * time.Second)
				fmt.Println("不能请求")
				return
			}
			fmt.Println("请求正常")
			time.Sleep(1 * time.Second)
		}(&wg)
	}
	wg.Wait()
}
