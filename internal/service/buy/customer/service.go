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

/*
package {subdomain}

import "github.com/ozonmp/omp-bot/internal/model/{domain}"

type {Subdomain}Service interface {
  Describe({subdomain}ID uint64) (*{domain}.{Subdomain}, error)
  List(cursor uint64, limit uint64) ([]{domain}.{Subdomain}, error)
  Create({domain}.{Subdomain}) (uint64, error)
  Update({subdomain}ID uint64, {subdomain} {domain}.{Subdomain}) error
  Remove({subdomain}ID uint64) (bool, error)
}

type Dummy{Subdomain}Service struct {}

func NewDummy{Subdomain}Service() *Dummy{Subdomain}Service {
  return &Dummy{Subdomain}Service{}
}
*/
