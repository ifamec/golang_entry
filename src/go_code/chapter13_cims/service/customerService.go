package service

import "go_code/chapter13_cims/model"

// add modify delete search
type CustomerService struct {
	customers []model.Customer
	// filed to represent slice length, used for new customer id
	customerNum int
}