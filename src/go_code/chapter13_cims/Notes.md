# Chapter 13 Notes - CIMS

Design Docs:  
charts, tables, db, modules etc.

Program Frame Diagram
- analyse how many files
- relationship between each file
- design structure 

customerView.go - UI V
- show UI
- accept input
- action per user requirements
    - call customerService methods to add del mod shw
- WILL CONTAINS CUSTOMERSERVICE

customerService.go - logic C
- modify
- delete
- show
- WILL DECLARE A CUSTOMER SLICE - including multiple customer variables

customer.go - M
- represent a customer
- fields of a customer

WHEN CODING< REVERSE THE ANALYSE


M (customer)
- Customer related fields:
```go
type Customer struct {
    Id int
    Name string
    Gender string
    Age int
    Phone string
}
```
0. add GetInfo method
2. add NewCustomerExceptId function with no id


C
0. NewCustomerService, return customerService variable
1. List method
    - customer list slice
2. Add method
    - append new customer to customers slice 
    - id will use the add order
3. Delete(id int) method
    - FindById(id int) index of the object in slice
4. Exit

V
1. show Lists
    - call customerService.List, show data
2. add method
    - call customerService.Add
3. delete method
    - call customerService.Delete
4. X
