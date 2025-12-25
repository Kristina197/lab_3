#pragma once

#include <iostream>
#include <vector>
#include <string>

using namespace std;

// Класс перечислений статуса ячейки
enum class HashStatus {
    EMPTY, DELETED, TAKEN
};

// Ячейка хеш-таблицы
struct HashEntry {
    string key;
    string value;
    HashStatus status;

    HashEntry(): status(HashStatus::EMPTY) {}
};

class OpenAddrHash {
    vector<HashEntry*> buckets;
    int size;
    int capacity;

    size_t hashFunc(const string& key);
    void clear();

public:
    OpenAddrHash(): capacity(10), size(0) {
        buckets.resize(capacity, nullptr);
    }
    OpenAddrHash(int cap): capacity(cap), size(0) {
        buckets.resize(capacity, nullptr);
    }
    OpenAddrHash(const OpenAddrHash&) = delete;
    OpenAddrHash& operator=(const OpenAddrHash&) = delete;
    ~OpenAddrHash() {
        clear();
    }

    bool put(const string& key, const string& val);
    HashEntry* get(const string& key);
    void remove(const string& key);
    void print();

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};