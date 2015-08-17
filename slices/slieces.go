package main
import "fmt"

func main() {

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)

	// Append element to sliece
	s2 = append(s2, 5)
	// Append sliece to sliece
	s3 := append(s1, s2...)

	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
}
