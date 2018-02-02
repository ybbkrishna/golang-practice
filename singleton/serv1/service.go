package serv1

import (
	"sync"
	"time"
)

type Service struct{}

var serv *Service
var once sync.Once

func GetService() *Service {
	once.Do(func() {
		time.Sleep(time.Second)
		serv = &Service{}
	})
	return serv
}
