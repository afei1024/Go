package main

import (
	"encapsulation_test/model"
	"fmt"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(*p) // 输出原值
	fmt.Println(p.Name, " age =", p.GetAge(), " sal =", p.GetSal())

}
