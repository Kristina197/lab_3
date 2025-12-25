package structs

import (
	"os"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	s := NewStack()

	s.Push("1")
	s.Push("2")
	s.Push("3")

	if s.Top() != "3" {
		t.Errorf("ожидается 3, получено %s", s.Top())
	}

	if popped := s.Pop(); popped != "3" {
		t.Errorf("ожидается 3, получено %s", popped)
	}

	if s.Top() != "2" {
		t.Errorf("ожидается 2, получено %s", s.Top())
	}

	s.Push("4")
	if s.Top() != "4" {
		t.Errorf("ожидается 4, получено %s", s.Top())
	}

	if popped := s.Pop(); popped != "4" {
		t.Errorf("ожидается 4, получено %s", popped)
	}

	if popped := s.Pop(); popped != "2" {
		t.Errorf("ожидается 2, получено %s", popped)
	}

	if popped := s.Pop(); popped != "1" {
		t.Errorf("ожидается 1, получено %s", popped)
	}

	if popped := s.Pop(); popped != "" {
		t.Errorf("ожидается пустая строка, получено %s", popped)
	}

	if s.Top() != "" {
		t.Errorf("ожидается пустая строка, получено %s", s.Top())
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Error("новый стек должен быть пустым")
	}

	s.Push("1")
	if s.IsEmpty() {
		t.Error("стек не должен быть пустым после добавления")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("стек должен быть пустым после удаления всех элементов")
	}
}

func TestStack_Size(t *testing.T) {
	s := NewStack()

	if size := s.Size(); size != 0 {
		t.Errorf("ожидается размер 0, получено %d", size)
	}

	s.Push("1")
	s.Push("2")
	s.Push("3")

	if size := s.Size(); size != 3 {
		t.Errorf("ожидается размер 3, получено %d", size)
	}

	s.Pop()
	if size := s.Size(); size != 2 {
		t.Errorf("ожидается размер 2, получено %d", size)
	}

	s.Pop()
	s.Pop()
	if size := s.Size(); size != 0 {
		t.Errorf("ожидается размер 0, получено %d", size)
	}
}

func TestStack_String(t *testing.T) {
	s := NewStack()

	expected := "top -> bottom"
	if result := s.String(); result != expected {
		t.Errorf("ожидается '%s', получено '%s'", expected, result)
	}

	s.Push("1")
	s.Push("2")
	s.Push("3")

	expected = "top -> 3 -> 2 -> 1 -> bottom"
	if result := s.String(); result != expected {
		t.Errorf("ожидается '%s', получено '%s'", expected, result)
	}
}

func TestStack_ToSlice(t *testing.T) {
	s := NewStack()

	if slice := s.ToSlice(); len(slice) != 0 {
		t.Errorf("ожидается пустой слайс, получено %v", slice)
	}

	s.Push("c")
	s.Push("b")
	s.Push("a")

	expected := []string{"a", "b", "c"}
	slice := s.ToSlice()

	if len(slice) != len(expected) {
		t.Errorf("ожидается длина слайса %d, получено %d", len(expected), len(slice))
	}

	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, slice[i])
		}
	}
}

//func TestStack_FromSlice(t *testing.T) {
//	s := NewStack()
//
//	elements := []string{"a", "b", "c", "d", "e"}
//	s.FromSlice(elements)
//
//	// Проверяем, что элементы находятся в правильном порядке в стеке
//	expectedOrder := []string{"e", "d", "c", "b", "a"}
//	for i := 0; i < len(expectedOrder); i++ {
//		if popped := s.Pop(); popped != expectedOrder[i] {
//			t.Errorf("ожидается '%s', получено '%s'", expectedOrder[i], popped)
//		}
//	}
//
//	s.FromSlice([]string{})
//	if !s.IsEmpty() {
//		t.Error("стек должен быть пустым после FromSlice с пустым слайсом")
//	}
//}

func TestStack_JSONSerialization(t *testing.T) {
	s := NewStack()
	s.Push("1")
	s.Push("2")
	s.Push("3")

	jsonStr, err := s.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON не удалось: %v", err)
	}

	s2 := NewStack()
	err = s2.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("FromJSON не удалось: %v", err)
	}

	// Восстанавливаем исходный стек для сравнения
	s3 := NewStack()
	s3.Push("1")
	s3.Push("2")
	s3.Push("3")

	for !s3.IsEmpty() {
		val1 := s3.Pop()
		val2 := s2.Pop()
		if val1 != val2 {
			t.Errorf("значения не совпадают: %s != %s", val1, val2)
		}
	}

	if !s2.IsEmpty() {
		t.Error("s2 должен быть пустым после сравнения")
	}

	emptyStack := NewStack()
	jsonStr, err = emptyStack.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON с пустым стеком не удалось: %v", err)
	}

	emptyStack2 := NewStack()
	err = emptyStack2.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("FromJSON с пустым стеком не удалось: %v", err)
	}

	if !emptyStack2.IsEmpty() {
		t.Error("десериализованный пустой стек должен быть пустым")
	}

	err = emptyStack2.FromJSON("некорректный json")
	if err == nil {
		t.Error("ожидается ошибка при некорректном JSON")
	}
}

