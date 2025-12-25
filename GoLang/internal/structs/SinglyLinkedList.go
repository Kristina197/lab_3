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

type FNode struct {
	val  string
	next *FNode
}

func NewFNode(value string) *FNode {
	return &FNode{
		val:  value,
		next: nil,
	}
}

type SinglyLinkedList struct {
	head *FNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
	}
}

func (s *SinglyLinkedList) InsertAtHead(val string) {
	newNode := NewFNode(val)
	newNode.next = s.head
	s.head = newNode
}

func (s *SinglyLinkedList) InsertAtTail(val string) {
	newNode := NewFNode(val)

	if s.head == nil {
		s.head = newNode
		return
	}

	tmp := s.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = newNode
}

func (s *SinglyLinkedList) InsertBefore(target string, val string) error {
	if s.head == nil {
		return fmt.Errorf("список пустой, вставка невозможна")
	}

	if s.head.val == target {
		s.InsertAtHead(val)
		return nil
	}

	curr := s.head
	for curr.next != nil && curr.next.val != target {
		curr = curr.next
	}

	if curr.next != nil && curr.next.val == target {
		newNode := NewFNode(val)
		newNode.next = curr.next
		curr.next = newNode
		return nil
	}

	return fmt.Errorf("заданное значение '%s' не найдено", target)
}

func (s *SinglyLinkedList) InsertAfter(target string, val string) error {
	if s.head == nil {
		return fmt.Errorf("список пустой, вставка невозможна")
	}

	curr := s.head
	for curr != nil && curr.val != target {
		curr = curr.next
	}

	if curr != nil && curr.val == target {
		newNode := NewFNode(val)
		newNode.next = curr.next
		curr.next = newNode
		return nil
	}

	return fmt.Errorf("заданное значение '%s' не найдено", target)
}

func (s *SinglyLinkedList) DeleteHead() error {
	if s.head == nil {
		return fmt.Errorf("список пуст")
	}

	s.head = s.head.next
	return nil
}

func (s *SinglyLinkedList) DeleteTail() error {
	if s.head == nil {
		return fmt.Errorf("список пуст")
	}

	if s.head.next == nil {
		s.head = nil
		return nil
	}

	curr := s.head
	for curr.next != nil && curr.next.next != nil {
		curr = curr.next
	}

	curr.next = nil
	return nil
}

func (s *SinglyLinkedList) DeleteBefore(target string) error {
	if s.head == nil || s.head.next == nil {
		return fmt.Errorf("недостаточно элементов для удаления")
	}

	if s.head.next.val == target {
		return s.DeleteHead()
	}

	curr := s.head
	for curr.next != nil && curr.next.next != nil && curr.next.next.val != target {
		curr = curr.next
	}

	if curr.next != nil && curr.next.next != nil && curr.next.next.val == target {
		curr.next = curr.next.next
		return nil
	}

	return fmt.Errorf("заданное значение '%s' не найдено", target)
}

func (s *SinglyLinkedList) DeleteAfter(target string) error {
	if s.head == nil {
		return fmt.Errorf("список пустой, удаление невозможно")
	}

	curr := s.head
	for curr != nil && curr.val != target {
		curr = curr.next
	}

	if curr != nil && curr.next != nil {
		curr.next = curr.next.next
		return nil
	}

	if curr != nil && curr.next == nil {
		return fmt.Errorf("нет элемента после '%s'", target)
	}

	return fmt.Errorf("заданное значение '%s' не найдено", target)
}

func (s *SinglyLinkedList) DeleteByValue(target string) error {
	if s.head == nil {
		return fmt.Errorf("список пуст")
	}

	if s.head.val == target {
		return s.DeleteHead()
	}

	curr := s.head
	for curr.next != nil && curr.next.val != target {
		curr = curr.next
	}

	if curr.next != nil && curr.next.val == target {
		curr.next = curr.next.next
		return nil
	}

	return fmt.Errorf("заданное значение '%s' отсутствует в списке", target)
}

func (s *SinglyLinkedList) SearchByValue(target string) int {
	if s.head == nil {
		return -1
	}

	curr := s.head
	targetIndex := 0

	for curr != nil {
		if curr.val == target {
			return targetIndex
		}
		curr = curr.next
		targetIndex++
	}
	return -1
}

