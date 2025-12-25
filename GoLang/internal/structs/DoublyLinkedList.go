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

type LNode struct {
	val  string
	prev *LNode
	next *LNode
}

func NewLNode(value string) *LNode {
	return &LNode{
		val:  value,
		prev: nil,
		next: nil,
	}
}

type DoublyLinkedList struct {
	head *LNode
	tail *LNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList) searchNode(val string) *LNode {
	if d.head == nil {
		return nil
	}

	tmp := d.head
	for tmp != nil {
		if tmp.val == val {
			return tmp
		}
		tmp = tmp.next
	}
	return nil
}

func (d *DoublyLinkedList) InsertAtHead(val string) {
	newNode := NewLNode(val)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	newNode.next = d.head
	d.head.prev = newNode
	d.head = newNode
}

func (d *DoublyLinkedList) InsertAtTail(val string) {
	newNode := NewLNode(val)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

func (d *DoublyLinkedList) InsertBefore(target string, val string) error {
	targetNode := d.searchNode(target)
	if targetNode == nil {
		return fmt.Errorf("заданное значение '%s' отсутствует в списке", target)
	}

	if targetNode == d.head {
		d.InsertAtHead(val)
		return nil
	}

	newNode := NewLNode(val)
	newNode.prev = targetNode.prev
	newNode.next = targetNode
	targetNode.prev.next = newNode
	targetNode.prev = newNode

	return nil
}

func (d *DoublyLinkedList) InsertAfter(target string, val string) error {
	targetNode := d.searchNode(target)
	if targetNode == nil {
		return fmt.Errorf("заданное значение '%s' отсутствует в списке", target)
	}

	newNode := NewLNode(val)
	newNode.prev = targetNode
	newNode.next = targetNode.next

	if targetNode.next != nil {
		targetNode.next.prev = newNode
	} else {
		d.tail = newNode
	}
	targetNode.next = newNode

	return nil
}

func (d *DoublyLinkedList) DeleteHead() error {
	if d.head == nil {
		return fmt.Errorf("список пуст")
	}

	d.head = d.head.next

	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}

	return nil
}

func (d *DoublyLinkedList) DeleteTail() error {
	if d.head == nil {
		return fmt.Errorf("список пуст")
	}

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.prev
		d.tail.next = nil
	}

	return nil
}

func (d *DoublyLinkedList) DeleteBefore(target string) error {
	targetNode := d.searchNode(target)
	if targetNode == nil {
		return fmt.Errorf("заданное значение '%s' отсутствует в списке", target)
	}
	if targetNode.prev == nil {
		return fmt.Errorf("перед заданным значением '%s' отсутствуют узлы. Удаление невозможно", target)
	}

	deleteNode := targetNode.prev
	if deleteNode == d.head {
		return d.DeleteHead()
	}

	targetNode.prev = deleteNode.prev
	deleteNode.prev.next = targetNode

	return nil
}

func (d *DoublyLinkedList) DeleteAfter(target string) error {
	targetNode := d.searchNode(target)
	if targetNode == nil {
		return fmt.Errorf("заданное значение '%s' отсутствует в списке", target)
	}
	if targetNode.next == nil {
		return fmt.Errorf("после заданного значения '%s' отсутствуют узлы. Удаление невозможно", target)
	}

	deleteNode := targetNode.next
	if deleteNode == d.tail {
		return d.DeleteTail()
	}

	targetNode.next = deleteNode.next
	deleteNode.next.prev = targetNode

	return nil
}

func (d *DoublyLinkedList) DeleteByValue(target string) error {
	if d.head == nil {
		return fmt.Errorf("список пустой, удалять нечего")
	}

	targetNode := d.searchNode(target)
	if targetNode == nil {
		return fmt.Errorf("заданное значение '%s' отсутствует в списке", target)
	}

	if targetNode == d.head {
		return d.DeleteHead()
	}
	if targetNode == d.tail {
		return d.DeleteTail()
	}

	targetNode.prev.next = targetNode.next
	targetNode.next.prev = targetNode.prev

	return nil
}

