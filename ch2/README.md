# Learning Thoughts

- The word idiomatic keeps popping up, I don't know what that means in the current context

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


