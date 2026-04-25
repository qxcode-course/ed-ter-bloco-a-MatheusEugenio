package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int //slice
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) At(index int) (int, error) {

	if index >= v.size || index < 0 {
		return 0, errors.New("index out of range")
	}

	return v.data[index], nil
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}

	v.capacity = newCapacity
}

func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) Clear() {
	v.size = 0
}

func (v *Vector) Contains(value int) bool {
	for _, val := range v.data {
		if val == value {
			return true
		}
	}
	return false
}

func (v *Vector) IndexOf(value int) int {

	for i, val := range v.data {
		if val == value {
			return i
		}
	}
	return -1
}

func (v *Vector) Insert(index int, value int) error {

	if v.size == v.capacity {

		if v.capacity == 0 {
			v.capacity = 1
		} else {
			v.capacity *= 2
		}

		newData := make([]int, v.capacity)
		for i, val := range v.data {
			newData[i] = val
		}

		v.data = newData
	}

	if index > v.capacity || index < 0 {
		return errors.New("vector is empty")
	}

	if index >= v.size {
		v.size++
		v.data[index] = value
		return nil
	}

	for i := v.size - 1; i >= index; i-- {
		v.data[i+1] = v.data[i]
	}
	v.size++
	v.data[index] = value
	return nil
}

func (v *Vector) PopBack() (int, error) {

	if v.size == 0 {
		return 0, errors.New("vector is empty")
	}

	indexUltimoVal := v.size - 1
	v.size--

	return v.data[indexUltimoVal], nil
}

func (v *Vector) PushBack(value int) {

	if v.size < v.capacity {
		v.data[v.size] = value
		v.size++
		return
	}

	v.capacity *= 2
	newData := make([]int, v.capacity)

	for i, val := range v.data {
		newData[i] = val
	}

	v.data = newData
	v.data[v.size] = value
	v.size++
}

func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}

func (v *Vector) Status() string {
	return "size:" + strconv.Itoa(v.size) + " capacity:" + strconv.Itoa(v.capacity)
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (v *Vector) Set(index int, value int) error {

	if index >= v.size || index < 0 {
		return errors.New("index out of range")
	}

	v.data[index] = value
	return nil
}

func (v *Vector) Erase(index int) error {

	if index >= v.size || index < 0 {
		return errors.New("index out of range")
	}

	if index == v.size-1 {
		v.size--
		return nil
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}

func (v *Vector) normalizaIndice(index int) int {

	if v.size == 0 {
		return 0
	}

	if index >= 0 && index < v.size {
		return index
	}

	// se o indice for negativo
	if index < 0 {
		return v.normalizaIndice(index + v.size)
	}

	// se o indice extrapola o size
	return v.normalizaIndice(index - v.size)
}

func (v *Vector) Slice(start int, end int) *Vector {

	normStart := v.normalizaIndice(start)
	normEnd := v.normalizaIndice(end)

	var novoTamanho int
	if normStart == normEnd && start != end {
		//volta completa
		novoTamanho = v.size
	} else if normStart <= normEnd {
		// normal
		novoTamanho = normEnd - normStart
	} else {
		// circular, quando start > end
		novoTamanho = (v.size - normStart) + normEnd
	}

	novaCapacidade := v.capacity - normStart

	return &Vector{
		data:     v.data[normStart:v.capacity],
		size:     novoTamanho,
		capacity: novaCapacidade,
	}
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
