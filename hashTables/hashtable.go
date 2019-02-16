package hashTable

import "fmt"

type HashTable struct {
	Size  int
	Count int
	Table map[string]string
}

func CreateHashTable(size int) *HashTable {
	return &HashTable{Size: size, Table: make(map[string]string, size), Count: 0}
}

func (ht *HashTable) Put(key string, val string) bool {
	if ht.Size == ht.Count {
		return false
	}

	if ht.Table[key] == "" {
		ht.Table[key] = val
		ht.Count++
		return true
	}
	return false
}

func (ht *HashTable) Remove(key string) bool {
	if ht.Table[key] != "" {
		delete(ht.Table, key)
		ht.Count--
		return true
	}
	return false
}

func (ht *HashTable) Contains(key string) (bool, string) {
	if ht.Table[key] != "" {
		return true, ht.Table[key]
	}
	return false, ""
}

func (ht *HashTable) Update(key string, value string) bool {
	if ht.Table[key] != "" {
		ht.Table[key] = value
	}
	return false
}

func (ht *HashTable) Print() {
	for key, value := range ht.Table {
		fmt.Printf("%v\t\t%v\n", key, value)
	}
	fmt.Println()
}
