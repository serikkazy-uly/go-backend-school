package main

import "fmt"

// Раскрытый ответ в README
func main() {
	var s = []string{"1", "2", "3"}

	modifySlice(s)

	fmt.Println(s)

}

func modifySlice(i []string) {
	i[0] = "3"
	// fmt.Println(len(i), cap(i)) // 3 3
	i = append(i, "4")               // modifies the slice header, not the underlying array
	fmt.Println("new slice:", i[:6]) // [3 2 3 4]

	i[1] = "5" // modifies the underlying array
	// fmt.Println(len(i), cap(i)) // 4 6 new underlying array allocated
	i = append(i, "6")
	// fmt.Println("new slice:", i[:6]) // [3 5 3 4 6]
}
