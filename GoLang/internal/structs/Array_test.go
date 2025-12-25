package structs

import (
	"os"
	"testing"
)

func TestArray_NewArray(t *testing.T) {
	a := NewArray()
	if a == nil {
		t.Error("NewArray вернул nil")
	}
	if a.Size != 0 {
		t.Errorf("ожидается размер 0, получено %d", a.Size)
	}
	if a.Capacity < 10 {
		t.Errorf("ожидается емкость не менее 10, получено %d", a.Capacity)
	}
}

func TestArray_NewArrayWithCapacity(t *testing.T) {
	a := NewArrayWithCapacity(5)
	if a == nil {
		t.Error("NewArrayWithCapacity вернул nil")
	}
	if a.Size != 0 {
		t.Errorf("ожидается размер 0, получено %d", a.Size)
	}
	if a.Capacity != 5 {
		t.Errorf("ожидается емкость 5, получено %d", a.Capacity)
	}
}

func TestArray_Cap(t *testing.T) {
	a := NewArrayWithCapacity(20)
	if cap := a.Cap(); cap != 20 {
		t.Errorf("ожидается емкость 20, получено %d", cap)
	}
}

func TestArray_Length(t *testing.T) {
	a := NewArray()
	if length := a.Length(); length != 0 {
		t.Errorf("ожидается длина 0, получено %d", length)
	}

	a.PushBack("1")
	if length := a.Length(); length != 1 {
		t.Errorf("ожидается длина 1, получено %d", length)
	}

	a.PushBack("2")
	if length := a.Length(); length != 2 {
		t.Errorf("ожидается длина 2, получено %d", length)
	}
}

func TestArray_PushBack(t *testing.T) {
	a := NewArrayWithCapacity(2)

	if !a.PushBack("1") {
		t.Error("PushBack должен вернуть true")
	}
	if a.Size != 1 {
		t.Errorf("ожидается размер 1, получено %d", a.Size)
	}
	if a.Data[0] != "1" {
		t.Errorf("ожидается '1', получено '%s'", a.Data[0])
	}

	if !a.PushBack("2") {
		t.Error("PushBack должен вернуть true")
	}
	if a.Size != 2 {
		t.Errorf("ожидается размер 2, получено %d", a.Size)
	}
	if a.Data[1] != "2" {
		t.Errorf("ожидается '2', получено '%s'", a.Data[1])
	}

	// Проверка увеличения емкости
	if !a.PushBack("3") {
		t.Error("PushBack должен вернуть true при переполнении")
	}
	if a.Size != 3 {
		t.Errorf("ожидается размер 3, получено %d", a.Size)
	}
	if a.Data[2] != "3" {
		t.Errorf("ожидается '3', получено '%s'", a.Data[2])
	}
	if a.Capacity <= 2 {
		t.Errorf("емкость должна увеличиться, получено %d", a.Capacity)
	}
}

func TestArray_PushByIndex(t *testing.T) {
	a := NewArrayWithCapacity(5)

	// Добавление в пустой массив
	if !a.PushByIndex("1", 0) {
		t.Error("PushByIndex должен вернуть true для индекса 0")
	}
	if a.Size != 1 || a.Data[0] != "1" {
		t.Errorf("ожидается ['1'], получено %v", a.Data[:a.Size])
	}

	// Добавление в конец
	if !a.PushByIndex("3", 1) {
		t.Error("PushByIndex должен вернуть true для индекса 1")
	}
	if a.Size != 2 || a.Data[1] != "3" {
		t.Errorf("ожидается ['1','3'], получено %v", a.Data[:a.Size])
	}

	// Добавление в середину
	if !a.PushByIndex("2", 1) {
		t.Error("PushByIndex должен вернуть true для индекса 1")
	}
	if a.Size != 3 {
		t.Errorf("ожидается размер 3, получено %d", a.Size)
	}
	expected := []string{"1", "2", "3"}
	for i, val := range expected {
		if a.Data[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, a.Data[i])
		}
	}

	// Некорректный индекс
	if a.PushByIndex("4", 5) {
		t.Error("PushByIndex должен вернуть false для индекса больше размера")
	}
	if a.Size != 3 {
		t.Errorf("размер не должен измениться, получено %d", a.Size)
	}

	// Проверка увеличения емкости
	a2 := NewArrayWithCapacity(2)
	a2.PushByIndex("1", 0)
	a2.PushByIndex("2", 1)
	if !a2.PushByIndex("3", 1) {
		t.Error("PushByIndex должен вернуть true при переполнении")
	}
	if a2.Size != 3 {
		t.Errorf("ожидается размер 3, получено %d", a2.Size)
	}
}

