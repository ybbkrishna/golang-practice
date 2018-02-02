package serv4

import (
	"fmt"
	"sync"
	"time"
)

type Service struct{}

var serv *Service
var once sync.Once

func GetService() *Service {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("in serv 3")
	once.Do(func() {
		fmt.Println("in serv 3 waiting")
		time.Sleep(5 * time.Second)
		serv = &Service{}
	})
	fmt.Println("in end serv 3")
	return serv
}
