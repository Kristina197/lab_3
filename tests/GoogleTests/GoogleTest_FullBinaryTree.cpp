#include <cmath>
#include <gtest/gtest.h>
#include "../../src/redirect/OutputRedirect.h"

#include "../../src/FBT/FullBinaryTree.h"

class BinaryTreeTest : public ::testing::Test {
protected:
    Tree fbt;
};

TEST_F(BinaryTreeTest, InsertTests) {
    EXPECT_TRUE(fbt.insertNode(15));
    fbt.insertNode(10);
    fbt.insertNode(25);
    fbt.insertNode(13);
    fbt.insertNode(20);
    {
        OutputRedirect redirect;
        fbt.preOrder();
        string output = redirect.getOutput();
        string expected = "15 10 13 25 20 \n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        fbt.inOrder();
        output = redirect.getOutput();
        expected = "10 13 15 20 25 \n";
        EXPECT_EQ(output, expected);

        redirect.clear();
        fbt.postOrder();
        output = redirect.getOutput();
        expected = "13 10 20 25 15 \n";
        EXPECT_EQ(output, expected);
    }

    EXPECT_FALSE(fbt.insertNode(20));
}

TEST_F(BinaryTreeTest, SearchIsFullTests) {
    fbt.insertNode(27);
    fbt.insertNode(20);
    fbt.insertNode(35);
    fbt.insertNode(14);
    fbt.insertNode(26);
    fbt.insertNode(30);
    fbt.insertNode(44);

    EXPECT_TRUE(fbt.isFullBinary());
    fbt.insertNode(29);
    EXPECT_FALSE(fbt.isFullBinary());

    EXPECT_TRUE(fbt.searchNode(20));
    EXPECT_FALSE(fbt.searchNode(100));
}

TEST_F(BinaryTreeTest, SerializeDeserializeTest) {
    Tree original;
    original.insertNode(50);
    original.insertNode(30);
    original.insertNode(70);
    original.insertNode(20);
    original.insertNode(40);

    EXPECT_TRUE(original.serialize("test_tree.bin"));

    Tree loaded;
    EXPECT_TRUE(loaded.deserialize("test_tree.bin"));

    // Проверяем структуру дерева
    EXPECT_TRUE(loaded.searchNode(50));
    EXPECT_TRUE(loaded.searchNode(30));
    EXPECT_TRUE(loaded.searchNode(70));
    EXPECT_TRUE(loaded.searchNode(20));
    EXPECT_TRUE(loaded.searchNode(40));
    EXPECT_FALSE(loaded.searchNode(100));
}

TEST_F(BinaryTreeTest, SerializeDeserializeEmptyTest) {
    Tree empty;
    EXPECT_TRUE(empty.serialize("empty_tree.bin"));

    Tree loadedEmpty;
    EXPECT_TRUE(loadedEmpty.deserialize("empty_tree.bin"));
    EXPECT_FALSE(loadedEmpty.searchNode(50));
}

TEST_F(BinaryTreeTest, SerializeDeserializeComplexTest) {
    Tree tree;
    tree.insertNode(100);
    tree.insertNode(50);
    tree.insertNode(150);
    tree.insertNode(25);
    tree.insertNode(75);
    tree.insertNode(125);
    tree.insertNode(175);

    EXPECT_TRUE(tree.serialize("complex_tree.bin"));

    Tree loaded;
    EXPECT_TRUE(loaded.deserialize("complex_tree.bin"));

    // Проверяем обходы
    {
        OutputRedirect redirect;
        tree.inOrder();
        string originalOutput = redirect.getOutput();

        redirect.clear();
        loaded.inOrder();
        string loadedOutput = redirect.getOutput();

        EXPECT_EQ(originalOutput, loadedOutput);
    }
}
