#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/ChainingHash/ChainingHash.h"


TEST_CASE("Chaining Hash Tests", "[chaining_hash]") {
    // Тест на добавление элементов
    SECTION("PutTest") {
        ChainingHash ch;

        REQUIRE(ch.put("1", "one"));
        ch.put("2", "two");
        ch.put("3", "three");

        REQUIRE(ch.put("2", "update"));
    }

    // Тест на получение элементов
    SECTION("GetTest") {
        ChainingHash ch;
        REQUIRE(ch.get("1") == nullptr);

        ch.put("1", "one");
        string* val = ch.get("1");
        REQUIRE(val);
        REQUIRE(*val == "one");
    }

    // Тест на удаление элементов
    SECTION("RemoveTest") {
        ChainingHash ch(3);
        ch.put("1", "one");
        ch.put("11", "eleven");
        ch.put("21", "twenty_one");

        ch.remove("11");
        REQUIRE(ch.get("1") != nullptr);
        REQUIRE(ch.get("11") == nullptr);
        REQUIRE(ch.get("21") != nullptr);

        ch.remove("1");
        REQUIRE(ch.get("1") == nullptr);
        REQUIRE(ch.get("21") != nullptr);

        ch.remove("21");
        REQUIRE(ch.get("21") == nullptr);
    }

    // Тест на вывод содержимого таблицы
    SECTION("PrintTest") {
        ChainingHash oah;
        REQUIRE_NOTHROW(oah.print());

        oah.put("1", "one");
        oah.put("2", "two");
        REQUIRE_NOTHROW(oah.print());

        ChainingHash smallTable(3);
        smallTable.put("1", "one");
        smallTable.put("4", "four");
        smallTable.put("7", "seven");
        REQUIRE_NOTHROW(smallTable.print());
    }

    // Тесты на создание таблиц с разной вместимостью, обновление существующих ключей, удаление элементов с коллизиями
    SECTION("AdditionalTests") {
        {
            ChainingHash defaultTable;
            ChainingHash customTable(20);
        }
        const int TEST_CAPACITY = 5;

        ChainingHash table(TEST_CAPACITY);
        REQUIRE(table.put("1", "10") == true);
        REQUIRE(table.put("2", "20") == true);
        REQUIRE(*table.get("1") == "10");
        REQUIRE(table.put("1", "100") == true);
        REQUIRE(*table.get("1") == "100");

        table.remove("1");
        REQUIRE(table.get("1") == nullptr);

        ChainingHash collisionTable(3);
        collisionTable.put("1", "100");
        collisionTable.put("4", "400");
        collisionTable.put("7", "700");

        REQUIRE(*collisionTable.get("1") == "100");
        REQUIRE(*collisionTable.get("4") == "400");
        REQUIRE(*collisionTable.get("7") == "700");

        collisionTable.remove("4");
        REQUIRE(collisionTable.get("4") == nullptr);
    }

    // Тест на сериализацию и десериализацию
    SECTION("SerializeDeserializeTest") {
        ChainingHash original;
        original.put("apple", "red");
        original.put("banana", "yellow");
        original.put("grape", "purple");
        original.put("lemon", "yellow");

        REQUIRE(original.serialize("test_chaining.bin"));

        ChainingHash loaded;
        REQUIRE(loaded.deserialize("test_chaining.bin"));

        REQUIRE(*original.get("apple") == *loaded.get("apple"));
        REQUIRE(*original.get("banana") == *loaded.get("banana"));
        REQUIRE(*original.get("grape") == *loaded.get("grape"));
        REQUIRE(*original.get("lemon") == *loaded.get("lemon"));
    }

    // Тест на сериализацию и десериализацию пустой таблицы
    SECTION("SerializeDeserializeEmptyTest") {
        ChainingHash empty;
        REQUIRE(empty.serialize("empty_chaining.bin"));

        ChainingHash loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_chaining.bin"));
        REQUIRE(loadedEmpty.get("anything") == nullptr);
    }

    // Тест на сериализацию и десериализацию таблицы с коллизиями
    SECTION("SerializeDeserializeWithCollisionsTest") {
        ChainingHash table(3);
        table.put("cat", "animal");
        table.put("dog", "pet");
        table.put("bird", "fly");

        REQUIRE(table.serialize("collision_chaining.bin"));

        ChainingHash loaded;
        REQUIRE(loaded.deserialize("collision_chaining.bin"));

        REQUIRE(*table.get("cat") == *loaded.get("cat"));
        REQUIRE(*table.get("dog") == *loaded.get("dog"));
        REQUIRE(*table.get("bird") == *loaded.get("bird"));
    }
}