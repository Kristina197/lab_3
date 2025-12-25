#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/DLL/DoublyLinkedList.h"

class DLLTest : public ::testing::Test {
protected:
    DoublyLinkedList dll;
};

TEST_F(DLLTest, EmptyListOperations) {
    EXPECT_EQ(dll.searchByValue("A"), -1);
    EXPECT_NO_THROW(dll.print());
    EXPECT_NO_THROW(dll.printReverse());

    EXPECT_NO_THROW(dll.deleteHead());
    EXPECT_NO_THROW(dll.deleteTail());
    EXPECT_NO_THROW(dll.deleteByValue("A"));
    EXPECT_NO_THROW(dll.deleteBefore("A"));
    EXPECT_NO_THROW(dll.deleteAfter("A"));
}

TEST_F(DLLTest, InsertHeadAndTail) {
    dll.insertAtHead("B");
    EXPECT_EQ(dll.searchByValue("B"), 0);

    dll.insertAtHead("A");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("B"), 1);

    dll.insertAtTail("C");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("B"), 1);
    EXPECT_EQ(dll.searchByValue("C"), 2);

    dll.insertAtTail("D");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("B"), 1);
    EXPECT_EQ(dll.searchByValue("C"), 2);
    EXPECT_EQ(dll.searchByValue("D"), 3);
}

TEST_F(DLLTest, InsertBeforeAndAfter) {
    dll.insertAtTail("A");
    dll.insertAtTail("C");
    dll.insertAtTail("D");

    // Вставка в середину
    dll.insertAfter("A", "B");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("B"), 1);
    EXPECT_EQ(dll.searchByValue("C"), 2);
    EXPECT_EQ(dll.searchByValue("D"), 3);

    // Вставка после хвоста
    dll.insertAfter("D", "E");
    EXPECT_EQ(dll.searchByValue("D"), 3);
    EXPECT_EQ(dll.searchByValue("E"), 4);

    // Вставка до головы
    dll.insertBefore("A", "Z");
    EXPECT_EQ(dll.searchByValue("Z"), 0);
    EXPECT_EQ(dll.searchByValue("A"), 1);

    // Вставка до элемента в середине
    dll.insertBefore("C", "X");
    EXPECT_EQ(dll.searchByValue("B"), 2);
    EXPECT_EQ(dll.searchByValue("X"), 3);
    EXPECT_EQ(dll.searchByValue("C"), 4);

    dll.insertBefore("NonExistent", "W");
    dll.insertAfter("NonExistent", "W");

    EXPECT_EQ(dll.searchByValue("W"), -1);
    EXPECT_EQ(dll.searchByValue("X"), 3);
}

TEST_F(DLLTest, DeleteHeadAndTail) {
    dll.insertAtTail("A");
    dll.insertAtTail("B");
    dll.insertAtTail("C");

    dll.deleteHead();
    EXPECT_EQ(dll.searchByValue("B"), 0);
    EXPECT_EQ(dll.searchByValue("C"), 1);
    EXPECT_EQ(dll.searchByValue("A"), -1);

    dll.deleteTail();
    EXPECT_EQ(dll.searchByValue("B"), 0);
    EXPECT_EQ(dll.searchByValue("C"), -1);

    dll.deleteHead();
    EXPECT_EQ(dll.searchByValue("B"), -1);
    EXPECT_EQ(dll.searchByValue("A"), -1);
    EXPECT_EQ(dll.searchByValue("C"), -1);

    EXPECT_NO_THROW(dll.deleteHead());
    EXPECT_NO_THROW(dll.deleteTail());
}

TEST_F(DLLTest, DeleteByValue) {
    dll.insertAtTail("A");
    dll.insertAtTail("B");
    dll.insertAtTail("C");
    dll.insertAtTail("D");
    dll.insertAtTail("E");

    dll.deleteByValue("C");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("B"), 1);
    EXPECT_EQ(dll.searchByValue("D"), 2);
    EXPECT_EQ(dll.searchByValue("E"), 3);
    EXPECT_EQ(dll.searchByValue("C"), -1);

    dll.deleteByValue("A");
    EXPECT_EQ(dll.searchByValue("B"), 0);
    EXPECT_EQ(dll.searchByValue("D"), 1);
    EXPECT_EQ(dll.searchByValue("E"), 2);
    EXPECT_EQ(dll.searchByValue("A"), -1);

    dll.deleteByValue("E");
    EXPECT_EQ(dll.searchByValue("B"), 0);
    EXPECT_EQ(dll.searchByValue("D"), 1);
    EXPECT_EQ(dll.searchByValue("E"), -1);

    dll.deleteByValue("Z");
    EXPECT_EQ(dll.searchByValue("B"), 0);
    EXPECT_EQ(dll.searchByValue("D"), 1);
}

TEST_F(DLLTest, DeleteBeforeAndAfter) {
    dll.insertAtTail("A");
    dll.insertAtTail("B");
    dll.insertAtTail("C");
    dll.insertAtTail("D");
    dll.insertAtTail("E");

    dll.deleteAfter("B");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("B"), 1);
    EXPECT_EQ(dll.searchByValue("D"), 2);
    EXPECT_EQ(dll.searchByValue("E"), 3);
    EXPECT_EQ(dll.searchByValue("C"), -1);

    dll.deleteAfter("D");
    EXPECT_EQ(dll.searchByValue("E"), -1);

    dll.deleteBefore("D");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("D"), 1);
    EXPECT_EQ(dll.searchByValue("B"), -1);

    EXPECT_NO_THROW(dll.deleteBefore("A"));
    EXPECT_EQ(dll.searchByValue("A"), 0);

    dll.deleteAfter("A");
    EXPECT_EQ(dll.searchByValue("A"), 0);
    EXPECT_EQ(dll.searchByValue("D"), -1);
}

TEST_F(DLLTest, PrintTests) {
    dll.insertAtTail("1");
    dll.insertAtTail("2");
    dll.insertAtTail("3");
    dll.insertAtTail("4");

    {
        OutputRedirect redirect;
        dll.print();
        string output = redirect.getOutput();
        string expected = "1 2 3 4 \n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        dll.printReverse();
        output = redirect.getOutput();
        expected = "4 3 2 1 \n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(DLLTest, SerializeDeserializeTest) {
    DoublyLinkedList original;
    original.insertAtTail("first");
    original.insertAtTail("second");
    original.insertAtTail("third");

    EXPECT_TRUE(original.serialize("test_dll.bin"));

    DoublyLinkedList loaded;
    EXPECT_TRUE(loaded.deserialize("test_dll.bin"));

    // Проверяем порядок элементов в прямом направлении
    EXPECT_EQ(original.searchByValue("first"), loaded.searchByValue("first"));
    EXPECT_EQ(original.searchByValue("second"), loaded.searchByValue("second"));
    EXPECT_EQ(original.searchByValue("third"), loaded.searchByValue("third"));
}

TEST_F(DLLTest, SerializeDeserializeEmptyTest) {
    DoublyLinkedList empty;
    EXPECT_TRUE(empty.serialize("empty_dll.bin"));

    DoublyLinkedList loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_dll.bin"));
    EXPECT_EQ(loadedEmpty.searchByValue("anything"), -1);
}
