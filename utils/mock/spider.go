package mock

//go:generate mockgen -destination mock_spider.go -package mock go-demo/mock Spider

type Spider interface {
	Init()
	GetBody() string
}

func GetGoVersion(s Spider) string {
	s.Init()
	body := s.GetBody()
	return body
}
