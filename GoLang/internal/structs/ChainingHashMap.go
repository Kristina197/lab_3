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

type HashNode struct {
	key   string
	value string
	next  *HashNode
}

func NewHashNode(key string, value string) *HashNode {
	return &HashNode{
		key:   key,
		value: value,
		next:  nil,
	}
}

type ChainingHash struct {
	buckets  []*HashNode
	size     int
	capacity int
}

func NewChainingHash() *ChainingHash {
	capacity := 10
	return &ChainingHash{
		buckets:  make([]*HashNode, capacity),
		size:     0,
		capacity: capacity,
	}
}

func NewChainingHashWithCapacity(cap int) *ChainingHash {
	if cap <= 0 {
		cap = 10
	}
	return &ChainingHash{
		buckets:  make([]*HashNode, cap),
		size:     0,
		capacity: cap,
	}
}

func (h *ChainingHash) hashFunc(key string) uint32 {
	hasher := fnv.New32a()

	keyStr := fmt.Sprintf("%v", key)
	hasher.Write([]byte(keyStr))

	return hasher.Sum32() % uint32(h.capacity)
}

func (h *ChainingHash) Put(key, value string) bool {
	hash := h.hashFunc(key)
	curr := h.buckets[hash]

	for curr != nil {
		if curr.key == key {
			curr.value = value
			return true
		}
		curr = curr.next
	}

	newNode := NewHashNode(key, value)
	newNode.next = h.buckets[hash]
	h.buckets[hash] = newNode
	h.size++

	return true
}

func (h *ChainingHash) Get(key string) (string, bool) {
	hash := h.hashFunc(key)
	curr := h.buckets[hash]

	for curr != nil {
		if curr.key == key {
			return curr.value, true
		}
		curr = curr.next
	}

	return "", false
}