func TestArray_GetByIndex(t *testing.T) {
	a := NewArray()
	a.PushBack("1")
	a.PushBack("2")
	a.PushBack("3")

	if val := a.GetByIndex(0); val != "1" {
		t.Errorf("ожидается '1', получено '%s'", val)
	}
	if val := a.GetByIndex(1); val != "2" {
		t.Errorf("ожидается '2', получено '%s'", val)
	}
	if val := a.GetByIndex(2); val != "3" {
		t.Errorf("ожидается '3', получено '%s'", val)
	}

	// Некорректные индексы
	if val := a.GetByIndex(-1); val != "" {
		t.Errorf("ожидается пустая строка для индекса -1, получено '%s'", val)
	}
	if val := a.GetByIndex(3); val != "" {
		t.Errorf("ожидается пустая строка для индекса 3, получено '%s'", val)
	}
	if val := a.GetByIndex(100); val != "" {
		t.Errorf("ожидается пустая строка для индекса 100, получено '%s'", val)
	}
}

func TestArray_DeleteByIndex(t *testing.T) {
	a := NewArray()
	a.PushBack("1")
	a.PushBack("2")
	a.PushBack("3")
	a.PushBack("4")
	a.PushBack("5")

	// Удаление из середины
	if !a.DeleteByIndex(2) {
		t.Error("DeleteByIndex должен вернуть true для существующего индекса")
	}
	if a.Size != 4 {
		t.Errorf("ожидается размер 4, получено %d", a.Size)
	}
	expected := []string{"1", "2", "4", "5"}
	for i, val := range expected {
		if a.Data[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, a.Data[i])
		}
	}

	// Удаление первого элемента
	if !a.DeleteByIndex(0) {
		t.Error("DeleteByIndex должен вернуть true для индекса 0")
	}
	if a.Size != 3 {
		t.Errorf("ожидается размер 3, получено %d", a.Size)
	}
	expected = []string{"2", "4", "5"}
	for i, val := range expected {
		if a.Data[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, a.Data[i])
		}
	}

	// Удаление последнего элемента
	if !a.DeleteByIndex(2) {
		t.Error("DeleteByIndex должен вернуть true для последнего индекса")
	}
	if a.Size != 2 {
		t.Errorf("ожидается размер 2, получено %d", a.Size)
	}
	expected = []string{"2", "4"}
	for i, val := range expected {
		if a.Data[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, a.Data[i])
		}
	}

	// Некорректные индексы
	if a.DeleteByIndex(-1) {
		t.Error("DeleteByIndex должен вернуть false для индекса -1")
	}
	if a.DeleteByIndex(2) {
		t.Error("DeleteByIndex должен вернуть false для индекса 2")
	}
	if a.DeleteByIndex(100) {
		t.Error("DeleteByIndex должен вернуть false для индекса 100")
	}
}

func TestArray_SwapByIndex(t *testing.T) {
	a := NewArray()
	a.PushBack("1")
	a.PushBack("2")
	a.PushBack("3")

	// Корректная замена
	if !a.SwapByIndex("новый", 1) {
		t.Error("SwapByIndex должен вернуть true для существующего индекса")
	}
	if a.Data[1] != "новый" {
		t.Errorf("ожидается 'новый', получено '%s'", a.Data[1])
	}
	if a.Size != 3 {
		t.Errorf("размер не должен измениться, получено %d", a.Size)
	}

	// Некорректные индексы
	if a.SwapByIndex("тест", -1) {
		t.Error("SwapByIndex должен вернуть false для индекса -1")
	}
	if a.SwapByIndex("тест", 3) {
		t.Error("SwapByIndex должен вернуть false для индекса 3")
	}
	if a.SwapByIndex("тест", 100) {
		t.Error("SwapByIndex должен вернуть false для индекса 100")
	}
}

func TestArray_Print(t *testing.T) {
	a := NewArray()
	a.PushBack("1")
	a.PushBack("2")
	a.PushBack("3")

	a.Print()
	if a.Size != 3 {
		t.Errorf("ожидается размер 3, получено %d", a.Size)
	}
}

func TestArray_doubleCapacity(t *testing.T) {
	a := NewArrayWithCapacity(2)
	initialCapacity := a.Capacity

	a.doubleCapacity()
	if a.Capacity != initialCapacity*2 {
		t.Errorf("емкость должна удвоиться, ожидается %d, получено %d", initialCapacity*2, a.Capacity)
	}

	// Проверка для емкости 0
	a2 := NewArrayWithCapacity(0)
	a2.Capacity = 0
	a2.doubleCapacity()
	if a2.Capacity != 1 {
		t.Errorf("при удвоении емкости 0 ожидается 1, получено %d", a2.Capacity)
	}
}

