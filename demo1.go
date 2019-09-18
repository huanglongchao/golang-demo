package main


import "reflect"

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("panic")
		}
	}()
	fv := reflect.ValueOf(func(int) struct{} { return struct{}{} })
	args := []reflect.Value{reflect.ValueOf(0)}

	for {
		fv.Call(args)
	}
}
