#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/Stack/Stack.h"


TEST_CASE("Stack Tests", "[stack]") {
    SECTION("PushPopTest") {
        Stack st;
        st.push("10");
        st.push("777");
        st.push("smth");

        REQUIRE(st.top() == "smth");
        REQUIRE(st.pop() == "smth");
        REQUIRE(st.pop() == "777");
    }

    SECTION("PrintTest") {
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

        REQUIRE(output == expected);
    }

    SECTION("SerializeDeserializeTest") {
        Stack original;
        original.push("first");
        original.push("second");
        original.push("third");

        REQUIRE(original.serialize("test_stack.bin"));

        Stack loaded;
        REQUIRE(loaded.deserialize("test_stack.bin"));

        // Проверяем порядок LIFO
        REQUIRE(original.top() == loaded.top());
        REQUIRE(original.pop() == loaded.pop());
        REQUIRE(original.pop() == loaded.pop());
        REQUIRE(original.pop() == loaded.pop());
        REQUIRE(original.top().empty());
        REQUIRE(loaded.top().empty());
    }

    SECTION("SerializeDeserializeEmptyTest") {
        Stack empty;
        REQUIRE(empty.serialize("empty_stack.bin"));

        Stack loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_stack.bin"));
        REQUIRE(loadedEmpty.top().empty());

        // Проверяем ошибку загрузки
        Stack invalid;
        REQUIRE_FALSE(invalid.deserialize("non_existent.bin"));
    }

}

