#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/OpenAddrHash/OpenAddrHash.h"

BOOST_AUTO_TEST_SUITE(OpenAddressingTests)

// Тест на операции с пустой таблицей
BOOST_AUTO_TEST_CASE(EmptyTableOperations) {
    OpenAddrHash oah(5);

    BOOST_CHECK(!oah.get("1"));
    BOOST_CHECK_NO_THROW(oah.remove("1"));
}


// Тест на базовые операции добавления и получения
BOOST_AUTO_TEST_CASE(BasicPutAndGet) {
    OpenAddrHash oah(10);

    BOOST_CHECK(oah.put("apple", "100"));

    HashEntry* entry = oah.get("apple");

    BOOST_REQUIRE_NE(entry, nullptr);
    BOOST_CHECK_EQUAL(entry->key, "apple");
    BOOST_CHECK_EQUAL(entry->value, "100");
    BOOST_CHECK(entry->status == HashStatus::TAKEN);
}


// Тест на переполнение таблицы
BOOST_AUTO_TEST_CASE(OverflowTest) {
    OpenAddrHash oah(3);
    oah.put("1", "100");
    oah.put("2", "200");
    oah.put("3", "300");
    {
        OutputRedirect redirect;
        oah.put("4", "400");
        string output = redirect.getOutput();
        BOOST_CHECK_EQUAL(output, "Хеш-таблица переполнена\n");
    }
}


// Тест на получение несуществующего элемента
BOOST_AUTO_TEST_CASE(GetNonExistent) {
    OpenAddrHash oah(10);
    oah.put("apple", "100");

    HashEntry* entry = oah.get("banana");
    BOOST_CHECK(entry == nullptr);
}


// Тест на обновление значения существующего ключа
BOOST_AUTO_TEST_CASE(UpdateValue) {
    OpenAddrHash oah;
    oah.put("apple", "100");
    oah.put("apple", "200");

    HashEntry* entry = oah.get("apple");
    BOOST_REQUIRE_NE(entry, nullptr);
    BOOST_CHECK_EQUAL(entry->value, "200");
}


// Тест на обработку коллизий
BOOST_AUTO_TEST_CASE(CollisionsTests) {
    OpenAddrHash oah(5);

    BOOST_CHECK(oah.put("0", "zero"));
    BOOST_CHECK(oah.put("5", "five"));
    BOOST_CHECK(oah.put("10", "ten"));

    HashEntry* entry0 = oah.get("0");
    BOOST_REQUIRE_NE(entry0, nullptr);
    BOOST_CHECK_EQUAL(entry0->value, "zero");

    HashEntry* entry5 = oah.get("5");
    BOOST_REQUIRE_NE(entry5, nullptr);
    BOOST_CHECK_EQUAL(entry5->value, "five");

    HashEntry* entry10 = oah.get("10");
    BOOST_REQUIRE_NE(entry10, nullptr);
    BOOST_CHECK_EQUAL(entry10->value, "ten");

    HashEntry* entry15 = oah.get("15");
    BOOST_CHECK(entry15 == nullptr);
}


// Тест на повторное использование удаленных ячеек
BOOST_AUTO_TEST_CASE(PutIntoDeletedSlot) {
    OpenAddrHash oah(3);

    oah.put("hello", "goodbye");
    oah.put("smth", "testing");
    oah.put("active", "lazy");
    oah.remove("smth");

    auto* entry = oah.get("smth");
    BOOST_CHECK(!entry);

    oah.put("epic", "comeback");
    entry = oah.get("epic");
    BOOST_CHECK(entry);
}


// Тест на вывод содержимого таблицы
BOOST_AUTO_TEST_CASE(PrintTest) {
    OpenAddrHash oah(5);

    BOOST_CHECK_NO_THROW(oah.print());

    oah.put("1", "one");
    oah.put("2", "two");
    BOOST_CHECK_NO_THROW(oah.print());

    oah.remove("1");
    BOOST_CHECK_NO_THROW(oah.print());
}


// Тест на сериализацию и десериализацию
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    OpenAddrHash original(10);
    BOOST_CHECK(original.put("moscow", "russia"));
    BOOST_CHECK(original.put("paris", "france"));
    BOOST_CHECK(original.put("tokyo", "japan"));
    BOOST_CHECK(original.put("berlin", "germany"));

    BOOST_CHECK(original.serialize("test_openaddr.bin"));

    OpenAddrHash loaded;
    BOOST_CHECK(loaded.deserialize("test_openaddr.bin"));

    HashEntry* moscow = original.get("moscow");
    HashEntry* loadedMoscow = loaded.get("moscow");
    BOOST_REQUIRE_NE(moscow, nullptr);
    BOOST_REQUIRE_NE(loadedMoscow, nullptr);
    BOOST_CHECK_EQUAL(moscow->value, loadedMoscow->value);

    HashEntry* paris = original.get("paris");
    HashEntry* loadedParis = loaded.get("paris");
    BOOST_REQUIRE_NE(paris, nullptr);
    BOOST_REQUIRE_NE(loadedParis, nullptr);
    BOOST_CHECK_EQUAL(paris->value, loadedParis->value);
}

BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    OpenAddrHash empty(5);
    BOOST_CHECK(empty.serialize("empty_openaddr.bin"));

    OpenAddrHash loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_openaddr.bin"));
    BOOST_CHECK(loadedEmpty.get("anything") == nullptr);
}

BOOST_AUTO_TEST_CASE(SerializeDeserializeBasicTest) {
    OpenAddrHash table(5);
    BOOST_CHECK(table.put("red", "color"));
    BOOST_CHECK(table.put("blue", "color"));

    BOOST_CHECK(table.serialize("basic_openaddr.bin"));

    OpenAddrHash loaded;
    BOOST_CHECK(loaded.deserialize("basic_openaddr.bin"));

    HashEntry* red = table.get("red");
    HashEntry* loadedRed = loaded.get("red");
    BOOST_REQUIRE_NE(red, nullptr);
    BOOST_REQUIRE_NE(loadedRed, nullptr);
    BOOST_CHECK_EQUAL(red->value, loadedRed->value);
}

BOOST_AUTO_TEST_SUITE_END()
