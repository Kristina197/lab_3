#include <gtest/gtest.h>

#include "GoogleTests/GoogleTest_Array.cpp"
#include "GoogleTests/GoogleTest_Stack.cpp"
#include "GoogleTests/GoogleTest_Queue.cpp"
#include "GoogleTests/GoogleTest_SinglyLinkedList.cpp"
#include "GoogleTests/GoogleTest_DoublyLinkedList.cpp"
#include "GoogleTests/GoogleTest_FullBinaryTree.cpp"
#include "GoogleTests/GoogleTest_ChainingHash.cpp"
#include "GoogleTests/GoogleTest_OpenAddrHash.cpp"

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}