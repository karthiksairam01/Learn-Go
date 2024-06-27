# Functions

> Yay the interesting part!

## ```main()```

- Starting point of every Go program
- Has no params or return vals

## ```func()```

- Basically looks like this:

```go
func mul(num int, den int) int {

    if den == 0 || num ==0 {
        return 0
    }

    return num * den
}
```

- Has 4 parts:
  - keyword ```func```
  - name of function
  - input params with the names of vars first, then data type
    - leave empty paranthesis is no params, like ```main()```
    - if you have multiple params of same type, you can list only var names and then its type in the end like ```func mul(num, den int) int {}```
  - return type
    - mandatory if function returns value
    - leave blank if nothing is returned

### Parameters

- Go needs all params to a function when called, unlike Python where you have named and optional params
- This is actually not a limitation because Go believes that a function should not have more than a few params anyway, and you dont need that funcitonality if so.
- If you still want to implement named and optional params, pass a struct as a param

```go
type mys struct {

    Name string
    Age int
    Married bool
}

func myfunc(person mys) error {

    // do something
}

func main() {

    myfunc(mys {

        Name: "Bob"
        Married: true
    })
}
```

### Variadic params

- Basically variable # of params
- Always has to be last in the input param list, indicated by ```...``` before the type
- 

