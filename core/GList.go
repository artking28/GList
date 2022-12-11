package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type GList[Any any] []Any

// NewGList method creates am instance of GList
func NewGList[Any any](obj ...Any) *GList[Any] {
	l := (*GList[Any])(&[]Any{})
	return l.Add(obj...)
}

// Reverse method reverses the postition of all elements, so the start will be the end, and the end, the begining
func (this *GList[Any]) Reverse() *GList[Any] {
	if this.Count() < 2 {
		return this
	} else {
		reordered := *this.Clone()
		for i := range *this {
			i = i
			l := reordered.Sublist(i, reordered.LastIndexOf())
			reordered = append(*reordered.Sublist(0, i), reordered.Last())
			reordered.Add(*l...)
		}
		*this = reordered
		return &reordered
	}
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

// DropIf method deletes elements if the given condition be satisfied
func (this *GList[Any]) DropIf(fn func(Any) bool) {
	//n := this.Clone()
	count := 0
	for index := this.Count() - 1; index > 0; index-- {
		if fn(this.Get(index)) == true {
			this.DropAt(index)
			count++
		}
	}
}

// IndexOf looks for the index of some object, if it ain't in the list IndexOf retuns '-1'
func (this *GList[Any]) IndexOf(obj Any) int {
	obj1, _ := json.Marshal(obj)
	for i, one := range *this {
		if &obj == &one {
			return i
		} else {
			if obj2, err := json.Marshal(one); err != nil {
				panic("Object cannot be compared")
			} else {
				if string(obj1) == string(obj2) {
					return i
				}
			}
		}
	}
	return -1
}

// ContainsAll method checks if all given elements are into the list
func (this *GList[Any]) ContainsAll(objs ...Any) bool {
	for _, one := range objs {
		if i := this.IndexOf(one); i == -1 {
			return false
		}
	}
	return true
}

// Clear method empty the list
func (this *GList[Any]) Clear() *GList[Any] {
	*this = *NewGList[Any]()
	return this
}

// Clone method creates a copy of the list
func (this *GList[Any]) Clone() *GList[Any] {
	return NewGList[Any](*this...)
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

func (this *GList[Any]) LastIndexOf() int {
	l := this.Count() - 1
	if l <= 0 {
		return 0
	} else {
		return l
	}
}

// IsEmpty method checks if the list is empty
func (this *GList[Any]) IsEmpty() bool {
	return len(*this) == 0
}

// First method reuturns the firt element os the list
func (this *GList[Any]) First() Any {
	if this.Count() == 0 {
		panic(errors.New("an empty list can't use 'First()' method"))
	} else {
		return (*this)[0]
	}
}

// Last method reuturns the last element os the list
func (this *GList[Any]) Last() Any {
	if this.Count() == 0 {
		panic(errors.New("an empty list can't use 'Last()' method"))
	} else {
		ret := (*this)[this.Count()-1]
		return ret
	}
}

// Sublist creates a new list based on the given range
func (this *GList[Any]) Sublist(from int, until int) *GList[Any] {
	sub := NewGList[Any]((*this)[from:until]...)
	return sub
}

// Partition method splits the list into 2 sublist and the elements will be put into each based on given condition
func (this *GList[Any]) Partition(fn func(Any) bool) (*GList[Any], *GList[Any]) {
	ret1 := NewGList[Any]()
	ret2 := NewGList[Any]()
	for _, one := range *this {
		if fn(one) == true {
			ret1.Add(one)
		} else {
			ret2.Add(one)
		}
	}
	return ret1, ret2
}

// QuickSort will order the list by default parameters
func (this *GList[Any]) QuickSort() *GList[Any] {
	kindList := NewGList[reflect.Kind](
		reflect.Invalid,
		reflect.Array,
		reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.UnsafePointer)
	if this.Count() < 2 {
		return this
	} else {
		pivot := this.First()
		smaller, greater := this.Drop(1).Partition(func(obj Any) bool {
			if kindList.ContainsAll(reflect.TypeOf(obj).Kind()) {
				panic("Unsortable Glist, this type cannot be processed")
			} else {
				typeId := reflect.ValueOf(obj).Kind()
				if typeId == reflect.Bool {
					return fmt.Sprintf("%v", obj)[0] == 'f'
				} else if typeId >= reflect.Int && typeId <= reflect.Complex128 {
					n1, _ := strconv.ParseFloat(fmt.Sprintf("%v", obj), 64)
					n2, _ := strconv.ParseFloat(fmt.Sprintf("%v", pivot), 64)
					return n1 <= n2
				} else if typeId == reflect.Interface || typeId == reflect.Struct {
					if reflect.ValueOf(obj).NumField() <= 0 {
						panic("Unsortable Glist, this type cannot be processed")
					} else {
						n1 := fmt.Sprintf("%v", reflect.ValueOf(obj).Field(0))[0]
						n2 := fmt.Sprintf("%v", reflect.ValueOf(pivot).Field(0))[0]
						return n1 <= n2
					}
				} else {
					return (int)(fmt.Sprintf("%v", obj)[0]) <= (int)(fmt.Sprintf("%v", pivot)[0])
				}
			}
		})
		return smaller.QuickSort().Add(pivot).Add(*greater.QuickSort()...)
	}
}

// ForEach will iterate your array in a more clean, confortable and readable way, just like Java or TypeScript
func (this *GList[Any]) ForEach(fn func(index int, obj Any)) {
	for i, one := range *this {
		fn(i, one)
	}
}

// Filter method will return a new list with elements satisfing the given condition
func (this *GList[Any]) Filter(fn func(obj Any) bool) *GList[Any] {
	ret := NewGList[Any]()
	this.ForEach(func(i int, one Any) {
		if fn(one) == true {
			ret.Add(one)
		}
	})
	return ret
}
