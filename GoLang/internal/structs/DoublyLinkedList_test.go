package structs

import (
	"os"
	"strings"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()
	if list.head != nil || list.tail != nil {
		t.Error("Новый двусвязный список должен быть пустым")
	}
}

func TestDLLInsertAtHead(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtHead("first")
	if list.head.val != "first" || list.tail.val != "first" {
		t.Error("Вставка в голову в пустой список не работает")
	}

	list.InsertAtHead("second")
	if list.head.val != "second" || list.tail.val != "first" {
		t.Error("Вставка в голову не работает")
	}
}

func TestDLLInsertAtTail(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("first")
	if list.head.val != "first" || list.tail.val != "first" {
		t.Error("Вставка в хвост в пустой список не работает")
	}

	list.InsertAtTail("second")
	if list.head.val != "first" || list.tail.val != "second" {
		t.Error("Вставка в хвост не работает")
	}
}

func TestDLLInsertBefore(t *testing.T) {
	list := NewDoublyLinkedList()

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")
	list.Print()

	err := list.InsertBefore("two", "new")
	if err != nil || list.head.next.val != "new" {
		t.Error("Вставка перед не работает")
	}

	err = list.InsertBefore("one", "head")
	if err != nil || list.head.val != "head" {
		t.Error("Вставка перед головой не работает")
	}

	err = list.InsertBefore("missing", "val")
	if err == nil {
		t.Error("Должна быть ошибка при вставке перед несуществующим элементом")
	}
}

func TestDLLInsertAfter(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.InsertAfter("two", "new")
	if err != nil || list.head.next.next.val != "new" {
		t.Error("Вставка после не работает")
	}

	err = list.InsertAfter("three", "tail")
	if err != nil || list.tail.val != "tail" {
		t.Error("Вставка после хвоста не работает")
	}

	err = list.InsertAfter("missing", "val")
	if err == nil {
		t.Error("Должна быть ошибка при вставке после несуществующего элемента")
	}
}

