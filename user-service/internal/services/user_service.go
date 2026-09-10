package services

type service struct {
	repository repository
}

func NewUserService(repo repository) *service {
	return &service{
		repository: repo,
	}
}
