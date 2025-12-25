#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/SLL/SinglyLinkedList.h"

BOOST_AUTO_TEST_SUITE(SLLTests)


// Тест на вставку в начало и конец списка
BOOST_AUTO_TEST_CASE(InsertAtHeadTailTest) {
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
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Тест на вставку до указанного элемента
BOOST_AUTO_TEST_CASE(InsertBeforeTests) {
    SinglyLinkedList sll;
    {
        OutputRedirect redirect;
        sll.insertBefore("1", "2");
        string output = redirect.getOutput();
        string expected = "Список пустой, вставка невозможна\n";
        BOOST_CHECK_EQUAL(output, expected);
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
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        sll.insertBefore("10", "20");
        output = redirect.getOutput();
        expected = "Заданное значение 10 не найдено\n";
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Тест на вставку после указанного элемента
BOOST_AUTO_TEST_CASE(InsertAfterTests) {
    SinglyLinkedList sll;
    {
        OutputRedirect redirect;
        sll.insertAfter("1", "2");
        string output = redirect.getOutput();
        string expected = "Список пустой, вставка невозможна\n";
        BOOST_CHECK_EQUAL(output, expected);
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
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Комплексный тест на удаление из односвязного списка
// Проверяет все виды удаления
BOOST_AUTO_TEST_CASE(DeleteTests) {
    SinglyLinkedList sll;
    {
        OutputRedirect redirect;
        sll.deleteHead();
        string output = redirect.getOutput();
        string expected = "Список пуст\n";
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        sll.deleteTail();
        output = redirect.getOutput();
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        sll.deleteBefore("2");
        output = redirect.getOutput();
        expected = "Недостаточно элементов для удаления\n";
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        sll.deleteAfter("2");
        output = redirect.getOutput();
        expected = "Список пустой, удаление невозможно\n";
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        sll.deleteByValue("2");
        output = redirect.getOutput();
        BOOST_CHECK_EQUAL(output, "");
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
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Тест на поиск элементов в списке
BOOST_AUTO_TEST_CASE(SearchTests) {
    SinglyLinkedList sll;

    int foundElement = sll.searchByValue("1");
    BOOST_CHECK_EQUAL(foundElement, -1);

    sll.insertAtTail("2");
    sll.insertAtTail("3");
    sll.insertAtTail("4");
    foundElement = sll.searchByValue("4");
    BOOST_CHECK_EQUAL(foundElement, 2);

    foundElement = sll.searchByValue("10");
    BOOST_CHECK_EQUAL(foundElement, -1);
}


// Тест на сериализацию и десериализацию списка
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    SinglyLinkedList original;
    original.insertAtTail("first");
    original.insertAtTail("second");
    original.insertAtTail("third");

    BOOST_CHECK(original.serialize("test_sll.bin"));

    SinglyLinkedList loaded;
    BOOST_CHECK(loaded.deserialize("test_sll.bin"));

    BOOST_CHECK_EQUAL(original.searchByValue("first"), loaded.searchByValue("first"));
    BOOST_CHECK_EQUAL(original.searchByValue("second"), loaded.searchByValue("second"));
    BOOST_CHECK_EQUAL(original.searchByValue("third"), loaded.searchByValue("third"));
}

BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    SinglyLinkedList empty;
    BOOST_CHECK(empty.serialize("empty_sll.bin"));

    SinglyLinkedList loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_sll.bin"));
    BOOST_CHECK_EQUAL(loadedEmpty.searchByValue("anything"), -1);
}

BOOST_AUTO_TEST_SUITE_END()