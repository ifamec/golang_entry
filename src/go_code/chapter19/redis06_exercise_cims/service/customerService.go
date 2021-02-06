package service

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_code/chapter19/redis06_exercise_cims/model"
)

// add modify delete search
type CustomerService struct {
	customers []model.Customer
	// filed to represent slice length, used for new customer id
	customerNum int
}

func NewCustomerService() *CustomerService {
	// mode customer data
	customerService := &CustomerService{}
	// customer := model.NewCustomer(1, "Tom", "M", 20, "000-0000-0000", "tom@gmail.com")
	// (*customerService).customerNum = 1
	// (*customerService).customers = append((*customerService).customers, customer)
	// (*customerService).Add(customer)

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
func (cs *CustomerService) Modify(customer model.Customer) bool {
	var index int = (*cs).FindById(customer.Id)
	if index == -1 {
		return false
	} else {
		// modify element in slice
		(*cs).customers[index] = customer
		return true
	}
}
func (cs *CustomerService) SaveData() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Redis Conn Error", err)
	}
	defer conn.Close()

	data, err := json.Marshal((*cs).customers)
	if err != nil {
		fmt.Println("Serialization Error", err)
	}
	_, err = conn.Do("SET", "customerNumber", (*cs).customerNum)
	if err != nil {
		fmt.Println("SET Error:", err)
	}
	_, err = conn.Do("SET", "customers", data)
	if err != nil {
		fmt.Println("SET Error:", err)
	}
	fmt.Println("Save Success")
}
func (cs *CustomerService) RecoverData() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Redis Conn Error", err)
	}
	defer conn.Close()

	customers, err := redis.Bytes(conn.Do("GET", "customers"))
	if err != nil {
		fmt.Println("Get Customer Error:", err)
	}
	err = json.Unmarshal(customers, &(*cs).customers)
	if err != nil {
		fmt.Println("Recover Customer Error:", err)
	}

	customerNumber, err := redis.Int(conn.Do("GET", "customerNumber"))
	if err != nil {
		fmt.Println("Get CustomerNumber Error:", err)
	}
	(*cs).customerNum = customerNumber
	fmt.Println("Recover Customer Success")
}