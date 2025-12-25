package structs

import (
	"os"
	"strings"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	if list.head != nil {
		t.Error("Новый список должен быть пустым")
	}
}

func TestSLLInsertAtHead(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtHead("first")
	if list.head.val != "first" {
		t.Error("Вставка в голову не работает")
	}

	list.InsertAtHead("second")
	if list.head.val != "second" {
		t.Error("Вставка в голову не работает")
	}
}

func TestSLLInsertAtTail(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtTail("first")
	if list.head.val != "first" {
		t.Error("Вставка в хвост не работает")
	}

	list.InsertAtTail("second")
	if list.head.next.val != "second" {
		t.Error("Вставка в хвост не работает")
	}
}

func TestSLLInsertBefore(t *testing.T) {
	list := NewSinglyLinkedList()

	if err := list.InsertBefore("nothing", "smth"); err == nil {
		t.Error("Ничего не должно вставиться, элемент не найден")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.InsertBefore("two", "new")
	if err != nil {
		t.Error("Вставка перед не должна возвращать ошибку")
	}
	if list.head.next.val != "new" {
		t.Error("Вставка перед не работает")
	}

	err = list.InsertBefore("one", "head")
	if err != nil {
		t.Error("Вставка перед головой не должна возвращать ошибку")
	}
	if list.head.val != "head" {
		t.Error("Вставка перед головой не работает")
	}

	err = list.InsertBefore("missing", "val")
	if err == nil {
		t.Error("Должна быть ошибка при вставке перед несуществующим элементом")
	}
}

func TestSLLInsertAfter(t *testing.T) {
	list := NewSinglyLinkedList()

	if err := list.InsertAfter("nothing", "smth"); err == nil {
		t.Error("Ничего не должно вставиться, элемент не найден")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.InsertAfter("two", "new")
	if err != nil {
		t.Error("Вставка после не должна возвращать ошибку")
	}
	if list.head.next.next.val != "new" {
		t.Error("Вставка после не работает")
	}

	err = list.InsertAfter("missing", "val")
	if err == nil {
		t.Error("Должна быть ошибка при вставке после несуществующего элемента")
	}
}

func TestSLLDeleteHead(t *testing.T) {
	list := NewSinglyLinkedList()
	err := list.DeleteHead()
	if err == nil {
		t.Error("Удаление из пустого списка должно возвращать ошибку")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")

	err = list.DeleteHead()
	if err != nil || list.head.val != "two" {
		t.Error("Удаление головы не работает")
	}

}

func TestSLLDeleteTail(t *testing.T) {
	list := NewSinglyLinkedList()
	err := list.DeleteTail()
	if err == nil {
		t.Error("Удаление из пустого списка должно возвращать ошибку")
	}

	list.InsertAtTail("one")
	err = list.DeleteTail()
	if err != nil || list.head != nil {
		t.Error("Удаление хвоста из списка с одним элементом не работает")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	err = list.DeleteTail()
	if err != nil || list.head.next != nil {
		t.Error("Удаление хвоста не работает")
	}
}

func TestSLLDeleteBefore(t *testing.T) {
	list := NewSinglyLinkedList()

	if err := list.DeleteBefore("nothing"); err == nil {
		t.Error("Недостаточно элементов для удаления")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.DeleteBefore("three")
	if err != nil || list.head.next.val != "three" {
		t.Error("Удаление перед не работает")
	}

	err = list.DeleteBefore("missing")
	if err == nil {
		t.Error("Должна быть ошибка при удалении перед несуществующим элементом")
	}

	err = list.DeleteBefore("two")
	if err == nil {
		t.Error("Должен был удалиться head")
	}
}

func TestSLLDeleteAfter(t *testing.T) {
	list := NewSinglyLinkedList()

	if err := list.DeleteAfter("nothing"); err == nil {
		t.Error("Ничего не должно удалиться, элемент не найден")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.DeleteAfter("one")
	if err != nil || list.head.next.val != "three" {
		t.Error("Удаление после не работает")
	}

	err = list.DeleteAfter("three")
	if err == nil {
		t.Error("Должна быть ошибка при удалении после последнего элемента")
	}

	err = list.DeleteAfter("missing")
	if err == nil {
		t.Error("Должна быть ошибка при удалении после несуществующего элемента")
	}
}

func TestSLLDeleteByValue(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.DeleteByValue("two")
	if err != nil || list.Length() != 2 {
		t.Error("Удаление по значению не работает")
	}

	err = list.DeleteByValue("one")
	if err != nil || list.head.val != "three" {
		t.Error("Удаление головы по значению не работает")
	}

	err = list.DeleteByValue("missing")
	if err == nil {
		t.Error("Должна быть ошибка при удалении несуществующего значения")
	}

	list.Print()
}

func TestSLLSearchByValue(t *testing.T) {
	list := NewSinglyLinkedList()
	if list.SearchByValue("test") != -1 {
		t.Error("Поиск в пустом списке должен возвращать -1")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	if list.SearchByValue("two") != 1 {
		t.Error("Поиск по значению не работает")
	}

	if list.SearchByValue("missing") != -1 {
		t.Error("Поиск несуществующего значения должен возвращать -1")
	}
}

func TestSLLLength(t *testing.T) {
	list := NewSinglyLinkedList()
	if list.Length() != 0 {
		t.Error("Длина пустого списка должна быть 0")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")

	if list.Length() != 2 {
		t.Error("Длина списка неверна")
	}
}

func TestSLLIsEmpty(t *testing.T) {
	list := NewSinglyLinkedList()
	if !list.IsEmpty() {
		t.Error("Новый список должен быть пустым")
	}

	list.InsertAtTail("test")
	if list.IsEmpty() {
		t.Error("Непустой список не должен быть пустым")
	}
}

func TestGetHead(t *testing.T) {
	list := NewSinglyLinkedList()
	if list.GetHead() != "" {
		t.Error("Голова пустого списка должна быть пустой строкой")
	}

	list.InsertAtTail("test")
	if list.GetHead() != "test" {
		t.Error("GetHead возвращает неверное значение")
	}
}

func TestSLLToSlice(t *testing.T) {
	list := NewSinglyLinkedList()
	slice := list.ToSlice()
	if len(slice) != 0 {
		t.Error("Пустой список должен возвращать пустой срез")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	slice = list.ToSlice()
	if len(slice) != 2 || slice[0] != "one" || slice[1] != "two" {
		t.Error("ToSlice возвращает неверный срез")
	}
}

func TestSLLClear(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")

	list.Clear()
	if list.head != nil || !list.IsEmpty() {
		t.Error("Clear не очищает список")
	}
}

func TestSLLJSONSerialization(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	jsonStr, err := list.ToJSON()
	if err != nil || jsonStr == "" {
		t.Error("ToJSON не работает")
	}

	newList := NewSinglyLinkedList()
	err = newList.FromJSON(jsonStr)
	if err != nil || newList.Length() != 3 {
		t.Error("FromJSON не работает")
	}

	err = list.SaveToJSON("test.json")
	if err != nil {
		t.Error("SaveToJSON не работает")
	}
	defer os.Remove("test.json")

	loadedList := NewSinglyLinkedList()
	err = loadedList.LoadFromJSON("test.json")
	if err != nil || loadedList.Length() != 3 {
		t.Error("LoadFromJSON не работает")
	}
}

func TestSLLBinarySerialization(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	binaryData, err := list.ToBinary()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinary не работает")
	}

	newList := NewSinglyLinkedList()
	err = newList.FromBinary(binaryData)
	if err != nil || newList.Length() != 3 {
		t.Error("FromBinary не работает")
	}

	err = list.SaveToBinary("test.bin")
	if err != nil {
		t.Error("SaveToBinary не работает")
	}
	defer os.Remove("test.bin")

	loadedList := NewSinglyLinkedList()
	err = loadedList.LoadFromBinary("test.bin")
	if err != nil || loadedList.Length() != 3 {
		t.Error("LoadFromBinary не работает")
	}
}

func TestString(t *testing.T) {
	list := NewSinglyLinkedList()
	str := list.String()
	if str != "Список пуст" {
		t.Error("String для пустого списка неверен")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	str = list.String()
	if !strings.Contains(str, "one -> two") {
		t.Error("String неверно форматирует список")
	}
}

func TestEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Print()

	err := list.DeleteByValue("yo")
	if err == nil {
		t.Error("Ничего не должно удалиться, т.к. список пустой")
	}
}
