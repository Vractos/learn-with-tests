package math

import (
	"fmt"
	"strconv"
	"testing"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int, y int) {
		res := Add(x, y)
		if res != x+y {
			t.Errorf("unexpected result: %d", res)
		}
	})
}

func TestAdd(t *testing.T) {

	for i, test := range addTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if output := Add(test.arg1, test.arg2); output != test.expected {
				t.Errorf("Output %v not equal to expected %v", output, test.expected)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}
