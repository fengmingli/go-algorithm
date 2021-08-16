/**
 * @Author LFM
 * @Date 2021/8/16 3:17 下午
 * @Since V1
 */

package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := NewList()
	list.RPush(NewNode(1))
	list.RPush(NewNode(2))
	list.RPush(NewNode(3))
	list.RPush(NewNode(4))
	list.RPush(NewNode(5))

	all := list.LPopAll()

	for i := 0; i < len(all); i++ {
		fmt.Println(all[i].GetData())
		fmt.Println(all[i].GetDataNew())
	}
}