func (h *ChainingHash) Remove(key string) bool {
	hash := h.hashFunc(key)
	curr := h.buckets[hash]
	var prev *HashNode = nil

	for curr != nil {
		if curr.key == key {
			if prev == nil {
				h.buckets[hash] = curr.next
			} else {
				prev.next = curr.next
			}
			h.size--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func (h *ChainingHash) Contains(key string) bool {
	_, found := h.Get(key)
	return found
}

func (h *ChainingHash) Size() int {
	return h.size
}

func (h *ChainingHash) IsEmpty() bool {
	return h.size == 0
}

func (h *ChainingHash) Capacity() int {
	return h.capacity
}

func (h *ChainingHash) Clear() {
	for i := 0; i < h.capacity; i++ {
		h.buckets[i] = nil
	}
	h.size = 0
}

func (h *ChainingHash) Keys() []string {
	keys := make([]string, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		curr := h.buckets[i]
		for curr != nil {
			keys = append(keys, curr.key)
			curr = curr.next
		}
	}
	return keys
}

func (h *ChainingHash) Values() []string {
	values := make([]string, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		curr := h.buckets[i]
		for curr != nil {
			values = append(values, curr.value)
			curr = curr.next
		}
	}
	return values
}

func (h *ChainingHash) Print() {
	for i := 0; i < h.capacity; i++ {
		fmt.Printf("[%d]: ", i)
		curr := h.buckets[i]
		var pairs []string
		for curr != nil {
			pairs = append(pairs, fmt.Sprintf("{%v: %v}", curr.key, curr.value))
			curr = curr.next
		}
		if len(pairs) == 0 {
			fmt.Println("null")
		} else {
			fmt.Printf("%s -> null\n", strings.Join(pairs, " -> "))
		}
	}
}

func (h *ChainingHash) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("HashTable (size: %d, capacity: %d)\n", h.size, h.capacity))

	for i := 0; i < h.capacity; i++ {
		builder.WriteString(fmt.Sprintf("[%d]: ", i))
		curr := h.buckets[i]
		var pairs []string
		for curr != nil {
			pairs = append(pairs, fmt.Sprintf("{%v: %v}", curr.key, curr.value))
			curr = curr.next
		}
		if len(pairs) == 0 {
			builder.WriteString("null\n")
		} else {
			builder.WriteString(fmt.Sprintf("%s -> null\n", strings.Join(pairs, " -> ")))
		}
	}
	return builder.String()
}

func (h *ChainingHash) LoadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

func (h *ChainingHash) Rehash(newCapacity int) {
	if newCapacity <= 0 {
		return
	}

	oldBuckets := h.buckets
	oldCapacity := h.capacity

	h.buckets = make([]*HashNode, newCapacity)
	h.capacity = newCapacity
	h.size = 0

	for i := 0; i < oldCapacity; i++ {
		curr := oldBuckets[i]
		for curr != nil {
			h.Put(curr.key, curr.value)
			curr = curr.next
		}
	}
}

func (h *ChainingHash) AutoRehash(threshold float64) {
	if h.LoadFactor() > threshold {
		newCapacity := h.capacity * 2
		h.Rehash(newCapacity)
	}
}

// СЕРИАЛИЗАЦИЯ / ДЕСЕРИАЛИЗАЦИЯ

// EntryJSON вспомогательная структура для JSON сериализации
type EntryJSON struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// MarshalJSON реализует интерфейс json.Marshaler
func (h *ChainingHash) MarshalJSON() ([]byte, error) {
	type HashTableJSON struct {
		Capacity   int         `json:"capacity"`
		Size       int         `json:"size"`
		Entries    []EntryJSON `json:"entries"`
		LoadFactor float64     `json:"load_factor"`
	}

	// Собираем все пары ключ-значение
	entries := make([]EntryJSON, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		curr := h.buckets[i]
		for curr != nil {
			entries = append(entries, EntryJSON{
				Key:   curr.key,
				Value: curr.value,
			})
			curr = curr.next
		}
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
func (h *ChainingHash) UnmarshalJSON(data []byte) error {
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

	// Восстанавливаем хеш-таблицу
	if hashTableJSON.Capacity > 0 {
		h.capacity = hashTableJSON.Capacity
	} else {
		h.capacity = 10 // значение по умолчанию
	}

	h.buckets = make([]*HashNode, h.capacity)
	h.size = 0

	// Добавляем все записи
	for _, entry := range hashTableJSON.Entries {
		h.Put(entry.Key, entry.Value)
	}

	return nil
}

// ToJSON сериализует хеш-таблицу в JSON строку
func (h *ChainingHash) ToJSON() (string, error) {
	bytes, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в хеш-таблицу
func (h *ChainingHash) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), h)
}

// SaveToJSON сохраняет хеш-таблицу в JSON файл
func (h *ChainingHash) SaveToJSON(filename string) error {
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
func (h *ChainingHash) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(h)
}

// ToBinary сериализует хеш-таблицу в бинарный формат
func (h *ChainingHash) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	// 1. Записываем capacity и size
	if err := binary.Write(&buf, binary.LittleEndian, int32(h.capacity)); err != nil {
		return nil, fmt.Errorf("ошибка записи capacity: %w", err)
	}
	if err := binary.Write(&buf, binary.LittleEndian, int32(h.size)); err != nil {
		return nil, fmt.Errorf("ошибка записи size: %w", err)
	}

	// 2. Записываем все пары ключ-значение
	for i := 0; i < h.capacity; i++ {
		curr := h.buckets[i]
		for curr != nil {
			// Записываем ключ
			keyBytes := []byte(curr.key)
			keyLen := int32(len(keyBytes))
			if err := binary.Write(&buf, binary.LittleEndian, keyLen); err != nil {
				return nil, fmt.Errorf("ошибка записи длины ключа: %w", err)
			}
			if _, err := buf.Write(keyBytes); err != nil {
				return nil, fmt.Errorf("ошибка записи ключа: %w", err)
			}

			// Записываем значение
			valueBytes := []byte(curr.value)
			valueLen := int32(len(valueBytes))
			if err := binary.Write(&buf, binary.LittleEndian, valueLen); err != nil {
				return nil, fmt.Errorf("ошибка записи длины значения: %w", err)
			}
			if _, err := buf.Write(valueBytes); err != nil {
				return nil, fmt.Errorf("ошибка записи значения: %w", err)
			}

			curr = curr.next
		}
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует хеш-таблицу из бинарного формата
func (h *ChainingHash) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	// 1. Читаем capacity и size
	var capacity, size int32
	if err := binary.Read(buf, binary.LittleEndian, &capacity); err != nil {
		return fmt.Errorf("ошибка чтения capacity: %w", err)
	}
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("ошибка чтения size: %w", err)
	}

	// 2. Инициализируем хеш-таблицу
	if capacity > 0 {
		h.capacity = int(capacity)
	} else {
		h.capacity = 10
	}
	h.buckets = make([]*HashNode, h.capacity)
	h.size = 0

	// 3. Читаем все пары ключ-значение
	for i := 0; i < int(size); i++ {
		// Читаем ключ
		var keyLen int32
		if err := binary.Read(buf, binary.LittleEndian, &keyLen); err != nil {

			return fmt.Errorf("ошибка чтения длины ключа %d: %w", i, err)
		}

		keyBytes := make([]byte, keyLen)
		if _, err := io.ReadFull(buf, keyBytes); err != nil {

			return fmt.Errorf("ошибка чтения ключа %d: %w", i, err)
		}
		key := string(keyBytes)

		// Читаем значение
		var valueLen int32
		if err := binary.Read(buf, binary.LittleEndian, &valueLen); err != nil {
			return fmt.Errorf("ошибка чтения длины значения %d: %w", i, err)
		}

		valueBytes := make([]byte, valueLen)
		if _, err := io.ReadFull(buf, valueBytes); err != nil {
			return fmt.Errorf("ошибка чтения значения %d: %w", i, err)
		}
		value := string(valueBytes)

		h.Put(key, value)
	}

	return nil
}

