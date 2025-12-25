package structs

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewChainingHash(t *testing.T) {
	hash := NewChainingHash()
	if hash.size != 0 || hash.capacity != 10 {
		t.Error("Новая хеш-таблица должна быть пустой с capacity 10")
	}

	hash2 := NewChainingHashWithCapacity(20)
	if hash2.capacity != 20 || hash2.size != 0 {
		t.Error("Хеш-таблица с заданной capacity должна иметь правильную capacity")
	}

	hash3 := NewChainingHashWithCapacity(0)
	if hash3.capacity != 10 {
		t.Error("Хеш-таблица с нулевой capacity должна использовать значение по умолчанию")
	}
}

func TestPutAndGet(t *testing.T) {
	hash := NewChainingHash()

	if !hash.Put("key1", "value1") {
		t.Error("Put должен возвращать true при успешной вставке")
	}

	if hash.size != 1 {
		t.Error("Размер должен увеличиваться после вставки")
	}

	value, found := hash.Get("key1")
	if !found || value != "value1" {
		t.Error("Get должен возвращать вставленное значение")
	}

	if !hash.Put("key1", "updated") {
		t.Error("Put должен обновлять существующий ключ")
	}

	value, found = hash.Get("key1")
	if !found || value != "updated" {
		t.Error("Get должен возвращать обновленное значение")
	}

	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	if hash.size != 3 {
		t.Error("Размер должен правильно учитывать количество элементов")
	}
}

func TestGetNonExistent(t *testing.T) {
	hash := NewChainingHash()

	value, found := hash.Get("nonexistent")
	if found || value != "" {
		t.Error("Get должен возвращать false для несуществующего ключа")
	}
}

func TestRemove(t *testing.T) {
	hash := NewChainingHash()

	if hash.Remove("nonexistent") {
		t.Error("Remove должен возвращать false для несуществующего ключа")
	}

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")

	if !hash.Remove("key1") {
		t.Error("Remove должен возвращать true при успешном удалении")
	}

	if hash.size != 1 {
		t.Error("Размер должен уменьшаться после удаления")
	}

	if hash.Contains("key1") {
		t.Error("Удаленный ключ не должен существовать")
	}

	if !hash.Contains("key2") {
		t.Error("Остальные ключи должны оставаться после удаления")
	}

	hash.Put("key1", "value1")
	hash.Put("key3", "value3")

	if !hash.Remove("key2") || !hash.Remove("key1") || !hash.Remove("key3") {
		t.Error("Remove должен работать для всех элементов")
	}

	if hash.size != 0 {
		t.Error("Все элементы должны быть удалены")
	}
}

func TestContains(t *testing.T) {
	hash := NewChainingHash()

	if hash.Contains("any") {
		t.Error("Contains должен возвращать false для пустой таблицы")
	}

	hash.Put("key1", "value1")
	if !hash.Contains("key1") {
		t.Error("Contains должен находить существующие ключи")
	}

	if hash.Contains("key2") {
		t.Error("Contains не должен находить несуществующие ключи")
	}
}

func TestSizeAndIsEmpty(t *testing.T) {
	hash := NewChainingHash()

	if !hash.IsEmpty() || hash.Size() != 0 {
		t.Error("Новая таблица должна быть пустой")
	}

	hash.Put("key1", "value1")
	if hash.IsEmpty() || hash.Size() != 1 {
		t.Error("Таблица не должна быть пустой после вставки")
	}

	hash.Remove("key1")
	if !hash.IsEmpty() || hash.Size() != 0 {
		t.Error("Таблица должна быть пустой после удаления всех элементов")
	}
}

func TestCapacity(t *testing.T) {
	hash := NewChainingHashWithCapacity(15)

	if hash.Capacity() != 15 {
		t.Error("Capacity должен возвращать правильную емкость")
	}
}

