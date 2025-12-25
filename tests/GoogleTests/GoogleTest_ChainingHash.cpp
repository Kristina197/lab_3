#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/ChainingHash/ChainingHash.h"

class ChainingHashTest : public ::testing::Test {
protected:
    ChainingHash ch;
};

// Тест на добавление элементов
TEST_F(ChainingHashTest, PutTest) {
    EXPECT_TRUE(ch.put("1", "one"));
    ch.put("2", "two");
    ch.put("3", "three");

    EXPECT_TRUE(ch.put("2", "update"));
}


// Тест на получение элементов
TEST_F(ChainingHashTest, GetTest) {
    EXPECT_EQ(ch.get("1"), nullptr);

    ch.put("1", "one");
    string* val = ch.get("1");
    EXPECT_NE(val, nullptr);
    EXPECT_EQ(*val, "one");
}


// Тест на удаление элементов
TEST_F(ChainingHashTest, RemoveTest) {
    ChainingHash ch(3);
    ch.put("1", "one");
    ch.put("11", "eleven");
    ch.put("21", "twenty_one");

    ch.remove("11");
    EXPECT_NE(ch.get("1"), nullptr);
    EXPECT_EQ(ch.get("11"), nullptr);
    EXPECT_NE(ch.get("21"), nullptr);

    ch.remove("1");
    EXPECT_EQ(ch.get("1"), nullptr);
    EXPECT_NE(ch.get("21"), nullptr);

    ch.remove("21");
    EXPECT_EQ(ch.get("21"), nullptr);
}


// Тест на вывод содержимого таблицы
TEST_F(ChainingHashTest, PrintTest) {
    ChainingHash oah;
    EXPECT_NO_THROW(oah.print());

    oah.put("1", "one");
    oah.put("2", "two");
    EXPECT_NO_THROW(oah.print());

    ChainingHash smallTable(3);
    smallTable.put("1", "one");
    smallTable.put("4", "four");
    smallTable.put("7", "seven");
    EXPECT_NO_THROW(smallTable.print());
}


// Тесты на создание таблиц с разной вместимостью, обновление существующих ключей, удаление элементов с коллизиями
TEST_F(ChainingHashTest, AdditionalTests) {
    {
        ChainingHash defaultTable;
        ChainingHash customTable(20);
    }
    const int TEST_CAPACITY = 5;

    ChainingHash table(TEST_CAPACITY);
    EXPECT_TRUE(table.put("1", "10"));
    EXPECT_TRUE(table.put("2", "20"));
    EXPECT_EQ(*table.get("1"), "10");
    EXPECT_TRUE(table.put("1", "100"));
    EXPECT_EQ(*table.get("1"), "100");

    table.remove("1");
    EXPECT_EQ(table.get("1"), nullptr);

    ChainingHash collisionTable(3);
    collisionTable.put("1", "100");
    collisionTable.put("4", "400");
    collisionTable.put("7", "700");

    EXPECT_EQ(*collisionTable.get("1"), "100");
    EXPECT_EQ(*collisionTable.get("4"), "400");
    EXPECT_EQ(*collisionTable.get("7"), "700");

    collisionTable.remove("4");
    EXPECT_EQ(collisionTable.get("4"), nullptr);
}


// Тест на сериализацию и десериализацию
TEST_F(ChainingHashTest, SerializeDeserializeTest) {
    ChainingHash original;
    original.put("apple", "red");
    original.put("banana", "yellow");
    original.put("grape", "purple");
    original.put("lemon", "yellow");

    EXPECT_TRUE(original.serialize("test_chaining.bin"));

    ChainingHash loaded;
    EXPECT_TRUE(loaded.deserialize("test_chaining.bin"));

    EXPECT_EQ(*original.get("apple"), *loaded.get("apple"));
    EXPECT_EQ(*original.get("banana"), *loaded.get("banana"));
    EXPECT_EQ(*original.get("grape"), *loaded.get("grape"));
    EXPECT_EQ(*original.get("lemon"), *loaded.get("lemon"));
}


// Тест на сериализацию и десериализацию пустой таблицы
TEST_F(ChainingHashTest, SerializeDeserializeEmptyTest) {
    ChainingHash empty;
    EXPECT_TRUE(empty.serialize("empty_chaining.bin"));

    ChainingHash loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_chaining.bin"));
    EXPECT_EQ(loadedEmpty.get("anything"), nullptr);
}


// Тест на сериализацию и десериализацию таблицы с коллизиями
TEST_F(ChainingHashTest, SerializeDeserializeWithCollisionsTest) {
    ChainingHash table(3);
    table.put("cat", "animal");
    table.put("dog", "pet");
    table.put("bird", "fly");

    EXPECT_TRUE(table.serialize("collision_chaining.bin"));

    ChainingHash loaded;
    EXPECT_TRUE(loaded.deserialize("collision_chaining.bin"));

    EXPECT_EQ(*table.get("cat"), *loaded.get("cat"));
    EXPECT_EQ(*table.get("dog"), *loaded.get("dog"));
    EXPECT_EQ(*table.get("bird"), *loaded.get("bird"));
}
