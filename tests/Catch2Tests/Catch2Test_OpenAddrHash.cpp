#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/OpenAddrHash/OpenAddrHash.h"


TEST_CASE("Open Addressing Tests", "[open_addressing]") {
    SECTION("EmptyTableOperations") {
        OpenAddrHash oah(5);

        REQUIRE(!oah.get("1"));
        REQUIRE_NOTHROW(oah.remove("1"));
    }

    SECTION("BasicPutAndGet") {
        OpenAddrHash oah(10);

        REQUIRE(oah.put("apple", "100"));

        HashEntry* entry = oah.get("apple");

        REQUIRE(entry != nullptr);
        REQUIRE(entry->key == "apple");
        REQUIRE(entry->value == "100");
        REQUIRE(entry->status == HashStatus::TAKEN);
    }

    SECTION("OverflowTest") {
        OpenAddrHash oah(3);
        oah.put("1", "100");
        oah.put("2", "200");
        oah.put("3", "300");
        {
            OutputRedirect redirect;
            oah.put("4", "400");
            string output = redirect.getOutput();
            REQUIRE(output == "Хеш-таблица переполнена\n");
        }
    }

    SECTION("GetNonExistent") {
        OpenAddrHash oah(10);
        oah.put("apple", "100");

        HashEntry* entry = oah.get("banana");
        REQUIRE(entry == nullptr);
    }

    SECTION("UpdateValue") {
        OpenAddrHash oah;
        oah.put("apple", "100");
        oah.put("apple", "200");

        HashEntry* entry = oah.get("apple");
        REQUIRE(entry != nullptr);
        REQUIRE(entry->value == "200");
    }

    SECTION("CollisionsTests") {
        OpenAddrHash oah(5);

        REQUIRE(oah.put("0", "zero"));
        REQUIRE(oah.put("5", "five"));
        REQUIRE(oah.put("10", "ten"));

        HashEntry* entry0 = oah.get("0");
        REQUIRE(entry0 != nullptr);
        REQUIRE(entry0->value == "zero");

        HashEntry* entry5 = oah.get("5");
        REQUIRE(entry5 != nullptr);
        REQUIRE(entry5->value == "five");

        HashEntry* entry10 = oah.get("10");
        REQUIRE(entry10 != nullptr);
        REQUIRE(entry10->value == "ten");

        HashEntry* entry15 = oah.get("15");
        REQUIRE(entry15 == nullptr);
    }

    SECTION("PutIntoDeletedSlot") {
        OpenAddrHash oah(3);

        oah.put("hello", "goodbye");
        oah.put("smth", "testing");
        oah.put("active", "lazy");
        oah.remove("smth");

        auto* entry = oah.get("smth");
        REQUIRE(!entry);

        oah.put("epic", "comeback");
        entry = oah.get("epic");
        REQUIRE(entry);
    }

    SECTION("PrintTest") {
        OpenAddrHash oah(5);

        REQUIRE_NOTHROW(oah.print());

        oah.put("1", "one");
        oah.put("2", "two");
        REQUIRE_NOTHROW(oah.print());

        oah.remove("1");
        REQUIRE_NOTHROW(oah.print());
    }

    SECTION("SerializeDeserializeTest") {
        OpenAddrHash original(10);
        REQUIRE(original.put("moscow", "russia"));
        REQUIRE(original.put("paris", "france"));
        REQUIRE(original.put("tokyo", "japan"));
        REQUIRE(original.put("berlin", "germany"));

        REQUIRE(original.serialize("test_openaddr.bin"));

        OpenAddrHash loaded;
        REQUIRE(loaded.deserialize("test_openaddr.bin"));

        // Проверяем данные с проверкой на nullptr
        HashEntry* moscow = original.get("moscow");
        HashEntry* loadedMoscow = loaded.get("moscow");
        REQUIRE(moscow != nullptr);
        REQUIRE(loadedMoscow != nullptr);
        REQUIRE(moscow->value == loadedMoscow->value);

        HashEntry* paris = original.get("paris");
        HashEntry* loadedParis = loaded.get("paris");
        REQUIRE(paris != nullptr);
        REQUIRE(loadedParis != nullptr);
        REQUIRE(paris->value == loadedParis->value);
    }

    SECTION("SerializeDeserializeEmptyTest") {
        OpenAddrHash empty(5);
        REQUIRE(empty.serialize("empty_openaddr.bin"));

        OpenAddrHash loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_openaddr.bin"));
        REQUIRE(loadedEmpty.get("anything") == nullptr);
    }

    SECTION("SerializeDeserializeBasicTest") {
        OpenAddrHash table(5);
        REQUIRE(table.put("red", "color"));
        REQUIRE(table.put("blue", "color"));

        REQUIRE(table.serialize("basic_openaddr.bin"));

        OpenAddrHash loaded;
        REQUIRE(loaded.deserialize("basic_openaddr.bin"));

        HashEntry* red = table.get("red");
        HashEntry* loadedRed = loaded.get("red");
        REQUIRE(red != nullptr);
        REQUIRE(loadedRed != nullptr);
        REQUIRE(red->value == loadedRed->value);
    }
}