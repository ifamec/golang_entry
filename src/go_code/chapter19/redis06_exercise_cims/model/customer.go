package model

import "fmt"

// declare a customer struct to represent a customer

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string, ) Customer {
	return Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}
func NewCustomerExceptId(name string, gender string,
	age int, phone string, email string, ) Customer {
	return Customer{
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

func (c *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", (*c).Id, (*c).Name, (*c).Gender, (*c).Age, (*c).Phone, (*c).Email)
}