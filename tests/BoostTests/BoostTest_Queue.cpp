#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/Queue/Queue.h"

BOOST_AUTO_TEST_SUITE(QueueTests)


// Тест на базовые операции очереди: добавление и удаление
BOOST_AUTO_TEST_CASE(EnqueueDequeueTests) {
    Queue q;
    q.enqueue("hello");
    q.enqueue("hi");
    q.enqueue("welcome");

    {
        OutputRedirect redirect;
        q.print();
        string output = redirect.getOutput();
        string expected = "hello hi welcome \n";
        BOOST_CHECK_EQUAL(output, expected);
    }

    {
        OutputRedirect redirect;
        q.dequeue();
        q.enqueue("wassup");
        q.print();
        string output = redirect.getOutput();
        string expected = "hi welcome wassup \n";
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Тест на операции с пустой очередью
BOOST_AUTO_TEST_CASE(EmptyQueueTests) {
    Queue q;
    q.enqueue("seeya");
    BOOST_CHECK_EQUAL(q.length(), 1);

    q.dequeue();

    {
        OutputRedirect redirect;
        q.print();
        string output = redirect.getOutput();
        string expected = "Очередь пустая\n";
        BOOST_CHECK_EQUAL(output, expected);
    }
}


// Тест на сериализацию и десериализацию очереди
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    Queue original;
    original.enqueue("first");
    original.enqueue("second");
    original.enqueue("third");

    BOOST_CHECK(original.serialize("test_queue.bin"));

    Queue loaded;
    BOOST_CHECK(loaded.deserialize("test_queue.bin"));

    BOOST_CHECK_EQUAL(original.dequeue(), loaded.dequeue());
    BOOST_CHECK_EQUAL(original.dequeue(), loaded.dequeue());
    BOOST_CHECK_EQUAL(original.dequeue(), loaded.dequeue());
    BOOST_CHECK(original.length() == 0 && loaded.length() == 0);
}



BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    Queue empty;
    BOOST_CHECK(empty.serialize("empty_queue.bin"));

    Queue loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_queue.bin"));
    BOOST_CHECK_EQUAL(loadedEmpty.length(), 0);

    Queue invalid;
    BOOST_CHECK(!invalid.deserialize("non_existent.bin"));
}

BOOST_AUTO_TEST_SUITE_END()
