package structs

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strings"
)

type HashStatus int

const (
	EMPTY HashStatus = iota
	DELETED
	TAKEN
)

type HashEntry struct {
	key    string
	value  string
	status HashStatus
}

func NewHashEntry() *HashEntry {
	return &HashEntry{
		status: EMPTY,
	}
}

type OpenAddrHash struct {
	buckets  []*HashEntry
	size     int
	capacity int
}

func NewOpenAddrHash() *OpenAddrHash {
	capacity := 10
	buckets := make([]*HashEntry, capacity)
	for i := range buckets {
		buckets[i] = NewHashEntry()
	}

	return &OpenAddrHash{
		buckets:  buckets,
		size:     0,
		capacity: capacity,
	}
}

func NewOpenAddrHashWithCapacity(cap int) *OpenAddrHash {
	if cap <= 0 {
		cap = 10
	}

	buckets := make([]*HashEntry, cap)
	for i := range buckets {
		buckets[i] = NewHashEntry()
	}

	return &OpenAddrHash{
		buckets:  buckets,
		size:     0,
		capacity: cap,
	}
}

func (h *OpenAddrHash) hashFunc(key string) uint32 {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return hasher.Sum32() % uint32(h.capacity)
}

func (h *OpenAddrHash) Put(key, value string) bool {
	if h.size >= h.capacity {
		fmt.Println("Хеш-таблица переполнена")
		return false
	}

	hash := h.hashFunc(key)
	deletedIndex := -1
	i := 0

	for i < h.capacity {
		index := (int(hash) + i) % h.capacity

		if h.buckets[index].status == EMPTY {
			h.buckets[index].key = key
			h.buckets[index].value = value
			h.buckets[index].status = TAKEN
			h.size++
			return true
		}

		if h.buckets[index].status == TAKEN && h.buckets[index].key == key {
			h.buckets[index].value = value
			return true
		}

		if deletedIndex == -1 && h.buckets[index].status == DELETED {
			deletedIndex = index
		}
		i++
	}

	if deletedIndex != -1 {
		h.buckets[deletedIndex].status = TAKEN
		h.buckets[deletedIndex].key = key
		h.buckets[deletedIndex].value = value
		h.size++
		return true
	}

	return false
}

func (h *OpenAddrHash) Get(key string) (string, bool) {
	hash := h.hashFunc(key)
	i := 0

	for i < h.capacity {
		index := (int(hash) + i) % h.capacity

		if h.buckets[index].status == EMPTY {
			return "", false
		}

		if h.buckets[index].status == DELETED {
			i++
			continue
		}

		if h.buckets[index].status == TAKEN && h.buckets[index].key == key {
			return h.buckets[index].value, true
		}
		i++
	}

	return "", false
}

func (h *OpenAddrHash) Remove(key string) bool {
	hash := h.hashFunc(key)
	i := 0

	for i < h.capacity {
		index := (int(hash) + i) % h.capacity

		if h.buckets[index].status == EMPTY {
			return false
		}

		if h.buckets[index].status == TAKEN && h.buckets[index].key == key {
			h.buckets[index].status = DELETED
			h.buckets[index].key = ""
			h.buckets[index].value = ""
			h.size--
			return true
		}
		i++
	}

	return false
}

func (h *OpenAddrHash) Contains(key string) bool {
	_, found := h.Get(key)
	return found
}

func (h *OpenAddrHash) Size() int {
	return h.size
}

func (h *OpenAddrHash) Capacity() int {
	return h.capacity
}

func (h *OpenAddrHash) IsEmpty() bool {
	return h.size == 0
}

func (h *OpenAddrHash) IsFull() bool {
	return h.size >= h.capacity
}

func (h *OpenAddrHash) LoadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

func (h *OpenAddrHash) Clear() {
	for i := 0; i < h.capacity; i++ {
		h.buckets[i].status = EMPTY
		h.buckets[i].key = ""
		h.buckets[i].value = ""
	}
	h.size = 0
}

func (h *OpenAddrHash) Keys() []string {
	keys := make([]string, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		if h.buckets[i].status == TAKEN {
			keys = append(keys, h.buckets[i].key)
		}
	}
	return keys
}

func (h *OpenAddrHash) Values() []string {
	values := make([]string, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		if h.buckets[i].status == TAKEN {
			values = append(values, h.buckets[i].value)
		}
	}
	return values
}

func (h *OpenAddrHash) Print() {
	for i := 0; i < h.capacity; i++ {
		fmt.Printf("[%d]: ", i)

		switch h.buckets[i].status {
		case TAKEN:
			fmt.Printf("{%s: %s}\n", h.buckets[i].key, h.buckets[i].value)
		case DELETED:
			fmt.Println("DELETED")
		case EMPTY:
			fmt.Println("null")
		}
	}
	fmt.Println()
}

func (h *OpenAddrHash) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("OpenAddrHash (size: %d, capacity: %d, load factor: %.2f)\n",
		h.size, h.capacity, h.LoadFactor()))

	for i := 0; i < h.capacity; i++ {
		builder.WriteString(fmt.Sprintf("[%d]: ", i))

		switch h.buckets[i].status {
		case TAKEN:
			builder.WriteString(fmt.Sprintf("{%s: %s}\n", h.buckets[i].key, h.buckets[i].value))
		case DELETED:
			builder.WriteString("DELETED\n")
		case EMPTY:
			builder.WriteString("null\n")
		}
	}
	return builder.String()
}

