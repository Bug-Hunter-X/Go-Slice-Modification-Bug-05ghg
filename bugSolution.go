func myFunc(a []int) {
b  := make([]int, len(a))
for i := range a {
  b[i] = a[i] + 1
}
 return b
}

func main() { x := []int{1,2,3}; y := myFunc(x);fmt.Println(x, y)}