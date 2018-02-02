package serv3

import (
	"fmt"
	"sync"
	"time"

	lol "github.com/ybbkrishna/golang-practice/singleton/serv3/ser4"
)

type Service struct {
	S *lol.Service
}

var serv *Service
var once sync.Once

func GetService() *Service {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// fmt.Println("in serv 3")
	once.Do(func() {
		fmt.Println("in serv 3 waiting")
		time.Sleep(5 * time.Second)
		serv = &Service{
			S: lol.GetService(),
		}
	})
	fmt.Println("in end serv 3")
	return serv
}
