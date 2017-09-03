### types
- int, uint : 8, 16, 32, 64 bit
- float : 32, 64 bit
- string : "literal interpolate", \`literal\`
- rune : (int32 alias) 'a'
- []<type> : is called a slice. Contains references to <type>s
- map[<key type>]<val type> : is called a map. Contains references to <key> and <val> types
- Array: `var x[<len>] <type>`. Arrays are static sized.

## functions
- if `retType` declared must have `return <return value>`
- if `retName` declared must have `return`
- args passed by value
- anonymous:
  ```go
  func(){ code }
  ```
  immediate call
  ```go
  func(){ code }()
  ```
- named:
  ```go
  func name(){ code }
  func name(argname argtype, ...) { code }
  func name(...) retType { code }
  func name(...) (retName retType) { code }
  ```

### pointers
- `*` declares ptr
  ```go
  var i *int
  ```
- `*` dereferences ptr
  ```go
  func getPrtVal(i *int) int{
    return *i
  }
  ```
- `&` get mem addr
  ```go
  i := 1
  getPtrVal(&i)
  ```

### Miscellaneous
- slicing and array/slice
  ```go
    slice = []int{0,1,2,3,4,5,6,7,8}
    chunk = slice[3:len(slice)-2]
  ```
  nb: the slicing takes up to *but not including* the limit index
  ```go
    => [3,4,5]
  ```
- can instantiate a slice and optionally give the capacity of the underlying array:
  ```go
    make([]<type>, <length>, <capacity>)
    make([]<type>, 50, 100) == make([100]<type>)[0:50]
  ```
  can check the capacity of a slice using `cap(slice)` but slices _are_ extensible, a new array is created (with copied data) if the capacity is exceeded. There is a performance hit for this, so initialize your slices with a decent capacity.
  omitting capacity arg sets underlying array to size of slice
