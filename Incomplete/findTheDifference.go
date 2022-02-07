import "fmt"

func findTheDifference(s string, t string) byte {
    var result byte

	for _, c := range []byte(s) {
		result ^= c
        fmt.Println(c)
        fmt.Println(result)
	}

	for _, c := range []byte(t) {
		result ^= c
        fmt.Println(c)
        fmt.Println(result)
	}

	return result
}
