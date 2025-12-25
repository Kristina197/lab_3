package structs

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewOpenAddrHash(t *testing.T) {
	hash := NewOpenAddrHash()
	if hash.size != 0 || hash.capacity != 10 {
		t.Error("Новая хеш-таблица должна быть пустой с capacity 10")
	}

	hash2 := NewOpenAddrHashWithCapacity(20)
	if hash2.capacity != 20 || hash2.size != 0 {
		t.Error("Хеш-таблица с заданной capacity должна иметь правильную capacity")
	}

	hash3 := NewOpenAddrHashWithCapacity(0)
	if hash3.capacity != 10 {
		t.Error("Хеш-таблица с нулевой capacity должна использовать значение по умолчанию")
	}
}

func TestOAHPutAndGet(t *testing.T) {
	hash := NewOpenAddrHash()

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
}

func TestPutWhenFull(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(2)

	if !hash.Put("key1", "value1") {
		t.Error("Первая вставка должна работать")
	}

	if !hash.Put("key2", "value2") {
		t.Error("Вторая вставка должна работать")
	}

	if hash.Put("key3", "value3") {
		t.Error("Вставка в полную таблицу должна возвращать false")
	}
}

func TestOAHGetNonExistent(t *testing.T) {
	hash := NewOpenAddrHash()

	value, found := hash.Get("nonexistent")
	if found || value != "" {
		t.Error("Get должен возвращать false для несуществующего ключа")
	}
}

func TestOAHRemove(t *testing.T) {
	hash := NewOpenAddrHash()

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

	if !hash.Put("key3", "value3") {
		t.Error("Должна быть возможность вставить новый ключ после удаления")
	}
}

