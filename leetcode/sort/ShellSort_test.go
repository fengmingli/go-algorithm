package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestShellSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		list = append(list, r.Intn(100))
	}
	fmt.Println("sort before", list)
	shellSort(list, len(list)/2)
	fmt.Println("sort after", list)
}
