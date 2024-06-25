# Control Structures

## Blocks

- It is any place where declaration occurs
- You can access any identifier in an outer block, in its any inner block

### Shadowing

- Happens when you declare same identifier both inside and outside a block. 

```go
func main(){

    x := 10

    if x > 5 {

        fmt.Println(x)
        x := 5
        fmt.Println(x)
    }

    fmt.Println(x)
}

    //output is: 10 5 10

```

- Such a var is called a shadowing var

- The ```x``` defined inside the ```if``` block is a shadowing variable, and ***shadowing lasts only until that block lasts***.

- There is **no way** to access the outer variable inside the inner block as long as shadowing variable is declared.

- Be careful to not shadow a package import, like declaring a var called main or fmt.

- You can use the shadow linter to detect shadowing vars.

### Universal Block

- Instead of defining int, nil, string as keywords, Go declares them as pre-defined identifiers in the universal block, which is the block that contains all other blocks.

- Hence you can shadow these too, like ```true := 10``` etc.

- Don't do it 👌


## If statement

- Same as other langs like Python, but no paranthesis around the conditions

- All vars are local to not only the if block, but the whole if and else blocks. (yes, if and else are 2 separate blocks in Go)

- Vars are local to also the condition, such as below

```go
if n := rand.Intn(10); n==0 {

    fmt.PrintLn("too low")
} else if n > 5 {

    fmt.PrintLn("too big")
} else {

    fmt.PrintLn("perfect")
}

```

- You can't access ```n``` outside the if else blocks, throws compilation error


## For statement

- Used for looping

- The *only* looping construct in Go, used in 4 ways:

### 1. Complete style

- Same as C, has 3 parts:

    - *initialization*: you **have** to use ```:=```, ```var``` is illegal. shadowing is allowed
    - *comparison*: must evaluate to a ```bool```
    - *increment*: any assignment, not just ```+=```, works here

```go
for i := 0; i < 10; i++ {

    fmt.PrintLn(i)
}
```

### 2. Condition-only style

- Acts like a while loop in Python

```go

for x < 10 {

    fmt.PrintLn(i)
    i = i * 2
}
```

### 3. Infinite style

- No condition too

- Just a simple infinite loop

```go
for {

    fmt.PrintLn("Hello")
}
```

### For-range style

- Iterates over some composite types like slices, maps etc.

- It actually gives 2 loop variables (interesting!) unlike Python:

    - **position** in the data structure being iterated
    - **value** at that position

```go
evenVals := []int{2, 4, 6, 8, 10, 12}

for i, v := range evenVals {

    fmt.Println(i, v)
}
```

- If you *dont* want to use i, which is the position variable (remember that you **HAVE** to use all declared variables in Go), then just use an underscore in place of the position variable declaration.

```go
evenVals := []int{2, 4, 6, 8, 10, 12}

for _, v := range evenVals {

    fmt.Println(v)
}
```

- It is also valid to completely leave off the value, but why would you want to? (unless you're using a map as a set, and then you only want the keys)


> Iterating over maps with for-range is interesting, because the order of a for-range iteration varies a bit every time you loop over it.


```go
var m map[string]int = {

    "a" : 1,
    "b" : 2,
    "c" : 3,
}

for k, v := range m {

    fmt,Println(k ,v)
}
```

- You can use this style to iterate over strings too, but remember that it iterates over runes and not bytes.

- Also remember that the variables you use for accessing the values in a for loop, like ```i```m\, ```k``` or ```v``` make a copy of the data structure you send in them, and hence modifying these values **do not** change the values in the original data structure.

### Break and continue statements

- The Go version of a do-while loop is an infinite loop with a break statement inside an if condition.

- Using continue makes your code more idiomatic because Go encourages short ```if``` bodies, as left-aligned as possible (means not too much nesting)

### Labeling for statements 

