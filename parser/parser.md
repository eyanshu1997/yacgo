## parser
this package is the parser that reads each statements and transfers them into ast 

### Supported syntax
#### let statments
```let a =5```
```let a =<expression>```
```let a =fn(){function defination}```

#### return statements
```return a```
```return expr```

#### expressions
they can be combination of any operations defined under
we will also consider functiondefinations as expressions

### operators

#### prefix operators
-5
!true
!false

#### binary operators

5 + 5
5 - 5
5 / 5
5 * 5

#### airthamatic operators

foo == bar
foo != bar
foo < bar
foo > bar


#### parenthesis
5 * (5 + 5)
((5 + 5) * 5) * 5

#### call expressions
add(2, 3)
add(add(2, 3), add(5, 10))
max(5, add(5, (5 * 5)))


### function expressions
fn(x, y) { return x + y }(5, 5)
(fn(x) { return x }(5) + 10 ) * 10

### if expressions
let result = if (10 > 5) { true } else { false };
result // => true