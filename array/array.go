package main

import "fmt"

func main(){

	var a [4]int //元素自动初始化为0
	b := [4]int{2,5} //为提供初始值的元素自动初始化为0
	c := [4]int{5,3:10} //指定索引初始化
	d := [...]int{1,2,3}
	e := [...]int{10,3:100}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	type user struct{
		name string
		age byte
	}
	f := [...]user{
		{"Tom",20},
		{"Marry",18},
	}
	fmt.Printf("%#v\n",f)

	x,y := 10,20
	w := [...]*int{&x,&y} //元素为指针的指针数组
	p := &w //存储数组地址的指针

	fmt.Printf("%T, %v\n",w,w)
	fmt.Printf("%T, %v\n",p,p)


	v := [2]int{10,20}
	var z [2]int
	z = v
	fmt.Printf("a: %p, %v\n",&v,v)
	fmt.Printf("a: %p, %v\n",&z,z)
	test(v)

}

func test(x [2]int){
	fmt.Printf("x: %p, %v\n",&x,x)
}