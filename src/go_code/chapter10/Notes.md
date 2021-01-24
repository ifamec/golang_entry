# Chapter 10 Notes - Object Oriented

1. Golang supports OOP feature
2. Golang uses `struct` to realise OOP feature
3. Golang has no inherit (extends), overload, constructor，destructor, hidden this pointer etc.
4. have inheritance, encapsulation, polymorphism feature
5. Interface Oriented Programming

USE

1. extract feature of certain category of objects, to create a new custom data type ->  struct
2. create objects with the struct -> represents an object

![struct_ram](img/struct_ram.png)

```go
type structName struct {
field1 type
field2 type
}
```

## FIELD:

- field == attributes
- field could be basic data type or reference type

### !

1. field declare: `fieldName type`
2. type of field could be basic/array/reference type
3. will use default value if not assigned after declaration
    - reference type default is `nil`
4. different struct object is identical, will not affect each other
    - struct is **VALUE** Type
    - pass by value
      ![struct_copy](img/struct_copy.png)

## Declare

1. `var cat Cat`
2. `var cat Cat = Cat{"Oolong", 4, "Orange"}`
3. `var cat *Cat = new(Cat)`
     ```go
      var cat3 *Cat = new(Cat)
      (*cat3).name = "C3"
      cat3.age = 4 // Simplified
      (*cat3).color = "Ame"
      cat3.hobby = "Climb" // Simplified
      ```
4. `var cat *Cat = &Cat{}`

## allocate ram

- method 3, 4 will return struct pointer
- to assign and get value, use (*object).field
- could also use object.field // golang do transfer while compiling

![struct_pointer_ram](img/struct_pointer_ram.png)

!! `(*cat2.name)` // ERROR "." operator is higher than "*"

### Detail

- all fields in a struct is continuous in RAM
    - ![struct_pointer_ram_2](img/struct_pointer_ram_2.png)
- struct is a custom type, if need to convert to other type, **WIHCH REQUIRES EXACT SAME FIELDS (amount, name, type)**
- rename the struct, Golang will treat that as new type, but can force convert
    - [Check Code](struct06/main.go)
- each field could add a `tag`, and could get with reflect mechanism
    - serialization(to JSON) & deserialization
    - [Check Code](struct07_tag/main.go)