### 3.2 运算符

**1.**

To declare an empty slice, with a non-fixed size,
is it better to do:

    mySlice1 := make([]int, 0)

or:

    mySlice2 := []int{}

Just wondering which one is the correct way.  

- The two alternative you gave are semantically identical, but using `make([]int, 0)` will result in an internal call to *runtime.makeslice* (Go 1.14).

  You also have the option to leave it with a `nil` value:

      var myslice []int

  As written in the [Golang.org blog](http://blog.golang.org/slices):

  > a nil slice is functionally equivalent to a zero-length slice, even though it points to nothing. It has length zero and can be appended to, with allocation.

  A `nil` slice will however `json.Marshal()` into `"null"` whereas an empty slice will marshal into `"[]"`, as pointed out by  @farwayer.

  None of the above options will cause any allocation, as pointed out by @ArmanOrdookhani.

- ```
  func main() {
  	var a struct{}
  	b := make([]int, 0)
  	c := []int{}
  	fmt.Println(&a==nil, &b==nil, &c==nil)
  }
  ```