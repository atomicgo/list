package list_test

import (
	"fmt"
	"math/rand"
	"strings"

	"atomicgo.dev/list"
)

func ExampleList_functional() {
	var l list.List[string]

	l.Append("a", "b", "c")

	l.Map(func(s string) string {
		return s + "!"
	}).Filter(func(s string) bool {
		return !strings.Contains(s, "b")
	})

	fmt.Println(l)
	// Output:
	// [a! c!]
}

func ExampleList_Append() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l)
	// Output:
	// [a b c]
}

func ExampleList_Filter() {
	var l list.List[string]

	l.Append("a", "b", "c")

	l.Filter(func(s string) bool {
		return s != "b"
	})

	fmt.Println(l)
	// Output:
	// [a c]
}

func ExampleList_Map() {
	var l list.List[string]

	l.Append("a", "b", "c").Map(strings.ToUpper)

	fmt.Println(l)
	// Output:
	// [A B C]
}

func ExampleList_Slice() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l.Slice())
	// Output:
	// [a b c]
}

func ExampleSliceToList() {
	l := list.SliceToList([]string{"a", "b", "c"})

	fmt.Println(l)
	// Output:
	// [a b c]
}

func ExampleList_Insert() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Insert(1, "inserted")
	fmt.Println(l)

	l.Insert(0, "a", "b", "c")
	fmt.Println(l)
	// Output:
	// [a inserted b c]
	// [a b c a inserted b c]
}

func ExampleList_Remove() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Remove(1)

	fmt.Println(l)
	// Output:
	// [a c]
}

func ExampleList_Set() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Set(1, "set")

	fmt.Println(l)
	// Output:
	// [a set c]
}

func ExampleList_Get() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l.Get(1))
	// Output:
	// b
}

func ExampleList_Len() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l.Len())
	// Output:
	// 3
}

func ExampleList_Sort() {
	var l list.List[int]

	l.Append(3, 2, 1)
	l.Sort(func(a, b int) bool {
		return a < b
	})

	fmt.Println(l)
	// Output:
	// [1 2 3]
}

func ExampleList_ForEach() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.ForEach(func(s string) {
		fmt.Println(s)
	})
	// Output:
	// a
	// b
	// c
}

func ExampleList_Contains() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l.Contains("b"))
	// Output:
	// true
}

func ExampleList_IndexOf() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l.IndexOf("b"))
	// Output:
	// 1
}

func ExampleList_Copy() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l.Copy().Append("appended"))
	fmt.Println(l.Copy().Append("appended2"))
	fmt.Println(l)
	// Output:
	// [a b c appended]
	// [a b c appended2]
	// [a b c]
}

func ExampleList_Clear() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Clear()

	fmt.Println(l)
	// Output:
	// []
}

func ExampleList_Reverse() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Reverse()

	fmt.Println(l)
	// Output:
	// [c b a]
}

func ExampleList_Shuffle() {
	var l list.List[string]

	rand.Seed(1337) // You should probably use time.Now().UnixNano() in your code.

	l.Append("a", "b", "c")
	l.Shuffle()

	fmt.Println(l)
	// Output:
	// [b a c]
}

func ExampleList_Swap() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Swap(0, 2)

	fmt.Println(l)
	// Output:
	// [c b a]
}

func ExampleList_Reduce() {
	var l list.List[int]

	l.Append(1, 2, 3)
	sum := l.Reduce(func(a, b int) int {
		return a + b
	})

	fmt.Println(sum)
	// Output:
	// 6
}

func ExampleList_String() {
	var l list.List[string]

	l.Append("a", "b", "c")

	fmt.Println(l)
	// Output:
	// [a b c]
}

func ExampleList_Prepend() {
	var l list.List[string]

	l.Append("a", "b", "c")
	l.Prepend("d", "e", "f")

	fmt.Println(l)
	// Output:
	// [d e f a b c]
}
