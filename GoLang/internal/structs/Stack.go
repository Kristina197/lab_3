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

type SNode struct {
	val  string
	next *SNode
}

func NewSNode(value string) *SNode {
	return &SNode{
		val:  value,
		next: nil,
	}
}

type Stack struct {
	head *SNode
}

func NewStack() *Stack {
	return &Stack{
		head: nil,
	}
}

func NewStackWithCapacity(capacity int) *Stack {
	return &Stack{
		head: nil,
	}
}

func (s *Stack) Push(val string) {
	newNode := NewSNode(val)
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Pop() string {
	if s.head == nil {
		return ""
	}

	deleteNode := s.head
	poppedElem := deleteNode.val

	s.head = s.head.next

	return poppedElem
}

func (s *Stack) Top() string {
	if s.head == nil {
		return ""
	}
	return s.head.val
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Print() {
	curr := s.head

	fmt.Print("top -> ")
	for curr != nil {
		fmt.Print(curr.val)
		if curr.next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.next
	}
	fmt.Println(" -> bottom")
}

func (s *Stack) String() string {
	if s.head == nil {
		return "top -> bottom"
	}

	var elements []string
	curr := s.head
	for curr != nil {
		elements = append(elements, curr.val)
		curr = curr.next
	}

	return "top -> " + strings.Join(elements, " -> ") + " -> bottom"
}

// Size возвращает количество элементов в стеке
func (s *Stack) Size() int {
	count := 0
	curr := s.head
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

// ToSlice преобразует стек в слайс (для сериализации)
func (s *Stack) ToSlice() []string {
	var result []string
	curr := s.head
	for curr != nil {
		result = append(result, curr.val)
		curr = curr.next
	}
	return result
}

// FromSlice создает стек из слайса
func (s *Stack) FromSlice(elements []string) {
	s.head = nil
	for i := len(elements) - 1; i >= 0; i-- {
		s.Push(elements[i])
	}
}

// MarshalJSON реализует интерфейс json.Marshaler
func (s *Stack) MarshalJSON() ([]byte, error) {
	type StackJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
	}

	stackJSON := StackJSON{
		Elements: s.ToSlice(),
		Size:     s.Size(),
	}

	return json.Marshal(stackJSON)
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (s *Stack) UnmarshalJSON(data []byte) error {
	type StackJSON struct {
		Elements []string `json:"elements"`
		Size     int      `json:"size"`
	}

	var stackJSON StackJSON
	if err := json.Unmarshal(data, &stackJSON); err != nil {
		return err
	}

	s.FromSlice(stackJSON.Elements)
	return nil
}

// ToJSON сериализует стек в JSON строку
func (s *Stack) ToJSON() (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в стек
func (s *Stack) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), s)
}

// SaveToJSON сохраняет стек в JSON файл
func (s *Stack) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(s)
}

// LoadFromJSON загружает стек из JSON файла
func (s *Stack) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(s)
}

// ToBinary сериализует стек в бинарный формат
func (s *Stack) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	size := int32(s.Size())
	if err := binary.Write(&buf, binary.LittleEndian, size); err != nil {
		return nil, fmt.Errorf("ошибка записи размера стека: %w", err)
	}

	curr := s.head
	for curr != nil {
		strBytes := []byte(curr.val)
		strLen := int32(len(strBytes))

		if err := binary.Write(&buf, binary.LittleEndian, strLen); err != nil {
			return nil, err
		}

		if _, err := buf.Write(strBytes); err != nil {
			return nil, fmt.Errorf("ошибка записи данных строки: %w", err)
		}

		curr = curr.next
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует стек из бинарного формата
func (s *Stack) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	var size int32
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("ошибка чтения размера стека: %w", err)
	}

	elements := make([]string, 0, size)

	for i := 0; i < int(size); i++ {
		var strLen int32
		if err := binary.Read(buf, binary.LittleEndian, &strLen); err != nil {
			if err == io.EOF {
				return fmt.Errorf("неожиданный конец данных при чтении длины строки %d. Ожидалось %d элементов", i, size)
			}
			return fmt.Errorf("ошибка чтения длины строки %d: %w", i, err)
		}

		strBytes := make([]byte, strLen)
		if _, err := io.ReadFull(buf, strBytes); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return fmt.Errorf("недостаточно данных для чтения строки %d длиной %d байт", i, strLen)
			}
			return fmt.Errorf("ошибка чтения данных строки %d: %w", i, err)
		}

		elements = append(elements, string(strBytes))
	}

	s.head = nil
	for i := len(elements) - 1; i >= 0; i-- {
		s.Push(elements[i])
	}

	return nil
}

// SaveToBinary сохраняет стек в бинарный файл
func (s *Stack) SaveToBinary(filename string) error {
	data, err := s.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает стек из бинарного файла
func (s *Stack) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return s.FromBinary(data)
}
