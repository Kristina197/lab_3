#pragma once

#include <string>
#include <vector>

using namespace std;

// Реализация узла
struct FNode {
    string val;
    FNode* next;
    FNode(string x) : val(x), next(nullptr) {}
};

class SinglyLinkedList {
    FNode* head;

public:
    SinglyLinkedList(): head(nullptr) {}
    ~SinglyLinkedList() = default;

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

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};

