# JIPL : Just an Interpreted Programming Language

## What is JIPL?

JIPL is an interpreted programming language. It is written in Go and is very easy to use.
It is also very easy to learn. It is a very simple language and is not very powerful. It is not meant to be used for large projects

## How to use JIPL (With go installed)?

1. Clone the repo

   ````
   git clone```

   ````

2. cd to the repo

   ````
   cd JIPL```

   ````

3. Run the main.go file

   ````
    go run ./cmd/main.go```

   ````

4. Now you can use JIPL in the terminal

# JIPL Documentation

1. Variables

   1. supported data types
      1. integers
      2. booleans
      3. undefined
   2. defining variables
      1. integers
         1. syntax
            1. `def <variable name> = <value>`
         2. example
            1. `def a = 10`
      2. booleans
         1. syntax
            1. `def <variable name> = <value>`
         2. example
            1. `def a = true`

2. Functions

   1. syntax
      1. `function function_name(arguments) { function_body ;})`
   2. example
      1. `function add(a,b) { return a+b ;})`
   3. calling functions
      1. syntax
         1. `<function_name>(arguments);`
      2. example
         1. `add(10,20);`

3. If statements

   1. syntax
      1. `if (condition) { body ;} else { else_body ;}`
   2. example
      1. `if (a == 10) { return true ;} else { return false ;}`

4. Loops
   1. for loops
      1. syntax
         1. `for (initialization; condition; increment) { body ;}`
      2. example
         1. `for (def i = 0; i <= 10; i++) { if(i==3){return 3;}}`
5.
