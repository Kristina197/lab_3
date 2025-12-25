#include <cmath>

#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/Stack/Stack.h"

using namespace std;

BOOST_AUTO_TEST_SUITE(StackTests)

BOOST_AUTO_TEST_CASE(PushPopTest) {
    Stack st;
    st.push("10");
    st.push("777");
    st.push("smth");

    BOOST_CHECK_EQUAL(st.top(), "smth");
    BOOST_CHECK_EQUAL(st.pop(), "smth");
    BOOST_CHECK_EQUAL(st.pop(), "777");
}


// Тест на печать стека
BOOST_AUTO_TEST_CASE(PrintTest) {
    Stack st;
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

    BOOST_CHECK_EQUAL(output, expected);
}

BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    Stack original;
    original.push("first");
    original.push("second");
    original.push("third");

    BOOST_CHECK(original.serialize("test_stack.bin"));

    Stack loaded;
    BOOST_CHECK(loaded.deserialize("test_stack.bin"));

    BOOST_CHECK_EQUAL(original.top(), loaded.top());
    BOOST_CHECK_EQUAL(original.pop(), loaded.pop());
    BOOST_CHECK_EQUAL(original.pop(), loaded.pop());
    BOOST_CHECK_EQUAL(original.pop(), loaded.pop());
    BOOST_CHECK(original.top().empty());
    BOOST_CHECK(loaded.top().empty());
}

BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    Stack empty;
    BOOST_CHECK(empty.serialize("empty_stack.bin"));

    Stack loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_stack.bin"));
    BOOST_CHECK(loadedEmpty.top().empty());

    Stack invalid;
    BOOST_CHECK(!invalid.deserialize("non_existent.bin"));
}

BOOST_AUTO_TEST_SUITE_END()