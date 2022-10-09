package list

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

// List is a generic list type.
type List[T any] struct {
	items []T
}

// SliceToList converts a slice to a list.
func SliceToList[T any](items []T) List[T] {
	return List[T]{items: items}
}

// Len returns the length of the list.
func (l *List[T]) Len() int {
	return len(l.items)
}

// Get returns the item at the given index.
func (l *List[T]) Get(i int) T {
	return l.items[i]
}

// Set sets the item at the given index.
func (l *List[T]) Set(i int, item T) *List[T] {
	l.items[i] = item
	return l
}

// Append adds items to the end of list.
func (l *List[T]) Append(items ...T) *List[T] {
	l.items = append(l.items, items...)
	return l
}

// Prepend adds items to the beginning of list.
func (l *List[T]) Prepend(items ...T) *List[T] {
	return l.Insert(0, items...)
}

// Remove removes the item at the given index.
func (l *List[T]) Remove(i int) *List[T] {
	l.items = append(l.items[:i], l.items[i+1:]...)
	return l
}

// Insert adds items at the given index.
func (l *List[T]) Insert(i int, items ...T) *List[T] {
	l.items = append(l.items[:i], append(items, l.items[i:]...)...)
	return l
}

// Slice returns the list as a slice.
func (l *List[T]) Slice() []T {
	return l.items
}

// Map overwrites the list with the result of applying the given function to each item.
func (l *List[T]) Map(f func(T) T) *List[T] {
	for i, item := range l.items {
		l.items[i] = f(item)
	}

	return l
}

// Filter removes all items from the list that do not match the given predicate.
func (l *List[T]) Filter(f func(T) bool) *List[T] {
	var filtered []T
	for _, item := range l.items {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	l.items = filtered

	return l
}

// Reduce reduces the list to a single value by applying the given function to each item.
func (l *List[T]) Reduce(f func(T, T) T) T {
	if len(l.items) == 0 {
		return reflect.Zero(reflect.TypeOf(*new(T))).Interface().(T)
	}

	acc := l.items[0]
	for _, item := range l.items[1:] {
		acc = f(acc, item)
	}

	return acc
}

// ForEach iterates over the list and calls the given function for each item.
func (l List[T]) ForEach(f func(T)) {
	for _, item := range l.items {
		f(item)
	}
}

// Contains returns true if the list contains the given item.
func (l List[T]) Contains(item T) bool {
	for _, i := range l.items {
		if reflect.DeepEqual(i, item) {
			return true
		}
	}

	return false
}

// IndexOf returns the index of the first occurrence of the given item.
func (l List[T]) IndexOf(item T) int {
	for j, i := range l.items {
		if reflect.DeepEqual(i, item) {
			return j
		}
	}

	return -1
}

// Reverse reverses the list.
func (l *List[T]) Reverse() *List[T] {
	for i := len(l.items)/2 - 1; i >= 0; i-- {
		opp := len(l.items) - 1 - i
		l.items[i], l.items[opp] = l.items[opp], l.items[i]
	}

	return l
}

// Clear removes all items from the list.
func (l *List[T]) Clear() *List[T] {
	l.items = l.items[:0]
	return l
}

// Copy returns a new copy of the list.
// Useful when you want to modify a list without modifying the original.
func (l *List[T]) Copy() *List[T] {
	items := make([]T, len(l.items))
	copy(items, l.items)

	return &List[T]{items: items}
}

// Sort sorts the list using the given function.
func (l *List[T]) Sort(f func(T, T) bool) *List[T] {
	sort.Slice(l.items, func(i, j int) bool {
		return f(l.items[i], l.items[j])
	})

	return l
}

// Shuffle shuffles the list.
// You need to seed the random number generator yourself.
func (l *List[T]) Shuffle() *List[T] {
	for i := range l.items {
		j := rand.Intn(i + 1)
		l.items[i], l.items[j] = l.items[j], l.items[i]
	}

	return l
}

// Swap swaps the items at the given indices.
func (l *List[T]) Swap(i, j int) *List[T] {
	l.items[i], l.items[j] = l.items[j], l.items[i]
	return l
}

func (l List[T]) String() string {
	return fmt.Sprint(l.items)
}
