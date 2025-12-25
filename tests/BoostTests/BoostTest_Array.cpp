#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/Array/Array.h"

BOOST_AUTO_TEST_SUITE(ArrayTests)


// Тест на основные операции: добавление в конец и удаление по индексу
BOOST_AUTO_TEST_CASE(PushBackDeleteTest) {
    Array arr;
    arr.pushBack("1");
    arr.pushBack("2");
    arr.pushBack("3");
    arr.pushBack("4");

    BOOST_CHECK(arr.pushBack("5"));
    BOOST_CHECK_EQUAL(arr.length(), 5);

    arr.deleteByIndex(1);
    BOOST_CHECK_EQUAL(arr.length(), 4);
    BOOST_CHECK(!arr.deleteByIndex(5));
}


// Тест на автоматическое увеличение емкости
BOOST_AUTO_TEST_CASE(DoubleCapacityTest) {
    Array arr(3);
    arr.pushBack("1");
    arr.pushBack("2");
    arr.pushBack("3");

    BOOST_CHECK_LE(arr.length(), arr.cap());

    arr.pushBack("4");
    arr.pushBack("5");

    BOOST_CHECK_LE(arr.length(), arr.cap());
}


// Тест на вставку элемента по индексу
BOOST_AUTO_TEST_CASE(PushByIndexTest) {
    Array arr(5);
    OutputRedirect redirect;
    arr.pushBack("1");
    arr.pushBack("2");

    BOOST_CHECK(!arr.pushByIndex("3", 3));
    BOOST_CHECK(arr.pushByIndex("3", 1));

    arr.pushBack("4");
    arr.pushByIndex("5", 0);
    BOOST_CHECK_LE(arr.length(), arr.cap());
}


// Тест на операции с индексами: замена элемента, получение элемента и вывод на печать
BOOST_AUTO_TEST_CASE(IndexTests) {
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
        BOOST_CHECK_EQUAL(output, expected);
    }

    BOOST_CHECK_EQUAL(arr.getByIndex(2), "3");
    BOOST_CHECK_EQUAL(arr.getByIndex(5), "");
    BOOST_CHECK(!arr.swapByIndex("???", 10));
}


// Тест на сериализацию и десериализацию данных
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    Array original;
    original.pushBack("first");
    original.pushBack("second");
    original.pushBack("third");
    original.pushBack("fourth");

    BOOST_CHECK(original.serialize("test_array.bin"));

    Array loaded;
    BOOST_CHECK(loaded.deserialize("test_array.bin"));

    BOOST_CHECK_EQUAL(original.length(), loaded.length());
    BOOST_CHECK_EQUAL(original.cap(), loaded.cap());

    for (int i = 0; i < original.length(); i++) {
        BOOST_CHECK_EQUAL(original.getByIndex(i), loaded.getByIndex(i));
    }

    Array empty;
    BOOST_CHECK(empty.serialize("empty_array.bin"));

    Array emptyLoaded;
    BOOST_CHECK(emptyLoaded.deserialize("empty_array.bin"));
    BOOST_CHECK_EQUAL(emptyLoaded.length(), 0);

    Array invalid;
    BOOST_CHECK(!invalid.deserialize("non_existent_file.bin"));
}


// Тест на сериализацию с операциями над массивом
BOOST_AUTO_TEST_CASE(SerializeDeserializeWithOperationsTest) {
    Array arr;
    arr.pushBack("apple");
    arr.pushBack("banana");
    arr.pushByIndex("orange", 1);
    arr.swapByIndex("grape", 0);

    BOOST_CHECK(arr.serialize("complex_array.bin"));

    Array newArr;
    BOOST_CHECK(newArr.deserialize("complex_array.bin"));

    BOOST_CHECK_EQUAL(newArr.length(), 3);
    BOOST_CHECK_EQUAL(newArr.getByIndex(0), "grape");
    BOOST_CHECK_EQUAL(newArr.getByIndex(1), "orange");
    BOOST_CHECK_EQUAL(newArr.getByIndex(2), "banana");

    BOOST_CHECK(newArr.pushBack("peach"));
    BOOST_CHECK_EQUAL(newArr.length(), 4);
    BOOST_CHECK_EQUAL(newArr.getByIndex(3), "peach");
}

BOOST_AUTO_TEST_SUITE_END()
