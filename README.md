# Jp
Jp is JSON(JavaScript Object Notation) parser.

## Syntax
EBNF(Extended Backus-Naur Form)

```
json = object | array .
object = "{" [ ( string ":" value ) { ","  string ":" value } ] "}" .
array = "[" [ value { "," value } ] "]" .
value = null | boolean | string | number | array | object .
null = "null" .
boolean = "true" | "false" .
string = { "a"..."z" | "A"..."Z" } .
number = [ "-" ] { "1"..."9" } "0"..."9" [ "." { "0"..."9" } "1"..."9" ] .
```
