package pkg01

type GreetService interface {
	Greeting()
}

func NewGreetService() GreetService {
	return &greetServiceImpl{}
}
