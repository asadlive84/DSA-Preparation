package main

import "fmt"

type Array struct {
	elements []int32
	length   int32
}

func newArray(size int32) *Array {
	return &Array{
		elements: make([]int32, size+1),
		length:   0,
	}
}

func (arr *Array) push(data int32) {
	// O(1)
	arrLen := len(arr.elements)
	if arrLen == int(arr.length) {
		return
	}
	arr.elements[arr.length] = data
	arr.length++
}

func (arr *Array) showByIndex(index int32) {
	// O(1)
	fmt.Println(arr.elements[index])

}

func (arr *Array) show() {
	// O(n)
	fmt.Println(arr.elements)
}

func (arr *Array) insert(index int32, data int32) {
	// O(1)
	arrLen := len(arr.elements)
	if index > int32(arrLen)-1 {
		return
	}
	arr.elements[index] = data

}

func (arr *Array) shiftInsert(index int32, data int32) {
	// O(n)
	arrLen := len(arr.elements)
	if index < 0 || index > int32(arrLen)-1 {
		return
	}

	newArr := make([]int32, arr.length+1)
	copy(newArr, arr.elements)

	// for i := index; i < arr.length; i++ {
	// 	temp := arr.elements[i]
	// 	newArr[i+1] = temp
	// }

	// data copied from right to left

	for i := arr.length - 1; i > index; i-- {
		newArr[i+1] = arr.elements[i]
	}

	newArr[index] = data
	arr.elements = newArr
	arr.length++

}

func (arr *Array) delete(data int32) {
	// O(n)
	for index, k := range arr.elements {
		if k == data {
			arr.elements[index] = 0
			arr.length--
			return
		}
	}
}

func (arr *Array) deleteWithIndex(index int32) {
	if index < 0 || (arr.length-1) < index {
		return
	}

	
	for i := index; i < arr.length-1; i++ {
		arr.elements[i] = arr.elements[i+1]
	}
	
	arr.length--
	arr.elements=arr.elements[:arr.length]
}



func (arr *Array) pop() {
	// O(1)
	if arr.length == 0 {
		return
	}
	lastItem := arr.length - 1
	arr.elements[lastItem] = 0
	arr.elements = arr.elements[:arr.length-1]
	arr.length--

}

func main() {

	arr := newArray(4)

	arr.push(5)
	arr.push(7)
	arr.push(8)
	arr.push(9)
	arr.push(10)

	arr.show()

	arr.deleteWithIndex(1)

	arr.show()
}
