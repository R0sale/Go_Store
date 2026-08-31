package service

type service struct {
	repository repo
}

func NewService(rep repo) *service {
	return &service{
		repository: rep,
	}
}
