#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/Stack/Stack.h"

class StackTest : public ::testing::Test {
protected:
    Stack st;
};

TEST_F(StackTest, PushPopTest) {
    st.push("10");
    st.push("777");
    st.push("smth");

    EXPECT_EQ(st.top(), "smth");
    EXPECT_EQ(st.pop(), "smth");
    EXPECT_EQ(st.pop(), "777");
}

TEST_F(StackTest, PrintTest) {
    OutputRedirect redirect;
    st.push("1");
    st.push("2");
    st.push("3");
    st.push("4");
    st.push("5");
    st.push("6");

    st.print();
    string output = redirect.getOutput();
    string expected = "top -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> bottom\n";

    EXPECT_EQ(output, expected);
}

TEST_F(StackTest, SerializeDeserializeTest) {
    Stack original;
    original.push("first");
    original.push("second");
    original.push("third");

    EXPECT_TRUE(original.serialize("test_stack.bin"));

    Stack loaded;
    EXPECT_TRUE(loaded.deserialize("test_stack.bin"));

    // Проверяем порядок LIFO - верхние элементы должны совпадать
    EXPECT_EQ(original.top(), loaded.top());

    // Проверяем все элементы
    EXPECT_EQ(original.pop(), "third");
    EXPECT_EQ(original.pop(), "second");
    EXPECT_EQ(original.pop(), "first");

    EXPECT_EQ(loaded.pop(), "third");
    EXPECT_EQ(loaded.pop(), "second");
    EXPECT_EQ(loaded.pop(), "first");

    EXPECT_TRUE(original.top().empty());
    EXPECT_TRUE(loaded.top().empty());
}

TEST_F(StackTest, SerializeDeserializeEmptyTest) {
    Stack empty;
    EXPECT_TRUE(empty.serialize("empty_stack.bin"));

    Stack loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_stack.bin"));
    EXPECT_TRUE(loadedEmpty.top().empty());

    // Проверяем ошибку загрузки
    Stack invalid;
    EXPECT_FALSE(invalid.deserialize("non_existent.bin"));
}