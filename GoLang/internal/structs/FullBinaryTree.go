package structs

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type TNode struct {
	val   int
	left  *TNode
	right *TNode
}

func NewTNode(value int) *TNode {
	return &TNode{
		val:   value,
		left:  nil,
		right: nil,
	}
}

func NewTNodeWithChildren(value int, leftPtr *TNode, rightPtr *TNode) *TNode {
	return &TNode{
		val:   value,
		left:  leftPtr,
		right: rightPtr,
	}
}

type Tree struct {
	root *TNode
}

func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
}

func (t *Tree) searchElement(node *TNode, val int) *TNode {
	if node == nil {
		return nil
	}
	if node.val == val {
		return node
	}

	leftSearch := t.searchElement(node.left, val)
	if leftSearch != nil {
		return leftSearch
	}

	return t.searchElement(node.right, val)
}

func (t *Tree) fullBinaryCheck(node *TNode) bool {
	if node == nil {
		return true
	}
	if node.left == nil && node.right == nil {
		return true
	}

	if node.left != nil && node.right != nil {
		return t.fullBinaryCheck(node.left) && t.fullBinaryCheck(node.right)
	}
	return false
}

func (t *Tree) InsertNode(val int) bool {
	newNode := NewTNode(val)

	if t.root == nil {
		t.root = newNode
		return true
	}

	curr := t.root
	for curr != nil {
		if val > curr.val {
			if curr.right != nil {
				curr = curr.right
			} else {
				curr.right = newNode
				return true
			}
		} else if val < curr.val {
			if curr.left != nil {
				curr = curr.left
			} else {
				curr.left = newNode
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func (t *Tree) SearchNode(val int) bool {
	foundNode := t.searchElement(t.root, val)
	return foundNode != nil
}

func (t *Tree) IsFullBinary() bool {
	return t.fullBinaryCheck(t.root)
}

func (t *Tree) PreOrder() []int {
	result := make([]int, 0)
	t.preorderTraverse(t.root, &result)
	return result
}

func (t *Tree) preorderTraverse(node *TNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.val)
	t.preorderTraverse(node.left, result)
	t.preorderTraverse(node.right, result)
}

func (t *Tree) InOrder() []int {
	result := make([]int, 0)
	t.inorderTraverse(t.root, &result)
	return result
}

func (t *Tree) inorderTraverse(node *TNode, result *[]int) {
	if node == nil {
		return
	}
	t.inorderTraverse(node.left, result)
	*result = append(*result, node.val)
	t.inorderTraverse(node.right, result)
}

func (t *Tree) PostOrder() []int {
	result := make([]int, 0)
	t.postorderTraverse(t.root, &result)
	return result
}

func (t *Tree) postorderTraverse(node *TNode, result *[]int) {
	if node == nil {
		return
	}
	t.postorderTraverse(node.left, result)
	t.postorderTraverse(node.right, result)
	*result = append(*result, node.val)
}

func (t *Tree) LevelOrder() []int {
	if t.root == nil {
		return []int{}
	}

	result := make([]int, 0)
	queue := []*TNode{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.val)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return result
}

func (t *Tree) Height() int {
	return t.calculateHeight(t.root)
}

func (t *Tree) calculateHeight(node *TNode) int {
	if node == nil {
		return 0
	}

	leftHeight := t.calculateHeight(node.left)
	rightHeight := t.calculateHeight(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (t *Tree) Size() int {
	return t.countNodes(t.root)
}

func (t *Tree) countNodes(node *TNode) int {
	if node == nil {
		return 0
	}
	return 1 + t.countNodes(node.left) + t.countNodes(node.right)
}

func (t *Tree) IsEmpty() bool {
	return t.root == nil
}

func (t *Tree) GetRootValue() (int, bool) {
	if t.root == nil {
		return 0, false
	}
	return t.root.val, true
}

func (t *Tree) PrintPreOrder() {
	elements := t.PreOrder()
	fmt.Println("PreOrder:", strings.Trim(fmt.Sprint(elements), "[]"))
}

func (t *Tree) PrintInOrder() {
	elements := t.InOrder()
	fmt.Println("InOrder:", strings.Trim(fmt.Sprint(elements), "[]"))
}

func (t *Tree) PrintPostOrder() {
	elements := t.PostOrder()
	fmt.Println("PostOrder:", strings.Trim(fmt.Sprint(elements), "[]"))
}

func (t *Tree) PrintLevelOrder() {
	elements := t.LevelOrder()
	fmt.Println("LevelOrder:", strings.Trim(fmt.Sprint(elements), "[]"))
}

func (t *Tree) String() string {
	if t.root == nil {
		return "Дерево пустое"
	}

	elements := t.LevelOrder()
	return fmt.Sprintf("Дерево (обход по уровням): %v", elements)
}

// СЕРИАЛИЗАЦИЯ / ДЕСЕРИАЛИЗАЦИЯ

// TreeNodeJSON вспомогательная структура для JSON сериализации
type TreeNodeJSON struct {
	Value int           `json:"value"`
	Left  *TreeNodeJSON `json:"left,omitempty"`
	Right *TreeNodeJSON `json:"right,omitempty"`
}

// MarshalJSON реализует интерфейс json.Marshaler
func (t *Tree) MarshalJSON() ([]byte, error) {
	if t.root == nil {
		return json.Marshal(nil)
	}

	rootJSON := t.serializeNodeToJSON(t.root)
	return json.Marshal(rootJSON)
}

// serializeNodeToJSON рекурсивно преобразует узел в JSON структуру
func (t *Tree) serializeNodeToJSON(node *TNode) *TreeNodeJSON {
	if node == nil {
		return nil
	}

	return &TreeNodeJSON{
		Value: node.val,
		Left:  t.serializeNodeToJSON(node.left),
		Right: t.serializeNodeToJSON(node.right),
	}
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (t *Tree) UnmarshalJSON(data []byte) error {
	// Проверяем, не пустое ли дерево
	if string(data) == "null" {
		t.root = nil
		return nil
	}

	var rootJSON TreeNodeJSON
	if err := json.Unmarshal(data, &rootJSON); err != nil {
		return err
	}

	t.root = t.deserializeNodeFromJSON(&rootJSON)
	return nil
}

// deserializeNodeFromJSON рекурсивно восстанавливает дерево из JSON
func (t *Tree) deserializeNodeFromJSON(nodeJSON *TreeNodeJSON) *TNode {
	if nodeJSON == nil {
		return nil
	}

	node := NewTNode(nodeJSON.Value)
	node.left = t.deserializeNodeFromJSON(nodeJSON.Left)
	node.right = t.deserializeNodeFromJSON(nodeJSON.Right)

	return node
}

// ToJSON сериализует дерево в JSON строку
func (t *Tree) ToJSON() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в дерево
func (t *Tree) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), t)
}

// SaveToJSON сохраняет дерево в JSON файл
func (t *Tree) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(t)
}

// LoadFromJSON загружает дерево из JSON файла
func (t *Tree) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(t)
}

// ToBinary сериализует дерево в бинарный формат
func (t *Tree) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	// Используем обход в ширину для сериализации
	if t.root == nil {
		if err := binary.Write(&buf, binary.LittleEndian, int32(0)); err != nil {
			return nil, fmt.Errorf("ошибка записи маркера пустого дерева: %w", err)
		}
		return buf.Bytes(), nil
	}

	// Записываем 1 как признак непустого дерева
	if err := binary.Write(&buf, binary.LittleEndian, int32(1)); err != nil {
		return nil, fmt.Errorf("ошибка записи маркера непустого дерева: %w", err)
	}

	queue := []*TNode{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Записываем значение узла
		if err := binary.Write(&buf, binary.LittleEndian, int32(node.val)); err != nil {
			return nil, fmt.Errorf("ошибка записи значения узла: %w", err)
		}

		hasLeft := int32(0)
		if node.left != nil {
			hasLeft = 1
		}
		hasRight := int32(0)
		if node.right != nil {
			hasRight = 1
		}

		if err := binary.Write(&buf, binary.LittleEndian, hasLeft); err != nil {
			return nil, fmt.Errorf("ошибка записи маркера левого потомка: %w", err)
		}
		if err := binary.Write(&buf, binary.LittleEndian, hasRight); err != nil {
			return nil, fmt.Errorf("ошибка записи маркера правого потомка: %w", err)
		}

		// Добавляем потомков в очередь
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует дерево из бинарного формата
func (t *Tree) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	// Читаем маркер пустого/непустого дерева
	var isEmpty int32
	if err := binary.Read(buf, binary.LittleEndian, &isEmpty); err != nil {
		return fmt.Errorf("ошибка чтения маркера дерева: %w", err)
	}

	if isEmpty == 0 {
		t.root = nil
		return nil
	}

	// Восстанавливаем дерево обходом в ширину
	type NodeWithParent struct {
		node   *TNode
		parent *TNode
		isLeft bool
	}

	// Читаем корень
	var rootVal int32
	if err := binary.Read(buf, binary.LittleEndian, &rootVal); err != nil {
		return fmt.Errorf("ошибка чтения значения корня: %w", err)
	}

	t.root = NewTNode(int(rootVal))

	// Читаем маркеры потомков для корня
	var rootHasLeft, rootHasRight int32
	if err := binary.Read(buf, binary.LittleEndian, &rootHasLeft); err != nil {
		return fmt.Errorf("ошибка чтения маркера левого потомка корня: %w", err)
	}
	if err := binary.Read(buf, binary.LittleEndian, &rootHasRight); err != nil {
		return fmt.Errorf("ошибка чтения маркера правого потомка корня: %w", err)
	}

	queue := []NodeWithParent{
		{node: t.root, parent: nil, isLeft: false},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Если это не корень, читаем значение узла
		if current.parent != nil {
			var nodeVal int32
			if err := binary.Read(buf, binary.LittleEndian, &nodeVal); err != nil {
				return fmt.Errorf("ошибка чтения значения узла: %w", err)
			}
			current.node.val = int(nodeVal)
		}

		// Читаем маркеры потомков для текущего узла
		var hasLeft, hasRight int32
		if current.parent == nil {
			// Для корня маркеры уже прочитаны
			hasLeft = rootHasLeft
			hasRight = rootHasRight
		} else {
			if err := binary.Read(buf, binary.LittleEndian, &hasLeft); err != nil {
				return fmt.Errorf("ошибка чтения маркера левого потомка: %w", err)
			}
			if err := binary.Read(buf, binary.LittleEndian, &hasRight); err != nil {
				return fmt.Errorf("ошибка чтения маркера правого потомка: %w", err)
			}
		}

		// Создаем левого потомка если нужно
		if hasLeft == 1 {
			leftNode := NewTNode(0) // Значение будет прочитано позже
			current.node.left = leftNode
			queue = append(queue, NodeWithParent{
				node:   leftNode,
				parent: current.node,
				isLeft: true,
			})
		}

		// Создаем правого потомка если нужно
		if hasRight == 1 {
			rightNode := NewTNode(0) // Значение будет прочитано позже
			current.node.right = rightNode
			queue = append(queue, NodeWithParent{
				node:   rightNode,
				parent: current.node,
				isLeft: false,
			})
		}
	}

	return nil
}

// SaveToBinary сохраняет дерево в бинарный файл
func (t *Tree) SaveToBinary(filename string) error {
	data, err := t.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает дерево из бинарного файла
func (t *Tree) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return t.FromBinary(data)
}

// ToBinaryCompact сериализует дерево в компактный бинарный формат (альтернативный вариант)
func (t *Tree) ToBinaryCompact() ([]byte, error) {
	var buf bytes.Buffer

	// Используем префиксный обход (pre-order) с маркерами null
	if err := t.serializePreOrderCompact(t.root, &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *Tree) serializePreOrderCompact(node *TNode, buf *bytes.Buffer) error {
	if node == nil {
		// Записываем специальное значение как маркер null
		nullMarker := int32(-1 << 31) // Минимальное значение int32
		if err := binary.Write(buf, binary.LittleEndian, nullMarker); err != nil {
			return fmt.Errorf("ошибка записи маркера null: %w", err)
		}
		return nil
	}

	// Записываем значение узла
	if err := binary.Write(buf, binary.LittleEndian, int32(node.val)); err != nil {
		return fmt.Errorf("ошибка записи значения узла: %w", err)
	}

	// Рекурсивно сериализуем левое и правое поддеревья
	if err := t.serializePreOrderCompact(node.left, buf); err != nil {
		return err
	}
	if err := t.serializePreOrderCompact(node.right, buf); err != nil {
		return err
	}

	return nil
}

// FromBinaryCompact десериализует дерево из компактного бинарного формата
func (t *Tree) FromBinaryCompact(data []byte) error {
	buf := bytes.NewReader(data)

	// Вспомогательная функция для рекурсивного восстановления
	var deserialize func() (*TNode, error)
	deserialize = func() (*TNode, error) {
		var val int32
		if err := binary.Read(buf, binary.LittleEndian, &val); err != nil {
			return nil, fmt.Errorf("ошибка чтения значения: %w", err)
		}

		// Проверяем маркер null
		if val == -1<<31 {
			return nil, nil
		}

		node := NewTNode(int(val))

		left, err := deserialize()
		if err != nil {
			return nil, err
		}
		node.left = left

		right, err := deserialize()
		if err != nil {
			return nil, err
		}
		node.right = right

		return node, nil
	}

	root, err := deserialize()
	if err != nil {
		return err
	}
	t.root = root

	return nil
}

// SaveToBinaryCompact сохраняет дерево в компактный бинарный файл
func (t *Tree) SaveToBinaryCompact(filename string) error {
	data, err := t.ToBinaryCompact()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinaryCompact загружает дерево из компактного бинарного файла
func (t *Tree) LoadFromBinaryCompact(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return t.FromBinaryCompact(data)
}
