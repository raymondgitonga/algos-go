package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable_Insert(t *testing.T) {

	type testCase struct {
		id     string
		name   string
		hash   *HashTable
		result bool
	}
	hashTable := Init()
	list := []string{"ERIC", "KENNY"}
	for _, v := range list {
		hashTable.Insert(v)
	}

	testCases := []testCase{
		{
			"ERIC exists",
			"ERIC",
			hashTable,
			true,
		},
		{
			"KENNY exists",
			"KENNY",
			hashTable,
			true,
		},
		{
			"JOHN exists",
			"JOHN",
			hashTable,
			false,
		},
	}

	for _, v := range testCases {
		t.Run(v.id, func(t *testing.T) {
			exists := hashTable.Search(v.name)

			assert.Equal(t, v.result, exists)
		})
	}

}

func TestHashTable_Search(t *testing.T) {
	type testCase struct {
		id     string
		name   string
		hash   *HashTable
		result bool
	}
	hashTable := Init()
	list := []string{"ERIC", "KENNY", "KYLE", "STAN", "RANDY", "BUTTERS", "TOKEN"}
	for _, v := range list {
		hashTable.Insert(v)
	}

	testCases := []testCase{
		{
			"ERIC exists",
			"ERIC",
			hashTable,
			true,
		},
		{
			"JOHN exists",
			"JOHN",
			hashTable,
			false,
		},
	}

	for _, v := range testCases {
		t.Run(v.id, func(t *testing.T) {
			exists := hashTable.Search(v.name)

			assert.Equal(t, v.result, exists)
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	type testCase struct {
		id     string
		name   string
		hash   *HashTable
		result bool
	}
	hashTable := Init()
	list := []string{"ERIC", "KENNY", "KYLE"}
	for _, v := range list {
		hashTable.Insert(v)
	}

	testCases := []testCase{
		{
			"ERIC exists",
			"ERIC",
			hashTable,
			true,
		},
		{
			"KENNY exists",
			"KENNY",
			hashTable,
			false,
		},
	}

	for _, v := range testCases {
		t.Run(v.id, func(t *testing.T) {
			hashTable.Delete("KENNY")
			exists := hashTable.Search(v.name)
			assert.Equal(t, v.result, exists)
		})
	}

}

func TestHashTable_Hash(t *testing.T) {
	name := "TOSH"
	hashValue := hash(name)
	assert.Equal(t, 3, hashValue)
}
