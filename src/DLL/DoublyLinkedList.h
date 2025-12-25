#pragma once

#include <string>
#include <vector>

using namespace std;

struct LNode {
    string val;
    LNode* prev, *next;
    LNode(string x) : val(x), prev(nullptr), next(nullptr) {}
};

class DoublyLinkedList {
    LNode* head;
    LNode* tail;

    LNode* searchNode(string) const;

public:
    DoublyLinkedList(): head(nullptr), tail(nullptr) {}
    ~DoublyLinkedList() = default;

    void insertAtHead(string);
    void insertAtTail(string);
    void insertBefore(string, string);
    void insertAfter(string, string);
    void deleteHead();
    void deleteTail();
    void deleteBefore(string);
    void deleteAfter(string);
    void deleteByValue(string);

    int searchByValue(string) const;
    void print() const;
    void printReverse() const;

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};