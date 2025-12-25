#pragma once

#include <string>

using namespace std;

struct SNode {
    string val;
    SNode* next;
    SNode(const string& x): val(x), next(nullptr) {}
};


class Stack {
    SNode* head;
    // int size;

public:
    Stack(): head(nullptr) {}
    Stack(int s): head(nullptr) {}

    void push(string);
    string pop();
    string top() const;
    void print() const;

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};