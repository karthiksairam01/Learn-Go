# Composite Types

## Arrays

- rarely used directly (apparently)
- ```var x [3]int``` creates array of 3 ints, all init to 0
- ```var x = [3]int{10, 20, 30}``` is an array literal
- If you ahve a sparse array and only want to populate certain indices then specify their indices like ```var x = [12]int{1, 5: 4, 6, 10:100, 15}``` this works like a pointer-based framework
- you don't have to specify the array size if you are initializing an array literal like this ```var x = [...]int{10, 20}```
- arrays can be compared 
- you can make 2D arrays ```var x [2][3]int``` means array of length 2, whose elements are of type array of ints of length 3. Go does not support matrices
- read and write to/from array by normal bracket notation ```x[0] = 10``` 
- no negative indices (like Python), 
- len() takes array and returns size

### Biggest drawback of arrays

> Go considers size of the array to be part of the type of the array
>
> Hence [3]int and [4]int are different types, not the same *array* type as you might think
>
> Hence cant use variable to specify size of array, because types are resolved at compiletime
>
> There's no option to change all of them to identical types, hence you can't write functions that works with array of any size. (WHA- ?) 

- Hence, dont use arrays unless specified. But why this limitation? Because arrays are a backbone for **slices** 

## Slices

- When you think of using an array in any other lang, use slice in Go. Basically, for a sequence of vals
- Obviously, length is not a part of the type for a slice
- we can grow them, iterate through them etc normally
- You **don't** specify the size for a slice when you declare them such as ```var x = []int{10,20,30}```. The subtle difference is that ```[...]``` is an array and ```[]``` is a slice. Hence. the above is a slice literal.
- Same rules apply for specifying sparse and multidimensional slices, reading and writing from a slice; as an array.
- ```var x []int``` is an empty slices initialized to the value ```nil```, that is its zero value. nil represents lack of value, has no type 
- slices cannot be compared, you can only compare a slice with nil, nothing else. not another slice, not an int etc.
- len() works for slices too, returns its length

### append()

- takes 2 params, slice and value(s). slice type and value type must be same
- it returns a slice of that type, hence needs to be initalized to a variable. remeber to always initalize, otherwise compile-time error 😄

```
var x []int
x = append(x, 10)

var x = []int{1,2,3,4}
x = append(x,5)

x = append(x, 6, 7, 8, 9)
```

- you can also append another slice to a slice, first you need to unroll it using ```...``` like 

```
y := []int{1,2,3}
x = append(x, y...)
```


