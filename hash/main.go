package main

import "fmt"

type set struct {
	key  int
	data int
}

const capacity = 10

var size = 0

var array = new([10]set)

func init_array() {

	for i := 0; i < capacity; i++ {
		array[i].key = 0
		array[i].data = 0
	}
}

func hashFunction(key int) int {

	return (key % capacity)
}

func insert(key, value int) {

	index := hashFunction(key)

	if array[index].key == 0 {

		array[index].key = key
		array[index].data = value
		size++
		fmt.Printf("Key (%d) has been inserted\n", key)
	} else if array[index].key == key {

		array[index].data = value
	} else {

		fmt.Println("Collision occured")
	}
}

func removeElement(key int) {

	index := hashFunction(key)

	if array[index].data == 0 {

		fmt.Println("This key does not exits")
	} else {

		array[index].key = 0
		array[index].data = 0
		size--
		fmt.Printf("Key (%d) has been removed", key)
	}
}

func sizeHashTable() int {

	return size
}

func Display() {

	for i := 0; i < capacity; i++ {

		if array[i].data == 0 {
			fmt.Printf("\narray[%d]:  /", i)
		} else {
			fmt.Printf("\nKey: %d array[%d]: %d\t", array[i].key, i, array[i].data)
		}
	}
}


func main() {

	var key, data, choice int

	init_array()

	c := 1

	for c != 0 {
		fmt.Println("1. Insert item in the Hash Table")
		fmt.Println("2. Remove item from in the Hash Table")
		fmt.Println("3. Check the size of Hash Table")
		fmt.Println("4. Display Hash Table")
		fmt.Printf("Please Enter you choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Enter key: ")
			fmt.Scan(&key)
			fmt.Printf("Enter key: ")
			fmt.Scan(&data)

			insert(key, data)
		case 2:
			fmt.Printf("Enter the key Remove: ")
			fmt.Scan(&key)

			removeElement(key)
		case 3:
			n := sizeHashTable()
			fmt.Println("Size Hash Table: ", n)

		case 4:
			Display()
		default:
			fmt.Println("Invalid input")
		}

		fmt.Printf("\nDo you want to continue (press 1 for yes):")
		fmt.Scanf("%d", &c)
	}

}
