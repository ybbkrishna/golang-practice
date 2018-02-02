package serv2

import (
	"sync"
	"time"

	"github.com/ybbkrishna/golang-practice/singleton/serv1"
)

type Service struct {
	S *serv1.Service
}

var serv *Service
var once sync.Once

func GetService() *Service {
	once.Do(func() {
		time.Sleep(time.Second)
		serv = &Service{
			S: serv1.GetService(),
		}
	})
	return serv
}