func (d *DoublyLinkedList) SearchByValue(target string) int {
	curr := d.head
	searchIndex := 0

	for curr != nil {
		if curr.val == target {
			return searchIndex
		}
		curr = curr.next
		searchIndex++
	}

	return -1
}

func (d *DoublyLinkedList) Length() int {
	count := 0
	curr := d.head

	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.head == nil
}

func (d *DoublyLinkedList) GetHead() string {
	if d.head == nil {
		return ""
	}
	return d.head.val
}

func (d *DoublyLinkedList) GetTail() string {
	if d.tail == nil {
		return ""
	}
	return d.tail.val
}

func (d *DoublyLinkedList) Print() {
	if d.head == nil {
		fmt.Println("Список пуст")
		return
	}

	var elements []string
	curr := d.head

	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.next
	}

	fmt.Println(strings.Join(elements, " "))
}

func (d *DoublyLinkedList) PrintReverse() {
	if d.head == nil {
		fmt.Println("Список пуст")
		return
	}

	var elements []string
	curr := d.tail

	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.prev
	}

	fmt.Println(strings.Join(elements, " "))
}

func (d *DoublyLinkedList) String() string {
	if d.head == nil {
		return "Список пуст"
	}

	var elements []string
	curr := d.head

	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.next
	}

	return "head -> " + strings.Join(elements, " <-> ") + " -> tail"
}

func (d *DoublyLinkedList) StringReverse() string {
	if d.head == nil {
		return "Список пуст"
	}

	var elements []string
	curr := d.tail

	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.prev
	}

	return "tail -> " + strings.Join(elements, " <-> ") + " -> head"
}

func (d *DoublyLinkedList) ToSlice() []string {
	var result []string
	curr := d.head

	for curr != nil {
		result = append(result, curr.val)
		curr = curr.next
	}

	return result
}

func (d *DoublyLinkedList) ToSliceReverse() []string {
	var result []string
	curr := d.tail

	for curr != nil {
		result = append(result, curr.val)
		curr = curr.prev
	}

	return result
}

func (d *DoublyLinkedList) Clear() {
	d.head = nil
	d.tail = nil
}

// MarshalJSON реализует интерфейс json.Marshaler
func (d *DoublyLinkedList) MarshalJSON() ([]byte, error) {
	type LinkedListJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
		Head     string   `json:"head,omitempty"`
		Tail     string   `json:"tail,omitempty"`
	}

	elements := d.ToSlice()
	linkedListJSON := LinkedListJSON{
		Elements: elements,
		Size:     d.Length(),
	}

	if d.head != nil {
		linkedListJSON.Head = d.head.val
	}
	if d.tail != nil {
		linkedListJSON.Tail = d.tail.val
	}

	return json.Marshal(linkedListJSON)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (d *DoublyLinkedList) UnmarshalJSON(data []byte) error {
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

	d.Clear()
	for _, elem := range linkedListJSON.Elements {
		d.InsertAtTail(elem)
	}

	return nil
}

// ToJSON сериализует список в JSON строку
func (d *DoublyLinkedList) ToJSON() (string, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в список
func (d *DoublyLinkedList) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), d)
}

// SaveToJSON сохраняет список в JSON файл
func (d *DoublyLinkedList) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(d)
}

// LoadFromJSON загружает список из JSON файла
func (d *DoublyLinkedList) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(d)
}

// ToBinary сериализует список в бинарный формат
func (d *DoublyLinkedList) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	// 1. Записываем размер списка
	size := int32(d.Length())
	if err := binary.Write(&buf, binary.LittleEndian, size); err != nil {
		return nil, fmt.Errorf("ошибка записи размера списка: %w", err)
	}

	// 2. Записываем каждый элемент (в порядке от head к tail)
	current := d.head
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
func (d *DoublyLinkedList) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	// 1. Считываем размер списка
	var size int32
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("ошибка чтения размера списка: %w", err)
	}

	// 2. Очищаем текущий список
	d.Clear()

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
		d.InsertAtTail(string(strBytes))
	}

	return nil
}

// SaveToBinary сохраняет список в бинарный файл
func (d *DoublyLinkedList) SaveToBinary(filename string) error {
	data, err := d.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает список из бинарного файла
func (d *DoublyLinkedList) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return d.FromBinary(data)
}