func (s *SinglyLinkedList) Length() int {
	count := 0
	curr := s.head

	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (s *SinglyLinkedList) IsEmpty() bool {
	return s.head == nil
}

func (s *SinglyLinkedList) GetHead() string {
	if s.head == nil {
		return ""
	}
	return s.head.val
}

func (s *SinglyLinkedList) Print() {
	if s.head == nil {
		fmt.Println("Список пуст")
		return
	}

	var elements []string
	curr := s.head

	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.next
	}

	fmt.Println(strings.Join(elements, " "))
}

func (s *SinglyLinkedList) String() string {
	if s.head == nil {
		return "Список пуст"
	}

	var elements []string
	curr := s.head

	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.next
	}

	return "head -> " + strings.Join(elements, " -> ") + " -> tail"
}

func (s *SinglyLinkedList) ToSlice() []string {
	var result []string
	curr := s.head

	for curr != nil {
		result = append(result, curr.val)
		curr = curr.next
	}

	return result
}

func (s *SinglyLinkedList) Clear() {
	s.head = nil
}

// MarshalJSON реализует интерфейс json.Marshaler
func (s *SinglyLinkedList) MarshalJSON() ([]byte, error) {
	type LinkedListJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
		Head     string   `json:"head,omitempty"`
		Tail     string   `json:"tail,omitempty"`
	}

	// Получаем все элементы
	elements := s.ToSlice()

	// Получаем хвост (последний элемент)
	var tail string
	if len(elements) > 0 {
		tail = elements[len(elements)-1]
	}

	linkedListJSON := LinkedListJSON{
		Elements: elements,
		Size:     s.Length(),
	}

	if s.head != nil {
		linkedListJSON.Head = s.head.val
		linkedListJSON.Tail = tail
	}

	return json.Marshal(linkedListJSON)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (s *SinglyLinkedList) UnmarshalJSON(data []byte) error {
	type LinkedListJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
		Head     string   `json:"head,omitempty"`
		Tail     string   `json:"tail,omitempty"`
	}

	var linkedListJSON LinkedListJSON
	if err := json.Unmarshal(data, &linkedListJSON); err != nil {
		return err
	}

	// Очищаем список и добавляем элементы по порядку
	s.Clear()
	for _, elem := range linkedListJSON.Elements {
		s.InsertAtTail(elem)
	}

	return nil
}

// ToJSON сериализует список в JSON строку
func (s *SinglyLinkedList) ToJSON() (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в список
func (s *SinglyLinkedList) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), s)
}

// SaveToJSON сохраняет список в JSON файл
func (s *SinglyLinkedList) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(s)
}

// LoadFromJSON загружает список из JSON файла
func (s *SinglyLinkedList) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(s)
}

// ToBinary сериализует список в бинарный формат
func (s *SinglyLinkedList) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	// 1. Записываем размер списка
	size := int32(s.Length())
	if err := binary.Write(&buf, binary.LittleEndian, size); err != nil {
		return nil, fmt.Errorf("ошибка записи размера списка: %w", err)
	}

	// 2. Записываем каждый элемент (в порядке от head к tail)
	current := s.head
	for current != nil {
		strBytes := []byte(current.val)
		strLen := int32(len(strBytes))

		// Записываем длину строки
		if err := binary.Write(&buf, binary.LittleEndian, strLen); err != nil {
			return nil, fmt.Errorf("ошибка записи длины строки: %w", err)
		}

		// Записываем байты строки
		if _, err := buf.Write(strBytes); err != nil {
			return nil, fmt.Errorf("ошибка записи данных строки: %w", err)
		}

		current = current.next
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует список из бинарного формата
func (s *SinglyLinkedList) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	// 1. Считываем размер списка
	var size int32
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("ошибка чтения размера списка: %w", err)
	}

	// 2. Очищаем текущий список
	s.Clear()

	// 3. Считываем каждый элемент
	for i := 0; i < int(size); i++ {
		// Считываем длину строки
		var strLen int32
		if err := binary.Read(buf, binary.LittleEndian, &strLen); err != nil {
			return fmt.Errorf("ошибка чтения длины строки %d: %w", i, err)
		}

		// Считываем байты строки
		strBytes := make([]byte, strLen)
		if _, err := io.ReadFull(buf, strBytes); err != nil {
			return fmt.Errorf("ошибка чтения данных строки %d: %w", i, err)
		}

		// Добавляем элемент в конец списка
		s.InsertAtTail(string(strBytes))
	}

	return nil
}

// SaveToBinary сохраняет список в бинарный файл
func (s *SinglyLinkedList) SaveToBinary(filename string) error {
	data, err := s.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает список из бинарного файла
func (s *SinglyLinkedList) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return s.FromBinary(data)
}
