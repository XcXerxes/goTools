package set

import (
	"fmt"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set1 := NewHashSet()
	set1.Add("a")
	set1.Add("a")
	fmt.Println(set1)
	fmt.Println(set1.Len())
	fmt.Println(set1.Elements())
}

func TestHashSet_Remove(t *testing.T) {
	set1 := NewHashSet()
	set1.Add("a")
	set1.Add("b")
	fmt.Println("before==============", set1.Elements())

	set1.Remove("a")
	fmt.Println(set1.Len())
	fmt.Println("after=============", set1.Elements())
}

func TestHashSet_Clear(t *testing.T) {
	set1 := NewHashSet()
	set1.Add("a")
	set1.Add("b")
	set1.Add("c")
	fmt.Println("before================", set1.Elements())

	set1.Clear()
	fmt.Println("after==================", set1.Elements())
}

func TestHashSet_Contains(t *testing.T) {
	set1 := NewHashSet()
	set1.Add("a")
	set1.Add("b")
	fmt.Println("a=============", set1.Contains("a"))
	fmt.Println("c==============", set1.Contains("c"))
}

func TestHashSet_Same(t *testing.T) {
	set1 := NewHashSet()
	set1.Add("a")
	set1.Add("b")

	set2 := NewHashSet()
	fmt.Println("set1=====?set2", set1.Same(set2))

	set2.Add("a")
	set2.Add("b")
	fmt.Println("set1==========?set2", set1.Same(set2))

	set2.Clear()
	set2.Add("b")
	set2.Add("a")
	fmt.Println("set1==========?set2", set1.Same(set2))
}
