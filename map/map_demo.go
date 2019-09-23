package main

import "fmt"

func main(){
	var capitalCountyMap map[string]string
	capitalCountyMap = make(map[string]string)
	capitalCountyMap["中国"] = "北京"
	capitalCountyMap["日本"] = "东京"
	capitalCountyMap["法国"] = "巴黎"
	capitalCountyMap["德国"] = "柏林"
	fmt.Println(capitalCountyMap)
	delete(capitalCountyMap,"日本")
	fmt.Println(capitalCountyMap)
}
