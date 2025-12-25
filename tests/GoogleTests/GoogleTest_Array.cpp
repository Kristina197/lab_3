#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/Array/Array.h"

class ArrayTest : public ::testing::Test {
protected:
    Array arr;
};

// Тест на основные операции: добавление в конец и удаление по индексу
TEST_F(ArrayTest, PushBackDeleteTest) {
    arr.pushBack("1");
    arr.pushBack("2");
    arr.pushBack("3");
    arr.pushBack("4");

    EXPECT_TRUE(arr.pushBack("5"));
    EXPECT_EQ(arr.length(), 5);

    arr.deleteByIndex(1);
    EXPECT_EQ(arr.length(), 4);
    EXPECT_FALSE(arr.deleteByIndex(5));
}


// Тест на автоматическое увеличение емкости
TEST_F(ArrayTest, DoubleCapacityTest) {
    Array arr(3);
    arr.pushBack("1");
    arr.pushBack("2");
    arr.pushBack("3");

    EXPECT_LE(arr.length(), arr.cap());

    arr.pushBack("4");
    arr.pushBack("5");

    EXPECT_LE(arr.length(), arr.cap());
}


// Тест на вставку элемента по индексу
TEST_F(ArrayTest, PushByIndexTest) {
    Array arr(5);
    OutputRedirect redirect;
    arr.pushBack("1");
    arr.pushBack("2");

    EXPECT_FALSE(arr.pushByIndex("3", 3));
    EXPECT_TRUE(arr.pushByIndex("3", 1));

    arr.pushBack("4");
    arr.pushByIndex("5", 0);
    EXPECT_LE(arr.length(), arr.cap());
}


// Тест на операции с индексами: замена элемента, получение элемента и вывод на печать
TEST_F(ArrayTest, IndexTests) {
    Array arr;
    arr.pushBack("1");
    arr.pushBack("2");
    arr.pushBack("3");
    arr.pushBack("4");
    arr.pushBack("5");

    arr.swapByIndex("100", 3);
    {
        OutputRedirect redirect;
        arr.print();
        string output = redirect.getOutput();
        string expected = "1 2 3 100 5\n";
        EXPECT_EQ(output, expected);
    }

    EXPECT_EQ(arr.getByIndex(2), "3");
    EXPECT_EQ(arr.getByIndex(5), "");
    EXPECT_FALSE(arr.swapByIndex("???", 10));
}


// Тест на сериализацию и десериализацию данных
TEST_F(ArrayTest, SerializeDeserializeTest) {
    Array original;
    original.pushBack("first");
    original.pushBack("second");
    original.pushBack("third");
    original.pushBack("fourth");

    EXPECT_TRUE(original.serialize("test_array.bin"));

    Array loaded;
    EXPECT_TRUE(loaded.deserialize("test_array.bin"));

    EXPECT_EQ(original.length(), loaded.length());
    EXPECT_EQ(original.cap(), loaded.cap());

    for (int i = 0; i < original.length(); i++) {
        EXPECT_EQ(original.getByIndex(i), loaded.getByIndex(i));
    }

    Array empty;
    EXPECT_TRUE(empty.serialize("empty_array.bin"));

    Array emptyLoaded;
    EXPECT_TRUE(emptyLoaded.deserialize("empty_array.bin"));
    EXPECT_EQ(emptyLoaded.length(), 0);

    Array invalid;
    EXPECT_FALSE(invalid.deserialize("non_existent_file.bin"));
}


// Тест на сериализацию с операциями над массивом
TEST_F(ArrayTest, SerializeDeserializeWithOperationsTest) {
    Array arr;
    arr.pushBack("apple");
    arr.pushBack("banana");
    arr.pushByIndex("orange", 1);
    arr.swapByIndex("grape", 0);

    EXPECT_TRUE(arr.serialize("complex_array.bin"));

    Array newArr;
    EXPECT_TRUE(newArr.deserialize("complex_array.bin"));

    EXPECT_EQ(newArr.length(), 3);
    EXPECT_EQ(newArr.getByIndex(0), "grape");
    EXPECT_EQ(newArr.getByIndex(1), "orange");
    EXPECT_EQ(newArr.getByIndex(2), "banana");

    EXPECT_TRUE(newArr.pushBack("peach"));
    EXPECT_EQ(newArr.length(), 4);
    EXPECT_EQ(newArr.getByIndex(3), "peach");
}