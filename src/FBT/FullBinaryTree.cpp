#include <iostream>
#include <fstream>
#include "FullBinaryTree.h"

#include <queue>

using namespace std;

// Поиск элемента в дереве
TNode* Tree::searchElement(TNode* TNode, int val) const {
    if (!TNode) return nullptr;
    if (TNode->val == val) return TNode;

    auto leftSearch = searchElement(TNode->left, val);
    if (leftSearch) return leftSearch;

    return searchElement(TNode->right, val);
}

// Проверка на Full Binary Tree (вспомогательная функция)
bool Tree::FullBinaryCheck(TNode* node) const {
    if (!node) return true;
    if (!node->left && !node->right) return true;

    if (node->left && node->right) {
        return FullBinaryCheck(node->left) && FullBinaryCheck(node->right);
    }
    return false;
}

// Обход дерева Preorder
void Tree::PreorderTraverse(TNode* TNode) const {
    if (!TNode) return;
    cout << TNode->val << " ";
    PreorderTraverse(TNode->left);
    PreorderTraverse(TNode->right);
}

// Обход дерева Inorder
void Tree::InorderTraverse(TNode* TNode) const {
    if (!TNode) return;
    InorderTraverse(TNode->left);
    cout << TNode->val << " ";
    InorderTraverse(TNode->right);
}

// Обход дерева Postorder
void Tree::PostorderTraverse(TNode* TNode) const {
    if (!TNode) return;
    PostorderTraverse(TNode->left);
    PostorderTraverse(TNode->right);
    cout << TNode->val << " ";
}




// Добавление узла в дерево
bool Tree::insertNode(int val) {
    TNode* newNode = new TNode(val);

    if (!root) {
        root = newNode;
        return true;
    }

    TNode* curr = root;
    while (curr) {
        if (val > curr->val) {
            if (curr->right) {
                curr = curr->right;
            } else {
                curr->right = newNode;
                return true;
            }
        } else if (val < curr->val) {
            if (curr->left) {
                curr = curr->left;
            } else {
                curr->left = newNode;
                return true;
            }
        } else {
            delete newNode;
            return false;
        }
    }
}

// Поиска узла в дереве
bool Tree::searchNode(int val) const {
    auto foundNode = searchElement(root, val);
    if (foundNode) {
        return true;
    }
    return false;
}

// Проверка на Full Binary Tree
bool Tree::isFullBinary() const {
    if (FullBinaryCheck(root)) {
        return true;
    }
    return false;
}


void Tree::preOrder() const {
    PreorderTraverse(root);
    cout << endl;
}

void Tree::inOrder() const {
    InorderTraverse(root);
    cout << endl;
}

void Tree::postOrder() const{
    PostorderTraverse(root);
    cout << endl;
}

void Tree::clearTree(TNode* node) {
    if (!node) return;
    clearTree(node->left);
    clearTree(node->right);
    delete node;
}


void Tree::serializeHelper(ofstream& file, TNode* node) const {
    if (!node) {
        int marker = -1;
        file.write(reinterpret_cast<const char*>(&marker), sizeof(marker));
        return;
    }

    file.write(reinterpret_cast<const char*>(&node->val), sizeof(node->val));
    serializeHelper(file, node->left);
    serializeHelper(file, node->right);
}

TNode* Tree::deserializeHelper(ifstream& file) {
    int val;
    file.read(reinterpret_cast<char*>(&val), sizeof(val));

    if (val == -1) {
        return nullptr;
    }

    TNode* node = new TNode(val);
    node->left = deserializeHelper(file);
    node->right = deserializeHelper(file);
    return node;
}

bool Tree::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    serializeHelper(file, root);
    file.close();
    return true;
}

bool Tree::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file.is_open()) {
        return false;
    }

    // Проверяем, не пустой ли файл
    if (file.peek() == ifstream::traits_type::eof()) {
        file.close();
        return false;
    }

    clearTree(root);
    root = deserializeHelper(file);

    // Проверяем, что чтение прошло успешно
    if (file.fail()) {
        clearTree(root);
        root = nullptr;
        file.close();
        return false;
    }

    file.close();
    return true;
}