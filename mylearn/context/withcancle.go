package main
import (
	"fmt"
	"context"
	"time"
)

/*
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
当调用CancelFunc函数时，context会被关闭，可用于防止goroutinue泄露
1. 创建根节点context---Background()
2. 将根context传入参数
3. 在子goroutinue中监控ctx.Done()是否取消,取消则退出goroutinue
4. 在父goroutinue中撤销context,则所有由它父context派生出来的context都会被取消（<-ctx.Done()）

*/


func withcancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				fmt.Println(time.Now())
				select {
				case <-ctx.Done():
					fmt.Println(ctx.Err())
					return // returning not to leaking goroutine
				case dst <- n:
					n ++
				}
			}
		}()
		return dst
	}
	//父context：context.Background()
	//子context: ctx
	ctx , cancel := context.WithCancel(context.Background())
	//defer cancel()  // cancel when we are finish consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	cancel()
	time.Sleep(5 * time.Second)
}

func main() {
	withcancel()
}