// SaveToBinary сохраняет хеш-таблицу в бинарный файл
func (h *ChainingHash) SaveToBinary(filename string) error {
	data, err := h.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает хеш-таблицу из бинарного файла
func (h *ChainingHash) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return h.FromBinary(data)
}

// ToBinaryWithBuckets сериализует хеш-таблицу с сохранением структуры бакетов (альтернативный вариант)
func (h *ChainingHash) ToBinaryWithBuckets() ([]byte, error) {
	var buf bytes.Buffer

	// 1. Записываем capacity
	if err := binary.Write(&buf, binary.LittleEndian, int32(h.capacity)); err != nil {
		return nil, fmt.Errorf("ошибка записи capacity: %w", err)
	}

	// 2. Для каждого бакета записываем его цепочку
	for i := 0; i < h.capacity; i++ {
		// Считаем количество элементов в цепочке
		chainLength := 0
		curr := h.buckets[i]
		for curr != nil {
			chainLength++
			curr = curr.next
		}

		// Записываем длину цепочки
		if err := binary.Write(&buf, binary.LittleEndian, int32(chainLength)); err != nil {
			return nil, fmt.Errorf("ошибка записи длины цепочки бакета %d: %w", i, err)
		}

		// Записываем элементы цепочки
		curr = h.buckets[i]
		for curr != nil {
			// Записываем ключ
			keyBytes := []byte(curr.key)
			keyLen := int32(len(keyBytes))
			if err := binary.Write(&buf, binary.LittleEndian, keyLen); err != nil {
				return nil, fmt.Errorf("ошибка записи длины ключа в бакете %d: %w", i, err)
			}
			if _, err := buf.Write(keyBytes); err != nil {
				return nil, fmt.Errorf("ошибка записи ключа в бакете %d: %w", i, err)
			}

			// Записываем значение
			valueBytes := []byte(curr.value)
			valueLen := int32(len(valueBytes))
			if err := binary.Write(&buf, binary.LittleEndian, valueLen); err != nil {
				return nil, fmt.Errorf("ошибка записи длины значения в бакете %d: %w", i, err)
			}
			if _, err := buf.Write(valueBytes); err != nil {
				return nil, fmt.Errorf("ошибка записи значения в бакете %d: %w", i, err)
			}

			curr = curr.next
		}
	}

	return buf.Bytes(), nil
}

// FromBinaryWithBuckets десериализует хеш-таблицу из формата с сохранением структуры бакетов
func (h *ChainingHash) FromBinaryWithBuckets(data []byte) error {
	buf := bytes.NewReader(data)

	// 1. Читаем capacity
	var capacity int32
	if err := binary.Read(buf, binary.LittleEndian, &capacity); err != nil {
		return fmt.Errorf("ошибка чтения capacity: %w", err)
	}

	// 2. Инициализируем хеш-таблицу
	if capacity > 0 {
		h.capacity = int(capacity)
	} else {
		h.capacity = 10
	}
	h.buckets = make([]*HashNode, h.capacity)
	h.size = 0

	// 3. Читаем каждый бакет
	for i := 0; i < h.capacity; i++ {
		var chainLength int32
		if err := binary.Read(buf, binary.LittleEndian, &chainLength); err != nil {

			return fmt.Errorf("ошибка чтения длины цепочки бакета %d: %w", i, err)
		}

		// Восстанавливаем цепочку в обратном порядке
		var head *HashNode
		for j := 0; j < int(chainLength); j++ {
			var keyLen int32
			if err := binary.Read(buf, binary.LittleEndian, &keyLen); err != nil {
				return fmt.Errorf("ошибка чтения длины ключа в бакете %d, элемент %d: %w", i, j, err)
			}

			keyBytes := make([]byte, keyLen)
			if _, err := io.ReadFull(buf, keyBytes); err != nil {
				return fmt.Errorf("ошибка чтения ключа в бакете %d, элемент %d: %w", i, j, err)
			}
			key := string(keyBytes)

			// Читаем значение
			var valueLen int32
			if err := binary.Read(buf, binary.LittleEndian, &valueLen); err != nil {
				return fmt.Errorf("ошибка чтения длины значения в бакете %d, элемент %d: %w", i, j, err)
			}

			valueBytes := make([]byte, valueLen)
			if _, err := io.ReadFull(buf, valueBytes); err != nil {
				return fmt.Errorf("ошибка чтения значения в бакете %d, элемент %d: %w", i, j, err)
			}
			value := string(valueBytes)

			// Создаем узел и добавляем в начало цепочки
			node := NewHashNode(key, value)
			node.next = head
			head = node
			h.size++
		}

		// Сохраняем цепочку в бакет
		h.buckets[i] = head
	}

	return nil
}