func TestClear(t *testing.T) {
	hash := NewChainingHash()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	hash.Clear()

	if hash.size != 0 || !hash.IsEmpty() {
		t.Error("Clear должен очищать таблицу")
	}

	if hash.Contains("key1") || hash.Contains("key2") || hash.Contains("key3") {
		t.Error("Clear должен удалять все элементы")
	}
}

func TestKeys(t *testing.T) {
	hash := NewChainingHash()

	keys := hash.Keys()
	if len(keys) != 0 {
		t.Error("Keys пустой таблицы должен быть пустым")
	}

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	keys = hash.Keys()
	if len(keys) != 3 {
		t.Error("Keys должен возвращать все ключи")
	}

	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}

	if !keyMap["key1"] || !keyMap["key2"] || !keyMap["key3"] {
		t.Error("Keys должен содержать все добавленные ключи")
	}
}

func TestValues(t *testing.T) {
	hash := NewChainingHash()

	values := hash.Values()
	if len(values) != 0 {
		t.Error("Values пустой таблицы должен быть пустым")
	}

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	values = hash.Values()
	if len(values) != 3 {
		t.Error("Values должен возвращать все значения")
	}

	valueMap := make(map[string]bool)
	for _, value := range values {
		valueMap[value] = true
	}

	if !valueMap["value1"] || !valueMap["value2"] || !valueMap["value3"] {
		t.Error("Values должен содержать все добавленные значения")
	}
}

func TestLoadFactor(t *testing.T) {
	hash := NewChainingHashWithCapacity(5)

	if hash.LoadFactor() != 0.0 {
		t.Error("LoadFactor пустой таблицы должен быть 0")
	}

	hash.Put("key1", "value1")
	if hash.LoadFactor() != 0.2 {
		t.Error("LoadFactor должен правильно вычисляться")
	}

	hash.Put("key2", "value2")
	hash.Put("key3", "value3")
	if hash.LoadFactor() != 0.6 {
		t.Error("LoadFactor должен правильно вычисляться")
	}
}

func TestRehash(t *testing.T) {
	hash := NewChainingHashWithCapacity(3)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	//oldCapacity := hash.capacity
	hash.Rehash(10)

	if hash.capacity != 10 || hash.size != 3 {
		t.Error("Rehash должен изменять capacity и сохранять элементы")
	}

	if !hash.Contains("key1") || !hash.Contains("key2") || !hash.Contains("key3") {
		t.Error("Rehash должен сохранять все элементы")
	}

	value, found := hash.Get("key1")
	if !found || value != "value1" {
		t.Error("Элементы должны быть доступны после rehash")
	}

	hash.Rehash(0)
	if hash.capacity != 10 {
		t.Error("Rehash с неположительной capacity не должен изменять таблицу")
	}
}

func TestAutoRehash(t *testing.T) {
	hash := NewChainingHashWithCapacity(3)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")

	//oldCapacity := hash.capacity
	hash.AutoRehash(0.5)

	//if hash.capacity != oldCapacity {
	//	t.Error("AutoRehash не должен выполняться при load factor <= threshold")
	//}

	hash.Put("key3", "value3")
	hash.AutoRehash(0.5)

	if hash.capacity != 6 {
		t.Error("AutoRehash должен удваивать capacity при превышении threshold")
	}

	if hash.size != 3 || !hash.Contains("key1") || !hash.Contains("key2") || !hash.Contains("key3") {
		t.Error("AutoRehash должен сохранять все элементы")
	}
}

func TestCollisionHandling(t *testing.T) {
	hash := NewChainingHashWithCapacity(2)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	if hash.size != 3 {
		t.Error("Таблица должна обрабатывать коллизии")
	}

	if !hash.Contains("key1") || !hash.Contains("key2") || !hash.Contains("key3") {
		t.Error("Все ключи должны быть доступны при коллизиях")
	}

	hash.Remove("key2")
	if hash.size != 2 || hash.Contains("key2") || !hash.Contains("key1") || !hash.Contains("key3") {
		t.Error("Удаление должно работать при коллизиях")
	}
}

