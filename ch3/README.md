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
- it returns a slice of that type, hence needs to be initalized to a variable. 
- remember to always initalize, otherwise compile-time error 😄 It's because go uses call-by-value, which means it makes a copy every time something is passed as a parameter, hence after making the copy it needs to reassign it to a var right? that's why it always needs to initialized

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

### cap(), the Capacity in go

- cap function allows you to see the capacity of a slice. it means how much you can append until the compiler actually copies all your old slice data, renames it and gives it more space. old slice is garbage collected. not used much, but good to know that there exists a cap in memory for slices, its not infinite (obviously)

```
cap(x)
```

### make()

- If you do know the size of a slice beforehand, its better to use ```make()``` because there is an overhead that comes with automtically growing the slice by appending.

- this is also the only way to initialize a slice that has a fixed size to it (remember from slices introduction that the syntax should not have the size, otherwise it becomes an array declaration)

```
p := make([]int, 10)
```

- note that length and capacity are set to the number that you give in make(), this means that the slice is not empty, its initialized to its null value, which is 0 for all int literals

- Hence, **don't append** to make() slices, because they are not empty slices. 

- **But**, make() also takes one param in the middle for its initial length. If you make that 0, then it means your slice is empty, has capacity 10. You **can** append to this safely.

```
x := make([]int, 0, 10)
```

Check lines 52-60 in [main.go](./main.go) to see the working example.

### Slicing Slices

```
x := []int{1, 2, 3, 4}
y := x[:2]
z := x[1:]
```

> You need to know that a slice and its slice will share memory (the book says 'sometimes', i dont know why that is). This means that x and y share memory and changes to an ele in a slice affects all slices of that slice that have that element.

- Hence, refrain from taking slice of a slice or modifying a slice of a slice.

- To prevent the above, you can use a three part slice such as:
```
y := x[:2:2]
```

- This limits the capacity of the subslices to their lengths and append() causes new slices to be created instead of old ones. (Remember that new slices are created once its capacity is gone over?)

## Strings and Runes

- String is not made out of runes (weird), its a sequence of bytes encoded in UTF-8

- They can be sliced too, but remember rhat when you're slicing something like an emoji (which is not in the English language) and since strings are stored as bytes, your slicing might not really work because one emoji can be represented as multiple bytes

```
var s string = "Hello"
var s2 string = s[2:4] 
```

### Conversions among rune, byte and string

- Remember that byte is the most common storage mechanism in Go

```
var a rune = 'x'
var s string = string(a)

var b byte = 'y'
var s2 string = string(b)
```

- Both of the above are allowed, but don't try converting an int to string, like in python.

- You can also convert strings to slices like this: (remember that you always convert these character types to bytes, hence slice of bytes are very common, not slice of runes or slice of strings since strings are immutable)

```
var s string = "Hello"
var bs []byte = []byte(s)
var rs []rune = []rune(s)

Output:
[72 101 108 108 111 44 32 240 159 140 158]
[72 101 108 108 111 44 32 127774]
```

- The entire string converted to UTF-8 is the []byte() slice; the second is unnecessary.

## Maps

- Like dict in Python, for k-v pairs

- The map that's built in to Go is a hash map

- To declare one:

```
//a nil map, cant write to it (apparently causes panic?)
var mymap map[string]int 

//or

// empty map, not a nil map
mymap := map[string]int{}

//map literal
mymap := map[string]int{

    "me" : 1,
    "you" : 2,
}

```

- Notice that even last value has a comma

- you can also use make() if you know the size of map you want to create

- Maps:

    - automatically grow
    - can be used with make()
    - len() tells # of k-v pairs
    - have a zero val of nil
    - are not comparable, except with nil
    - are unordered (obviously, just like dict in Python)

### Hash Maps

- we already know what a hash map, collision, key- value are.

- you insert k-v, k becomes a non-unique number using hash algo, this number is used as index to an array, each ele in array is a bucket and k-v is stored in it. bviously, the bucket is also an array, hence stores multiple ele, but when a bucket has more than one elemts there is a collision. stores both k-vs in that bucket.

### r/w into maps

- assign values using ```=```, not ```:=```

- if a value is not set for a key, its set to that data type's zero value (0 for int)

```
mymap := map[string]int{}
mymap["karthik"] = 15
```

### comma ok idiom

- used to tell the difference between a key that's associated with a zero value and key not in a map

- you assign the results of a map read to 2 variables, first gets the value associated with the key and second is a bool. if ```ok``` is true then key is present in map, otherwise no.

```
m := map[string]int{

    "hello" : 5,
    "world" : 0,
}

v, ok := m["hello"]
fmt.printLn(v, ok)

v, ok := m["world"]
fmt.printLn(v, ok)

v, ok := m["there]
fmt.printLn(v, ok)


//output: 

5 true
0 true
0 false
```

