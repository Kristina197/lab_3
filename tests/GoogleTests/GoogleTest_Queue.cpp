#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/Queue/Queue.h"

class QueueTest : public ::testing::Test {
protected:
    Queue q;
};

TEST_F(QueueTest, EnqueueDequeueTests) {
    q.enqueue("hello");
    q.enqueue("hi");
    q.enqueue("welcome");

    {
        OutputRedirect redirect;
        q.print();
        string output = redirect.getOutput();
        string expected = "hello hi welcome \n";
        EXPECT_EQ(output, expected);
    }

    {
        OutputRedirect redirect;
        q.dequeue();
        q.enqueue("wassup");
        q.print();
        string output = redirect.getOutput();
        string expected = "hi welcome wassup \n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(QueueTest, EmptyQueueTests) {
    q.enqueue("seeya");
    EXPECT_EQ(q.length(), 1);

    q.dequeue();

    {
        OutputRedirect redirect;
        q.print();
        string output = redirect.getOutput();
        string expected = "Очередь пустая\n";
        EXPECT_EQ(output, expected);
    }
}

TEST_F(QueueTest, SerializeDeserializeTest) {
    Queue original;
    original.enqueue("first");
    original.enqueue("second");
    original.enqueue("third");

    EXPECT_TRUE(original.serialize("test_queue.bin"));

    Queue loaded;
    EXPECT_TRUE(loaded.deserialize("test_queue.bin"));

    // Проверяем порядок FIFO
    EXPECT_EQ(original.dequeue(), loaded.dequeue());
    EXPECT_EQ(original.dequeue(), loaded.dequeue());
    EXPECT_EQ(original.dequeue(), loaded.dequeue());
    EXPECT_TRUE(original.length() == 0 && loaded.length() == 0);
}

TEST_F(QueueTest, SerializeDeserializeEmptyTest) {
    Queue empty;
    EXPECT_TRUE(empty.serialize("empty_queue.bin"));

    Queue loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_queue.bin"));
    EXPECT_EQ(loadedEmpty.length(), 0);

    // Проверяем ошибку загрузки
    Queue invalid;
    EXPECT_FALSE(invalid.deserialize("non_existent.bin"));
}