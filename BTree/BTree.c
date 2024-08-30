#include "BTree.h"
#include <stdio.h>
#include <stdlib.h>

// Function to create a B-tree node
BNode *createBNode(const int isLeaf) {
    BNode *node = malloc(sizeof(BNode));
    if (node == NULL) {
        fprintf(stderr, "Failed to allocate memory for node\n");
        return NULL;
    }
    node->numKeys = 0;
    node->isLeaf = isLeaf;
    for (int i = 0; i < MAX_KEYS; i++) {
        node->keys[i] = 0;
    }
    for (int i = 0; i < MAX_KEYS + 1; i++) {
        node->children[i] = NULL;
    }
    return node;
}

// Function to insert a key into the B-tree
BNode *insertBNode(BNode **root, const int key) {
    if (*root == NULL) {
        *root = createBNode(1);
        (*root)->keys[0] = key;
        (*root)->numKeys = 1;
        return *root;
    }

    // If there is space in the current node
    if ((*root)->numKeys < MAX_KEYS) {
        // Insert key into the current node
        // Loop through the keys to find the correct position
        int i = (*root)->numKeys - 1;
        while (i >= 0 && key < (*root)->keys[i]) {
            (*root)->keys[i + 1] = (*root)->keys[i];
            i--;
        }
        (*root)->keys[i + 1] = key;
        (*root)->numKeys++;
        return *root;
    }

    // If the current node is a leaf
    // Split the node
    BNode *newNode = createBNode(1);
    int tempKeys[MAX_KEYS + 1];
    int i = 0;
    while (i < MAX_KEYS && key > (*root)->keys[i]) {
        tempKeys[i] = (*root)->keys[i];
        i++;
    }
    tempKeys[i] = key;
    while (i < MAX_KEYS) {
        tempKeys[i + 1] = (*root)->keys[i];
        i++;
    }

    // Update the current node
    (*root)->numKeys = (MAX_KEYS + 1) / 2;
    for (i = 0; i < (*root)->numKeys; i++) {
        (*root)->keys[i] = tempKeys[i];
    }

    // Update the new node
    newNode->numKeys = MAX_KEYS / 2;
    for (i = 0; i < newNode->numKeys; i++) {
        newNode->keys[i] = tempKeys[i + (*root)->numKeys];
    }

    return newNode;
}

// Function to search for a key in the B-tree
BNode *searchBNode(BNode *root, const int key) {
    if (root == NULL) {
        return NULL;
    }

    int i = 0;
    while (i < root->numKeys && key > root->keys[i]) {
        i++;
    }

    if (i < root->numKeys && key == root->keys[i]) {
        return root;
    }

    if (root->isLeaf) {
        return NULL;
    }

    return searchBNode(root->children[i], key);
}


// Function to free the memory allocated for the B-tree
void freeBNode(BNode *root) {
    if (root == NULL) {
        return;
    }

    for (int i = 0; i < root->numKeys + 1; i++) {
        freeBNode(root->children[i]);
    }
    free(root);
}