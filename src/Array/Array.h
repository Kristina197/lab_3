#pragma once

#include <string>

using namespace std;

class Array {
    string* data;
    int size;
    int capacity;

    void doubleCapacity();

public:
    Array(): size(0), capacity(10) {
        data = new string[capacity];
    }
    Array(int cap): size(0), capacity(cap) {
        data = new string[capacity];
    }

    int cap() const;
    int length() const;
    bool pushBack(string);
    bool pushByIndex(string, int);
    string getByIndex(int) const;
    bool deleteByIndex(int);
    bool swapByIndex(string, int);
    void print() const;

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};