func TestArray_MarshalUnmarshalJSON(t *testing.T) {
	a := NewArrayWithCapacity(5)
	a.PushBack("1")
	a.PushBack("2")
	a.PushBack("3")

	data, err := a.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON не удалось: %v", err)
	}

	a2 := NewArray()
	err = a2.UnmarshalJSON(data)
	if err != nil {
		t.Fatalf("UnmarshalJSON не удалось: %v", err)
	}

	if a.Size != a2.Size {
		t.Errorf("размеры не совпадают: %d != %d", a.Size, a2.Size)
	}

	for i := 0; i < a.Size; i++ {
		if a.Data[i] != a2.Data[i] {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, a.Data[i], a2.Data[i])
		}
	}

	// Тест с некорректными данными
	err = a2.UnmarshalJSON([]byte("некорректный json"))
	if err == nil {
		t.Error("ожидается ошибка при некорректном JSON")
	}
}

func TestArray_ToFromJSON(t *testing.T) {
	a := NewArray()
	a.PushBack("тест1")
	a.PushBack("тест2")
	a.PushBack("тест3")

	jsonStr, err := a.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON не удалось: %v", err)
	}

	a2 := NewArray()
	err = a2.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("FromJSON не удалось: %v", err)
	}

	if a.Size != a2.Size {
		t.Errorf("размеры не совпадают: %d != %d", a.Size, a2.Size)
	}

	for i := 0; i < a.Size; i++ {
		if a.Data[i] != a2.Data[i] {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, a.Data[i], a2.Data[i])
		}
	}

	// Тест с пустым массивом
	a3 := NewArray()
	jsonStr, err = a3.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON пустого массива не удалось: %v", err)
	}

	a4 := NewArray()
	err = a4.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("FromJSON пустого массива не удалось: %v", err)
	}

	if a4.Size != 0 {
		t.Errorf("ожидается пустой массив, размер %d", a4.Size)
	}

	// Некорректный JSON
	err = a4.FromJSON("некорректный json")
	if err == nil {
		t.Error("ожидается ошибка при некорректном JSON")
	}
}

func TestArray_SaveLoadJSON(t *testing.T) {
	a := NewArray()
	a.PushBack("файл1")
	a.PushBack("файл2")
	a.PushBack("файл3")

	jsonFile := "test_array.json"
	defer func() {
		if err := os.Remove(jsonFile); err != nil {
			t.Logf("не удалось удалить файл: %v", err)
		}
	}()

	err := a.SaveToJSON(jsonFile)
	if err != nil {
		t.Fatalf("SaveToJSON не удалось: %v", err)
	}

	a2 := NewArray()
	err = a2.LoadFromJSON(jsonFile)
	if err != nil {
		t.Fatalf("LoadFromJSON не удалось: %v", err)
	}

	if a.Size != a2.Size {
		t.Errorf("размеры не совпадают: %d != %d", a.Size, a2.Size)
	}

	for i := 0; i < a.Size; i++ {
		if a.Data[i] != a2.Data[i] {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, a.Data[i], a2.Data[i])
		}
	}

	// Тест с несуществующим файлом
	a3 := NewArray()
	err = a3.LoadFromJSON("не_существующий.json")
	if err == nil {
		t.Error("ожидается ошибка при загрузке несуществующего файла")
	}
}

func TestArray_ToFromBinary(t *testing.T) {
	a := NewArray()
	a.PushBack("бинарный1")
	a.PushBack("бинарный2")
	a.PushBack("бинарный3")
	a.PushBack("") // пустая строка

	data, err := a.ToBinary()
	if err != nil {
		t.Fatalf("ToBinary не удалось: %v", err)
	}

	a2 := NewArray()
	err = a2.FromBinary(data)
	if err != nil {
		t.Fatalf("FromBinary не удалось: %v", err)
	}

	if a.Size != a2.Size {
		t.Errorf("размеры не совпадают: %d != %d", a.Size, a2.Size)
	}

	for i := 0; i < a.Size; i++ {
		if a.Data[i] != a2.Data[i] {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, a.Data[i], a2.Data[i])
		}
	}

	// Тест с пустым массивом
	a3 := NewArray()
	data, err = a3.ToBinary()
	if err != nil {
		t.Fatalf("ToBinary пустого массива не удалось: %v", err)
	}

	a4 := NewArray()
	err = a4.FromBinary(data)
	if err != nil {
		t.Fatalf("FromBinary пустого массива не удалось: %v", err)
	}

	if a4.Size != 0 {
		t.Errorf("ожидается пустой массив, размер %d", a4.Size)
	}

	// Тест с некорректными данными
	err = a4.FromBinary([]byte{})
	if err == nil {
		t.Error("ожидается ошибка при пустых бинарных данных")
	}

	// Тест с неполными данными
	invalidData := []byte{0x01, 0x00, 0x00, 0x00} // size = 1, но нет данных строки
	err = a4.FromBinary(invalidData)
	if err == nil {
		t.Error("ожидается ошибка при некорректных бинарных данных")
	}
}

