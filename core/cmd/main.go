package main

import (
	"github.com/artking28/GList/core"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPeople(name string, age int) People {
	p := People{}
	p.Name = name
	p.Age = age
	return p
}

func main() {
	list := core.NewGList[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	//list := core.NewGList[People]()
	//list.Add(NewPeople("Arthur", 20))
	//list.Add(NewPeople("Isadora", 23))
	//list.Add(NewPeople("Eduarda", 20))
	//list.Add(NewPeople("Helena", 18))
	//println(r1.Stringify())

	println(list.Reverse().Stringify())
}
