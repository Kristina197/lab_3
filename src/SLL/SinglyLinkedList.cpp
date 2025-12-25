#include <iostream>
#include <fstream>
#include "SinglyLinkedList.h"

using namespace std;


// Добавление в начало
void SinglyLinkedList::insertAtHead(string val) {
    FNode* newNode = new FNode(val);
    newNode->next = head;
    head = newNode;
}

// Добавление в конец
void SinglyLinkedList::insertAtTail(string val) {
    FNode* newNode = new FNode(val);

    // Список пустой - новый узел становится головой
    if (!head) {
        head = newNode;
        return;
    }

    FNode* tmp = head;
    while (tmp->next) {
        tmp = tmp->next;
    }
    tmp->next = newNode;
}

// Добавление до элемента
void SinglyLinkedList::insertBefore(string target, string val) {
    if (!head) {
        cout << "Список пустой, вставка невозможна" << endl;
        return ;
    }
    if (head->val == target) {
        insertAtHead(val);
        return ;
    }

    FNode* curr = head;
    while (curr->next && curr->next->val != target)
        curr = curr->next;

    if (curr->next && curr->next->val == target) {
        FNode* newNode = new FNode(val);
        newNode->next = curr->next;
        curr->next = newNode;
        return ;
    }

    cout << "Заданное значение " << target << " не найдено" << endl;
}

// Добавление после элемента
void SinglyLinkedList::insertAfter(string target, string val) {
    if (!head) {
        cout << "Список пустой, вставка невозможна" << endl;
        return ;
    }

    FNode* curr = head;
    while (curr && curr->val != target) {
        curr = curr->next;
    }

    if (curr && curr->val == target) {
        FNode* newNode = new FNode(val);
        newNode->next = curr->next;
        curr->next = newNode;
        return ;
    }

    cout << "Заданное значение " << target << " не найдено" << endl;
}

// Удаление головы
void SinglyLinkedList::deleteHead() {
    if (!head) {
        cout << "Список пуст" << endl;
        return ;
    }
    FNode* tmp = head;
    head = head->next;
    delete tmp;
}

// Удаление хвоста
void SinglyLinkedList::deleteTail() {
    if (head == nullptr) {
        cout << "Список пуст" << endl;
        return ;
    }

    if (head->next == nullptr) {
        delete head;
        head = nullptr;
        return ;
    }

    FNode* curr = head;
    while (curr->next && curr->next->next) {
        curr = curr->next;
    }
    FNode* tail = curr->next;
    curr->next = nullptr;
    delete tail;
}

// Удаление до элемента
void SinglyLinkedList::deleteBefore(string target) {
    if (!head || !head->next) {
        cout << "Недостаточно элементов для удаления" << endl;
        return ;
    }
    if (head->next->val == target) deleteHead();

    FNode* curr = head;
    while (curr->next && curr->next->next && curr->next->next->val != target) {
        curr = curr->next;
    }

    if (curr->next && curr->next->next && curr->next->next->val == target) {
        FNode* deleteNode = curr->next;
        curr->next = curr->next->next;
        delete deleteNode;
        return ;
    }

    cout << "Заданное значение " << target << " не найдено" << endl;
}

// Удаление после элемента
void SinglyLinkedList::deleteAfter(string target) {
    if (!head) {
        cout << "Список пустой, удаление невозможно" << endl;
        return ;
    }

    FNode* curr = head;
    while (curr && curr->val != target) {
        curr = curr->next;
    }

    if (curr && curr->next) {
        FNode* deleteNode = curr->next;
        curr->next = deleteNode->next;
        delete deleteNode;
        return ;
    }
    if (curr && !curr->next) {
        cout << "Нет элемента после " << target << endl;
        return ;
    }

    cout << "Заданное значение " << target << " не найдено" << endl;
}

// Удаление элемента по значению
void SinglyLinkedList::deleteByValue(string target) {
    if (!head) return ;

    if (head->val == target) {
        deleteHead();
        return ;
    }

    FNode* curr = head;
    while (curr->next && curr->next->val != target)
        curr = curr->next;

    if (curr->next && curr->next->val == target) {
        FNode* deleteNode = curr->next;
        curr->next = deleteNode->next;
        delete deleteNode;
        return ;
    }
    cout << "Заданное значение " << target << " отсутствует в списке" << endl;
}

// Печать списка
void SinglyLinkedList::print() const {
    FNode* tmp = head;
    while (tmp != nullptr) {
        cout << tmp->val << " ";
        tmp = tmp->next;
    }
    cout << endl;
}

// Поиск элемента по значению
int SinglyLinkedList::searchByValue(string target) const {
    if (!head) {
        return -1;
    }

    FNode* curr = head;
    int targetIndex {};

    while (curr) {
        if (curr->val == target)
            return targetIndex;
        curr = curr->next;
        targetIndex++;
    }
    return -1;
}


bool SinglyLinkedList::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    int count = 0;
    FNode* current = head;
    while (current) {
        count++;
        current = current->next;
    }

    file.write(reinterpret_cast<const char*>(&count), sizeof(count));

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

bool SinglyLinkedList::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    while (head) {
        deleteHead();
    }

    int count;
    file.read(reinterpret_cast<char*>(&count), sizeof(count));

    for (int i = 0; i < count; i++) {
        size_t strLength;
        file.read(reinterpret_cast<char*>(&strLength), sizeof(strLength));

        char* buffer = new char[strLength + 1];
        file.read(buffer, strLength);
        buffer[strLength] = '\0';

        insertAtTail(string(buffer));
        delete[] buffer;
    }

    file.close();
    return true;
}