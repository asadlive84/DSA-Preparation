package main

import "fmt"

type KeyValue struct {
	key   string
	value string
}

type HashTable struct {
	buckets []*KeyValue
	Size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*KeyValue, size),
		Size:    size,
	}
}

func (h *HashTable) HasFunction(key string) int {
	hash := 0
	for _, ch := range key {
		hash += int(ch)
	}
	return hash % h.Size
}

func (h *HashTable) Insert(key, value string) {
	// O(1)
	index := h.HasFunction(key)

	if h.buckets[index] != nil {
		if h.buckets[index].key == key {
			h.buckets[index].value = value
			return
		}

		index = (index + 1) % h.Size
	}

	h.buckets[index] = &KeyValue{
		key:   key,
		value: value,
	}

}

func (h *HashTable) Get(key string) (string, error) {
	// O(n)

	index := h.HasFunction(key)
	for h.buckets[index] != nil {
		if h.buckets[index].key == key {
			return h.buckets[index].value, nil
		}
		index = (index + 1) % h.Size
	}
	return "", fmt.Errorf("key %s is not found", key)
}

func (h *HashTable) Remove(key string) error {
	// O(n)

	index := h.HasFunction(key)
	for h.buckets[index] != nil {
		if h.buckets[index].key == key {
			h.buckets[index] = nil
			return nil
		}
		index = (index + 1) % h.Size
	}
	return fmt.Errorf("key %s is not found", key)
}

func (h *HashTable) Show() {
	// O(n)

	for i, kv := range h.buckets {
		if kv != nil {
			fmt.Printf("Bucket %d : %s ---> %s \n", i, kv.key, kv.value)
		} else {
			fmt.Printf("Bucket is empty\n")
		}
	}
}

func main() {
	ht := NewHashTable(1)

	// Insert some key-value pairs
	ht.Insert("name", "John")
	ht.Insert("age", "25")
	ht.Insert("city", "New York")

	// Display the hash table
	ht.Show()

	fmt.Println("========GET data==========")

	getData, err := ht.Get("city")

	fmt.Printf("err %+v\n", err)
	fmt.Printf("getData %+v\n", getData)

	fmt.Println("========Remove==========")

	err = ht.Remove("city")
	fmt.Printf("err %+v\n", err)
	fmt.Println("========Show data==========")
	// Display the hash table
	ht.Show()

}