func TestStack_BinarySerialization(t *testing.T) {
	s := NewStack()
	s.Push("1")
	s.Push("22")
	s.Push("333")

	data, err := s.ToBinary()
	if err != nil {
		t.Fatalf("ToBinary не удалось: %v", err)
	}

	s2 := NewStack()
	err = s2.FromBinary(data)
	if err != nil {
		t.Fatalf("FromBinary не удалось: %v", err)
	}

	// Восстанавливаем исходный стек для сравнения
	s3 := NewStack()
	s3.Push("1")
	s3.Push("22")
	s3.Push("333")

	for !s3.IsEmpty() {
		val1 := s3.Pop()
		val2 := s2.Pop()
		if val1 != val2 {
			t.Errorf("значения не совпадают: %s != %s", val1, val2)
		}
	}

	if !s2.IsEmpty() {
		t.Error("s2 должен быть пустым после сравнения")
	}

	emptyStack := NewStack()
	data, err = emptyStack.ToBinary()
	if err != nil {
		t.Fatalf("ToBinary с пустым стеком не удалось: %v", err)
	}

	emptyStack2 := NewStack()
	err = emptyStack2.FromBinary(data)
	if err != nil {
		t.Fatalf("FromBinary с пустым стеком не удалось: %v", err)
	}

	if !emptyStack2.IsEmpty() {
		t.Error("десериализованный пустой стек должен быть пустым")
	}

	err = emptyStack2.FromBinary([]byte{})
	if err == nil {
		t.Error("ожидается ошибка при пустых бинарных данных")
	}

	invalidData := []byte{0x01, 0x00, 0x00, 0x00}
	err = emptyStack2.FromBinary(invalidData)
	if err == nil {
		t.Error("ожидается ошибка при некорректных бинарных данных")
	}
}

func TestStack_FileSerialization(t *testing.T) {
	s := NewStack()
	s.Push("файл")
	s.Push("тест")
	s.Push("данные")

	jsonFile := "test_stack.json"
	defer func() {
		if err := os.Remove(jsonFile); err != nil {
			t.Logf("не удалось удалить файл: %v", err)
		}
	}()

	err := s.SaveToJSON(jsonFile)
	if err != nil {
		t.Fatalf("SaveToJSON не удалось: %v", err)
	}

	s2 := NewStack()
	err = s2.LoadFromJSON(jsonFile)
	if err != nil {
		t.Fatalf("LoadFromJSON не удалось: %v", err)
	}

	// Восстанавливаем исходный стек для сравнения
	s3 := NewStack()
	s3.Push("файл")
	s3.Push("тест")
	s3.Push("данные")

	for !s3.IsEmpty() {
		val1 := s3.Pop()
		val2 := s2.Pop()
		if val1 != val2 {
			t.Errorf("значения из JSON файла не совпадают: %s != %s", val1, val2)
		}
	}

	binFile := "test_stack.bin"
	defer func() {
		if err := os.Remove(binFile); err != nil {
			t.Logf("не удалось удалить файл: %v", err)
		}
	}()

	s4 := NewStack()
	s4.Push("бинарный")
	s4.Push("файл")
	s4.Push("тест")

	err = s4.SaveToBinary(binFile)
	if err != nil {
		t.Fatalf("SaveToBinary не удалось: %v", err)
	}

	s5 := NewStack()
	err = s5.LoadFromBinary(binFile)
	if err != nil {
		t.Fatalf("LoadFromBinary не удалось: %v", err)
	}

	// Восстанавливаем для сравнения
	s6 := NewStack()
	s6.Push("бинарный")
	s6.Push("файл")
	s6.Push("тест")

	for !s6.IsEmpty() {
		val1 := s6.Pop()
		val2 := s5.Pop()
		if val1 != val2 {
			t.Errorf("значения из бинарного файла не совпадают: %s != %s", val1, val2)
		}
	}

	nonExistentStack := NewStack()
	err = nonExistentStack.LoadFromJSON("не_существующий.json")
	if err == nil {
		t.Error("ожидается ошибка при загрузке несуществующего JSON файла")
	}

	err = nonExistentStack.LoadFromBinary("не_существующий.bin")
	if err == nil {
		t.Error("ожидается ошибка при загрузке несуществующего бинарного файла")
	}
}

