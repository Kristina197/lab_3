#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/Queue/Queue.h"

TEST_CASE("Queue Tests", "[queue]") {
    SECTION("EnqueueDequeueTests") {
        Queue q;
        q.enqueue("hello");
        q.enqueue("hi");
        q.enqueue("welcome");

        {
            OutputRedirect redirect;
            q.print();
            string output = redirect.getOutput();
            string expected = "hello hi welcome \n";
            REQUIRE(output == expected);
        }

        {
            OutputRedirect redirect;
            q.dequeue();
            q.enqueue("wassup");
            q.print();
            string output = redirect.getOutput();
            string expected = "hi welcome wassup \n";
            REQUIRE(output == expected);
        }
    }

    SECTION("EmptyQueueTests") {
        Queue q;
        q.enqueue("seeya");
        REQUIRE(q.length() == 1);

        q.dequeue();

        {
            OutputRedirect redirect;
            q.print();
            string output = redirect.getOutput();
            string expected = "Очередь пустая\n";
            REQUIRE(output == expected);
        }
    }

    SECTION("SerializeDeserializeTest") {
        Queue original;
        original.enqueue("first");
        original.enqueue("second");
        original.enqueue("third");

        REQUIRE(original.serialize("test_queue.bin"));

        Queue loaded;
        REQUIRE(loaded.deserialize("test_queue.bin"));

        // Проверяем порядок FIFO
        REQUIRE(original.dequeue() == loaded.dequeue());
        REQUIRE(original.dequeue() == loaded.dequeue());
        REQUIRE(original.dequeue() == loaded.dequeue());
        REQUIRE(original.length() == 0);
        REQUIRE(loaded.length() == 0);
    }

    SECTION("SerializeDeserializeEmptyTest") {
        Queue empty;
        REQUIRE(empty.serialize("empty_queue.bin"));

        Queue loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_queue.bin"));
        REQUIRE(loadedEmpty.length() == 0);

        // Проверяем ошибку загрузки
        Queue invalid;
        REQUIRE_FALSE(invalid.deserialize("non_existent.bin"));
    }
}
