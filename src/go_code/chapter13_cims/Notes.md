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