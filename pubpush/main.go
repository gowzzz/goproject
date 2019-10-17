package main

import (
	"fmt"
	"goproject/pubpush/pubsub"
	"strings"
	"time"
)

// 天气预报之类的应用就可以应用这个并发模式
func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()
	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("hello,world!")
	p.Publish("hello,golang!")
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()
	time.Sleep(10 * time.Second)

}
