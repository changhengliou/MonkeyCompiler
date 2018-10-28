## Monkey lang Compiler
### follow the design from the book `Writing an interpreter in Go`

```
let a = 1;
let b = "text";
let c = 3 * (2 + 1);
let arr = [ 1, 2, 3 ];
let obj = { w: 1, x: 2 };
arr[0] // = 1
obj["w"] = 1;
let add = func(x, y) { x + y; };
add(1, 2);

let twice = func(f, y) {
    reutrn f(f(x));
}
```

KEYWORD, IDENTIFIER, NUMBER, STRING, OPERATOR

IDENTIFIER: must start with `[A-Za-z_] or [0-9]`
STRINGLITERAL: start with `"` and end with `"` 

```
let x = 5;
let y = 10;
let foo = add(1,2);
let bar = 1 + (2 + 3) * 4 + add(2, 3)
let foobar = func(x) { x + 3; }
```

### Parsing let statement
```
let <identifier> = <expression>;
```

### Parsing return statement
```
return <expression>;
``` 

### Parsing expression

**Example of expression type**
```$xslt
-1
!true
-5 - 3
2 + 3 * 4 - 5
5 * (5 + 5)
add(2,3)
add(2, add(3, add(4, 5)))
foo * bar + foo
add(2, foo) + bar - 6
fn() {return 3}()
x = if (2 > 3) return 4
x
```
**Prefix, infix & postfix operator**
```$xslt
--3   // prefix
3++   // postfix
2 + 3 // infix
```

**Prefix operator**
```$xslt
-5
!foo
6 + -3
<prefix op><expression>
```
**Infix operator**
```$xslt
<expression><infix op><expression>
```
**Expression produces value, Statement doesn't.**
```
let x = 3;
        <statement>
        /         \
  <Identifier> = <Expression>
```

## Simple recursive descent parsing
> pseudo-code for number `[0-9]` and  operator `+ - * /` only
```$xslt
3 + 2 * 2 * 1 * 6 - 5
```
```$xslt
const parseNumber = (str) => number;
const parseProductAndDivide = function() {
    let num1 = parseNumber()
    while (currPtr == '*' || currPtr == '/') {
        const num2 = parseNumber()
        if (currPtr == '*') {
            num1 *= num2
        } else {
            num1 /= num2
        }
    }
    return num1
}
const parsePlusAndMinus = function() {
    let num1 = parseProductAndDivide()
    
    while (currPtr == '+') {
        const num2 = parseNumber()
        num1 += num2
    }   
    return num1 
}

const evaluate = () => {
    return parsePlusAndMinus()
}
```
