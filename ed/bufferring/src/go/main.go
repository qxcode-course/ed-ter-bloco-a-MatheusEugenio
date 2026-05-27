package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data     []int
	front    int
	size     int
	capacity int
}

func (d *Deque) PopBack() error {

	if d.size == 0 {
		return errors.New("fail: buffer vazio")
	}

	d.size--
	return nil
}

func (d *Deque) PopFront() error {

	if d.size == 0 {
		return errors.New("fail: buffer vazio")
	}

	d.front = (d.front + 1) % d.capacity

	d.size--
	return nil
}

func (d *Deque) PushFront(val int) {

	if d.size == d.capacity {
		if d.capacity == 0 {
			d.capacity++
		}

		d.resize(d.capacity * 2)
	}

	d.front = (d.front - 1 + d.capacity) % d.capacity

	d.data[d.front] = val
	d.size++
}

func (d *Deque) PushBack(val int) {

	if d.size == d.capacity {
		if d.capacity == 0 {
			d.capacity++
		}

		d.resize(d.capacity * 2)
	}

	indexBack := (d.front + d.size) % d.capacity

	d.data[indexBack] = val
	d.size++
}

func (d *Deque) Front() (int, error) {

	if d.size == 0 {
		return -1, errors.New("fail: buffer vazio")
	}

	return d.data[d.front], nil
}

func (d *Deque) Back() (int, error) {

	if d.size == 0 {
		return -1, errors.New("fail: buffer vazio")
	}

	index := (d.front + d.size - 1 + d.capacity) % d.capacity
	return d.data[index], nil
}

func (d *Deque) resize(newCapacity int) {

	if newCapacity <= d.size ||
		newCapacity <= len(d.data) {
		return
	}

	copyArray := make([]int, newCapacity)

	for i, _ := range d.data {
		indexCircular := (d.front + i) % d.capacity
		copyArray[i] = d.data[indexCircular]
	}

	d.front = 0
	d.capacity = newCapacity
	d.data = copyArray
}

func (d *Deque) Clear() {
	d.size = 0
	d.front = 0
}

func (d *Deque) Len() int {
	return d.size
}

func (b *Deque) String() string {
	result := []string{}
	for i := range b.size {
		val := b.data[(b.front+i)%b.capacity]
		result = append(result, fmt.Sprint(val))
	}
	return "[" + strings.Join(result, ", ") + "]"
}

func (b *Deque) Debug() string {
	result := make([]string, b.capacity)
	for i := range result {
		result[i] = " _"
		if i == b.front {
			result[i] = ">_"
		}
	}
	for i := range b.size {
		index := (b.front + i) % b.capacity
		val := b.data[index]
		prefix := " "
		if i == 0 {
			prefix = ">"
		}
		result[index] = fmt.Sprintf("%s%d", prefix, val)
	}
	return strings.Join(result, " |")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := &Deque{data: make([]int, 4), capacity: 4}

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(buf.String())
		case "debug":
			fmt.Println(buf.Debug())
		case "size":
			fmt.Println(buf.Len())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushFront(num)
			}
		case "pop_back":
			if err := buf.PopBack(); err != nil {
				fmt.Println(err)
			}
		case "pop_front":
			if err := buf.PopFront(); err != nil {
				fmt.Println(err)
			}
		case "front":
			if val, err := buf.Front(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "back":
			if val, err := buf.Back(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "clear":
			buf.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
