#include <iostream>
#include <fstream>
#include "DoublyLinkedList.h"

using namespace std;

// Поиск узла по значению (вспомогательная функция)
LNode* DoublyLinkedList::searchNode(string val) const {
    if (!head) {
        cout << "Список пустой" << endl;
        return nullptr;
    }

    LNode* tmp = head;
    while (tmp != nullptr) {
        if (tmp->val == val)
            return tmp;
        tmp = tmp->next;
    }
    return nullptr;
}


// Добавление в начало
void DoublyLinkedList::insertAtHead(string val) {
    LNode* newNode = new LNode(val);

    if (head == nullptr) {
        head = newNode;
        tail = newNode;
        return;
    }

    newNode->next = head;
    head->prev = newNode;
    head = newNode;
}

// Добавление в конец
void DoublyLinkedList::insertAtTail(string val) {
    LNode* newNode = new LNode(val);

    if (head == nullptr) {
        head = newNode;
        tail = newNode;
        return ;
    }

    tail->next = newNode;
    newNode->prev = tail;
    tail = newNode;
}

// Добавление до элемента
void DoublyLinkedList::insertBefore(string target, string val) {
    LNode* targetNode = searchNode(target);
    if (!targetNode) {
        cout << "Заданное значение " << target << " отсутствует в списке" << endl;
        return ;
    }

    if (targetNode == head) {
        insertAtHead(val);
        return ;
    }

    LNode* newNode = new LNode(val);
    newNode->prev = targetNode->prev;
    newNode->next = targetNode;
    targetNode->prev->next = newNode;
    targetNode->prev = newNode;
}

// Добавление после элемента
void DoublyLinkedList::insertAfter(string target, string val) {
    LNode* targetNode = searchNode(target);
    if (!targetNode) {
        cout << "Заданное значение " << target << " отсутствует в списке" << endl;
        return ;
    }

    LNode* newNode = new LNode(val);
    newNode->prev = targetNode;
    newNode->next = targetNode->next;

    // Проверка на позицию найденного узла
    if (targetNode->next) {
        targetNode->next->prev = newNode;
    } else {
        tail = newNode;
    }
    targetNode->next = newNode;
}

// Удаление головы
void DoublyLinkedList::deleteHead( ) {
    if (!head) {
        cout << "Список пуст" << endl;
        return ;
    }

    LNode* tmp = head;
    head = head->next;

    if (head) {
        head->prev = nullptr;
    } else {
        tail = nullptr;    
    }

    delete tmp;
}

// Удаление хвоста
void DoublyLinkedList::deleteTail() {
    if (!head) {
        cout << "Список пуст" << endl;
        return ;
    }

    LNode* tmp = tail;
    if (head == tail) {
        head = nullptr;
        tail = nullptr;
    } else {
        tail = tail->prev;
        tail->next = nullptr;
    }

    delete tmp;
}

// Удаление до элемента
void DoublyLinkedList::deleteBefore(string target) {
    LNode* targetNode = searchNode(target);
    if (!targetNode) {
        // cerr << "Заданное значение " << target << " отсутствует в списке" << endl;
        return ;
    }
    if (!targetNode->prev) {
        // cerr << "Перед заданным значением " << target << " отсутствуют узлы. Удаление невозможно" << endl;
        return ;
    }
    
    LNode* deleteNode = targetNode->prev;
    if (deleteNode == head) {
        return deleteHead();
    }

    targetNode->prev = deleteNode->prev;
    deleteNode->prev->next = targetNode;
    
    delete deleteNode;
}

// Удаление после элемента
void DoublyLinkedList::deleteAfter(string target) {
    LNode* targetNode = searchNode(target);
    if (!targetNode) {
        // cerr << "Заданное значение " << target << " отсутствует в списке" << endl;
        return ;
    }
    if (!targetNode->next) {
        // cerr << "После заданного значения " << target << " отсутствуют узлы. Удаление невозможно" << endl;
        return ;
    }
    
    LNode* deleteNode = targetNode->next;
    if (deleteNode == tail) {
        return deleteTail();
    }

    targetNode->next = deleteNode->next;
    deleteNode->next->prev = targetNode;
    
    delete deleteNode;
}

// Удаление по значению
void DoublyLinkedList::deleteByValue(string target) {
    if (head == nullptr) {
        cout << "Список пустой, удалять нечего" << endl;
        return ;
    }

    LNode* targetNode = searchNode(target);
    if (!targetNode) {
        cout << "Заданное значение " << target << " отсутствует в списке" << endl;
        return ;
    }

    if (targetNode == head) {
        return deleteHead();
    }
    if (targetNode == tail) {
        return deleteTail();
    }

    targetNode->prev->next = targetNode->next;
    targetNode->next->prev = targetNode->prev;

    delete targetNode;
}

// Печать списка
void DoublyLinkedList::print() const {
    LNode* tmp = head;
    if (tmp == nullptr) {
        cout << "Список пуст" << endl;
        return ;
    }
    while (tmp) {
        cout << tmp->val << " ";
        tmp = tmp->next;
    }
    cout << endl;
}

// Печать списка наоборот
void DoublyLinkedList::printReverse() const {
    if (!head) {
        cout << "Список пуст" << endl;
        return;
    }

    LNode* tmp = tail;

    while (tmp != nullptr) {
        cout << tmp->val << " ";
        tmp = tmp->prev;
    }
    cout << endl;
}

// Поиск элемента по значению
int DoublyLinkedList::searchByValue(string target) const {
    LNode* curr = head;
    int searchIndex {};
    while (curr) {
        if (curr->val == target)
            return searchIndex;
        curr = curr->next;
        searchIndex++;
    }

    return -1;
}




bool DoublyLinkedList::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        cerr << "Не удалось открыть файл для записи: " << filename << endl;
        return false;
    }

    int count = 0;
    LNode* current = head;
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

bool DoublyLinkedList::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        cerr << "Не удалось открыть файл для чтения: " << filename << endl;
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