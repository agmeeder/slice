// This package provides Javascript like methods to manipulate generic slices.
//
// Use the generic Slice[T any] type to define slices of type T = []T.
//
// For example: listOfString := Slice[string]{} or tasks := Slice[Task]{}
package slice

import (
	"fmt"
	"sort"
	"strings"
)

// Defining a generic slice type
type Slice[T any] []T

// METHODS
// The At() method returns an indexed element from an slice.
func (s *Slice[T]) At(index int) T {
	return (*s)[index]
}

// The Concat() method does not change the existing slice. It returns a new slice.
func (s *Slice[T]) Concat(elements ...Slice[T]) *Slice[T] {
	newArray := make(Slice[T], 0)
	newArray = append(newArray, *s...)
	for _, element := range elements {
		newArray = append(newArray, element...)
	}
	return &newArray
}

// The join() method joins all slice elements into a string.
func (s Slice[T]) Join(separator string) string {
	var builder strings.Builder
	for i, v := range s {
		builder.WriteString(fmt.Sprint(v))
		if i < len(s)-1 {
			builder.WriteString(separator)
		}
	}
	return builder.String()
}

// The Length() method returns the length of the slice.
func (s *Slice[T]) Length() int {
	return len(*s)
}

// The Merge() method modifies the existing slice instead of returning a new one.
func (s *Slice[T]) Merge(items ...Slice[T]) *Slice[T] {
	for _, item := range items {
		*s = append(*s, item...)
	}
	return s
}

// The Pop() method removes the last element from an slice.
func (s *Slice[T]) Pop() T {
	var e T
	l := len(*s)
	if l > 0 {
		e = (*s)[l-1]
		*s = (*s)[:l-1]
	}
	return e
}

// The Push() method adds a new element to an slice (at the end).
func (s *Slice[T]) Push(item T) *Slice[T] {
	*s = append(*s, item)
	return s
}

// The Shift() method removes the first slice element and "shifts" all other elements to a lower index.
func (s *Slice[T]) Shift() T {
	var e T
	l := len(*s)
	if l > 0 {
		e = (*s)[0]
		*s = (*s)[1:]
	}
	return e
}

/*
The Slice() method can be used to remove an item from a slice and return the modified slice.

The first argument (index) is the index of the element to remove.
*/
func (s *Slice[T]) Slice(index int) *Slice[T] {
	return s.Splice(index, 1)
}

/*
The Splice() method can be used to add new or remove elements to an slice

The first argument (start) is the index at which to start adding/removing items.

The second argument (deleteCount) is the number of elements to remove from original the slice.

The remaining argument (elements) are the new elements to add to the slice.
*/
func (s *Slice[T]) Splice(start, deleteCount int, elements ...T) *Slice[T] {
	l := len(*s)
	if start > l {
		start = l
	}
	if start < 0 {
		start = 0
	}
	end := start + deleteCount
	if end > l {
		end = l
	}
	result := Slice[T]{}
	result = append(result, (*s)[:start]...)
	result = append(result, elements...)
	result = append(result, (*s)[end:]...)
	*s = result
	return s
}

/*
The ToSliced() method returns a new slice from an existing slice where the element with the given index is removed.
*/
func (s *Slice[T]) ToSliced(index int) *Slice[T] {
	result := Slice[T]{}
	result = append(result, *s.ToSpliced(index, 1)...)
	return &result
}

/*
The ToSpliced() method can be used to add new items to an slice without changing the original slice.

The first argument (start) is the index at which to start adding/removing items.

The second argument (deleteCount) is the number of elements to remove from original the slice.

The remaining argument (elements) are the new elements to add to the slice.
*/
func (s *Slice[T]) ToSpliced(start, deleteCount int, elements ...T) *Slice[T] {
	l := len(*s)
	if start > l {
		start = l
	}
	if start < 0 {
		start = 0
	}
	end := start + deleteCount
	if end > l {
		end = l
	}
	result := Slice[T]{}
	result = append(result, (*s)[:start]...)
	result = append(result, elements...)
	result = append(result, (*s)[end:]...)
	return &result
}

// The ToString() method joins all slice elements into a string.
func (s Slice[T]) ToString() string {
	var builder strings.Builder
	for i, v := range s {
		builder.WriteString(fmt.Sprint(v))
		if i < len(s)-1 {
			builder.WriteString("")
		}
	}
	return builder.String()
}

