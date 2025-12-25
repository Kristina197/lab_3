#pragma once

#include <string>
#include <vector>

using namespace std;

struct TNode {
    int val;
    TNode* left, *right;
    TNode(int x): val(x), left(nullptr), right(nullptr) {}
    TNode(int x, TNode* leftptr, TNode* rightptr): val(x), left(leftptr), right(rightptr) {}
    ~TNode() = default;
};


class Tree {
    TNode* root;

    TNode* searchElement(TNode*, int) const;
    bool FullBinaryCheck(TNode*) const;
    void PreorderTraverse(TNode*) const;
    void InorderTraverse(TNode*) const;
    void PostorderTraverse(TNode*) const;

    void clearTree(TNode* node);
    void serializeHelper(ofstream& file, TNode* node) const;
    TNode* deserializeHelper(ifstream& file);

public:
    Tree(): root(nullptr) {}
    ~Tree() {
        clearTree(root);
    }

    bool insertNode(int);
    bool searchNode(int) const;
    bool isFullBinary() const;
    void preOrder() const;
    void inOrder() const;
    void postOrder() const;

    bool serialize(const string& filename) const;
    bool deserialize(const string& filename);
};

