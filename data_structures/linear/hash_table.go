package linear

// Size of the hashtable array
const arraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [arraySize]*bucket
}

// bucket is a linked list in each slot of the HashTable
type bucket struct {
	head *bucketNode
}

// bucket node is a node in the bucket linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

func (h *HashTable) Insert(key string) {
	hashKey := hash(key)

	h.array[hashKey].insert(key)
}

func (h *HashTable) Search(key string) bool {
	hashKey := hash(key)

	return h.array[hashKey].search(key)
}

func (h *HashTable) Delete(key string) {
	hashKey := hash(key)

	h.array[hashKey].delete(key)
}

func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(key string) bool {
	currentHead := b.head
	for currentHead != nil {
		if currentHead.key == key {
			return true
		}
		currentHead = currentHead.next
	}
	return false
}

func (b *bucket) delete(key string) {
	currentHead := b.head

	if currentHead.key == key {
		b.head = currentHead.next
		return
	}

	//3 -> 4 -> 5 -> 6
	for currentHead.next != nil {
		if currentHead.next.key == key {
			currentHead.next = currentHead.next.next
			return
		}
		currentHead = currentHead.next
	}
}

func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % arraySize
}

// Init will create a bucket in each slot in the hash table
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}
