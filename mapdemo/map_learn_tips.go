package mapdemo

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

/** golang学习笔记 之 字典
 */

func main(){
	ct()
	op()
	iter()
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	//ex6()
	//ex7()
	ex8()
}
//创建演示
func ct(){
	//创建
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct{
		x int
	}{
		1: {x:100},
		2: {x:200},
	}
	fmt.Println(m,m2)
}
//基本操作演示
func op(){
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	m["a"] = 10 //修改
	m["c"] = 30 //新增

	if v,ok := m["d"]; ok{ //使用ok-idiom 判断key是否存在，返回值
		println(v)
	}
	delete(m,"d") //删除键值对，不存在时，不会报错
}
//迭代
func iter(){
	m := make(map[string]int)
	for i:=0; i<8; i++{
		m[string('a'+i)] = i
	}
	for i:=0; i<4; i++{ //每次迭代返回的顺序不一样
		for k,v := range m{
			print(k,":",v," ")
		}
		println()
	}
}
//map被设计成 not addressable（因内存访问安全和哈希算法的缘故）,故不能直接修改value成员
func ex1(){
	type user struct {
		name string
		age byte
	}
	m := map[int]user{
		1: {"Tom",19},
	}
	//m[1].age += 1 //报错
	fmt.Println(m)
}
//正确的做法是 返回整个value 修改后再设置回去，或者直接用指针类型
func ex2(){

	type user struct {
		name string
		age byte
	}
	m := map[int]user{
		1:{"Tom",19},
	}
	u := m[1]
	u.age += 1
	m[1] = u

	m2 := map[int]*user{
		1:&user{"Jack",20},
	}
	m2[1].age++
}
//不能对nil字典进行写操作，但是却能读
func ex3(){
	var m map[string]int
	println(m)
	//m["a"] = 1 //panic: assignment to entry in nil mapdemo
}
//内容为空的字典，与nil是不同的
func ex4(){
	var m1 map[string]int
	m2 := map[string]int{}
	println(m1 == nil, m2 == nil) //true false
}

//安全
//在迭代期间删除或新增键值是安全的
func ex5(){
	m := make(map[int]int)
	for i:=0; i<10; i++{
		m[i] = i+10
	}
	for k:= range m{ //range 一个Map的时候，也可以使用一个返回值，这个默认的返回值就是Map的键
		if k==5{
			m[100] = 1000
		}
		delete(m,k)
		fmt.Println(k,m)
	}
}
//并发读写map 进程崩溃 无法recovery
//fatal error: concurrent mapdemo read and mapdemo write
func ex6(){

	m := make(map[string]int)

	go func() {
		for{
			m["a"] += 1 //写操作
			time.Sleep(time.Microsecond)
		}
	}()
	go func() {
		for{
			_ = m["b"] //读操作
			time.Sleep(time.Microsecond)
		}
	}()
	select{}
}
//go run -race test.go 启用数据竞争检查之泪问题

//用sync.RWMutex实现同步，避免读写操作同时进行

func ex7(){
	var lock sync.RWMutex
	m := make(map[string]int)

	go func() {
		lock.Lock() //注意锁的粒度
		m["a"] +=1
		lock.Unlock() //不能使用defer
		time.Sleep(time.Microsecond)
	}()
	go func() {
		lock.RLock()
		_ = m["b"]
		lock.RUnlock()

		time.Sleep(time.Microsecond)
	}()
	select {}
}

//性能

//字典对象本身就是指针包装，传参时无须再次取地址
func test(x map[string]int){
	fmt.Printf("x: %p\n",x)
}
func ex8(){
	m := make(map[string]int)
	test(m)
	fmt.Printf("m: %p, %d\n",m,unsafe.Sizeof(m))

	m2 := map[string]int{}
	test(m2)
	fmt.Printf("m2: %p, %d\n",m2,unsafe.Sizeof(m2))
}

//在创建map时，预先准备足够空间有助于提升性能，减少扩张时的内存分配和重新哈希操作
//字典不会收缩，适当替换为新对象是必要的
func Ex9() map[int]int{
	m := make(map[int]int)
	for i:=0; i<1000;i++{
		m[i] = i
	}
	return m
}
func Ex10() map[int]int{
	m:=make(map[int]int,1000)
	for i:=0; i<1000;i++{
		m[i]=i
	}
	return m
}