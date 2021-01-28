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
	(*customerService).customerNum = 1
	(*customerService).customers = append((*customerService).customers, customer)

	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return (*cs).customers
}