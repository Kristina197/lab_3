package structs

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Array struct {
	Data     []string `json:"data"`
	Size     int      `json:"size"`
	Capacity int      `json:"capacity,omitempty"`
}

func NewArray() *Array {
	return &Array{
		Data:     make([]string, 10),
		Size:     0,
		Capacity: 10,
	}
}

func NewArrayWithCapacity(cap int) *Array {
	return &Array{
		Data:     make([]string, cap),
		Size:     0,
		Capacity: cap,
	}
}

func (a *Array) doubleCapacity() {
	newCapacity := a.Capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}

	newData := make([]string, newCapacity)
	copy(newData, a.Data[:a.Size])
	a.Data = newData
	a.Capacity = newCapacity
}

func (a *Array) Cap() int {
	return a.Capacity
}

func (a *Array) Length() int {
	return a.Size
}

func (a *Array) PushBack(value string) bool {
	if a.Size >= a.Capacity {
		a.doubleCapacity()
	}
	a.Data[a.Size] = value
	a.Size++
	return true
}

func (a *Array) PushByIndex(value string, index int) bool {
	if index > a.Size {
		return false
	}

	if a.Size >= a.Capacity {
		a.doubleCapacity()
	}

	for i := a.Size; i > index; i-- {
		a.Data[i] = a.Data[i-1]
	}

	a.Data[index] = value
	a.Size++
	return true
}

func (a *Array) GetByIndex(index int) string {
	if index >= a.Size || index < 0 {
		return ""
	}
	return a.Data[index]
}

func (a *Array) DeleteByIndex(index int) bool {
	if index >= a.Size || index < 0 {
		return false
	}

	for i := index; i < a.Size-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	a.Size--
	return true
}

func (a *Array) SwapByIndex(value string, index int) bool {
	if index >= a.Size || index < 0 {
		return false
	}
	a.Data[index] = value
	return true
}

func (a *Array) Print() {
	for i := 0; i < a.Size; i++ {
		fmt.Print(a.Data[i])
		if i < (a.Size - 1) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// MarshalJSON реализует интерфейс json.Marshaler
func (a *Array) MarshalJSON() ([]byte, error) {
	type Alias Array
	return json.Marshal(&struct {
		*Alias
		ActualData []string `json:"actual_data"`
	}{
		Alias:      (*Alias)(a),
		ActualData: a.Data[:a.Size],
	})
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (a *Array) UnmarshalJSON(data []byte) error {
	type Alias Array
	aux := &struct {
		*Alias
		ActualData []string `json:"actual_data"`
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	a.Data = make([]string, a.Capacity)
	copy(a.Data, aux.ActualData)
	a.Size = len(aux.ActualData)

	return nil
}

// ToJSON сериализует массив в JSON строку
func (a *Array) ToJSON() (string, error) {
	bytes, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON десериализует JSON строку в массив
func (a *Array) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), a)
}

// SaveToJSON сохраняет в JSON файл
func (a *Array) SaveToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(a)
}

// LoadFromJSON загружает данные из JSON
func (a *Array) LoadFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(a)
}

// ToBinary сериализует массив в бинарный формат
func (a *Array) ToBinary() ([]byte, error) {
	// Используем LittleEndian для согласованности (выбор BigEndian также возможен)
	var buf bytes.Buffer

	// 1. Записываем фактический размер (Size) массива (как int32 или int64 для безопасности)
	size := int32(a.Size)
	if err := binary.Write(&buf, binary.LittleEndian, size); err != nil {
		return nil, fmt.Errorf("ошибка записи размера массива: %w", err)
	}

	// 2. Записываем каждую строку: сначала ее длину, затем байты строки.
	for i := 0; i < a.Size; i++ {
		strBytes := []byte(a.Data[i])
		strLen := int32(len(strBytes))

		// Записываем длину строки
		if err := binary.Write(&buf, binary.LittleEndian, strLen); err != nil {
			return nil, fmt.Errorf("ошибка записи длины строки %d: %w", i, err)
		}

		// Записываем байты строки
		if _, err := buf.Write(strBytes); err != nil {
			return nil, fmt.Errorf("ошибка записи данных строки %d: %w", i, err)
		}
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует массив из бинарного формата
func (a *Array) FromBinary(data []byte) error {
	buf := bytes.NewReader(data)

	// 1. Считываем фактический размер (Size) массива
	var size int32
	if err := binary.Read(buf, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("ошибка чтения размера массива: %w", err)
	}

	// Обновляем Size и создаем новый слайс Data
	a.Size = int(size)
	// Устанавливаем Capacity равным Size для начала
	a.Capacity = a.Size
	if a.Capacity == 0 {
		a.Capacity = 10 // или любой другой разумный минимальный размер
	}
	a.Data = make([]string, a.Capacity)

	// 2. Считываем каждую строку
	for i := 0; i < a.Size; i++ {
		// Считываем длину строки
		var strLen int32
		if err := binary.Read(buf, binary.LittleEndian, &strLen); err != nil {
			if err == io.EOF {
				return fmt.Errorf("неожиданный конец данных при чтении длины строки %d. Ожидалось %d строк", i, a.Size)
			}
			return fmt.Errorf("ошибка чтения длины строки %d: %w", i, err)
		}

		// Считываем байты строки
		strBytes := make([]byte, strLen)
		if _, err := io.ReadFull(buf, strBytes); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return fmt.Errorf("недостаточно данных для чтения строки %d длиной %d байт", i, strLen)
			}
			return fmt.Errorf("ошибка чтения данных строки %d: %w", i, err)
		}

		a.Data[i] = string(strBytes)
	}

	return nil
}

func (a *Array) SaveToBinary(filename string) error {
	data, err := a.ToBinary()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromBinary загружает массив из бинарного файла
func (a *Array) LoadFromBinary(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return a.FromBinary(data)
}
