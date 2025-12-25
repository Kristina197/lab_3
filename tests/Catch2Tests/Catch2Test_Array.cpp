#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/Array/Array.h"

TEST_CASE("Array Tests", "[array]") {
    // Тест на основные операции: добавление в конец и удаление по индексу
    SECTION("PushBackDeleteTest") {
        Array arr;
        arr.pushBack("1");
        arr.pushBack("2");
        arr.pushBack("3");
        arr.pushBack("4");

        REQUIRE(arr.pushBack("5"));
        REQUIRE(arr.length() == 5);

        arr.deleteByIndex(1);
        REQUIRE(arr.length() == 4);
        REQUIRE_FALSE(arr.deleteByIndex(5));
    }

    // Тест на автоматическое увеличение емкости
    SECTION("DoubleCapacityTest") {
        Array arr(3);
        arr.pushBack("1");
        arr.pushBack("2");
        arr.pushBack("3");

        REQUIRE(arr.length() <= arr.cap());

        arr.pushBack("4");
        arr.pushBack("5");

        REQUIRE(arr.length() <= arr.cap());
    }

    // Тест на вставку элемента по индексу
    SECTION("PushByIndexTest") {
        Array arr(5);
        OutputRedirect redirect;
        arr.pushBack("1");
        arr.pushBack("2");

        REQUIRE_FALSE(arr.pushByIndex("3", 3));
        REQUIRE(arr.pushByIndex("3", 1));

        arr.pushBack("4");
        arr.pushByIndex("5", 0);
        REQUIRE(arr.length() <= arr.cap());
    }

    // Тест на операции с индексами: замена элемента, получение элемента и вывод на печать
    SECTION("IndexTests") {
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
            REQUIRE(output == expected);
        }

        REQUIRE(arr.getByIndex(2) == "3");
        REQUIRE(arr.getByIndex(5) == "");
        REQUIRE_FALSE(arr.swapByIndex("???", 10));
    }

    // Тест на сериализацию и десериализацию данных
    SECTION("SerializeDeserializeTest") {
        Array original;
        original.pushBack("first");
        original.pushBack("second");
        original.pushBack("third");
        original.pushBack("fourth");

        REQUIRE(original.serialize("test_array.bin"));

        Array loaded;
        REQUIRE(loaded.deserialize("test_array.bin"));

        REQUIRE(original.length() == loaded.length());
        REQUIRE(original.cap() == loaded.cap());

        for (int i = 0; i < original.length(); i++) {
            REQUIRE(original.getByIndex(i) == loaded.getByIndex(i));
        }

        Array empty;
        REQUIRE(empty.serialize("empty_array.bin"));

        Array emptyLoaded;
        REQUIRE(emptyLoaded.deserialize("empty_array.bin"));
        REQUIRE(emptyLoaded.length() == 0);

        Array invalid;
        REQUIRE_FALSE(invalid.deserialize("non_existent_file.bin"));
    }

    // Тест на сериализацию с операциями над массивом
    SECTION("SerializeDeserializeWithOperationsTest") {
        Array arr;
        arr.pushBack("apple");
        arr.pushBack("banana");
        arr.pushByIndex("orange", 1);
        arr.swapByIndex("grape", 0);

        REQUIRE(arr.serialize("complex_array.bin"));

        Array newArr;
        REQUIRE(newArr.deserialize("complex_array.bin"));

        REQUIRE(newArr.length() == 3);
        REQUIRE(newArr.getByIndex(0) == "grape");
        REQUIRE(newArr.getByIndex(1) == "orange");
        REQUIRE(newArr.getByIndex(2) == "banana");

        REQUIRE(newArr.pushBack("peach"));
        REQUIRE(newArr.length() == 4);
        REQUIRE(newArr.getByIndex(3) == "peach");
    }
}