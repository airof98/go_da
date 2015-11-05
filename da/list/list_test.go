package list

import (
	"../list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	l.PushBack(100)
	l.PushBack(200)
	l.PushBack("test")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
