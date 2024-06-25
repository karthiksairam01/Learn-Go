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

