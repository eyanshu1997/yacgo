# YACGO - Yet Another Compiler in GO

This is a Compiler written in go. 
First Creating a interpretter then compiler


## YAPL - Yet Another Programming Language


this is sudo languge that we will be writing the code for

### Sample Syntax
```
let age =1;
let name = "YAL"
let result = 10 * (20/2)
let myarray = [1,2,3,4,5];
let myhash= {"name":"yal","type":"lang"}
let add = fn(a,b) {return a+b};

```
For a full list [refer here](parser/parser.md#supported-syntax)

### Components
- lexer 
- parser
- ast
- symbol table
- evaluator



## Some Notes

- i wanted to undertsand how interpreters work, by creating one from scratch.
- only supports ASCII
- for supported operators [refer here](parser/parser.md#operators)