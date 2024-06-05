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
> Hence [3]int and [4]int are different types, not the same *array* type as you might think
> Hence cant use variable to specify size of array, because types are resolved at compiletime
> There's no option to change all of them to identical types, hence you can't write functions that works with array of any size. (WHA- ?) 

