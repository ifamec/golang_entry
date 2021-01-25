# Chapter 11 Notes - OOP 2

## abstract

- extract common attributes and behavior from one kind of things to build a physic model
- attributes - fields
- behaviors - methods

## 3 features

inheritance, encapsulation, polymorphism

### encapsulation

encapsulate fields and methods together, and protect them in the package, use authorised method to handle the fields

- Pros
    - hide realising details
    - validate the data to ensure the data is reasonable
- Where reflect the encapsulation
    - encapsulate the fields in struct
    - use method, package

#### realise

- lowercase the struct fields the first letter (private like) BUT could access in current package
- provide a factory mode function (start with uppercase)
- provide a `Set` function to assign and estimate to a field
- provide a `Get` function to get the field value
- (Simplified Encapsulation)

### inheritance

use inheritance if two struct are almost alike

- less code redundancy [Check code](inheritance01/main.go)
- more maintainable
- more extendable

inheritance is use for

- solving code redundancy
- extract common fields and methods into another struct, inherit the common struct to get common fields and method. and
  add more unique things
- use "anonymous struct" to realise that

```go
type Goods struct{
    Name string
    Price float64
}
type Books struct{
    Goods
    Writer string
}
```

[Check code](inheritance02/main.go)

#### Details


- struct could nest all fields and methods to realise inheritance
- simplify visiting fields & methods
- when struct and anonymous struct have same naming fields and method
    - USE THE CLOSEST ONE
    - if need to use fields and methods in the anonymous struct, use the anonymous struct to approach that
    - [CHECK CODE](inheritance03/main.go)
- if embed two or more anonymous struct, if both of them contains same name fields/methods
    - NEED TO IDENTIFY THE ANONYMOUS STRUCT NAME TO USE  
    - [CHECK CODE](inheritance04/main.go)
- if a struct embed a non-anonymous struct, THIS CALLED COMBINATION
    - NEED TO IDENTIFY THE STRUCT NAME TO USE FIELDS  
    - [CHECK CODE](inheritance04/main.go)
- assign anonymous struct value when creating variables with embedded struct
    - [CHECK CODE](inheritance04/main.go)



If a struct nest multiple anonymous struct, this struct could visit nested fields and methods in anonymous struct 
- [CHECK CODE](inheritance04/main.go) Goods Brand Example
- if nested anonymous structs have same fields name or methods name, NEED TO IDENTIFY THE STRUCT NAME TO USE FIELDS AND METHODS
- to keep the code simple, do not use multiple inheritance
    ```go
	 var tv = &TV{ Goods{"Television", 1800.00}, Brand{"Sony", "Tokyo Japan"}}
	 fmt.Println((*tv).Goods.Name, (*tv).Brand.Name, (*tv).Price)
    ```

### Interface

polymorphism is realised by interface
[Interface](interface01/main.go)

Interface could declare a group of methods, but do not need to realise. Interface does not contain any variable.
When some custom type using that, the custom struct will realise those methods

```go
// Interface
type interfaceName interface{
    method1(params) (returnvalue)
    method2(params) (returnvalue)
}

// to realise the interface 
func (customVariable, customType) method1(params) (returnvalue) {
    // realise
}
func (customVariable, customType) method2(params) (returnvalue) {
    // realise
}
```

to realise the interface, need to realise all methods in interface
- all methods in an interface does not have a body 
    - represents "polymorphism and High cohesion, low coupling"
- interface in golang does not require a visible identifier
    - if one variable, contains all methods in the interface
    - then we say the variable realise the interface

#### Details
1. Interface cannot create a variable, but it could point to a variable that created by a custom type which realise the interface   
   - only if a custom type realise an interface, then the variable of the custom type could be assigned to interface type variable
2. Interface does not have body, which means no method realisation
3. A custom type that realise all method in an interface, we say that custom type realise the interface
4. Custom type could realise the interface, not only struct
5. Custom type could realise multiple interface
6. NO VARIABLES IN INTERFACE
7. An interface "A" could inherit multiple other interface ("B", "C"), you need to realise all methods in "B", "C" to realise the interface "A"
8. interface is a pointer pass by reference, if no init, will return nil
9. empty interface does not have any method - all type realise empty interface we could assign any type to empty interface

Interface inherit - if inherit two same name method // Error [Duplicate Define Demo](interface04/main.go)  
**?? Might be an update in later version, No error in 1.15.6**


### Interface vs Inheritance
[Interface vs Inheritance Example](interface_vs_inheritance/main.go)
- When struct B inherit A, B inherit all fields and methods, could use directly
- B want to have new methods but don't want to break the inheritance, use interface
- Interface is a supplement of inheritance

- Interface and Inheritance solving different issues
    - inheritance: code re-usability and maintainability
    - interface: design, design rules, have custom type to realise methods
- Interface is more flexible than Inheritance
- Interface could realise low coupling in some aspect

### polymorphism








