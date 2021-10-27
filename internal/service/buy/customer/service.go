package customer

import (
	"fmt"
	"sort"

	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type DummyCustomerService struct {
	curId      uint64
	customers  map[uint64]buy.Customer
	listCached bool
	cache      []buy.Customer
}

func NewDummyCustomerService() *DummyCustomerService {
	s := &DummyCustomerService{}
	s.customers = make(map[uint64]buy.Customer)
	return s
}

func (s *DummyCustomerService) Describe(customerID uint64) (*buy.Customer, error) {
	customer, ok := s.customers[customerID]
	if !ok {
		return &buy.Customer{}, fmt.Errorf("no customer with id %v found", customerID)
	}

	return &customer, nil
}

func (s *DummyCustomerService) List(cursor uint64, limit uint64) ([]buy.Customer, error) {
	if cursor > uint64(len(s.customers)) {
		return nil, fmt.Errorf("index %v is out of range", cursor)
	}

	cache := s.GetCache()
	l := min(limit+cursor, uint64(len(cache)))
	return cache[cursor:l], nil
}

func (s *DummyCustomerService) Create(customer buy.Customer) (uint64, error) {
	id := s.curId
	s.curId++
	s.customers[id] = buy.Customer{
		Id:   id,
		Name: customer.Name,
		Age:  customer.Age,
	}

	s.listCached = false
	return id, nil
}

func (s *DummyCustomerService) Update(customerID uint64, customer buy.Customer) error {
	_, ok := s.customers[customerID]
	if !ok {
		return fmt.Errorf("no customer with id %v found", customerID)
	}

	s.customers[customerID] = buy.Customer{
		Id:   customerID,
		Name: customer.Name,
		Age:  customer.Age,
	}

	s.listCached = false
	return nil
}

func (s *DummyCustomerService) Remove(customerID uint64) (bool, error) {
	_, ok := s.customers[customerID]
	if !ok {
		return false, fmt.Errorf("no customer with id %v found", customerID)
	}

	delete(s.customers, customerID)

	s.listCached = false
	return true, nil
}

func (s *DummyCustomerService) GetCache() []buy.Customer {
	if s.listCached {
		return s.cache
	}

	len := len(s.customers)

	keys := make([]uint64, 0, len)
	for k := range s.customers {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	s.cache = make([]buy.Customer, 0, len)
	for _, k := range keys {
		s.cache = append(s.cache, s.customers[k])
	}

	s.listCached = true
	return s.cache
}

func min(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
