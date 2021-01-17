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
