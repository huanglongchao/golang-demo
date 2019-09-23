package main

import "fmt"

func main(){
	m := map[string]int{
		"abc": 123,
	}

	key := []byte("abc")
	x,ok := m[string(key)]

	println(x,ok)

	s := "雨痕"
	for i:=0; i < len(s); i++ {	//byte
		fmt.Printf("%d: [%c]\n",i , s[i])
	}
	for i,c := range s{			//rune
		fmt.Printf("%d: [%c]\n", i, c)
	}
}
