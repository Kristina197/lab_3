package structs

import (
	"os"
	"testing"
)

func TestQueue_BasicOperations(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Error("новая очередь должна быть пустой")
	}

	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	if q.Length() != 3 {
		t.Errorf("ожидается длина 3, получено %d", q.Length())
	}

	if q.Peek() != "a" {
		t.Errorf("ожидается 'a', получено %s", q.Peek())
	}

	if dequeued := q.Dequeue(); dequeued != "a" {
		t.Errorf("ожидается 'a', получено %s", dequeued)
	}

	if q.Length() != 2 {
		t.Errorf("ожидается длина 2, получено %d", q.Length())
	}

	if q.Peek() != "b" {
		t.Errorf("ожидается 'b', получено %s", q.Peek())
	}

	q.Dequeue()
	q.Dequeue()

	if !q.IsEmpty() {
		t.Error("очередь должна быть пустой")
	}

	if dequeued := q.Dequeue(); dequeued != "" {
		t.Errorf("ожидается пустая строка, получено %s", dequeued)
	}
}

func TestQueue_ToFromSlice(t *testing.T) {
	q := NewQueue()
	q.Enqueue("x")
	q.Enqueue("y")
	q.Enqueue("z")

	slice := q.ToSlice()
	if len(slice) != 3 {
		t.Errorf("ожидается длина 3, получено %d", len(slice))
	}

	expected := []string{"x", "y", "z"}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("на позиции %d: ожидается '%s', получено '%s'", i, val, slice[i])
		}
	}

	q2 := NewQueue()
	q2.FromSlice(slice)

	if q2.Length() != 3 {
		t.Errorf("ожидается длина 3, получено %d", q2.Length())
	}

	for _, val := range expected {
		if dequeued := q2.Dequeue(); dequeued != val {
			t.Errorf("ожидается '%s', получено '%s'", val, dequeued)
		}
	}
}

func TestQueue_JSON(t *testing.T) {
	q := NewQueue()
	q.Enqueue("test1")
	q.Enqueue("test2")

	jsonStr, err := q.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON не удалось: %v", err)
	}

	q2 := NewQueue()
	err = q2.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("FromJSON не удалось: %v", err)
	}

	if q2.Length() != 2 {
		t.Errorf("ожидается длина 2, получено %d", q2.Length())
	}

	if q2.Dequeue() != "test1" || q2.Dequeue() != "test2" {
		t.Error("значения не совпадают")
	}
}

func TestQueue_FileJSON(t *testing.T) {
	q := NewQueue()
	q.Enqueue("file1")
	q.Enqueue("file2")

	filename := "test_queue.json"
	defer os.Remove(filename)

	err := q.SaveToJSON(filename)
	if err != nil {
		t.Fatalf("SaveToJSON не удалось: %v", err)
	}

	q2 := NewQueue()
	err = q2.LoadFromJSON(filename)
	if err != nil {
		t.Fatalf("LoadFromJSON не удалось: %v", err)
	}

	if q2.Length() != 2 {
		t.Errorf("ожидается длина 2, получено %d", q2.Length())
	}
}

func TestQueue_Binary(t *testing.T) {
	q := NewQueue()
	q.Enqueue("bin1")
	q.Enqueue("bin2")

	data, err := q.ToBinary()
	if err != nil {
		t.Fatalf("ToBinary не удалось: %v", err)
	}

	q2 := NewQueue()
	err = q2.FromBinary(data)
	if err != nil {
		t.Fatalf("FromBinary не удалось: %v", err)
	}

	if q2.Length() != 2 {
		t.Errorf("ожидается длина 2, получено %d", q2.Length())
	}
}

func TestQueue_FileBinary(t *testing.T) {
	q := NewQueue()
	q.Enqueue("data1")
	q.Enqueue("data2")

	filename := "test_queue.bin"
	defer os.Remove(filename)

	err := q.SaveToBinary(filename)
	if err != nil {
		t.Fatalf("SaveToBinary не удалось: %v", err)
	}

	q2 := NewQueue()
	err = q2.LoadFromBinary(filename)
	if err != nil {
		t.Fatalf("LoadFromBinary не удалось: %v", err)
	}

	if q2.Length() != 2 {
		t.Errorf("ожидается длина 2, получено %d", q2.Length())
	}
}

func TestQueue_Clear(t *testing.T) {
	q := NewQueue()
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	q.Clear()

	if !q.IsEmpty() {
		t.Error("очередь должна быть пустой после Clear")
	}

	if q.Length() != 0 {
		t.Errorf("ожидается длина 0, получено %d", q.Length())
	}
}

func TestQueue_String(t *testing.T) {
	q := NewQueue()
	q.Print()

	if q.String() != "Очередь пустая" {
		t.Errorf("неверное строковое представление пустой очереди")
	}

	q.Enqueue("first")
	q.Enqueue("second")

	q.Print()
	expected := "front -> first -> second -> rear"
	if q.String() != expected {
		t.Errorf("ожидается '%s', получено '%s'", expected, q.String())
	}
}
