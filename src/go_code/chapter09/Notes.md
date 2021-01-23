# Chapter 09 Notes - Map

key-value data structure. like collection

`var mapName map[keytype]valuetype`

**key**

Could use `bool` `number` `string` `pointer` `channel`, ALSO `interface`, `struct`, `array`

BUT generally, will use `int`, `string` as the key

CANNOT use `slice` `map` `function` as key

**value**

same as keytype,

BUT generally, will use `number`, `string`, `map`, `struct` in value

**Declare**  
`var a map[string]string`  
`var a map[string]int`  
`var a map[int]string`  
`var a map[string]map[string]string`  
Declare will not allocate RAM, need use `make`, then can assign value and use

- need to `make` before use
    - could leave size in `make`
- each key is unique
- value could be identical
- key-value is unordered

## add delete modify search

`map["key"] = value`

- if key exist, -> modify
- if key not exist, -> add

`delete(map, key)`

- if key exist, -> delete
- if key not exist, -> no change
- to delete all key
    - iterate and delete
    - make a new map

`val, find := m[3]`

- find is a bool
- if find is true, val could be the value of the key
- if find is false, val could be the default value of the dataType

## traverse

use `for range` only

get length `len(map)`

## slice

if the datatype of the slice is map, then its called "slice of map". in this case, the amount of map is dynamic

## sort

- no builtin method to sort the map by key
- map is unordered by default
- sort by key
    - key to slice
    - sort slice
    - traverse slice and print value

## Details
