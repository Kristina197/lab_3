#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/ChainingHash/ChainingHash.h"

BOOST_AUTO_TEST_SUITE(ChainingHashTests)

// Тест на добавление элементов
BOOST_AUTO_TEST_CASE(PutTest) {
    ChainingHash ch;

    BOOST_CHECK(ch.put("1", "one"));
    ch.put("2", "two");
    ch.put("3", "three");

    BOOST_CHECK(ch.put("2", "update"));
}


// Тест на получение элементов
BOOST_AUTO_TEST_CASE(GetTest) {
    ChainingHash ch;
    BOOST_TEST(ch.get("1") == nullptr);

    ch.put("1", "one");
    string* val = ch.get("1");
    BOOST_CHECK(val);
    BOOST_TEST(*val == "one");
}


// Тест на удаление элементов
BOOST_AUTO_TEST_CASE(RemoveTest) {
    ChainingHash ch(3);
    ch.put("1", "one");
    ch.put("11", "eleven");
    ch.put("21", "twenty_one");

    ch.remove("11");
    BOOST_TEST(ch.get("1") != nullptr);
    BOOST_TEST(ch.get("11") == nullptr);
    BOOST_TEST(ch.get("21") != nullptr);

    ch.remove("1");
    BOOST_TEST(ch.get("1") == nullptr);
    BOOST_TEST(ch.get("21") != nullptr);

    ch.remove("21");
    BOOST_TEST(ch.get("21") == nullptr);
}


// Тест на вывод содержимого таблицы
BOOST_AUTO_TEST_CASE(PrintTest) {
    ChainingHash oah;
    BOOST_CHECK_NO_THROW(oah.print());

    oah.put("1", "one");
    oah.put("2", "two");
    BOOST_CHECK_NO_THROW(oah.print());

    ChainingHash smallTable(3);
    smallTable.put("1", "one");
    smallTable.put("4", "four");
    smallTable.put("7", "seven");
    BOOST_CHECK_NO_THROW(smallTable.print());
}


// Тесты на создание таблиц с разной вместимостью, обновление существующих ключей, удаление элементов с коллизиями
BOOST_AUTO_TEST_CASE(AdditionalTests) {
    {
        ChainingHash defaultTable;
        ChainingHash customTable(20);
    }
    const int TEST_CAPACITY = 5;

    ChainingHash table(TEST_CAPACITY);
    BOOST_TEST(table.put("1", "10") == true);
    BOOST_TEST(table.put("2", "20") == true);
    BOOST_TEST(*table.get("1") == "10");
    BOOST_TEST(table.put("1", "100") == true);
    BOOST_TEST(*table.get("1") == "100");

    table.remove("1");
    BOOST_TEST(table.get("1") == nullptr);

    ChainingHash collisionTable(3);
    collisionTable.put("1", "100");
    collisionTable.put("4", "400");
    collisionTable.put("7", "700");

    BOOST_CHECK_EQUAL(*collisionTable.get("1"), "100");
    BOOST_CHECK_EQUAL(*collisionTable.get("4"), "400");
    BOOST_CHECK_EQUAL(*collisionTable.get("7"), "700");

    collisionTable.remove("4");
    BOOST_TEST(collisionTable.get("4") == nullptr);
}


// Тест на сериализацию и десериализацию
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    ChainingHash original;
    original.put("apple", "red");
    original.put("banana", "yellow");
    original.put("grape", "purple");
    original.put("lemon", "yellow");

    BOOST_CHECK(original.serialize("test_chaining.bin"));

    ChainingHash loaded;
    BOOST_CHECK(loaded.deserialize("test_chaining.bin"));

    BOOST_CHECK_EQUAL(*original.get("apple"), *loaded.get("apple"));
    BOOST_CHECK_EQUAL(*original.get("banana"), *loaded.get("banana"));
    BOOST_CHECK_EQUAL(*original.get("grape"), *loaded.get("grape"));
    BOOST_CHECK_EQUAL(*original.get("lemon"), *loaded.get("lemon"));
}


// Тест на сериализацию и десериализацию пустой таблицы
BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    ChainingHash empty;
    BOOST_CHECK(empty.serialize("empty_chaining.bin"));

    ChainingHash loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_chaining.bin"));
    BOOST_CHECK(loadedEmpty.get("anything") == nullptr);
}


// Тест на сериализацию и десериализацию таблицы с коллизиями
BOOST_AUTO_TEST_CASE(SerializeDeserializeWithCollisionsTest) {
    ChainingHash table(3);
    table.put("cat", "animal");
    table.put("dog", "pet");
    table.put("bird", "fly");

    BOOST_CHECK(table.serialize("collision_chaining.bin"));

    ChainingHash loaded;
    BOOST_CHECK(loaded.deserialize("collision_chaining.bin"));

    BOOST_CHECK_EQUAL(*table.get("cat"), *loaded.get("cat"));
    BOOST_CHECK_EQUAL(*table.get("dog"), *loaded.get("dog"));
    BOOST_CHECK_EQUAL(*table.get("bird"), *loaded.get("bird"));
}

BOOST_AUTO_TEST_SUITE_END()
