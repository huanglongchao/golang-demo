package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

/**
	doc: https://books.studygolang.com/advanced-go-programming-book/ch1-basic/ch1-06-goroutine.html
	常见的并发模式
 */

func main(){
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	//ex5()
	//ex6()

	//pc1()
	//pc2()
	ps1()
}

//并发版本的hello world  ex1-ex6
//我们不能对一个未加锁状态的sync.Mutex进行解锁
func ex1(){
	var mu sync.Mutex

	go func() {
		fmt.Println("hello world")
		mu.Lock()
	}()
	mu.Unlock()
	//fatal error: sync: unlock of unlocked mutex
}

func ex2(){
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()
	mu.Lock()
}
//根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前
//发送才能接收
//无缓存管道
func ex3(){
	done := make(chan int)
	go func() {
		fmt.Println("hello world")
		<-done
	}()
	done<-1
}
//对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前
//有缓存管道
func ex4(){
	done := make(chan int,2)
	go func() {
		fmt.Println("hello world")
		done <- 1
	}()
	<-done
}
//等待N个线程完成后再进行下一步的同步操作
func ex5(){
	done := make(chan int,10)
	for i:=0; i < cap(done); i++{
		go func() {
			fmt.Println("hello world")
			done <- 1
		}()
	}
	for i:=0; i< cap(done); i++{
		<-done
	}
}
//等待N个线程完成后再进行下一步的同步操作 简单做法 使用sync.WaitGroup

func ex6(){
	var wg sync.WaitGroup

	for i:=0; i<10;i++{
		wg.Add(1)
		go func() {
			fmt.Println("hello world")
			wg.Done()
		}()
	}
	wg.Wait()
}
//生产者消费者模型
func Producer(factor int,out chan<- int){
	for i:=0;;i++{
		out <- i*factor
	}
}
func Consumer(in <-chan int){
	for v:= range in{
		fmt.Println(v)
	}
}

func pc1(){
	ch := make(chan int, 64)

	go Producer(3,ch)
	go Producer(5,ch)
	go Consumer(ch)

	time.Sleep(5*time.Microsecond)

}
func pc2(){
	ch := make(chan int,64)

	go Producer(3,ch)
	go Producer(5,ch)
	go Consumer(ch)

	//Ctrl+C退出
	sig := make(chan os.Signal,1)
	signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
//发布订阅模型 pub/sub

func ps1(){
	p := NewPublisher(100*time.Millisecond,10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool{
		if s,ok := v.(string); ok{
			return strings.Contains(s,"golang")
		}
		return false
	})

	p.Pulish("hello world!")
	p.Pulish("hello, golang!")

	go func() {
		for msg := range all{
			fmt.Println("all:",msg)
		}
	}()

	go func() {
		for msg := range golang{
			fmt.Println("golang:",msg)
		}
	}()

	time.Sleep(3*time.Second)
}

type(   //定义类型的用法
	subscriber chan interface{}  //订阅者为一个管道
	topicFunc func(v interface{}) bool //主题为一个过滤器
)

//发布者对象
type Publisher struct{
	m sync.RWMutex //读写锁
	buffer int //订阅队列的缓冲大小
	timeout time.Duration //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息
}

//构建一个发布者对象，可以设置发布超时时间和缓冲队列的长度
func NewPublisher(publishTimeout time.Duration,buffer int) *Publisher{
	return &Publisher{
		buffer: buffer,
		timeout:publishTimeout,
		subscribers:make(map[subscriber]topicFunc),
	}
}
//添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{}{
	ch := make(chan interface{},p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}
//添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{}{
	return p.SubscribeTopic(nil)
}
//退出订阅
func (p *Publisher) Evict(sub chan interface{}){
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers,sub)
	close(sub)
}
//发布一个主题
func (p *Publisher) Pulish(v interface{}){
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub,topic := range p.subscribers{
		wg.Add(1)
		go p.sendTopic(sub,topic,v,&wg)
	}
	wg.Wait()
}
//发送主题，可以容忍一定的超时
func (p *Publisher) sendTopic(sub subscriber,topic topicFunc,v interface{},wg *sync.WaitGroup){
	defer wg.Done()
	if topic != nil && !topic(v){
		return
	}
	select{
		case sub <- v:
		case <-time.After(p.timeout):
	}
}

//关闭发布者对象，同时关闭所有的订阅者管道
func (p *Publisher) Close(){
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers{
		delete(p.subscribers,sub)
		close(sub)
	}
}
//控制并发数
//赢者为王
//素数筛