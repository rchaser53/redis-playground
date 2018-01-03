package main

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
}
