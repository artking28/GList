package GList

import (
	"fmt"
)

type GList[Any any] []Any

// NewGList method creates am instance of GList
func NewGList[Any any](obj ...Any) *GList[Any] {
	l := (*GList[Any])(&[]Any{})
	return l.Add(obj...)
}

// Add method adds one or many objects on the end of the list.
func (this *GList[Any]) Add(obj ...Any) *GList[Any] {
	*this = append(*this, obj...)
	return this
}

// Push method adds one or many objects on the starting of the list.
func (this *GList[Any]) Push(obj ...Any) *GList[Any] {
	*this = append(obj, *this...)
	return this
}

// Drop method delete elements from list by counting, if it's negative, the count starts from the ending
func (this *GList[Any]) Drop(n int) *GList[Any] {
	if (n > len(*this)) || (n*(-1) > len(*this)) {
		this.Clear()
		return this
	}
	if n > 0 {
		*this = (*this)[n:len(*this)]
	} else {
		*this = (*this)[0 : len(*this)+n]
	}
	return this
}

// DropAt method deletes an element present at the given index
func (this *GList[Any]) DropAt(index int) *GList[Any] {
	if index > len(*this) || index < 0 {
		return this
	}
	*this = append((*this)[0:index], (*this)[index+1:len(*this)]...)
	return this
}

// DropChunk method deletes elements from list if its into specified range
func (this *GList[Any]) DropChunk(from int, until int) *GList[Any] {
	start := from
	end := until
	if from > len(*this) || until < 0 {
		return this
	}
	if from < 0 {
		start = 0
	}
	if until > len(*this) {
		end = len(*this)
	}
	*this = append((*this)[0:start], (*this)[end:len(*this)]...)
	return this
}

// Clear method empty the list
func (this *GList[Any]) Clear() {
	*this = (*this)[:0]
}

// Stringify method converts the list into string
func (this *GList[Any]) Stringify() string {
	return fmt.Sprintf("%v", *this)
}

// Get method brings the element present at the given index
func (this *GList[Any]) Get(index int) Any {
	return (*this)[index]
}

// Count method returns the size of the list
func (this *GList[Any]) Count() int {
	return len(*this)
}

// IsEmpty method checks if the list is empty
func (this *GList[Any]) IsEmpty() bool {
	return len(*this) == 0
}

// First method reuturns the firt element os the list
func (this *GList[Any]) First() Any {
	return (*this)[0]
}

// Last method reuturns the last element os the list
func (this *GList[Any]) Last() Any {
	return (*this)[this.Count()-1]
}
