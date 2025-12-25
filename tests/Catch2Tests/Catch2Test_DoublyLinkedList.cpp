#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/DLL/DoublyLinkedList.h"

TEST_CASE("Doubly Linked List Tests", "[dll]") {
    // Тест на операции с пустым списком
    SECTION("EmptyListOperations") {
        DoublyLinkedList dll;

        REQUIRE(dll.searchByValue("A") == -1);
        REQUIRE_NOTHROW(dll.print());
        REQUIRE_NOTHROW(dll.printReverse());

        REQUIRE_NOTHROW(dll.deleteHead());
        REQUIRE_NOTHROW(dll.deleteTail());
        REQUIRE_NOTHROW(dll.deleteByValue("A"));
        REQUIRE_NOTHROW(dll.deleteBefore("A"));
        REQUIRE_NOTHROW(dll.deleteAfter("A"));
    }

    // Тест на вставку в начало и конец списка
    SECTION("InsertHeadAndTail") {
        DoublyLinkedList dll;

        dll.insertAtHead("B");
        REQUIRE(dll.searchByValue("B") == 0);

        dll.insertAtHead("A");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("B") == 1);

        dll.insertAtTail("C");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("B") == 1);
        REQUIRE(dll.searchByValue("C") == 2);

        dll.insertAtTail("D");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("B") == 1);
        REQUIRE(dll.searchByValue("C") == 2);
        REQUIRE(dll.searchByValue("D") == 3);
    }

    // Тест на вставку до и после указанного элемента
    SECTION("InsertBeforeAndAfter") {
        DoublyLinkedList dll;
        dll.insertAtTail("A");
        dll.insertAtTail("C");
        dll.insertAtTail("D");

        // Вставка в середину
        dll.insertAfter("A", "B");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("B") == 1);
        REQUIRE(dll.searchByValue("C") == 2);
        REQUIRE(dll.searchByValue("D") == 3);

        // Вставка после хвоста
        dll.insertAfter("D", "E");
        REQUIRE(dll.searchByValue("D") == 3);
        REQUIRE(dll.searchByValue("E") == 4);

        // Вставка до головы
        dll.insertBefore("A", "Z");
        REQUIRE(dll.searchByValue("Z") == 0);
        REQUIRE(dll.searchByValue("A") == 1);

        // Вставка до элемента в середине
        dll.insertBefore("C", "X");
        REQUIRE(dll.searchByValue("B") == 2);
        REQUIRE(dll.searchByValue("X") == 3);
        REQUIRE(dll.searchByValue("C") == 4);

        dll.insertBefore("NonExistent", "W");
        dll.insertAfter("NonExistent", "W");

        REQUIRE(dll.searchByValue("W") == -1);
        REQUIRE(dll.searchByValue("X") == 3);
    }


    // Тест на удаление головы и хвоста списка
    SECTION("DeleteHeadAndTail") {
        DoublyLinkedList dll;
        dll.insertAtTail("A");
        dll.insertAtTail("B");
        dll.insertAtTail("C");

        dll.deleteHead();
        REQUIRE(dll.searchByValue("B") == 0);
        REQUIRE(dll.searchByValue("C") == 1);
        REQUIRE(dll.searchByValue("A") == -1);

        dll.deleteTail();
        REQUIRE(dll.searchByValue("B") == 0);
        REQUIRE(dll.searchByValue("C") == -1);

        dll.deleteHead();
        REQUIRE(dll.searchByValue("B") == -1);
        REQUIRE(dll.searchByValue("A") == -1);
        REQUIRE(dll.searchByValue("C") == -1);

        REQUIRE_NOTHROW(dll.deleteHead());
        REQUIRE_NOTHROW(dll.deleteTail());
    }


    // Тест на удаление элемента по значению
    SECTION("DeleteByValue") {
        DoublyLinkedList dll;
        dll.insertAtTail("A");
        dll.insertAtTail("B");
        dll.insertAtTail("C");
        dll.insertAtTail("D");
        dll.insertAtTail("E");

        dll.deleteByValue("C");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("B") == 1);
        REQUIRE(dll.searchByValue("D") == 2);
        REQUIRE(dll.searchByValue("E") == 3);
        REQUIRE(dll.searchByValue("C") == -1);

        dll.deleteByValue("A");
        REQUIRE(dll.searchByValue("B") == 0);
        REQUIRE(dll.searchByValue("D") == 1);
        REQUIRE(dll.searchByValue("E") == 2);
        REQUIRE(dll.searchByValue("A") == -1);

        dll.deleteByValue("E");
        REQUIRE(dll.searchByValue("B") == 0);
        REQUIRE(dll.searchByValue("D") == 1);
        REQUIRE(dll.searchByValue("E") == -1);

        dll.deleteByValue("Z");
        REQUIRE(dll.searchByValue("B") == 0);
        REQUIRE(dll.searchByValue("D") == 1);
    }


    // Тест на удаление до и после указанного элемента
    SECTION("DeleteBeforeAndAfter") {
        DoublyLinkedList dll;
        dll.insertAtTail("A");
        dll.insertAtTail("B");
        dll.insertAtTail("C");
        dll.insertAtTail("D");
        dll.insertAtTail("E");

        dll.deleteAfter("B");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("B") == 1);
        REQUIRE(dll.searchByValue("D") == 2);
        REQUIRE(dll.searchByValue("E") == 3);
        REQUIRE(dll.searchByValue("C") == -1);

        dll.deleteAfter("D");
        REQUIRE(dll.searchByValue("E") == -1);

        dll.deleteBefore("D");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("D") == 1);
        REQUIRE(dll.searchByValue("B") == -1);

        REQUIRE_NOTHROW(dll.deleteBefore("A"));
        REQUIRE(dll.searchByValue("A") == 0);

        dll.deleteAfter("A");
        REQUIRE(dll.searchByValue("A") == 0);
        REQUIRE(dll.searchByValue("D") == -1);
    }


    // Тест на вывод списка в прямом и обратном порядке
    SECTION("PrintTests") {
        DoublyLinkedList dll;
        dll.insertAtTail("1");
        dll.insertAtTail("2");
        dll.insertAtTail("3");
        dll.insertAtTail("4");

        {
            OutputRedirect redirect;
            dll.print();
            string output = redirect.getOutput();
            string expected = "1 2 3 4 \n";
            REQUIRE(output == expected);

            redirect.clear();
            dll.printReverse();
            output = redirect.getOutput();
            expected = "4 3 2 1 \n";
            REQUIRE(output == expected);
        }
    }

    SECTION("SerializeDeserializeTest") {
        DoublyLinkedList original;
        original.insertAtTail("first");
        original.insertAtTail("second");
        original.insertAtTail("third");

        REQUIRE(original.serialize("test_dll.bin"));

        DoublyLinkedList loaded;
        REQUIRE(loaded.deserialize("test_dll.bin"));

        REQUIRE(original.searchByValue("first") == loaded.searchByValue("first"));
        REQUIRE(original.searchByValue("second") == loaded.searchByValue("second"));
        REQUIRE(original.searchByValue("third") == loaded.searchByValue("third"));
    }

    SECTION("SerializeDeserializeEmptyTest") {
        DoublyLinkedList empty;
        REQUIRE(empty.serialize("empty_dll.bin"));

        DoublyLinkedList loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_dll.bin"));
        REQUIRE(loadedEmpty.searchByValue("anything") == -1);
    }
}