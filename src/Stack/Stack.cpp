#include <iostream>
#include <fstream>
#include "Stack.h"

#include <vector>

using namespace std;


// Добавление элемента в стек
void Stack::push(string val) {
    SNode* newNode = new SNode(val);
    newNode->next = head;
    head = newNode;
    // size++;
}

// Удаление элемента из стека
string Stack::pop() {
    if (!head) {
        return "";
    }

    SNode* deleteNode = head;
    auto poppedElem = deleteNode->val;

    head = head->next;
    delete deleteNode;
    // size--;

    return poppedElem;
}

// Получение верхнего элемента без удаления
string Stack::top() const {
    if (head == nullptr) return "";
    return head->val;
}

// Печать стека
void Stack::print() const {
    SNode* curr = head;

    cout << "top -> ";
    while (curr) {
        cout << curr->val;
        if (curr->next) cout << " -> ";
        curr = curr->next;
    }
    cout << " -> bottom" << endl;
}


bool Stack::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    // Считаем реальный размер
    int actualSize = 0;
    SNode* current = head;
    while (current) {
        actualSize++;
        current = current->next;
    }

    file.write(reinterpret_cast<const char*>(&actualSize), sizeof(actualSize));

    // Сохраняем элементы в правильном порядке (сверху вниз)
    current = head;
    while (current) {
        size_t strLength = current->val.length();
        file.write(reinterpret_cast<const char*>(&strLength), sizeof(strLength));
        file.write(current->val.c_str(), strLength);
        current = current->next;
    }

    file.close();
    return true;
}

bool Stack::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    // Очищаем текущий стек
    while (head) {
        pop();
    }

    int stackSize;
    file.read(reinterpret_cast<char*>(&stackSize), sizeof(stackSize));

    // Временный вектор для хранения элементов в правильном порядке
    vector<string> elements;

    for (int i = 0; i < stackSize; i++) {
        size_t strLength;
        file.read(reinterpret_cast<char*>(&strLength), sizeof(strLength));

        char* buffer = new char[strLength + 1];
        file.read(buffer, strLength);
        buffer[strLength] = '\0';

        elements.push_back(string(buffer));
        delete[] buffer;
    }

    // Добавляем элементы в стек в обратном порядке (чтобы сохранить LIFO)
    for (auto it = elements.rbegin(); it != elements.rend(); ++it) {
        push(*it);
    }

    file.close();
    return true;
}
