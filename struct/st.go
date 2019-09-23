package main

import (
	"fmt"
	"unsafe"
)

type point struct {
	x,y int
}

type value struct {
	id int	//基本类型
	name string	//字符串类型
	data []byte //引用类型
	next *value //指针类型
	point	//匿名字段
}

func main(){
	v := value{
		id:1,
		name:"test",
		data:[]byte{1,2,3,4},
		point:point{x:100,y:200},
	}
	s :=`
		v: %p ~ %x, size: %d, align: %d
		field	address	offset	size
		-------|-------|-------|------
		id		%p		%d		%d
		name 	%p		%d		%d
		data	%p		%d		%d
		next 	%p		%d		%d
		x		%p		%d		%d
		y		%p		%d		%d
		`
	fmt.Printf(s,&v,uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v),unsafe.Sizeof(v),unsafe.Alignof(v),
		&v.id,unsafe.Offsetof(v.id),unsafe.Sizeof(v.id),
		&v.name,unsafe.Offsetof(v.name),unsafe.Sizeof(v.name),
		&v.data,unsafe.Offsetof(v.data),unsafe.Sizeof(v.data),
		&v.next,unsafe.Offsetof(v.next),unsafe.Sizeof(v.next),
		&v.x,unsafe.Offsetof(v.x),unsafe.Sizeof(v.x),
		&v.x,unsafe.Offsetof(v.y),unsafe.Sizeof(v.y))

	v1 := struct{
		a byte
		b byte
		c int32
	}{}
	v2 := struct{
		a byte
		b byte
	}{}
	v3 := struct{
		a byte
		b []int
		c byte
	}{}
	v4 := struct{
		a byte
		c byte
		b []int
	}{}
	fmt.Printf("\nv1:%d,%d\n",unsafe.Alignof(v1),unsafe.Sizeof(v1))
	fmt.Printf("v2:%d,%d\n",unsafe.Alignof(v2),unsafe.Sizeof(v2))
	fmt.Printf("v3:%d,%d\n",unsafe.Alignof(v3),unsafe.Sizeof(v3))
	fmt.Printf("v3:%d,%d\n",unsafe.Alignof(v4),unsafe.Sizeof(v4))

	v5 := struct{
		a struct{}
		b int
		c struct{}
	}{}
	s1 := `
			v: %p ~ %x, size: %d, align: %d
			field	address		offset		size
			-------|----------|---------|---------
			a		%p			%d			%d
			b		%p			%d			%d
			c		%p			%d			%d
		`
	fmt.Printf(s1,&v5,uintptr(unsafe.Pointer(&v5))+unsafe.Sizeof(v5),unsafe.Sizeof(v5),unsafe.Alignof(v5),
				&v5.a, unsafe.Offsetof(v5.a),unsafe.Sizeof(v5.a),
				&v5.b, unsafe.Offsetof(v5.b),unsafe.Sizeof(v5.b),
				&v5.c, unsafe.Offsetof(v5.c),unsafe.Sizeof(v5.c),
				)

	v6 := struct {
		a struct{}
	}{}
	fmt.Printf("\n%p, %d, %d\n",&v, unsafe.Sizeof(v6), unsafe.Alignof(v6))
}
