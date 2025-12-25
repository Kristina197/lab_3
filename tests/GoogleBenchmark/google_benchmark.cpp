#include <benchmark/benchmark.h>
#include "../../src/Array/Array.h"
#include "../../src/Stack/Stack.h"
#include "../../src/Queue/Queue.h"
#include "../../src/SLL/SinglyLinkedList.h"
#include "../../src/DLL/DoublyLinkedList.h"
#include "../../src/ChainingHash/ChainingHash.h"
#include "../../src/OpenAddrHash/OpenAddrHash.h"
#include "../../src/FBT/FullBinaryTree.h"


#include <string>

// Бенчмарк для  массива - добавление в конец
static void BM_Array_PushBack(benchmark::State& state) {
    for (auto _ : state) {
        Array arr;
        for (int i = 0; i < state.range(0); i++) {
            arr.pushBack("elem_" + to_string(i));
        }
    }
}
BENCHMARK(BM_Array_PushBack)->Range(1000, 100000);

// Бенчмарк для двусвязного списка - добавление в конец
static void BM_DoublyLinkedList_InsertAtTail(benchmark::State& state) {
    for (auto _ : state) {
        DoublyLinkedList list;
        for (int i = 0; i < state.range(0); i++) {
            list.insertAtTail("elem_" + to_string(i));
        }
    }
}
BENCHMARK(BM_DoublyLinkedList_InsertAtTail)->Range(1000, 100000);

// Бенчмарк для односвязного списка - добавление в конец
static void BM_SinglyLinkedList_InsertAtTail(benchmark::State& state) {
    for (auto _ : state) {
        SinglyLinkedList list;
        for (int i = 0; i < state.range(0); i++) {
            list.insertAtTail("elem_" + to_string(i));
        }
    }
}
BENCHMARK(BM_SinglyLinkedList_InsertAtTail)->Range(1000, 100000);

// Бенчмарк для хеш-таблицы с цепочками - добавление элементов
static void BM_ChainingHash_Put(benchmark::State& state) {
    for (auto _ : state) {
        ChainingHash hashTable(state.range(0) * 2);
        for (int i = 0; i < state.range(0); i++) {
            string key = "key_" + to_string(i);
            string value = "value_" + to_string(i);
            hashTable.put(key, value);
        }
    }
}
BENCHMARK(BM_ChainingHash_Put)->Range(1000, 100000);

// Бенчмарк для хеш-таблицы с открытой адресацией - добавление элементов
static void BM_OpenAddrHash_Put(benchmark::State& state) {
    for (auto _ : state) {
        OpenAddrHash hashTable(state.range(0) * 2);
        for (int i = 0; i < state.range(0); i++) {
            string key = "key_" + to_string(i);
            string value = "value_" + to_string(i);
            hashTable.put(key, value);
        }
    }
}
BENCHMARK(BM_OpenAddrHash_Put)->Range(1000, 100000);

// Бенчмарк для стека - добавление элементов
static void BM_Stack_Push(benchmark::State& state) {
    for (auto _ : state) {
        Stack stack;
        for (int i = 0; i < state.range(0); i++) {
            stack.push("elem_" + to_string(i));
        }
    }
}
BENCHMARK(BM_Stack_Push)->Range(1000, 100000);

// Бенчмарк для очереди - добавление элементов (enqueue)
static void BM_Queue_Enqueue(benchmark::State& state) {
    for (auto _ : state) {
        Queue queue;
        for (int i = 0; i < state.range(0); i++) {
            queue.enqueue("elem_" + to_string(i));
        }
    }
}
BENCHMARK(BM_Queue_Enqueue)->Range(1000, 100000);

// Бенчмарк для бинарного дерева - добавление узлов
static void BM_Tree_InsertNode(benchmark::State& state) {
    for (auto _ : state) {
        Tree tree;
        for (int i = 0; i < state.range(0); i++) {
            int val = rand() % 1000000;
            tree.insertNode(val);
        }
    }
}
BENCHMARK(BM_Tree_InsertNode)->Range(1000, 100000);


BENCHMARK_MAIN();