func TestDLLDeleteHead(t *testing.T) {
	list := NewDoublyLinkedList()
	err := list.DeleteHead()
	if err == nil {
		t.Error("Удаление из пустого списка должно возвращать ошибку")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")

	err = list.DeleteHead()
	if err != nil || list.head.val != "two" || list.head.prev != nil {
		t.Error("Удаление головы не работает")
	}

	err = list.DeleteHead()
	if err != nil || list.head != nil || list.tail != nil {
		t.Error("Удаление последнего элемента не работает")
	}
}

func TestDLLDeleteTail(t *testing.T) {
	list := NewDoublyLinkedList()
	err := list.DeleteTail()
	if err == nil {
		t.Error("Удаление из пустого списка должно возвращать ошибку")
	}

	list.InsertAtTail("one")
	err = list.DeleteTail()
	if err != nil || list.head != nil || list.tail != nil {
		t.Error("Удаление хвоста из списка с одним элементом не работает")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	err = list.DeleteTail()
	if err != nil || list.tail.val != "one" || list.tail.next != nil {
		t.Error("Удаление хвоста не работает")
	}
}

func TestDLLDeleteBefore(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.DeleteBefore("one")
	if err == nil {
		t.Error("Перед головой не может удалиться элемент")
	}

	err = list.DeleteBefore("three")
	if err != nil || list.head.next.val != "three" {
		t.Error("Удаление перед не работает")
	}

	err = list.DeleteBefore("missing")
	if err == nil {
		t.Error("Должна быть ошибка при удалении перед несуществующим элементом")
	}

	err = list.DeleteBefore("two")
	if err == nil {
		t.Error("Должна быть ошибка при удалении перед головой")
	}
}

func TestDLLDeleteAfter(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.DeleteAfter("one")
	if err != nil || list.head.next.val != "three" {
		t.Error("Удаление после не работает")
	}

	err = list.DeleteAfter("one")
	if err != nil || list.tail.val != "one" {
		t.Error("Удаление после хвоста должно удалять хвост")
	}

	err = list.DeleteAfter("one")
	if err == nil {
		t.Error("После последнего элемента ничего не должно удаляться")
	}

	err = list.DeleteAfter("missing")
	if err == nil {
		t.Error("Должна быть ошибка при удалении после несуществующего элемента")
	}

	err = list.DeleteAfter("two")
	if err == nil {
		t.Error("Должна быть ошибка при удалении после последнего элемента")
	}
}

func TestDLLDeleteByValue(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	err := list.DeleteByValue("two")
	if err != nil || list.Length() != 2 || list.head.next.val != "three" {
		t.Error("Удаление по значению не работает")
	}

	err = list.DeleteByValue("one")
	if err != nil || list.head.val != "three" || list.head.prev != nil {
		t.Error("Удаление головы по значению не работает")
	}

	err = list.DeleteByValue("three")
	if err != nil || list.head != nil || list.tail != nil {
		t.Error("Удаление последнего элемента по значению не работает")
	}

	err = list.DeleteByValue("missing")
	if err == nil {
		t.Error("Должна быть ошибка при удалении несуществующего значения")
	}
}

func TestDLLSearchByValue(t *testing.T) {
	list := NewDoublyLinkedList()
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

func TestDLLLength(t *testing.T) {
	list := NewDoublyLinkedList()
	if list.Length() != 0 {
		t.Error("Длина пустого списка должна быть 0")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")

	if list.Length() != 2 {
		t.Error("Длина списка неверна")
	}
}

func TestDLLIsEmpty(t *testing.T) {
	list := NewDoublyLinkedList()
	if !list.IsEmpty() {
		t.Error("Новый список должен быть пустым")
	}

	list.InsertAtTail("test")
	if list.IsEmpty() {
		t.Error("Непустой список не должен быть пустым")
	}
}

func TestDLLGetHeadAndTail(t *testing.T) {
	list := NewDoublyLinkedList()
	if list.GetHead() != "" || list.GetTail() != "" {
		t.Error("Голова и хвост пустого списка должны быть пустыми")
	}

	list.InsertAtTail("first")
	if list.GetHead() != "first" || list.GetTail() != "first" {
		t.Error("GetHead/GetTail не работают для одного элемента")
	}

	list.InsertAtTail("second")
	if list.GetHead() != "first" || list.GetTail() != "second" {
		t.Error("GetHead/GetTail не работают")
	}
}

func TestDLLToSlice(t *testing.T) {
	list := NewDoublyLinkedList()
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

func TestDLLToSliceReverse(t *testing.T) {
	list := NewDoublyLinkedList()
	slice := list.ToSliceReverse()
	if len(slice) != 0 {
		t.Error("Пустой список должен возвращать пустой обратный срез")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")
	slice = list.ToSliceReverse()
	if len(slice) != 3 || slice[0] != "three" || slice[1] != "two" || slice[2] != "one" {
		t.Error("ToSliceReverse возвращает неверный срез")
	}
}

func TestDLLPrintReverse(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PrintReverse()

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.PrintReverse()
}

func TestDLLStringReverse(t *testing.T) {
	list := NewDoublyLinkedList()
	str := list.StringReverse()
	if str != "Список пуст" {
		t.Error("StringReverse для пустого списка неверен")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	str = list.StringReverse()
	if !strings.Contains(str, "three") && !strings.Contains(str, "two") && !strings.Contains(str, "one") {
		t.Error("StringReverse неверно форматирует список")
	}
}

func TestDLLClear(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")

	list.Clear()
	if list.head != nil || list.tail != nil || !list.IsEmpty() {
		t.Error("Clear не очищает список")
	}
}

func TestDLLJSONSerialization(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	jsonStr, err := list.ToJSON()
	if err != nil || jsonStr == "" {
		t.Error("ToJSON не работает")
	}

	newList := NewDoublyLinkedList()
	err = newList.FromJSON(jsonStr)
	if err != nil || newList.Length() != 3 {
		t.Error("FromJSON не работает")
	}

	err = list.SaveToJSON("test_doubly.json")
	if err != nil {
		t.Error("SaveToJSON не работает")
	}
	defer os.Remove("test_doubly.json")

	loadedList := NewDoublyLinkedList()
	err = loadedList.LoadFromJSON("test_doubly.json")
	if err != nil || loadedList.Length() != 3 {
		t.Error("LoadFromJSON не работает")
	}
}

func TestBinarySerialization(t *testing.T) {
	list := NewDoublyLinkedList()
	list.InsertAtTail("one")
	list.InsertAtTail("two")
	list.InsertAtTail("three")

	binaryData, err := list.ToBinary()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinary не работает")
	}

	newList := NewDoublyLinkedList()
	err = newList.FromBinary(binaryData)
	if err != nil || newList.Length() != 3 {
		t.Error("FromBinary не работает")
	}

	err = list.SaveToBinary("test_doubly.bin")
	if err != nil {
		t.Error("SaveToBinary не работает")
	}
	defer os.Remove("test_doubly.bin")

	loadedList := NewDoublyLinkedList()
	err = loadedList.LoadFromBinary("test_doubly.bin")
	if err != nil || loadedList.Length() != 3 {
		t.Error("LoadFromBinary не работает")
	}
}

func TestDLLString(t *testing.T) {
	list := NewDoublyLinkedList()
	str := list.String()
	if str != "Список пуст" {
		t.Error("String для пустого списка неверен")
	}

	list.InsertAtTail("one")
	list.InsertAtTail("two")
	str = list.String()
	if !strings.Contains(str, "one <-> two") {
		t.Error("String неверно форматирует двусвязный список")
	}
}
