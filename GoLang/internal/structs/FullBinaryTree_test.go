package structs

import (
	"os"
	"strings"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree()
	if tree.root != nil {
		t.Error("Новое дерево должно быть пустым")
	}
}

func TestInsertNode(t *testing.T) {
	tree := NewTree()

	if !tree.InsertNode(10) {
		t.Error("Вставка корня не работает")
	}
	if tree.root.val != 10 {
		t.Error("Значение корня неверное")
	}

	if !tree.InsertNode(5) {
		t.Error("Вставка левого потомка не работает")
	}
	if !tree.InsertNode(15) {
		t.Error("Вставка правого потомка не работает")
	}

	if tree.InsertNode(5) {
		t.Error("Дубликаты не должны вставляться")
	}

	if tree.root.left.val != 5 || tree.root.right.val != 15 {
		t.Error("Вставка в BST не работает")
	}
}

func TestSearchNode(t *testing.T) {
	tree := NewTree()

	if tree.SearchNode(10) {
		t.Error("Поиск в пустом дереве должен возвращать false")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)

	if !tree.SearchNode(10) {
		t.Error("Поиск корня не работает")
	}
	if !tree.SearchNode(5) {
		t.Error("Поиск левого потомка не работает")
	}
	if !tree.SearchNode(15) {
		t.Error("Поиск правого потомка не работает")
	}
	if tree.SearchNode(99) {
		t.Error("Поиск несуществующего значения должен возвращать false")
	}
}

func TestIsFullBinary(t *testing.T) {
	tree := NewTree()

	if !tree.IsFullBinary() {
		t.Error("Пустое дерево должно быть полным бинарным")
	}

	tree.InsertNode(10)
	if !tree.IsFullBinary() {
		t.Error("Дерево с одним узлом должно быть полным бинарным")
	}

	tree.InsertNode(5)
	if tree.IsFullBinary() {
		t.Error("Дерево с одним потомком не должно быть полным бинарным")
	}

	tree.InsertNode(15)
	if !tree.IsFullBinary() {
		t.Error("Дерево с двумя потомками должно быть полным бинарным")
	}

	tree.InsertNode(3)
	tree.InsertNode(7)
	if !tree.IsFullBinary() {
		t.Error("Полное бинарное дерево должно определяться корректно")
	}

	tree2 := NewTree()
	tree2.root = NewTNodeWithChildren(10,
		NewTNodeWithChildren(5,
			NewTNode(3),
			NewTNode(7)),
		NewTNodeWithChildren(15,
			nil,
			NewTNode(20)))
	if tree2.IsFullBinary() {
		t.Error("Неполное дерево не должно определяться как полное")
	}
}

func TestPreOrder(t *testing.T) {
	tree := NewTree()

	if len(tree.PreOrder()) != 0 {
		t.Error("PreOrder пустого дерева должен быть пустым")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)
	tree.InsertNode(3)
	tree.InsertNode(7)

	result := tree.PreOrder()
	expected := []int{10, 5, 3, 7, 15}

	if len(result) != len(expected) {
		t.Error("PreOrder возвращает неверное количество элементов")
	}
	for i, v := range expected {
		if result[i] != v {
			t.Error("PreOrder возвращает неверный порядок")
		}
	}
}

func TestInOrder(t *testing.T) {
	tree := NewTree()

	if len(tree.InOrder()) != 0 {
		t.Error("InOrder пустого дерева должен быть пустым")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)
	tree.InsertNode(3)
	tree.InsertNode(7)

	result := tree.InOrder()
	expected := []int{3, 5, 7, 10, 15}

	if len(result) != len(expected) {
		t.Error("InOrder возвращает неверное количество элементов")
	}
	for i, v := range expected {
		if result[i] != v {
			t.Error("InOrder возвращает неверный порядок")
		}
	}
}

func TestPostOrder(t *testing.T) {
	tree := NewTree()

	if len(tree.PostOrder()) != 0 {
		t.Error("PostOrder пустого дерева должен быть пустым")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)
	tree.InsertNode(3)
	tree.InsertNode(7)

	result := tree.PostOrder()
	expected := []int{3, 7, 5, 15, 10}

	if len(result) != len(expected) {
		t.Error("PostOrder возвращает неверное количество элементов")
	}
	for i, v := range expected {
		if result[i] != v {
			t.Error("PostOrder возвращает неверный порядок")
		}
	}
}

func TestLevelOrder(t *testing.T) {
	tree := NewTree()

	if len(tree.LevelOrder()) != 0 {
		t.Error("LevelOrder пустого дерева должен быть пустым")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)
	tree.InsertNode(3)
	tree.InsertNode(7)

	result := tree.LevelOrder()
	expected := []int{10, 5, 15, 3, 7}

	if len(result) != len(expected) {
		t.Error("LevelOrder возвращает неверное количество элементов")
	}
	for i, v := range expected {
		if result[i] != v {
			t.Error("LevelOrder возвращает неверный порядок")
		}
	}
}

func TestHeight(t *testing.T) {
	tree := NewTree()

	if tree.Height() != 0 {
		t.Error("Высота пустого дерева должна быть 0")
	}

	tree.InsertNode(10)
	if tree.Height() != 1 {
		t.Error("Высота дерева с одним узлом должна быть 1")
	}

	tree.InsertNode(5)
	if tree.Height() != 2 {
		t.Error("Высота дерева с двумя узлами должна быть 2")
	}

	tree.InsertNode(15)
	tree.InsertNode(3)
	if tree.Height() != 3 {
		t.Error("Высота дерева должна вычисляться корректно")
	}
}

