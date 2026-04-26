package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(newCacpacity int) *Set {
	return &Set{
		data:     make([]int, newCacpacity),
		size:     0,
		capacity: newCacpacity,
	}
}

// privado
func (v *Set) reserve(newCacpacity int) {

	if newCacpacity == 0 {
		v.capacity = 1
	} else if newCacpacity <= v.capacity {
		return
	} else {
		v.capacity *= 2
	}

	dataTemp := make([]int, v.capacity)

	for i, val := range v.data {
		dataTemp[i] = val
	}

	v.data = dataTemp
}

// privado
func (v *Set) binarySearch(value int) int { //retorna o indice do numero, se existe

	inicio, final := 0, v.size-1

	if value == v.data[0] {
		return 0
	} else if value == v.data[v.size-1] {
		return v.size - 1
	}

	for inicio <= final {

		indexMeio := (inicio + final) / 2

		if value > v.data[indexMeio] {
			inicio = indexMeio + 1
		} else if value < v.data[indexMeio] {
			final = indexMeio - 1
		} else {
			return indexMeio
		}
	}

	return -1
}

// privado
func (v *Set) insert(value int, index int) error {

	if v.size == v.capacity {
		v.reserve(v.size + 1)
	}

	if index >= v.capacity || index < 0 {
		return errors.New("value not found")
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

func (v *Set) Insert(value int) { //usa busca binaria para saber se o elemento já existe

	if v.size == 0 {
		v.size++
		v.data[0] = value
		return
	}

	indice := v.binarySearch(value)
	if indice != -1 {
		return //então o valor existe
	}

	if v.size == 0 {
		v.size++
		v.data[0] = value
		return
	}

	index := 0

	for index < v.size {
		if v.data[index] > value {
			break
		}
		index++
	}

	v.insert(value, index)
}

// privado
func (v *Set) erase(index int) error {

	if index < 0 || index >= v.size {
		return errors.New("value not found")
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

func (v *Set) Erase(value int) bool {

	index := v.binarySearch(value)

	if index == -1 {
		return false
	}

	res := v.erase(index)
	if res == nil {
		return true
	} else {
		return false
	}
}

func (v *Set) Contains(value int) bool { //usa busca binaria para saber se o elemento existe
	existe := v.binarySearch(value)
	if existe != -1 {
		return true
	} else {
		return false
	}
}

func (v *Set) String() string {
	return "[" + Join(v.data[:v.size], ", ") + "]"
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

func (v *Set) Clear() {
	v.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			v = NewSet(value)

		case "insert":

			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}

		case "show":

			fmt.Println(v)

		case "erase":

			value, _ := strconv.Atoi(parts[1])

			if !v.Erase(value) {
				fmt.Println("value not found")
			}

		case "contains":

			value, _ := strconv.Atoi(parts[1])

			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
