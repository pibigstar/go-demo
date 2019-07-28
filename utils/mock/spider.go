package mock

type Spider interface {
	Init()
	GetBody() string
}

func GetGoVersion(s Spider) string {
	s.Init()
	body := s.GetBody()
	return body
}
