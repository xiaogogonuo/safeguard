package service

type IService interface {
	Start()
	Stop()
	Restart()
	Listen()
}