// The UnShift() method adds a new element to an array (at the beginning), and "UnShifts" older elements.
func (s *Slice[T]) UnShift(item T) *Slice[T] {
	*s = append(Slice[T]{item}, *s...)
	return s
}

// SORT

/*
The Sort() method sorts the elements of a slice in-place.

The less function parameter specifies a less-than comparison between two elements that returns true if the first argument is less than the second
*/
func (s *Slice[T]) Sort(less func(T, T) bool) *Slice[T] {
	sort.SliceStable(*s, func(i, j int) bool {
		return less((*s)[i], (*s)[j])
	})
	return s
}

/*
The ToSorted() method sorts the elements of a slice and returns a new slice instead of modifying the original one.

The less function parameter specifies a less-than comparison between two elements that returns true if the first argument is less than the second
*/
func (s *Slice[T]) ToSorted(less func(T, T) bool) *Slice[T] {
	result := Slice[T]{}
	result = append(result, *s...)
	sort.SliceStable(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return &result
}

/*
The Reverse() method reverse sorts the elements of a slice in-place.

The less function parameter specifies a less-than comparison between two elements that returns true if the first argument is less than the second
*/
func (s *Slice[T]) Reverse(less func(T, T) bool) *Slice[T] {
	sort.SliceStable(*s, func(i, j int) bool {
		return less((*s)[j], (*s)[i])
	})
	return s
}

/*
The ToReversed() method reverse sorts the elements of a slice and returns a new slice instead of modifying the original one.

The less function parameter specifies a less-than comparison between two elements that returns true if the first argument is less than the second
*/
func (s *Slice[T]) ToReversed(less func(T, T) bool) *Slice[T] {
	result := Slice[T]{}
	result = append(result, *s...)
	sort.SliceStable(result, func(i, j int) bool {
		return less(result[j], result[i])
	})
	return &result
}

// ITERATION

/*
The Every() method tests whether all elements in the slice pass the provided test function.
*/
func (s Slice[T]) Every(match func(T) bool) bool {
	for _, v := range s {
		if !match(v) {
			return false
		}
	}
	return true
}

// filter
func (s Slice[T]) Filter(match func(T) bool) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		if match(v) {
			result = append(result, v)
		}
	}
	return result
}

/*
The Map() method iterates over each element and applies the provided change function to each element, returning a new slice containing the results.
*/
func (s Slice[T]) Map(change func(T) T) Slice[T] {
	result := make(Slice[T], len(s))
	for i := range s {
		result[i] = change(s[i])
	}
	return result
}

/*
The Reduce() function applies an accumulator function over the elements of the slice, returning a single value.
*/
func Reduce[T any, U any](s Slice[T], initial U, reducer func(U, T) U) U {
	accumulator := initial
	for _, element := range s {
		accumulator = reducer(accumulator, element)
	}
	return accumulator
}

/*
The Some() method tests whether some elements in the slice pass the provided test function.
*/
func (s Slice[T]) Some(match func(T) bool) bool {
	for _, v := range s {
		if match(v) {
			return true
		}
	}
	return false
}

// SEARCH
/*
The Find() method returns the first element and true if the slice satisfies the provided function, otherwise it returns the zero value and false.
*/
func (s Slice[T]) Find(match func(T) bool) (T, bool) {
	var result T
	for _, element := range s {
		if match(element) {
			return element, true
		}
	}
	return result, false
}

/*
The FindLast() method returns the last element and true if the slice satisfies the provided function, otherwise it returns the zero value and false.
*/
func (s Slice[T]) FindLast(match func(T) bool) (T, bool) {
	var result T
	for i := len(s) - 1; i >= 0; i-- {
		if match(s[i]) {
			return s[i], true
		}
	}
	return result, false
}

/*
The Includes() method returns true if the slice contains an element that satisfies the provided testing function.
*/
func (s Slice[T]) Includes(match func(T) bool) bool {
	for _, element := range s {
		if match(element) {
			return true
		}
	}
	return false
}

/*
The IndexOf() method returns the index of the first element in the slice that satisfies the provided testing function.
*/
func (s Slice[T]) IndexOf(match func(T) bool) int {
	for i, element := range s {
		if match(element) {
			return i
		}
	}
	return -1
}

/*
The LastIndexOf() method returns the index of the last element in the slice that satisfies the provided testing function.
*/
func (s Slice[T]) LastIndexOf(match func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if match(s[i]) {
			return i
		}
	}
	return -1
}