func TestStack_MarshalUnmarshalJSON(t *testing.T) {
	s := NewStack()
	s.Push("1")
	s.Push("2")
	s.Push("3")

	data, err := s.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON не удалось: %v", err)
	}

	s2 := NewStack()
	err = s2.UnmarshalJSON(data)
	if err != nil {
		t.Fatalf("UnmarshalJSON не удалось: %v", err)
	}

	// Восстанавливаем для сравнения
	s3 := NewStack()
	s3.Push("1")
	s3.Push("2")
	s3.Push("3")

	for !s3.IsEmpty() {
		val1 := s3.Pop()
		val2 := s2.Pop()
		if val1 != val2 {
			t.Errorf("значения не совпадают: %s != %s", val1, val2)
		}
	}

	err = s2.UnmarshalJSON([]byte("некорректный json"))
	if err == nil {
		t.Error("ожидается ошибка при некорректном JSON")
	}
}

func TestStack_EmptyStackOperations(t *testing.T) {
	s := NewStack()

	if val := s.Pop(); val != "" {
		t.Errorf("ожидается пустая строка, получено %s", val)
	}

	if val := s.Top(); val != "" {
		t.Errorf("ожидается пустая строка, получено %s", val)
	}

	if !s.IsEmpty() {
		t.Error("стек должен быть пустым")
	}

	if size := s.Size(); size != 0 {
		t.Errorf("ожидается размер 0, получено %d", size)
	}
}

func TestStack_MultipleOperations(t *testing.T) {
	s := NewStack()

	for i := 0; i < 1000; i++ {
		s.Push("элемент")
	}

	if size := s.Size(); size != 1000 {
		t.Errorf("ожидается размер 1000, получено %d", size)
	}

	for i := 0; i < 1000; i++ {
		s.Pop()
	}

	if !s.IsEmpty() {
		t.Error("стек должен быть пустым после удаления всех элементов")
	}
}

func TestStack_ConcurrentPushPop(t *testing.T) {
	s := NewStack()

	s.Push("1")
	s.Push("2")
	s.Pop()
	s.Push("3")
	s.Push("4")
	s.Pop()
	s.Pop()

	if s.Top() != "1" {
		t.Errorf("ожидается 1, получено %s", s.Top())
	}

	if size := s.Size(); size != 1 {
		t.Errorf("ожидается размер 1, получено %d", size)
	}
}

func TestNewStackWithCapacity(t *testing.T) {
	s := NewStackWithCapacity(10)

	if s == nil {
		t.Error("NewStackWithCapacity вернул nil")
	}

	s.Push("тест")
	if s.Top() != "тест" {
		t.Errorf("ожидается 'тест', получено %s", s.Top())
	}
}

func TestSNode(t *testing.T) {
	node := NewSNode("тест")
	if node.val != "тест" {
		t.Errorf("ожидается значение узла 'тест', получено %s", node.val)
	}
	if node.next != nil {
		t.Error("новый узел должен иметь nil указатель next")
	}

	node2 := NewSNode("следующий")
	node.next = node2
	if node.next != node2 {
		t.Error("не удалось связать узлы")
	}
}

func TestStack_Print(t *testing.T) {
	s := NewStack()

	s.Push("1")
	s.Push("2")
	s.Push("3")

	s.Print()
	expected := "top -> 3 -> 2 -> 1 -> bottom"
	if result := s.String(); result != expected {
		t.Errorf("состояние стека некорректно. Ожидается '%s', получено '%s'", expected, result)
	}
}

func TestStack_BinaryEdgeCases(t *testing.T) {
	s := NewStack()

	s.Push("")
	s.Push("длинная строка с пробелами")
	s.Push("спецсимволы: !@#$%^&*()")

	data, err := s.ToBinary()
	if err != nil {
		t.Fatalf("ToBinary не удалось: %v", err)
	}

	s2 := NewStack()
	err = s2.FromBinary(data)
	if err != nil {
		t.Fatalf("FromBinary не удалось: %v", err)
	}

	if s.Size() != s2.Size() {
		t.Errorf("размеры не совпадают: %d != %d", s.Size(), s2.Size())
	}

	// Восстанавливаем для сравнения
	s3 := NewStack()
	s3.Push("")
	s3.Push("длинная строка с пробелами")
	s3.Push("спецсимволы: !@#$%^&*()")

	for !s3.IsEmpty() {
		val1 := s3.Pop()
		val2 := s2.Pop()
		if val1 != val2 {
			t.Errorf("значения не совпадают: %s != %s", val1, val2)
		}
	}
}
