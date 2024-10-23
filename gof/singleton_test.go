package gof

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once
var instance *singleton

type singleton struct{ value string }

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	fmt.Println(instance1.value)
	instance2 := GetInstance()
	fmt.Println(instance2.value)
	fmt.Println(instance1 == instance2)

}

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{value: "Instance created only once"}
	})
	return instance
}