func TestSize(t *testing.T) {
	tree := NewTree()

	if tree.Size() != 0 {
		t.Error("Размер пустого дерева должен быть 0")
	}

	tree.InsertNode(10)
	if tree.Size() != 1 {
		t.Error("Размер дерева с одним узлом должен быть 1")
	}

	tree.InsertNode(5)
	tree.InsertNode(15)
	if tree.Size() != 3 {
		t.Error("Размер дерева должен вычисляться корректно")
	}
}

func TestIsEmpty(t *testing.T) {
	tree := NewTree()

	if !tree.IsEmpty() {
		t.Error("Новое дерево должно быть пустым")
	}

	tree.InsertNode(10)
	if tree.IsEmpty() {
		t.Error("Дерево с узлами не должно быть пустым")
	}
}

func TestGetRootValue(t *testing.T) {
	tree := NewTree()

	val, ok := tree.GetRootValue()
	if ok || val != 0 {
		t.Error("GetRootValue пустого дерева должен возвращать false")
	}

	tree.InsertNode(42)
	val, ok = tree.GetRootValue()
	if !ok || val != 42 {
		t.Error("GetRootValue должен возвращать значение корня")
	}
}

func TestPrintFunctions(t *testing.T) {
	tree := NewTree()

	tree.PrintPreOrder()
	tree.PrintInOrder()
	tree.PrintPostOrder()
	tree.PrintLevelOrder()

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)

	tree.PrintPreOrder()
	tree.PrintInOrder()
	tree.PrintPostOrder()
	tree.PrintLevelOrder()
}

func TestFBTString(t *testing.T) {
	tree := NewTree()

	str := tree.String()
	if str != "Дерево пустое" {
		t.Error("String пустого дерева неверен")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	str = tree.String()
	if !strings.Contains(str, "10") || !strings.Contains(str, "5") {
		t.Error("String неверно форматирует дерево")
	}
}

func TestFBTJSONSerialization(t *testing.T) {
	tree := NewTree()

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)
	tree.InsertNode(3)
	tree.InsertNode(7)

	jsonStr, err := tree.ToJSON()
	if err != nil || jsonStr == "" {
		t.Error("ToJSON не работает")
	}

	newTree := NewTree()
	err = newTree.FromJSON(jsonStr)
	if err != nil || newTree.Size() != 5 {
		t.Error("FromJSON не работает")
	}

	err = tree.SaveToJSON("test_tree.json")
	if err != nil {
		t.Error("SaveToJSON не работает")
	}
	defer os.Remove("test_tree.json")

	loadedTree := NewTree()
	err = loadedTree.LoadFromJSON("test_tree.json")
	if err != nil || loadedTree.Size() != 5 {
		t.Error("LoadFromJSON не работает")
	}
}

func TestFBTBinarySerialization(t *testing.T) {
	tree := NewTree()

	binaryData, err := tree.ToBinary()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinary для пустого дерева не работает")
	}

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)

	binaryData, err = tree.ToBinary()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinary не работает")
	}

	newTree := NewTree()
	err = newTree.FromBinary(binaryData)
	if err != nil || newTree.Size() != 3 {
		t.Error("FromBinary не работает")
	}

	err = tree.SaveToBinary("test_tree.bin")
	if err != nil {
		t.Error("SaveToBinary не работает")
	}
	defer os.Remove("test_tree.bin")

	loadedTree := NewTree()
	err = loadedTree.LoadFromBinary("test_tree.bin")
	if err != nil || loadedTree.Size() != 3 {
		t.Error("LoadFromBinary не работает")
	}
}

func TestBinaryCompactSerialization(t *testing.T) {
	tree := NewTree()

	tree.InsertNode(10)
	tree.InsertNode(5)
	tree.InsertNode(15)
	tree.InsertNode(3)
	tree.InsertNode(7)

	binaryData, err := tree.ToBinaryCompact()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinaryCompact не работает")
	}

	newTree := NewTree()
	err = newTree.FromBinaryCompact(binaryData)
	if err != nil || newTree.Size() != 5 {
		t.Error("FromBinaryCompact не работает")
	}

	err = tree.SaveToBinaryCompact("test_tree_compact.bin")
	if err != nil {
		t.Error("SaveToBinaryCompact не работает")
	}
	defer os.Remove("test_tree_compact.bin")

	loadedTree := NewTree()
	err = loadedTree.LoadFromBinaryCompact("test_tree_compact.bin")
	if err != nil || loadedTree.Size() != 5 {
		t.Error("LoadFromBinaryCompact не работает")
	}
}

func TestFullBinaryComplex(t *testing.T) {
	tree := NewTree()

	tree.root = NewTNodeWithChildren(1,
		NewTNodeWithChildren(2,
			NewTNode(4),
			NewTNode(5)),
		NewTNodeWithChildren(3,
			NewTNode(6),
			NewTNode(7)))

	if !tree.IsFullBinary() {
		t.Error("Полное бинарное дерево должно определяться как полное")
	}

	tree.root.left.left = nil

	if tree.IsFullBinary() {
		t.Error("Неполное дерево не должно определяться как полное")
	}
}
