# Learning Thoughts

> The word idiomatic keeps popping up, I don't know what that means in the current context

- vars have a default zero value that's only declared, every type has its own zero value

- literals can be 
    - ints, 
    - floats, 
    - rune (chars): only have to be in single quotes, can be used to represent ocatl, hex, UNicode, escape characters
    - string: zero or more rune literals, double quotes, cannot have double quotes, backslashes and newlines that are unescaped
    - literals are untyped in Go, meaning they can interact with compatible vars
    - there's a default type for literals

- built-in data types
    - ```bool```: ```true``` or ```false``` , defaults to ```false```
    - numeric: there are 12 types, some are:
        - int: 32b or 64b signed depending on processor
        - unit: unsigned unit
        - rune
        - uintptr
        - all int types default to 0
    - float: uses IEEE 754 standard, involving matissa, sign and exponent
        - dividing float by Zero gives ```+Inf``` or ```-Inf```
        - 0 by 0 gives ```NaN```
        dont compare floats using == and !=, because a seemingly harmless number is stored as its exact value in more expanded format in bits 
    - complex (learn later, if I even want to)
    - string:
        - zero val is empty str
        - concat using +
        - immutable, meaning can reassign but cannot change value of str once assigned to it


## Operators on int:

- same +, -, *, / and % for modulus
- dont divide by 0
- int/int is int truncated, use type conversion otherwise
- has shorthand +=, -= and so on
- also has Bitwise operators, << and >> 
- has logical operators too, &=, |= etc

## Explicit type conversion

- look at this following code, Go does not automatic type promotion

```
var x int = 10
var y float64 = 30.2
var z float64 = float64(x) + y
var d int = x + int(y)
fmt.Println(z, d)
```

 - cannot interpret a nonzero int or nonempty str as ```true``` because of no Implicit type conversion
 - in fact, can NEVER convert anything to ```bool```, it's shielded

## Variables

Many ways to declare vars:

- ```var x int = 10```
- ```var x = 10``` also works, since int literals are of type int by default
- ```var x, y int = 10, 20``` is valid, so is ```var x, y int```
- ```var x, y = 10, "hello"``` is valid too, different types works
- If you are declaring multiple vars, wrap them in a *declaration list*. All the below are valid declarations. Remember to tabspace after the variable name

```
var (
    x   int
    y       = 20
    z   int = 30
    d, e    int = 40, 50
    f, g    = 60, "hello"
    h, i string
)
```

### Short Declaration Format

- The short declaration format is using ```:=``` to replace ```var```. These are valid:
    - ```x := 10```
    - ```x, y := 10, "hello"```
- There's one extra thing that this can do, allows you to assign values to existing vars as long you assign value to one new variable with it such as:

```
x := 10
x, y := 30, "hello"
```

- Use ```:=``` only inside functions, its not legal to otherwise. But you dont need to declare vars outside functions mostly anyway.
- Use ```var``` when initializing a var to its zero value
- You might run into shadow variables using ```:=``` because it allows reassignment, hence better to use var.

## const 

- Like other languages, it is used for immutable value declaration
- But, it is very limited. It cannot store values that need to be stored at runtime
- Hence const is just to give names ot literals in Go, there's no way to declare that a var is immutable.

## Unused Variables

- It is a *compile time error* to not have every declared local var read, this applies to function-level vars only
- As long as a variable is read once, its okay, even if the var is redeclared a lot of times.

> Go uses Camel case for identifiers
> 
