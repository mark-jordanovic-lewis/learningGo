### types
- int, uint : 8, 16, 32, 64 bit
- float : 32, 64 bit
- string : "literal interpolate", \`literal\`
- rune : (int32 alias) 'a'
- []<type> : is called a slice. Contains references to <type>s
- map[<key type>]<val type> : is called a map. Contains references to <key> and <val> types


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
