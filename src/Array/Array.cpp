#include <iostream>
#include <fstream>
#include "Array.h"

using namespace std;


void Array::doubleCapacity() {
    int newCapacity = capacity * 2;
    if (newCapacity == 0) newCapacity = 1;

    string* newData = new string[newCapacity];
    for (int i = 0; i < size; i++) {
        newData[i] = data[i];
    }
    delete[] data;
    data = newData;
    capacity = newCapacity;
}

int Array::cap() const {
    return capacity;
}

int Array::length() const  {
    return size;
}

bool Array::pushBack(string value) {
    if (size >= capacity) {
        doubleCapacity();
    }
    data[size] = value;
    size++;
    return true;
}

bool Array::pushByIndex(string value, int index) {
    if (index > size) {
        return false;
    }

    for (int i = size; i > index; i--) {
        data[i] = data[i - 1];
    }

    data[index] = value;
    size++;
    if (size >= capacity) {
        doubleCapacity();
    }

    return true;
}


string Array::getByIndex(int index) const {
    if (index >= size || index < 0) {
        return "";
    }
    return data[index];
}

bool Array::deleteByIndex(int index) {
    if (index > size) {
        return false;
    }

    for (int i = index; i < size - 1; i++) {
        data[i] = data[i + 1];
    }
    size--;
    return true;
}

bool Array::swapByIndex(string value, int index) {
    if (index > size) {
        return false;
    }
    data[index] = value;
    return true;
}

void Array::print() const {
    for (int i = 0; i < size; i++) {
        cout << data[i];
        if (i < (size - 1)) {
            cout << " ";
        }
    }
    cout << endl;
}


bool Array::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    file.write(reinterpret_cast<const char*>(&capacity), sizeof(capacity));

    for (int i = 0; i < size; i++) {
        size_t strLength = data[i].length();
        file.write(reinterpret_cast<const char*>(&strLength), sizeof(strLength));

        file.write(data[i].c_str(), strLength);
    }

    file.close();
    return true;
}

bool Array::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    file.read(reinterpret_cast<char*>(&capacity), sizeof(capacity));
    delete[] data;
    data = new string[capacity];

    for (int i = 0; i < size; i++) {
        size_t strLength;
        file.read(reinterpret_cast<char*>(&strLength), sizeof(strLength));

        char* buffer = new char[strLength + 1];
        file.read(buffer, strLength);
        buffer[strLength] = '\0';

        data[i] = string(buffer);
        delete[] buffer;
    }

    file.close();
    return true;
}

