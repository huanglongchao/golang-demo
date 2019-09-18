package main

import "fmt"

func main(){
	s1 := make([]int,5)
	fmt.Println("length of s1",len(s1))
	fmt.Println("cap of s1",cap(s1))
	fmt.Println("value of s1",s1)


	s2 := make([]int,5,8)
	fmt.Println("length of s2",len(s2))
	fmt.Println("cap of s2",cap(s2))
	fmt.Println("value of s2",s2)

	s3 := []int{}
	fmt.Println("length of s3",len(s3))
	fmt.Println("cap of s3",cap(s3))
	fmt.Println("value of s3",s3)

	s4 := []int{1,2,3,4,5}
	fmt.Println("length of s4",len(s4))
	fmt.Println("cap of s4",cap(s4))
	fmt.Println("value of s4",s4)

	s5 := s4[1:3]
	fmt.Println("length of s5",len(s5))
	fmt.Println("cap of s5",cap(s5))
	fmt.Println("value of s5",s5)

	//当我们通过切片表达式基于某个数组或切片生成新切片的时候，他们底层数组是同一个
	s4[1] = 100
	fmt.Println("value of s4",s4)
	fmt.Println("value of s5",s5)

	//在无需扩容时，append函数返回的是指向原底层数组的新切片...
	s4 = append(s4,200)
	s4[1] = 300
	fmt.Println("value of s4",s4)
	fmt.Println("value of s5",s5)

}
