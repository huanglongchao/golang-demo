package main

import "fmt"

func main(){
	go func() {
		for i:=0; i<20;i++{
			fmt.Println(i)
		}
	}()
	select {
	}
}
