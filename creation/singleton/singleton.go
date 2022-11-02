package main

import (
	"log"
	"reflect"
)

func main() {
	obj1 := GetXcInstance()
	obj2 := GetXcInstance()
	log.Println(reflect.DeepEqual(obj2, obj1))

	obj3 := GetLhsInstance()
	obj4 := GetLhsInstance()
	log.Println(reflect.DeepEqual(obj3, obj4))
}
