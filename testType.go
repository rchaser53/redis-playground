package main

import (
	abc "redisPlayground/PackageA"
)

type Nyan struct {
	myan int16
}

func (n Nyan) hoge() {
	// println(n.myan)
	println("naaaaaa")
}

func main() {
	nyan := &Nyan{1}
	nyan.hoge()

	// it's a cast in go
	// cannot using xi out of ok func
	if xi, ok := tetetete().(int); ok {
		println(xi)
	}

	println(tetetete().(int))

	client := abc.CreateRedisClient()
	abc.ExampleClient(client)
}

type a interface{}

func tetetete() a {
	// func tetetete() int8 {
	return 2
}
