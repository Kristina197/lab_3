#include <fstream>
#include "OpenAddrHash.h"

using namespace std;


void OpenAddrHash::clear() {
    for (int i = 0; i < capacity; ++i) {
        if (buckets[i]) {
            delete buckets[i];
            buckets[i] = nullptr;
        }
    }
    size = 0;
}

size_t OpenAddrHash::hashFunc(const string& key) {
    hash<string> h;
    return h(key) % capacity;
}

bool OpenAddrHash::put(const string& key, const string& val) {
    if (size >= capacity) {
        cout << "Хеш-таблица переполнена" << endl;
        return false;
    }

    size_t hash = hashFunc(key);
    int deletedIndex = -1;
    int i {};

    while (i < capacity) {
        size_t index = (hash + i) % capacity;

        if (!buckets[index]) {
            buckets[index] = new HashEntry();
            buckets[index]->key = key;
            buckets[index]->value = val;
            buckets[index]->status = HashStatus::TAKEN;
            size++;
            return true;
        }

        if (buckets[index]->status == HashStatus::TAKEN &&
            buckets[index]->key == key) {
            buckets[index]->value = val;
            return true;
        }

        if (deletedIndex == -1 && buckets[index]->status == HashStatus::DELETED) {
            deletedIndex = index;
        }
        i++;
    }

    if (deletedIndex != -1) {
        buckets[deletedIndex]->status = HashStatus::TAKEN;
        buckets[deletedIndex]->key = key;
        buckets[deletedIndex]->value = val;
        size++;
        return true;
    }

    return false;
}

HashEntry* OpenAddrHash::get(const string& key) {
    const size_t hash = hashFunc(key);
    int i {};

    while (i < capacity) {
        size_t index = (hash + i) % capacity;

        if (!buckets[index]) return nullptr;
        if (buckets[index]->status == HashStatus::DELETED) {
            i++;
            continue;
        }

        if (buckets[index]->status == HashStatus::TAKEN
            && buckets[index]->key == key) {
            return buckets[index];
        }
        i++;
    }

    return nullptr;
}

void OpenAddrHash::remove(const string& key) {
    size_t hash = hashFunc(key);
    size_t index = hash;
    int i {};

    while (i < capacity) {
        index = (hash + i) % capacity;

        if (buckets[index] &&
            buckets[index]->status == HashStatus::TAKEN &&
            buckets[index]->key == key) {
                buckets[index]->status = HashStatus::DELETED;
                buckets[index]->key = "";
                buckets[index]->value = "";
                size--;
                return;
            }
        i++;
    }
}

void OpenAddrHash::print() {
    for (int i = 0; i < capacity; ++i) {
        cout << "[" << i << "]: ";

        if (!buckets[i]) {
            cout << "null" << endl;
            continue;
        }

        switch (buckets[i]->status) {
            case HashStatus::TAKEN:
                cout << "{" << buckets[i]->key << ": " << buckets[i]->value << "}" << endl;
            break;
            case HashStatus::DELETED:
                cout << "DELETED" << endl;
            break;
            case HashStatus::EMPTY:
                cout << "EMPTY" << endl;
            break;
        }
    }
    cout << endl;
}

bool OpenAddrHash::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    file.write(reinterpret_cast<const char*>(&capacity), sizeof(capacity));
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));

    for (int i = 0; i < capacity; ++i) {
        int statusValue = 0; // EMPTY по умолчанию
        if (buckets[i]) {
            statusValue = static_cast<int>(buckets[i]->status);
        }
        file.write(reinterpret_cast<const char*>(&statusValue), sizeof(statusValue));

        if (buckets[i] && buckets[i]->status == HashStatus::TAKEN) {
            size_t keySize = buckets[i]->key.size();
            file.write(reinterpret_cast<const char*>(&keySize), sizeof(keySize));
            file.write(buckets[i]->key.c_str(), keySize);

            size_t valueSize = buckets[i]->value.size();
            file.write(reinterpret_cast<const char*>(&valueSize), sizeof(valueSize));
            file.write(buckets[i]->value.c_str(), valueSize);
        }
    }

    file.close();
    return true;
}

bool OpenAddrHash::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    // Очистка текущих данных
    clear();

    // Чтение новой capacity
    int newCapacity;
    file.read(reinterpret_cast<char*>(&newCapacity), sizeof(newCapacity));

    // Изменяем capacity и пересоздаем buckets
    capacity = newCapacity;
    buckets.resize(capacity, nullptr);

    file.read(reinterpret_cast<char*>(&size), sizeof(size));

    for (int i = 0; i < capacity; ++i) {
        int statusValue;
        file.read(reinterpret_cast<char*>(&statusValue), sizeof(statusValue));

        if (statusValue != 0) { // Если не EMPTY
            buckets[i] = new HashEntry();
            buckets[i]->status = static_cast<HashStatus>(statusValue);

            if (buckets[i]->status == HashStatus::TAKEN) {
                size_t keySize;
                file.read(reinterpret_cast<char*>(&keySize), sizeof(keySize));

                char* keyBuffer = new char[keySize + 1];
                file.read(keyBuffer, keySize);
                keyBuffer[keySize] = '\0';
                buckets[i]->key = string(keyBuffer);
                delete[] keyBuffer;

                size_t valueSize;
                file.read(reinterpret_cast<char*>(&valueSize), sizeof(valueSize));

                char* valueBuffer = new char[valueSize + 1];
                file.read(valueBuffer, valueSize);
                valueBuffer[valueSize] = '\0';
                buckets[i]->value = string(valueBuffer);
                delete[] valueBuffer;
            }
        }
    }

    file.close();
    return true;
}