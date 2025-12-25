#include <iostream>
#include <fstream>
#include "Queue.h"

using namespace std;


// Получение размера очереди
int Queue::length() const {
    return size;
}

// Добавление элемента в очередь
void Queue::enqueue(const string& value) {
    QNode* newNode = new QNode(value);

    if (!rear) {
        front = newNode;
        rear = newNode;
    } else {
        rear->next = newNode;
        rear = newNode;
    }
    size++;
}

// Удаление элемента из очереди
string Queue::dequeue() {
    if (!front) return "";

    auto dequeuedElem = front->data;
    QNode* tmp = front;
    front = front->next;

    if (!front) {
        rear = nullptr;
    }

    delete tmp;
    size--;

    return dequeuedElem;
}

// Печать очереди
void Queue::print() const {
    if (!front) {
        cout << "Очередь пустая" << endl;
        return;
    }

    QNode* curr = front;
    while (curr) {
        cout << curr->data << " ";
        curr = curr->next;
    }
    cout << endl;
}

// Очистка очереди
void Queue::clear() {
    while (front) {
        dequeue();
    }
}


bool Queue::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    file.write(reinterpret_cast<const char*>(&size), sizeof(size));

    QNode* current = front;
    while (current) {
        size_t strLength = current->data.length();
        file.write(reinterpret_cast<const char*>(&strLength), sizeof(strLength));
        file.write(current->data.c_str(), strLength);

        current = current->next;
    }

    file.close();
    return true;
}



bool Queue::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }
    clear();

    int queueSize;
    file.read(reinterpret_cast<char*>(&queueSize), sizeof(queueSize));

    for (int i = 0; i < queueSize; i++) {
        size_t strLength;
        file.read(reinterpret_cast<char*>(&strLength), sizeof(strLength));

        char* buffer = new char[strLength + 1];
        file.read(buffer, strLength);
        buffer[strLength] = '\0';

        enqueue(string(buffer));
        delete[] buffer;
    }

    file.close();
    return true;
}