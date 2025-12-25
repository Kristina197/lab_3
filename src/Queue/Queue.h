#pragma once

#include <string>

using namespace std;

struct QNode {
    string data;
    QNode* next;

    QNode(): data(""), next(nullptr) {}
    QNode(string x): data(x), next(nullptr) {}
};

class Queue {
    QNode* front;
    QNode* rear;
    int size;

    void clear();

public:
    Queue(): front(nullptr), rear(nullptr), size() {}
    ~Queue() {
        clear();
    }

    int length() const;
    void enqueue(const string&);
    string dequeue();
    void print() const;

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};