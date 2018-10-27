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

```
let <identifier> = <expression>;
```

## Parsing let statement
**Expression produces value, Statement doesn't.**
```
let x = 3;
        <statement>
        /         \
  <Identifier> = <Expression>
```

```
function parseProgram() {
  program = new ASTRoot()
  advanceToken()

  for (currentToken != EOF) {
    statement = null

    if (currentToken() == LET) {
      statement = parseLetStatement()
    } else if (currentToken() == RETURN) {
      statement = parseReturnStatement()
    } else if (currentToken() == IF) {
      statement = parseIfStatement()
    }
    if (statement != null) {
      program.Statements.push(statement)
    } 
    advanceToken()
  }
}

function parseLetStatement() {
  advanceToken()
  identifier = parseIdentifier()

  advanceToken()

  if (currentToken() != EQUAL) {
    throw new Error("Missing = sign")
  }
  advanceToken()
  value = parseExpression()
  variableStatement = new variableStatementASTNode()
  variableStatement.identifier = identifier
  variableStatement.value = value
  return variableStatement
}

function parseIdentifier() {
  identifier = new IdentifierASTNode()
  identifier.token = currentToken()
  return identifier
}

function parseExpression() {
  if (currentToken() == INTEGER) {
    if (nextToken() == PLUS) {
      return parseOperatorExpression()
    } else if (nextToken() == SEMICOLON) {
      return parseIntegerLiteral()
    }
  } else if (currentToken() == LEFT) {
    return parseGroupedExpression()
  }
}

function parseOperatorExpression() {
  operatorExpression = new OperatorExpression()

  operatorExpression.left = parseIntegerLiteral()
  operatorExpression.operator = currentToken()
  operatorExpression.right = parseExpression()

  return operatorExpression()
}

```