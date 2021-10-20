package customer

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Customer {
	return allEntities
}

func (s *Service) Get(idx int) (*Customer, error) {
	return &allEntities[idx], nil
}
