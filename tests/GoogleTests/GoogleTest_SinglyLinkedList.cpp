#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/SLL/SinglyLinkedList.h"

class SLLTest : public ::testing::Test {
protected:
    SinglyLinkedList sll;
};

TEST_F(SLLTest, InsertAtHeadTailTest) {
    sll.insertAtHead("1");
    sll.insertAtHead("2");
    sll.insertAtTail("3");
    sll.insertAtTail("4");
    sll.insertAtHead("5");

    {
        OutputRedirect redirect;
        sll.print();
        string output = redirect.getOutput();
        string expected = "5 2 1 3 4 \n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(SLLTest, InsertBeforeTests) {
    {
        OutputRedirect redirect;
        sll.insertBefore("1", "2");
        string output = redirect.getOutput();
        string expected = "Список пустой, вставка невозможна\n";
        EXPECT_EQ(output, expected);
    }

    sll.insertAtTail("1");
    sll.insertAtTail("2");
    sll.insertAtTail("3");
    sll.insertAtTail("4");
    sll.insertBefore("3", "5");
    sll.insertBefore("1", "6");     // Вставка перед головой

    {
        OutputRedirect redirect;
        sll.print();
        string output = redirect.getOutput();
        string expected = "6 1 2 5 3 4 \n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        sll.insertBefore("10", "20");
        output = redirect.getOutput();
        expected = "Заданное значение 10 не найдено\n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(SLLTest, InsertAfterTests) {
    {
        OutputRedirect redirect;
        sll.insertAfter("1", "2");
        string output = redirect.getOutput();
        string expected = "Список пустой, вставка невозможна\n";
        EXPECT_EQ(output, expected);
    }

    sll.insertAtTail("1");
    sll.insertAtTail("2");
    sll.insertAtTail("3");
    sll.insertAtTail("4");
    sll.insertAfter("3", "5");
    sll.insertAfter("5", "6");

    {
        OutputRedirect redirect;
        sll.print();
        string output = redirect.getOutput();
        string expected = "1 2 3 5 6 4 \n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(SLLTest, DeleteTests) {
    {
        OutputRedirect redirect;
        sll.deleteHead();
        string output = redirect.getOutput();
        string expected = "Список пуст\n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        sll.deleteTail();
        output = redirect.getOutput();
        EXPECT_EQ(output, expected);

        redirect.clear();
        sll.deleteBefore("2");
        output = redirect.getOutput();
        expected = "Недостаточно элементов для удаления\n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        sll.deleteAfter("2");
        output = redirect.getOutput();
        expected = "Список пустой, удаление невозможно\n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        sll.deleteByValue("2");
        output = redirect.getOutput();
        EXPECT_EQ(output, "");
    }

    sll.insertAtHead("1");      // *
    sll.deleteTail();

    sll.insertAtTail("2");      // *
    sll.insertAtTail("3");
    sll.insertAtTail("4");      // *
    sll.insertAtTail("5");
    sll.insertAtTail("6");      // *
    sll.insertAtTail("7");      // *
    sll.insertAtTail("8");
    sll.insertAtTail("9");      // *

    sll.deleteHead();
    sll.deleteTail();
    sll.deleteBefore("3");
    sll.deleteBefore("8");
    sll.deleteAfter("3");
    sll.deleteByValue("6");
    {
        OutputRedirect redirect;
        sll.print();
        string output = redirect.getOutput();
        string expected = "3 5 8 \n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(SLLTest, SearchTests) {
    int foundElement = sll.searchByValue("1");
    EXPECT_EQ(foundElement, -1);

    sll.insertAtTail("2");
    sll.insertAtTail("3");
    sll.insertAtTail("4");
    foundElement = sll.searchByValue("4");
    EXPECT_EQ(foundElement, 2);

    foundElement = sll.searchByValue("10");
    EXPECT_EQ(foundElement, -1);
}

TEST_F(SLLTest, SerializeDeserializeTest) {
    SinglyLinkedList original;
    original.insertAtTail("first");
    original.insertAtTail("second");
    original.insertAtTail("third");

    EXPECT_TRUE(original.serialize("test_sll.bin"));

    SinglyLinkedList loaded;
    EXPECT_TRUE(loaded.deserialize("test_sll.bin"));

    // Проверяем порядок элементов
    EXPECT_EQ(original.searchByValue("first"), loaded.searchByValue("first"));
    EXPECT_EQ(original.searchByValue("second"), loaded.searchByValue("second"));
    EXPECT_EQ(original.searchByValue("third"), loaded.searchByValue("third"));
}

TEST_F(SLLTest, SerializeDeserializeEmptyTest) {
    SinglyLinkedList empty;
    EXPECT_TRUE(empty.serialize("empty_sll.bin"));

    SinglyLinkedList loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_sll.bin"));
    EXPECT_EQ(loadedEmpty.searchByValue("anything"), -1);
}