func TestOAHContains(t *testing.T) {
	hash := NewOpenAddrHash()

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

func TestOAHSizeAndIsEmpty(t *testing.T) {
	hash := NewOpenAddrHash()

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

func TestOAHCapacity(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(15)

	if hash.Capacity() != 15 {
		t.Error("Capacity должен возвращать правильную емкость")
	}
}

func TestIsFull(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(2)

	if hash.IsFull() {
		t.Error("Пустая таблица не должна быть полной")
	}

	hash.Put("key1", "value1")
	if hash.IsFull() {
		t.Error("Таблица с одним элементом не должна быть полной")
	}

	hash.Put("key2", "value2")
	if !hash.IsFull() {
		t.Error("Полная таблица должна возвращать true")
	}
}

func TestOAHClear(t *testing.T) {
	hash := NewOpenAddrHash()

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

func TestOAHKeys(t *testing.T) {
	hash := NewOpenAddrHash()

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

func TestOAHValues(t *testing.T) {
	hash := NewOpenAddrHash()

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

func TestOAHLoadFactor(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(5)

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

func TestOAHRehash(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(3)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	if hash.Put("key4", "value4") {
		t.Error("В полную таблицу нельзя вставить элемент без rehash")
	}

	//oldCapacity := hash.capacity
	hash.Rehash(10)

	if hash.capacity != 10 || hash.size != 3 {
		t.Error("Rehash должен изменять capacity и сохранять элементы")
	}

	if !hash.Contains("key1") || !hash.Contains("key2") || !hash.Contains("key3") {
		t.Error("Rehash должен сохранять все элементы")
	}

	if !hash.Put("key4", "value4") {
		t.Error("После rehash должна быть возможность вставлять новые элементы")
	}

	if hash.size != 4 {
		t.Error("Размер должен увеличиваться после вставки после rehash")
	}

	hash.Rehash(5)
	//if hash.capacity != 5 {
	//	t.Error("Rehash с меньшей capacity не должен изменять таблицу")
	//}
}

func TestOAHAutoRehash(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(4)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")

	oldCapacity := hash.capacity
	hash.AutoRehash(0.5)

	if hash.capacity != oldCapacity {
		t.Error("AutoRehash не должен выполняться при load factor <= threshold")
	}

	hash.Put("key3", "value3")
	hash.AutoRehash(0.5)

	if hash.capacity != 8 {
		t.Error("AutoRehash должен удваивать capacity при превышении threshold")
	}

	if hash.size != 3 || !hash.Contains("key1") || !hash.Contains("key2") || !hash.Contains("key3") {
		t.Error("AutoRehash должен сохранять все элементы")
	}
}

func TestOAHCollisionHandling(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(3)

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

	if !hash.Put("key4", "value4") {
		t.Error("Должна быть возможность вставить элемент на место удаленного")
	}
}

func TestOAHString(t *testing.T) {
	hash := NewOpenAddrHash()

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

func TestOAHPrint(t *testing.T) {
	hash := NewOpenAddrHash()

	hash.Print()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	hash.Print()
}

func TestOAHJSONSerialization(t *testing.T) {
	hash := NewOpenAddrHash()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	jsonStr, err := hash.ToJSON()
	if err != nil || jsonStr == "" {
		t.Error("ToJSON не работает")
	}

	newHash := NewOpenAddrHash()
	err = newHash.FromJSON(jsonStr)
	if err != nil || newHash.Size() != 3 {
		t.Error("FromJSON не работает")
	}

	if !newHash.Contains("key1") || !newHash.Contains("key2") || !newHash.Contains("key3") {
		t.Error("FromJSON должен восстанавливать все элементы")
	}

	err = hash.SaveToJSON("test_openaddr.json")
	if err != nil {
		t.Error("SaveToJSON не работает")
	}
	defer os.Remove("test_openaddr.json")

	loadedHash := NewOpenAddrHash()
	err = loadedHash.LoadFromJSON("test_openaddr.json")
	if err != nil || loadedHash.Size() != 3 {
		t.Error("LoadFromJSON не работает")
	}
}

func TestOAHBinarySerialization(t *testing.T) {
	hash := NewOpenAddrHash()

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	binaryData, err := hash.ToBinary()
	if err != nil || len(binaryData) == 0 {
		t.Error("ToBinary не работает")
	}

	newHash := NewOpenAddrHash()
	err = newHash.FromBinary(binaryData)
	if err != nil || newHash.Size() != 3 {
		t.Error("FromBinary не работает")
	}

	if !newHash.Contains("key1") || !newHash.Contains("key2") || !newHash.Contains("key3") {
		t.Error("FromBinary должен восстанавливать все элементы")
	}

	err = hash.SaveToBinary("test_openaddr.bin")
	if err != nil {
		t.Error("SaveToBinary не работает")
	}
	defer os.Remove("test_openaddr.bin")

	loadedHash := NewOpenAddrHash()
	err = loadedHash.LoadFromBinary("test_openaddr.bin")
	if err != nil || loadedHash.Size() != 3 {
		t.Error("LoadFromBinary не работает")
	}
}

func TestOAHComplexOperations(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(10)

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		if !hash.Put(key, value) {
			t.Error("Должна поддерживаться вставка множества элементов")
		}
	}

	if hash.Size() != 8 || hash.LoadFactor() != 0.8 {
		t.Error("Размер и load factor должны быть правильными")
	}

	for i := 0; i < 4; i++ {
		key := fmt.Sprintf("key%d", i)
		if !hash.Remove(key) {
			t.Error("Удаление должно работать для множества элементов")
		}
	}

	if hash.Size() != 4 {
		t.Error("Размер должен правильно обновляться при массовом удалении")
	}

	for i := 4; i < 8; i++ {
		key := fmt.Sprintf("key%d", i)
		if !hash.Contains(key) {
			t.Error("Оставшиеся элементы должны быть доступны")
		}
	}

	for i := 0; i < 4; i++ {
		key := fmt.Sprintf("newkey%d", i)
		value := fmt.Sprintf("newvalue%d", i)
		if !hash.Put(key, value) {
			t.Error("Должна быть возможность вставлять на места удаленных элементов")
		}
	}

	if hash.Size() != 8 || hash.IsFull() {
		t.Error("Таблица должна правильно управлять размером")
	}

	hash.Clear()
	if hash.Size() != 0 || !hash.IsEmpty() {
		t.Error("Clear должен полностью очищать таблицу")
	}
}

func TestDeletedSlotReuse(t *testing.T) {
	hash := NewOpenAddrHashWithCapacity(5)

	hash.Put("key1", "value1")
	hash.Put("key2", "value2")
	hash.Put("key3", "value3")

	hash.Remove("key2")

	if hash.Contains("key2") {
		t.Error("Удаленный ключ не должен существовать")
	}

	if !hash.Put("key4", "value4") {
		t.Error("Должна быть возможность использовать удаленный слот")
	}

	if hash.Size() != 3 {
		t.Error("Размер должен быть правильным после переиспользования удаленного слота")
	}

	if !hash.Contains("key1") || !hash.Contains("key3") || !hash.Contains("key4") {
		t.Error("Все существующие ключи должны быть доступны")
	}
}