func TestCHMString(t *testing.T) {
	hash := NewChainingHash()

	str := hash.String()
	if !strings.Contains(str, "size: 0") || !strings.Contains(str, "capacity: 10") {
		t.Error("String должен содержать информацию о таблице")
	}

	hash.Put("key1", "value1")
	str = hash.String()
	if !strings.Contains(str, "key1") || !strings.Contains(str, "value1") {
		t.Error("String должен содержать элементы таблицы")
	}
}

func TestPrint(t *testing.T) {
	hash := NewChainingHash()

	hash.Print()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	hash.Print()
}

func TestJSONSerialization(t *testing.T) {
	hash := NewChainingHash()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	jsonStr, err := hash.ToJSON()
	if err != nil || jsonStr == "" {
		t.Error("ToJSON не работает")
	}

	newHash := NewChainingHash()
	err = newHash.FromJSON(jsonStr)
	if err != nil || newHash.Size() != 3 {
		t.Error("FromJSON не работает")
	}

	if !newHash.Contains("key1") || !newHash.Contains("key2") || !newHash.Contains("key3") {
		t.Error("FromJSON должен восстанавливать все элементы")
	}

	err = hash.SaveToJSON("test_hash.json")
	if err != nil {
		t.Error("SaveToJSON не работает")
	}
	defer os.Remove("test_hash.json")

	loadedHash := NewChainingHash()
	err = loadedHash.LoadFromJSON("test_hash.json")
	if err != nil || loadedHash.Size() != 3 {
		t.Error("LoadFromJSON не работает")
	}
}

func TestCHMBinarySerialization(t *testing.T) {
	hash := NewChainingHash()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	binaryData, err := hash.ToBinary()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinary не работает")
	}

	newHash := NewChainingHash()
	err = newHash.FromBinary(binaryData)
	if err != nil || newHash.Size() != 3 {
		t.Error("FromBinary не работает")
	}

	if !newHash.Contains("key1") || !newHash.Contains("key2") || !newHash.Contains("key3") {
		t.Error("FromBinary должен восстанавливать все элементы")
	}

	err = hash.SaveToBinary("test_hash.bin")
	if err != nil {
		t.Error("SaveToBinary не работает")
	}
	defer os.Remove("test_hash.bin")

	loadedHash := NewChainingHash()
	err = loadedHash.LoadFromBinary("test_hash.bin")
	if err != nil || loadedHash.Size() != 3 {
		t.Error("LoadFromBinary не работает")
	}
}

func TestBinaryWithBucketsSerialization(t *testing.T) {
	hash := NewChainingHashWithCapacity(3)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")
	hash.Put("key4", "value4")

	binaryData, err := hash.ToBinaryWithBuckets()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinaryWithBuckets не работает")
	}

	newHash := NewChainingHash()
	err = newHash.FromBinaryWithBuckets(binaryData)
	if err != nil || newHash.Size() != 4 {
		t.Error("FromBinaryWithBuckets не работает")
	}

	if !newHash.Contains("key1") || !newHash.Contains("key2") || !newHash.Contains("key3") || !newHash.Contains("key4") {
		t.Error("FromBinaryWithBuckets должен восстанавливать все элементы")
	}
}

func TestComplexOperations(t *testing.T) {
	hash := NewChainingHashWithCapacity(5)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		hash.Put(key, value)
	}

	if hash.Size() != 100 {
		t.Error("Должна поддерживаться вставка множества элементов")
	}

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		if !hash.Remove(key) {
			t.Error("Удаление должно работать для множества элементов")
		}
	}

	if hash.Size() != 50 {
		t.Error("Размер должен правильно обновляться при массовом удалении")
	}

	for i := 50; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		if !hash.Contains(key) {
			t.Error("Оставшиеся элементы должны быть доступны")
		}
	}

	hash.Clear()
	if hash.Size() != 0 || !hash.IsEmpty() {
		t.Error("Clear должен полностью очищать таблицу")
	}
}
