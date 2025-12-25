#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/DLL/DoublyLinkedList.h"

BOOST_AUTO_TEST_SUITE(DLLTests)

// Тест на операции с пустым списком
BOOST_AUTO_TEST_CASE(EmptyListOperations) {
    DoublyLinkedList dll;

    BOOST_CHECK_EQUAL(dll.searchByValue("A"), -1);
    BOOST_CHECK_NO_THROW(dll.print());
    BOOST_CHECK_NO_THROW(dll.printReverse());

    BOOST_CHECK_NO_THROW(dll.deleteHead());
    BOOST_CHECK_NO_THROW(dll.deleteTail());
    BOOST_CHECK_NO_THROW(dll.deleteByValue("A"));
    BOOST_CHECK_NO_THROW(dll.deleteBefore("A"));
    BOOST_CHECK_NO_THROW(dll.deleteAfter("A"));
}


// Тест на вставку в начало и конец списка
BOOST_AUTO_TEST_CASE(InsertHeadAndTail) {
    DoublyLinkedList dll;

    dll.insertAtHead("B");
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 0);

    dll.insertAtHead("A");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 1);

    dll.insertAtTail("C");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), 2);

    dll.insertAtTail("D");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), 2);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 3);
}


// Тест на вставку до и после указанного элемента
BOOST_AUTO_TEST_CASE(InsertBeforeAndAfter) {
    DoublyLinkedList dll;
    dll.insertAtTail("A");
    dll.insertAtTail("C");
    dll.insertAtTail("D");

    // Вставка в середину
    dll.insertAfter("A", "B");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), 2);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 3);

    // Вставка после хвоста
    dll.insertAfter("D", "E");
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 3);
    BOOST_CHECK_EQUAL(dll.searchByValue("E"), 4);

    // Вставка до головы
    dll.insertBefore("A", "Z");
    BOOST_CHECK_EQUAL(dll.searchByValue("Z"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 1);

    // Вставка до элемента в середине
    dll.insertBefore("C", "X");
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 2);
    BOOST_CHECK_EQUAL(dll.searchByValue("X"), 3);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), 4);

    dll.insertBefore("NonExistent", "W");
    dll.insertAfter("NonExistent", "W");

    BOOST_CHECK_EQUAL(dll.searchByValue("W"), -1);
    BOOST_CHECK_EQUAL(dll.searchByValue("X"), 3);
}


// Тест на удаление головы и хвоста списка
BOOST_AUTO_TEST_CASE(DeleteHeadAndTail) {
    DoublyLinkedList dll;
    dll.insertAtTail("A");
    dll.insertAtTail("B");
    dll.insertAtTail("C");

    dll.deleteHead();
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), -1);

    dll.deleteTail();
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), -1);

    dll.deleteHead();
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), -1);
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), -1);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), -1);


    BOOST_CHECK_NO_THROW(dll.deleteHead());
    BOOST_CHECK_NO_THROW(dll.deleteTail());
}


// Тест на удаление элемента по значению
BOOST_AUTO_TEST_CASE(DeleteByValue) {
    DoublyLinkedList dll;
    dll.insertAtTail("A");
    dll.insertAtTail("B");
    dll.insertAtTail("C");
    dll.insertAtTail("D");
    dll.insertAtTail("E");

    dll.deleteByValue("C");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 2);
    BOOST_CHECK_EQUAL(dll.searchByValue("E"), 3);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), -1);

    dll.deleteByValue("A");
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("E"), 2);
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), -1);

    dll.deleteByValue("E");
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("E"), -1);

    dll.deleteByValue("Z");
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 1);
}


// Тест на удаление до и после указанного элемента
BOOST_AUTO_TEST_CASE(DeleteBeforeAndAfter) {
    DoublyLinkedList dll;
    dll.insertAtTail("A");
    dll.insertAtTail("B");
    dll.insertAtTail("C");
    dll.insertAtTail("D");
    dll.insertAtTail("E");

    dll.deleteAfter("B");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 2);
    BOOST_CHECK_EQUAL(dll.searchByValue("E"), 3);
    BOOST_CHECK_EQUAL(dll.searchByValue("C"), -1);

    dll.deleteAfter("D");
    BOOST_CHECK_EQUAL(dll.searchByValue("E"), -1);

    dll.deleteBefore("D");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), 1);
    BOOST_CHECK_EQUAL(dll.searchByValue("B"), -1);

    BOOST_CHECK_NO_THROW(dll.deleteBefore("A"));
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);

    dll.deleteAfter("A");
    BOOST_CHECK_EQUAL(dll.searchByValue("A"), 0);
    BOOST_CHECK_EQUAL(dll.searchByValue("D"), -1);
}


// Тест на вывод списка в прямом и обратном порядке
BOOST_AUTO_TEST_CASE(PrintTests) {
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
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        dll.printReverse();
        output = redirect.getOutput();
        expected = "4 3 2 1 \n";
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Тест на сериализацию и десериализацию списка
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    DoublyLinkedList original;
    original.insertAtTail("first");
    original.insertAtTail("second");
    original.insertAtTail("third");

    BOOST_CHECK(original.serialize("test_dll.bin"));

    DoublyLinkedList loaded;
    BOOST_CHECK(loaded.deserialize("test_dll.bin"));

    BOOST_CHECK_EQUAL(original.searchByValue("first"), loaded.searchByValue("first"));
    BOOST_CHECK_EQUAL(original.searchByValue("second"), loaded.searchByValue("second"));
    BOOST_CHECK_EQUAL(original.searchByValue("third"), loaded.searchByValue("third"));
}

BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    DoublyLinkedList empty;
    BOOST_CHECK(empty.serialize("empty_dll.bin"));

    DoublyLinkedList loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_dll.bin"));
    BOOST_CHECK_EQUAL(loadedEmpty.searchByValue("anything"), -1);
}

BOOST_AUTO_TEST_SUITE_END()