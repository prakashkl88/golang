// GO program

package main
import "fmt"

func main(){
  a := 3
  b := true

  //another way of initializing variables
  var c int32
  c = 5
  var d float64
  d = 5.64

  //list or array
  arr := []int{3, 6 ,8, 11, 15}

  //dict or hashMap or map
  hash := map[string]int{"a": 1, "b": 2, "c": 3}

  fmt.Println(a, b, c, d, arr, hash)
}
