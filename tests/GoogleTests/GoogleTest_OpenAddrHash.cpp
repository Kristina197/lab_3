#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/OpenAddrHash/OpenAddrHash.h"

class OpenAddressingTest : public ::testing::Test {
protected:
    OpenAddrHash oah{5};
};

TEST_F(OpenAddressingTest, EmptyTableOperations) {
    EXPECT_FALSE(oah.get("1"));
    EXPECT_NO_THROW(oah.remove("1"));
}

TEST_F(OpenAddressingTest, BasicPutAndGet) {
    OpenAddrHash oah(10);

    EXPECT_TRUE(oah.put("apple", "100"));

    HashEntry* entry = oah.get("apple");

    ASSERT_NE(entry, nullptr);
    EXPECT_EQ(entry->key, "apple");
    EXPECT_EQ(entry->value, "100");
    EXPECT_EQ(entry->status, HashStatus::TAKEN);
}

TEST_F(OpenAddressingTest, OverflowTest) {
    OpenAddrHash oah(3);
    oah.put("1", "100");
    oah.put("2", "200");
    oah.put("3", "300");
    {
        OutputRedirect redirect;
        oah.put("4", "400");
        string output = redirect.getOutput();
        EXPECT_EQ(output, "Хеш-таблица переполнена\n");
    }
}

TEST_F(OpenAddressingTest, GetNonExistent) {
    OpenAddrHash oah(10);
    oah.put("apple", "100");

    HashEntry* entry = oah.get("banana");
    EXPECT_EQ(entry, nullptr);
}

TEST_F(OpenAddressingTest, UpdateValue) {
    OpenAddrHash oah;
    oah.put("apple", "100");
    oah.put("apple", "200");

    HashEntry* entry = oah.get("apple");
    ASSERT_NE(entry, nullptr);
    EXPECT_EQ(entry->value, "200");
}

TEST_F(OpenAddressingTest, CollisionsTests) {
    OpenAddrHash oah(5);

    EXPECT_TRUE(oah.put("0", "zero"));
    EXPECT_TRUE(oah.put("5", "five"));
    EXPECT_TRUE(oah.put("10", "ten"));

    HashEntry* entry0 = oah.get("0");
    ASSERT_NE(entry0, nullptr);
    EXPECT_EQ(entry0->value, "zero");

    HashEntry* entry5 = oah.get("5");
    ASSERT_NE(entry5, nullptr);
    EXPECT_EQ(entry5->value, "five");

    HashEntry* entry10 = oah.get("10");
    ASSERT_NE(entry10, nullptr);
    EXPECT_EQ(entry10->value, "ten");

    HashEntry* entry15 = oah.get("15");
    EXPECT_EQ(entry15, nullptr);
}

TEST_F(OpenAddressingTest, PutIntoDeletedSlot) {
    OpenAddrHash oah(3);

    oah.put("hello", "goodbye");
    oah.put("smth", "testing");
    oah.put("active", "lazy");
    oah.remove("smth");

    auto* entry = oah.get("smth");
    EXPECT_FALSE(entry);

    oah.put("epic", "comeback");
    entry = oah.get("epic");
    EXPECT_TRUE(entry);
}

TEST_F(OpenAddressingTest, PrintTest) {
    OpenAddrHash oah(5);

    EXPECT_NO_THROW(oah.print());

    oah.put("1", "one");
    oah.put("2", "two");
    EXPECT_NO_THROW(oah.print());

    oah.remove("1");
    EXPECT_NO_THROW(oah.print());
}

TEST_F(OpenAddressingTest, SerializeDeserializeTest) {
    OpenAddrHash original(10);
    EXPECT_TRUE(original.put("moscow", "russia"));
    EXPECT_TRUE(original.put("paris", "france"));
    EXPECT_TRUE(original.put("tokyo", "japan"));
    EXPECT_TRUE(original.put("berlin", "germany"));

    EXPECT_TRUE(original.serialize("test_openaddr.bin"));

    OpenAddrHash loaded;
    EXPECT_TRUE(loaded.deserialize("test_openaddr.bin"));

    HashEntry* moscow = original.get("moscow");
    HashEntry* loadedMoscow = loaded.get("moscow");
    ASSERT_NE(moscow, nullptr);
    ASSERT_NE(loadedMoscow, nullptr);
    EXPECT_EQ(moscow->value, loadedMoscow->value);

    HashEntry* paris = original.get("paris");
    HashEntry* loadedParis = loaded.get("paris");
    ASSERT_NE(paris, nullptr);
    ASSERT_NE(loadedParis, nullptr);
    EXPECT_EQ(paris->value, loadedParis->value);
}

TEST_F(OpenAddressingTest, SerializeDeserializeEmptyTest) {
    OpenAddrHash empty(5);
    EXPECT_TRUE(empty.serialize("empty_openaddr.bin"));

    OpenAddrHash loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_openaddr.bin"));
    EXPECT_EQ(loadedEmpty.get("anything"), nullptr);
}

TEST_F(OpenAddressingTest, SerializeDeserializeBasicTest) {
    OpenAddrHash table(5);
    EXPECT_TRUE(table.put("red", "color"));
    EXPECT_TRUE(table.put("blue", "color"));

    EXPECT_TRUE(table.serialize("basic_openaddr.bin"));

    OpenAddrHash loaded;
    EXPECT_TRUE(loaded.deserialize("basic_openaddr.bin"));

    HashEntry* red = table.get("red");
    HashEntry* loadedRed = loaded.get("red");
    ASSERT_NE(red, nullptr);
    ASSERT_NE(loadedRed, nullptr);
    EXPECT_EQ(red->value, loadedRed->value);
}
