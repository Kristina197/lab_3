#include <cmath>
#include <catch2/catch_all.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/FBT/FullBinaryTree.h"

TEST_CASE("Binary Tree Tests", "[tree]") {
    SECTION("InsertTests") {
        Tree fbt;

        REQUIRE(fbt.insertNode(15));
        fbt.insertNode(10);
        fbt.insertNode(25);
        fbt.insertNode(13);
        fbt.insertNode(20);
        {
            OutputRedirect redirect;
            fbt.preOrder();
            string output = redirect.getOutput();
            string expected = "15 10 13 25 20 \n";
            REQUIRE(output == expected);

            redirect.clear();
            fbt.inOrder();
            output = redirect.getOutput();
            expected = "10 13 15 20 25 \n";
            REQUIRE(output == expected);

            redirect.clear();
            fbt.postOrder();
            output = redirect.getOutput();
            expected = "13 10 20 25 15 \n";
            REQUIRE(output == expected);
        }

        REQUIRE_FALSE(fbt.insertNode(20));
    }

    SECTION("SearchIsFullTests") {
        Tree fbt;
        fbt.insertNode(27);
        fbt.insertNode(20);
        fbt.insertNode(35);
        fbt.insertNode(14);
        fbt.insertNode(26);
        fbt.insertNode(30);
        fbt.insertNode(44);

        REQUIRE(fbt.isFullBinary());
        fbt.insertNode(29);
        REQUIRE_FALSE(fbt.isFullBinary());

        REQUIRE(fbt.searchNode(20));
        REQUIRE_FALSE(fbt.searchNode(100));
    }

    SECTION("SerializeDeserializeTest") {
        Tree original;
        original.insertNode(50);
        original.insertNode(30);
        original.insertNode(70);
        original.insertNode(20);
        original.insertNode(40);

        REQUIRE(original.serialize("test_tree.bin"));

        Tree loaded;
        REQUIRE(loaded.deserialize("test_tree.bin"));

        // Проверяем структуру дерева
        REQUIRE(loaded.searchNode(50));
        REQUIRE(loaded.searchNode(30));
        REQUIRE(loaded.searchNode(70));
        REQUIRE(loaded.searchNode(20));
        REQUIRE(loaded.searchNode(40));
        REQUIRE_FALSE(loaded.searchNode(100));
    }

    SECTION("SerializeDeserializeEmptyTest") {
        Tree empty;
        REQUIRE(empty.serialize("empty_tree.bin"));

        Tree loadedEmpty;
        REQUIRE(loadedEmpty.deserialize("empty_tree.bin"));
        REQUIRE_FALSE(loadedEmpty.searchNode(50));
    }

    SECTION("SerializeDeserializeComplexTest") {
        Tree tree;
        tree.insertNode(100);
        tree.insertNode(50);
        tree.insertNode(150);
        tree.insertNode(25);
        tree.insertNode(75);
        tree.insertNode(125);
        tree.insertNode(175);

        REQUIRE(tree.serialize("complex_tree.bin"));

        Tree loaded;
        REQUIRE(loaded.deserialize("complex_tree.bin"));

        // Проверяем обходы
        {
            OutputRedirect redirect;
            tree.inOrder();
            string originalOutput = redirect.getOutput();

            redirect.clear();
            loaded.inOrder();
            string loadedOutput = redirect.getOutput();

            REQUIRE(originalOutput == loadedOutput);
        }
    }
}