func (h *OpenAddrHash) Rehash(newCapacity int) {
	if newCapacity <= h.capacity {
		return
	}

	oldBuckets := h.buckets
	oldCapacity := h.capacity

	h.buckets = make([]*HashEntry, newCapacity)
	for i := range h.buckets {
		h.buckets[i] = NewHashEntry()
	}
	h.capacity = newCapacity
	h.size = 0

	for i := 0; i < oldCapacity; i++ {
		if oldBuckets[i].status == TAKEN {
			h.Put(oldBuckets[i].key, oldBuckets[i].value)
		}
	}
}

func (h *OpenAddrHash) AutoRehash(threshold float64) {
	if h.LoadFactor() > threshold {
		newCapacity := h.capacity * 2
		h.Rehash(newCapacity)
	}
}

// СЕРИАЛИЗАЦИЯ / ДЕСЕРИАЛИЗАЦИЯ

// MarshalJSON реализует интерфейс json.Marshaler
func (h *OpenAddrHash) MarshalJSON() ([]byte, error) {
	type EntryJSON struct {
		Key    string `json:"key"`
		Value  string `json:"value"`
		Status int    `json:"status"`
	}

	type HashTableJSON struct {
		Capacity   int         `json:"capacity"`
		Size       int         `json:"size"`
		Entries    []EntryJSON `json:"entries"`
		LoadFactor float64     `json:"load_factor"`
	}

	entries := make([]EntryJSON, 0, h.capacity)
	for i := 0; i < h.capacity; i++ {
		entries = append(entries, EntryJSON{
			Key:    h.buckets[i].key,
			Value:  h.buckets[i].value,
			Status: int(h.buckets[i].status),
		})
	}

	hashTableJSON := HashTableJSON{
		Capacity:   h.capacity,
		Size:       h.size,
		Entries:    entries,
		LoadFactor: h.LoadFactor(),
	}

	return json.Marshal(hashTableJSON)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (h *OpenAddrHash) UnmarshalJSON(data []byte) error {
	type EntryJSON struct {
		Key    string `json:"key"`
		Value  string `json:"value"`
		Status int    `json:"status"`
	}

	type HashTableJSON struct {
		Capacity   int         `json:"capacity"`
		Size       int         `json:"size"`
		Entries    []EntryJSON `json:"entries"`
		LoadFactor float64     `json:"load_factor"`
	}

	var hashTableJSON HashTableJSON
	if err := json.Unmarshal(data, &hashTableJSON); err != nil {
		return err
	}

	h.capacity = hashTableJSON.Capacity
	h.size = 0
	h.buckets = make([]*HashEntry, h.capacity)

	for i := range h.buckets {
		h.buckets[i] = NewHashEntry()
	}

	for i, entry := range hashTableJSON.Entries {
		h.buckets[i].key = entry.Key
		h.buckets[i].value = entry.Value
		h.buckets[i].status = HashStatus(entry.Status)

		if h.buckets[i].status == TAKEN {
			h.size++
		}
	}

	return nil
}

// ToJSON сериализует хеш-таблицу в JSON строку
func (h *OpenAddrHash) ToJSON() (string, error) {
	bytes, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в хеш-таблицу
func (h *OpenAddrHash) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), h)
}

// SaveToJSON сохраняет хеш-таблицу в JSON файл
func (h *OpenAddrHash) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(h)
}

// LoadFromJSON загружает хеш-таблицу из JSON файла
func (h *OpenAddrHash) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(h)
}

// ToBinary сериализует хеш-таблицу в бинарный формат
func (h *OpenAddrHash) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, int32(h.capacity)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, int32(h.size)); err != nil {
		return nil, err
	}

	for i := 0; i < h.capacity; i++ {
		entry := h.buckets[i]

		if err := binary.Write(&buf, binary.LittleEndian, int32(entry.status)); err != nil {
			return nil, err
		}

		keyBytes := []byte(entry.key)
		keyLen := int32(len(keyBytes))
		if err := binary.Write(&buf, binary.LittleEndian, keyLen); err != nil {
			return nil, err
		}
		if _, err := buf.Write(keyBytes); err != nil {
			return nil, err
		}

		valueBytes := []byte(entry.value)
		valueLen := int32(len(valueBytes))
		if err := binary.Write(&buf, binary.LittleEndian, valueLen); err != nil {
			return nil, err
		}
		if _, err := buf.Write(valueBytes); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует хеш-таблицу из бинарного формата
func (h *OpenAddrHash) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	var capacity, size int32
	if err := binary.Read(buf, binary.LittleEndian, &capacity); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return err
	}

	h.capacity = int(capacity)
	h.size = 0
	h.buckets = make([]*HashEntry, h.capacity)

	for i := 0; i < h.capacity; i++ {
		h.buckets[i] = NewHashEntry()

		var status int32
		if err := binary.Read(buf, binary.LittleEndian, &status); err != nil {
			return err
		}
		h.buckets[i].status = HashStatus(status)

		var keyLen int32
		if err := binary.Read(buf, binary.LittleEndian, &keyLen); err != nil {
			return err
		}
		keyBytes := make([]byte, keyLen)
		if _, err := io.ReadFull(buf, keyBytes); err != nil {
			return err
		}
		h.buckets[i].key = string(keyBytes)

		var valueLen int32
		if err := binary.Read(buf, binary.LittleEndian, &valueLen); err != nil {
			return err
		}
		valueBytes := make([]byte, valueLen)
		if _, err := io.ReadFull(buf, valueBytes); err != nil {
			return err
		}
		h.buckets[i].value = string(valueBytes)

		if h.buckets[i].status == TAKEN {
			h.size++
		}
	}

	return nil
}

// SaveToBinary сохраняет хеш-таблицу в бинарный файл
func (h *OpenAddrHash) SaveToBinary(filename string) error {
	data, err := h.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает хеш-таблицу из бинарного файла
func (h *OpenAddrHash) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return h.FromBinary(data)
}
