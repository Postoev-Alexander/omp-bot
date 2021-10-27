package buy

import (
	"fmt"
	//	"time"
)

type Customer struct {
	Id   uint64
	Name string
	//Date      time.Time
	//AddressId uint64
	Age uint32
}

func (o Customer) String() string {
	return fmt.Sprintf(`Customer{id: %v, name: %v,  age: %v}`,
		o.Id, o.Name, o.Age)
}
