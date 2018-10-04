## Monkey lang Compiler
### follow the design from the book `Writing an interpreter in Go`

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

KEYWORD, IDENTIFIER, NUMBER, STRING, OPERATOR

IDENTIFIER: must start with [A-Za-z_] + [0-9]
STRINGLITERAL: start with " end with " 