func TestArray_SaveLoadBinary(t *testing.T) {
	a := NewArray()
	a.PushBack("бинарный_файл1")
	a.PushBack("бинарный_файл2")
	a.PushBack("спецсимволы: !@#$%^&*()")

	binFile := "test_array.bin"
	defer func() {
		if err := os.Remove(binFile); err != nil {
			t.Logf("не удалось удалить файл: %v", err)
		}
	}()

	err := a.SaveToBinary(binFile)
	if err != nil {
		t.Fatalf("SaveToBinary не удалось: %v", err)
	}

	a2 := NewArray()
	err = a2.LoadFromBinary(binFile)
	if err != nil {
		t.Fatalf("LoadFromBinary не удалось: %v", err)
	}

	if a.Size != a2.Size {
		t.Errorf("размеры не совпадают: %d != %d", a.Size, a2.Size)
	}

	for i := 0; i < a.Size; i++ {
		if a.Data[i] != a2.Data[i] {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, a.Data[i], a2.Data[i])
		}
	}

	// Тест с несуществующим файлом
	a3 := NewArray()
	err = a3.LoadFromBinary("не_существующий.bin")
	if err == nil {
		t.Error("ожидается ошибка при загрузке несуществующего файла")
	}
}

func TestArray_MultipleOperations(t *testing.T) {
	a := NewArrayWithCapacity(2)

	// Добавление элементов
	for i := 0; i < 100; i++ {
		if !a.PushBack("элемент") {
			t.Errorf("PushBack не удалось на итерации %d", i)
		}
	}

	if a.Size != 100 {
		t.Errorf("ожидается размер 100, получено %d", a.Size)
	}

	// Удаление элементов
	for i := 0; i < 50; i++ {
		if !a.DeleteByIndex(0) {
			t.Errorf("DeleteByIndex не удалось на итерации %d", i)
		}
	}

	if a.Size != 50 {
		t.Errorf("ожидается размер 50, получено %d", a.Size)
	}

	// Добавление по индексу
	for i := 0; i < 25; i++ {
		if !a.PushByIndex("новый", i*2) {
			t.Errorf("PushByIndex не удалось на итерации %d", i)
		}
	}

	if a.Size != 75 {
		t.Errorf("ожидается размер 75, получено %d", a.Size)
	}
}

func TestArray_EdgeCases(t *testing.T) {
	// Пустой массив
	a := NewArray()
	if a.Size != 0 {
		t.Errorf("ожидается размер 0, получено %d", a.Size)
	}

	// Получение из пустого массива
	if val := a.GetByIndex(0); val != "" {
		t.Errorf("ожидается пустая строка, получено '%s'", val)
	}

	// Удаление из пустого массива
	if a.DeleteByIndex(0) {
		t.Error("DeleteByIndex должен вернуть false для пустого массива")
	}

	// Замена в пустом массиве
	if a.SwapByIndex("тест", 0) {
		t.Error("SwapByIndex должен вернуть false для пустого массива")
	}

	// Массив с 0 емкостью
	a2 := NewArrayWithCapacity(0)
	if !a2.PushBack("тест") {
		t.Error("PushBack должен работать с емкостью 0")
	}
	if a2.Size != 1 || a2.Capacity < 1 {
		t.Errorf("после добавления: ожидается размер 1 и емкость >=1, получено размер=%d емкость=%d", a2.Size, a2.Capacity)
	}
}

func TestArray_CapacityManagement(t *testing.T) {
	a := NewArrayWithCapacity(3)

	// Заполняем до предела
	a.PushBack("1")
	a.PushBack("2")
	a.PushBack("3")

	initialCapacity := a.Capacity

	// Добавляем еще, должна увеличиться емкость
	a.PushBack("4")

	if a.Capacity <= initialCapacity {
		t.Errorf("емкость должна увеличиться, было %d, стало %d", initialCapacity, a.Capacity)
	}

	if a.Size != 4 {
		t.Errorf("ожидается размер 4, получено %d", a.Size)
	}

	// Проверяем, что данные сохранились
	expected := []string{"1", "2", "3", "4"}
	for i, val := range expected {
		if a.Data[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, a.Data[i])
		}
	}
}
