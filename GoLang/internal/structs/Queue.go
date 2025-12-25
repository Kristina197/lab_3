package structs

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type QNode struct {
	data string
	next *QNode
}

func NewQNodeWithValue(value string) *QNode {
	return &QNode{
		data: value,
		next: nil,
	}
}

type Queue struct {
	front *QNode
	rear  *QNode
	size  int
}

func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Enqueue(value string) {
	newNode := NewQNodeWithValue(value)

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *Queue) Dequeue() string {
	if q.front == nil {
		return ""
	}

	dequeuedElem := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.size--
	return dequeuedElem
}

func (q *Queue) Peek() string {
	if q.front == nil {
		return ""
	}
	return q.front.data
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Clear() {
	for q.front != nil {
		q.Dequeue()
	}
}

func (q *Queue) Print() {
	if q.front == nil {
		fmt.Println("Очередь пустая")
		return
	}

	current := q.front
	var elements []string

	for current != nil {
		elements = append(elements, current.data)
		current = current.next
	}

	fmt.Println(strings.Join(elements, " "))
}

func (q *Queue) String() string {
	if q.front == nil {
		return "Очередь пустая"
	}

	current := q.front
	var elements []string

	for current != nil {
		elements = append(elements, current.data)
		current = current.next
	}

	return "front -> " + strings.Join(elements, " -> ") + " -> rear"
}

// ToSlice преобразует очередь в слайс (для сериализации)
func (q *Queue) ToSlice() []string {
	result := make([]string, 0, q.size)
	current := q.front

	for current != nil {
		result = append(result, current.data)
		current = current.next
	}

	return result
}

// FromSlice создает очередь из слайса
func (q *Queue) FromSlice(elements []string) {
	q.Clear() // Очищаем текущую очередь
	for _, elem := range elements {
		q.Enqueue(elem)
	}
}

// MarshalJSON реализует интерфейс json.Marshaler
func (q *Queue) MarshalJSON() ([]byte, error) {
	type QueueJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
		Front    string   `json:"front,omitempty"`
		Rear     string   `json:"rear,omitempty"`
	}

	queueJSON := QueueJSON{
		Elements: q.ToSlice(),
		Size:     q.size,
	}

	if q.front != nil {
		queueJSON.Front = q.front.data
	}
	if q.rear != nil {
		queueJSON.Rear = q.rear.data
	}

	return json.Marshal(queueJSON)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (q *Queue) UnmarshalJSON(data []byte) error {
	type QueueJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
		Front    string   `json:"front,omitempty"`
		Rear     string   `json:"rear,omitempty"`
	}

	var queueJSON QueueJSON
	if err := json.Unmarshal(data, &queueJSON); err != nil {
		return err
	}

	q.FromSlice(queueJSON.Elements)
	return nil
}

// ToJSON сериализует очередь в JSON строку
func (q *Queue) ToJSON() (string, error) {
	bytes, err := json.Marshal(q)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в очередь
func (q *Queue) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), q)
}

// SaveToJSON сохраняет очередь в JSON файл
func (q *Queue) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(q)
}

// LoadFromJSON загружает очередь из JSON файла
func (q *Queue) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(q)
}

// ToBinary сериализует очередь в бинарный формат
func (q *Queue) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	size := int32(q.size)
	if err := binary.Write(&buf, binary.LittleEndian, size); err != nil {
		return nil, fmt.Errorf("ошибка записи размера очереди: %w", err)
	}

	current := q.front
	for current != nil {
		strBytes := []byte(current.data)
		strLen := int32(len(strBytes))

		if err := binary.Write(&buf, binary.LittleEndian, strLen); err != nil {
			return nil, fmt.Errorf("ошибка записи длины строки: %w", err)
		}

		if _, err := buf.Write(strBytes); err != nil {
			return nil, fmt.Errorf("ошибка записи данных строки: %w", err)
		}

		current = current.next
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует очередь из бинарного формата
func (q *Queue) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	var size int32
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("ошибка чтения размера очереди: %w", err)
	}

	q.Clear()

	for i := 0; i < int(size); i++ {
		var strLen int32
		if err := binary.Read(buf, binary.LittleEndian, &strLen); err != nil {
			return fmt.Errorf("ошибка чтения длины строки %d: %w", i, err)
		}

		strBytes := make([]byte, strLen)
		if _, err := io.ReadFull(buf, strBytes); err != nil {
			return fmt.Errorf("ошибка чтения данных строки %d: %w", i, err)
		}

		q.Enqueue(string(strBytes))
	}

	return nil
}

// SaveToBinary сохраняет очередь в бинарный файл
func (q *Queue) SaveToBinary(filename string) error {
	data, err := q.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает очередь из бинарного файла
func (q *Queue) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return q.FromBinary(data)
}
