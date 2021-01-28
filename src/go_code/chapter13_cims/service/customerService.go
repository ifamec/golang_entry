package service

import "go_code/chapter13_cims/model"

// add modify delete search
type CustomerService struct {
	customers []model.Customer
	// filed to represent slice length, used for new customer id
	customerNum int
}

func NewCustomerService() *CustomerService {
	// mode customer data
	customerService := &CustomerService{}
	customer := model.NewCustomer(1, "Tom", "M", 20, "000-0000-0000", "tom@gmail.com")
	// (*customerService).customerNum = 1
	// (*customerService).customers = append((*customerService).customers, customer)
	(*customerService).Add(customer)

	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return (*cs).customers
}
// use pointer
func (cs *CustomerService) Add(customer model.Customer) bool {
	// id = add order
	(*cs).customerNum ++
	customer.Id = (*cs).customerNum
	(*cs).customers = append((*cs).customers, customer)
	return true
}
func (cs *CustomerService) Delete(id int) bool {
	var index int = (*cs).FindById(id)
	if index == -1 {
		return false
	} else {
		// delete element in slice
		(*cs).customers = append((*cs).customers[:index], (*cs).customers[index+1:]...)
		return true
	}
}
func (cs *CustomerService) FindById(id int) (index int) {
	index = -1
	for i, v := range (*cs).customers {
		if id == v.Id {
			index = i
			return
		}
	}
	return
}