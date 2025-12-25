#pragma once

#include <iostream>
#include <algorithm>
#include <vector>
#include <string>

using namespace std;

struct Node {
    string key;
    string value;
    Node* next;

    Node(string k, string v): key(move(k)), value(move(v)), next(nullptr) {}
};

class ChainingHash {
    vector<Node*> buckets;
    int size;
    int capacity;

    size_t hashFunc(const string& key);

public:
    ChainingHash(): capacity(10), size(0) {
        buckets.resize(capacity, nullptr);
    }
    ChainingHash(int cap): capacity(cap), size(0) {
        buckets.resize(capacity, nullptr);
    }
    ~ChainingHash() {
        for (int i = 0; i < capacity; ++i) {
            Node* curr = buckets[i];
            while (curr) {
                Node* deleteNode = curr;
                curr = curr->next;
                delete deleteNode;
            }
        }
    }

    bool put(const string& key, const string& val);
    string* get(const string& key);
    void remove(const string& key);
    void print();

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};