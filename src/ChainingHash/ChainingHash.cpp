#include <fstream>
#include "ChainingHash.h"

using namespace std;

size_t ChainingHash::hashFunc(const string& key) {
    hash<string> h;
    return h(key) % capacity;
}

bool ChainingHash::put(const string& key, const string& val) {
    size_t hash = hashFunc(key);
    Node* curr = buckets[hash];

    // Проверка на наличие значения
    while (curr) {
        if (curr->key == key) {
            curr->value = val;
            return true;
        }
        curr = curr->next;
    }

    Node* newNode = new Node(key, val);
    newNode->next = buckets[hash];
    buckets[hash] = newNode;
    size++;

    return true;
}

string* ChainingHash::get(const string& key) {
    size_t hash = hashFunc(key);
    Node* curr = buckets[hash];

    while (curr) {
        if (curr->key == key)
            return &(curr->value);
        curr = curr->next;
    }
    return nullptr;
}

void ChainingHash::remove(const string& key) {
    size_t hash = hashFunc(key);
    Node* curr = buckets[hash];
    Node* prev = nullptr;

    while (curr) {
        if (curr->key == key) {
            if (!prev) {
                buckets[hash] = curr->next;
            } else {
                prev->next = curr->next;
            }
            delete curr;
            size--;
            return;
        }
        prev = curr;
        curr = curr->next;
    }
}

void ChainingHash::print() {
    for (int i = 0; i < capacity; ++i) {
        cout << "[" << i << "]: ";
        Node* curr = buckets[i];
        while (curr) {
            cout << "{" << curr->key << ": " << curr->value << "} -> ";
            curr = curr->next;
        }
        cout << "null" << endl;
    }
}

bool ChainingHash::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    file.write(reinterpret_cast<const char*>(&capacity), sizeof(capacity));
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));

    for (int i = 0; i < capacity; ++i) {
        Node* curr = buckets[i];
        int chainSize = 0;
        Node* temp = curr;
        while (temp) {
            chainSize++;
            temp = temp->next;
        }

        file.write(reinterpret_cast<const char*>(&chainSize), sizeof(chainSize));

        while (curr) {
            size_t keySize = curr->key.size();
            file.write(reinterpret_cast<const char*>(&keySize), sizeof(keySize));
            file.write(curr->key.c_str(), keySize);

            size_t valueSize = curr->value.size();
            file.write(reinterpret_cast<const char*>(&valueSize), sizeof(valueSize));
            file.write(curr->value.c_str(), valueSize);

            curr = curr->next;
        }
    }

    file.close();
    return true;
}

bool ChainingHash::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    // Очистка текущих данных
    for (int i = 0; i < capacity; ++i) {
        Node* curr = buckets[i];
        while (curr) {
            Node* deleteNode = curr;
            curr = curr->next;
            delete deleteNode;
        }
        buckets[i] = nullptr;
    }
    size = 0;

    // Чтение новой capacity и изменение размера вектора
    file.read(reinterpret_cast<char*>(&capacity), sizeof(capacity));
    file.read(reinterpret_cast<char*>(&size), sizeof(size));

    buckets.resize(capacity, nullptr);

    for (int i = 0; i < capacity; ++i) {
        int chainSize;
        file.read(reinterpret_cast<char*>(&chainSize), sizeof(chainSize));

        Node** curr = &buckets[i];
        for (int j = 0; j < chainSize; ++j) {
            size_t keySize;
            file.read(reinterpret_cast<char*>(&keySize), sizeof(keySize));

            char* keyBuffer = new char[keySize + 1];
            file.read(keyBuffer, keySize);
            keyBuffer[keySize] = '\0';
            string key = string(keyBuffer);
            delete[] keyBuffer;

            size_t valueSize;
            file.read(reinterpret_cast<char*>(&valueSize), sizeof(valueSize));

            char* valueBuffer = new char[valueSize + 1];
            file.read(valueBuffer, valueSize);
            valueBuffer[valueSize] = '\0';
            string value = string(valueBuffer);
            delete[] valueBuffer;

            *curr = new Node(key, value);
            curr = &((*curr)->next);
        }
        *curr = nullptr;
    }

    file.close();
    return true;
}