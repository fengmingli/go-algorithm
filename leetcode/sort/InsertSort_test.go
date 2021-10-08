/**
 * @Author LFM
 * @Date 2021/9/30 4:46 下午
 * @Since V1
 */

package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		list = append(list, r.Intn(10000))
	}
	fmt.Println("sort before", list)
	InsertSort(list)
	fmt.Println("sort after", list)
}
