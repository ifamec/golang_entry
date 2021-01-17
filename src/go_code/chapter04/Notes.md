# Chapter 04 Notes - Operator

- arithmetic
- assignment
- comparison
- logical
- binary
- other & *

## arithmetic
```
+ -
+ - * /
% ++ --
```

### Details
1. "/" notice data type, int/int -> keep int part only 
2. `a % b = a - a / b * b`
3. "++" "--" could just use by itself, cannot assign after ++/--
4. NO "++a" "--a" (NICE!)

## comparison

- return value is either `true` or `false`
- use in `if` / `for` statements

```
== !=
> < 
>= <=
```

### Details

1. value is `bool`
2. called relational expression
3. "==" is comparison operator, "=" is assignment operator

## logical

```
&&, ||, !
```

### Details
1. Statement A && Statement B - if A is false, B will not be executed
2. Statement A || Statement B - if A is true,  B will not be executed

## assignment

```
=
+= -= *= /= %=
"a ()= b" == "a = a () b"

binary Check Other Section
<<= >>=
&= ^= !=
```

- calculation order: right side to left side
- left side: variable, right: variable or const or statement
- "a ()= b" == "a = a () b"

## Precedence
- unary operation, assignment operation: r->l
- other: l->r

| Category       | Operator                          | Associativity | 
|----------------|-----------------------------------|---------------| 
| Postfix        | () [] -> . ++ - -                 | Left to right | 
| Unary          | + - ! ~ ++ - - (type)* & sizeof   | Right to left | 
| Multiplicative | * / %                             | Left to right | 
| Additive       | + -                               | Left to right | 
| Shift          | << >>                             | Left to right | 
| Relational     | < <= > >=                         | Left to right | 
| Equality       | == !=                             | Left to right | 
| Bitwise AND    | &                                 | Left to right | 
| Bitwise XOR    | ^                                 | Left to right | 
| Bitwise OR     | \|                                | Left to right | 
| Logical AND    | &&                                | Left to right | 
| Logical OR     | \|\|                              | Left to right | 
| Assignment     | = += -= *= /= %=>>= <<= &= ^= \|= | Right to left | 
| Comma          | ,                                 | Left to right | 

## bitwise

Check Binary section

## Other

`&` - get address  
`*` - get value

**NO TERNARY OPERATOR IN GOLANG**

## Get User Input

`fmt.Scanln` - get one line input  
`fmt.Scanf` - get value per format

## Binary

- base-2 
- base-8 (077) 
- base-10 
- base-16(0x7FFF/0X7FFF)

## bitwise 

```
& | ^ << >>
```

Sign-Magnitude
one's Complement
two's Complement

1. first bit of signed int: 0 +, 1 -
2. positive int - sign-magnitude == one's complement == two's complement
3. negative int
    - one's complement = no change on sign bit (1st bit), flip other bits
    - two's complement = one's complement + 1
4. 0's one's and two's complement are 0
5. all calculation uses two's complement

`&` both 1   -> 1  
`|` have 1   -> 1  
`^` not same -> 1 

```
two's complement
2   0000 0010
3   0000 0011

2&3 0000 0010 = 2
2|3 0000 0011 = 3
2^3 0000 0001 = 1

-2  1000 0010 - sign magnitute
    1111 1101 - one's

two's complement
-2  1111 1110 - two's
2   0000 0010
-2^2:
    1111 1100 - two's
    1111 1011 - one's
    1000 0100 = -4

```


`>>` low bit overflow, sign no change, fill overflowed high bit with sign bit  
`<<` sign no change, fill low bit with 0

```
1>>2
0000 0001 - two's
0000 0000 = 0

1<<2
0000 0001 - two's
0000 0100 = 4


```