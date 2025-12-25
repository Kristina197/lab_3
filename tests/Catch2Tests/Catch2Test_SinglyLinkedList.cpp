#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/SLL/SinglyLinkedList.h"

TEST_CASE("Singly Linked List Tests", "[sll]") {
    SECTION("InsertAtHeadTailTest") {
        SinglyLinkedList sll;
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
            REQUIRE(output == expected);
        }
    }

    SECTION("InsertBeforeTests") {
        SinglyLinkedList sll;
        {
            OutputRedirect redirect;
            sll.insertBefore("1", "2");
            string output = redirect.getOutput();
            string expected = "Список пустой, вставка невозможна\n";
            REQUIRE(output == expected);
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
            REQUIRE(output == expected);

            redirect.clear();
            sll.insertBefore("10", "20");
            output = redirect.getOutput();
            expected = "Заданное значение 10 не найдено\n";
            REQUIRE(output == expected);
        }
    }

    SECTION("InsertAfterTests") {
        SinglyLinkedList sll;
        {
            OutputRedirect redirect;
            sll.insertAfter("1", "2");
            string output = redirect.getOutput();
            string expected = "Список пустой, вставка невозможна\n";
            REQUIRE(output == expected);
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
            REQUIRE(output == expected);
        }
    }

    SECTION("DeleteTests") {
        SinglyLinkedList sll;
        {
            OutputRedirect redirect;
            sll.deleteHead();
            string output = redirect.getOutput();
            string expected = "Список пуст\n";
            REQUIRE(output == expected);

            redirect.clear();
            sll.deleteTail();
            output = redirect.getOutput();
            REQUIRE(output == expected);

            redirect.clear();
            sll.deleteBefore("2");
            output = redirect.getOutput();
            expected = "Недостаточно элементов для удаления\n";
            REQUIRE(output == expected);

            redirect.clear();
            sll.deleteAfter("2");
            output = redirect.getOutput();
            expected = "Список пустой, удаление невозможно\n";
            REQUIRE(output == expected);

            redirect.clear();
            sll.deleteByValue("2");
            output = redirect.getOutput();
            REQUIRE(output == "");
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
            REQUIRE(output == expected);
        }
    }

    SECTION("SearchTests") {
        SinglyLinkedList sll;

        int foundElement = sll.searchByValue("1");
        REQUIRE(foundElement == -1);

        sll.insertAtTail("2");
        sll.insertAtTail("3");
        sll.insertAtTail("4");
        foundElement = sll.searchByValue("4");
        REQUIRE(foundElement == 2);

        foundElement = sll.searchByValue("10");
        REQUIRE(foundElement == -1);
    }

    SECTION("SerializeDeserializeTest") {
        SinglyLinkedList original;
        original.insertAtTail("first");
        original.insertAtTail("second");
        original.insertAtTail("third");

        REQUIRE(original.serialize("test_sll.bin"));

        SinglyLinkedList loaded;
        REQUIRE(loaded.deserialize("test_sll.bin"));

        // Проверяем порядок элементов
        REQUIRE(original.searchByValue("first") == loaded.searchByValue("first"));
        REQUIRE(original.searchByValue("second") == loaded.searchByValue("second"));
        REQUIRE(original.searchByValue("third") == loaded.searchByValue("third"));
    }

    SECTION("SerializeDeserializeEmptyTest") {
        SinglyLinkedList empty;
        REQUIRE(empty.serialize("empty_sll.bin"));

        SinglyLinkedList loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_sll.bin"));
        REQUIRE(loadedEmpty.searchByValue("anything") == -1);
    }
}