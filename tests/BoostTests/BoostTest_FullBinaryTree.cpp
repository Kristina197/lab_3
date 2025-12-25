#include <boost/test/unit_test.hpp>
#include "../../src/redirect/OutputRedirect.h"
#include "../../src/FBT/FullBinaryTree.h"

BOOST_AUTO_TEST_SUITE(BinaryTreeTests)

// Тест на вставку узлов и проверку обходов дерева
BOOST_AUTO_TEST_CASE(InsertTests) {
    Tree fbt;

    BOOST_CHECK(fbt.insertNode(15));
    fbt.insertNode(10);
    fbt.insertNode(25);
    fbt.insertNode(13);
    fbt.insertNode(20);
    {
        OutputRedirect redirect;
        fbt.preOrder();
        string output = redirect.getOutput();
        string expected = "15 10 13 25 20 \n";
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        fbt.inOrder();
        output = redirect.getOutput();
        expected = "10 13 15 20 25 \n";
        BOOST_CHECK_EQUAL(output, expected);

        redirect.clear();
        fbt.postOrder();
        output = redirect.getOutput();
        expected = "13 10 20 25 15 \n";
        BOOST_CHECK_EQUAL(output, expected);
    }

    BOOST_CHECK(!fbt.insertNode(20));
}


// Тест на поиск узлов и проверку свойства Full Binary Tree
BOOST_AUTO_TEST_CASE(SearchIsFullTests) {
    Tree fbt;
    fbt.insertNode(27);
    fbt.insertNode(20);
    fbt.insertNode(35);
    fbt.insertNode(14);
    fbt.insertNode(26);
    fbt.insertNode(30);
    fbt.insertNode(44);

    BOOST_CHECK(fbt.isFullBinary());
    fbt.insertNode(29);
    BOOST_CHECK(!fbt.isFullBinary());

    BOOST_CHECK(fbt.searchNode(20));
    BOOST_CHECK(!fbt.searchNode(100));
}


// Тест на сериализацию и десериализацию дерева
BOOST_AUTO_TEST_CASE(SerializeDeserializeTest) {
    Tree original;
    original.insertNode(50);
    original.insertNode(30);
    original.insertNode(70);
    original.insertNode(20);
    original.insertNode(40);

    BOOST_CHECK(original.serialize("test_tree.bin"));

    Tree loaded;
    BOOST_CHECK(loaded.deserialize("test_tree.bin"));

    BOOST_CHECK(loaded.searchNode(50));
    BOOST_CHECK(loaded.searchNode(30));
    BOOST_CHECK(loaded.searchNode(70));
    BOOST_CHECK(loaded.searchNode(20));
    BOOST_CHECK(loaded.searchNode(40));
    BOOST_CHECK(!loaded.searchNode(100));
}


// Тест на сериализацию и десериализацию пустого дерева
BOOST_AUTO_TEST_CASE(SerializeDeserializeEmptyTest) {
    Tree empty;
    BOOST_CHECK(empty.serialize("empty_tree.bin"));

    Tree loadedEmpty;
    BOOST_CHECK(loadedEmpty.deserialize("empty_tree.bin"));
    BOOST_CHECK(!loadedEmpty.searchNode(50));
}



BOOST_AUTO_TEST_CASE(SerializeDeserializeComplexTest) {
    Tree tree;
    tree.insertNode(100);
    tree.insertNode(50);
    tree.insertNode(150);
    tree.insertNode(25);
    tree.insertNode(75);
    tree.insertNode(125);
    tree.insertNode(175);

    BOOST_CHECK(tree.serialize("complex_tree.bin"));

    Tree loaded;
    BOOST_CHECK(loaded.deserialize("complex_tree.bin"));

    {
        OutputRedirect redirect;
        tree.inOrder();
        string originalOutput = redirect.getOutput();

        redirect.clear();
        loaded.inOrder();
        string loadedOutput = redirect.getOutput();

        BOOST_CHECK_EQUAL(originalOutput, loadedOutput);
    }
}

BOOST_AUTO_TEST_SUITE_END()