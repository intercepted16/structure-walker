#include <stdio.h>
#include <stdlib.h>
#include "MaxHeapTree.h"

// Function to create a new node in the max-heap tree
MaxHeapNode* createMaxHeapNode(const int val) {
    MaxHeapNode *node = malloc(sizeof(MaxHeapNode));
    if (node == NULL) {
        fprintf(stderr, "Failed to allocate memory for node\n");
        return NULL;
    }
    node->val = val;
    node->left = NULL;
    node->right = NULL;
    return node;
}

// Function to perform a "bubble-up" operation to maintain the max-heap property
void bubbleUp(MaxHeapNode* node) {
    while (node->parent != NULL && node->val > node->parent->val) {
        // Swap values with parent
        int temp = node->val;
        node->val = node->parent->val;
        node->parent->val = temp;

        // Move up to parent
        node = node->parent;
    }
}

// Function to insert a value into the max-heap BST
void insertMaxHeapNode(MaxHeapNode **root, const int val) {
    if (*root == NULL) {
        *root = createMaxHeapNode(val);
        printf("Inserted %d as the root node.\n", val);
        return;
    }

    // Perform level-order insertion to maintain the complete tree property
    MaxHeapNode* queue[100]; // Fixed-size queue for simplicity
    int front = 0, rear = 0;

    // Enqueue the root node
    queue[rear++] = *root;

    // Find the first empty spot
    while (front < rear) {
        MaxHeapNode* current = queue[front++];

        // Insert in the left child
        if (current->left == NULL) {
            current->left = createMaxHeapNode(val);
            printf("Inserted %d as the left child of %d.\n", val, current->val);
            // Bubble up to maintain max-heap property
            bubbleUp(current->left);
            return;
        } else {
            // Enqueue the left child
            queue[rear++] = current->left;
        }

        // Insert in the right child
        if (current->right == NULL) {
            current->right = createMaxHeapNode(val);
            printf("Inserted %d as the right child of %d.\n", val, current->val);
            // Bubble up to maintain max-heap property
            bubbleUp(current->right);
            return;
        } else {
            // Enqueue the right child
            queue[rear++] = current->right;
        }
    }
}

// Function to search for a value in the max-heap BST
MaxHeapNode* searchMaxHeapNode(MaxHeapNode *root, int val) {
    if (root == NULL || root->val == val) {
        return root;
    }
    if (val > root->val) {
        return searchMaxHeapNode(root->left, val);
    }
    return searchMaxHeapNode(root->right, val);
}

// Function to delete a value from the max-heap BST
MaxHeapNode* deleteMaxHeapNode(MaxHeapNode *root, int val) {
    if (root == NULL) {
        return root;
    }

    if (val > root->val) {
        root->left = deleteMaxHeapNode(root->left, val);
    } else if (val < root->val) {
        root->right = deleteMaxHeapNode(root->right, val);
    } else {
        if (root->left == NULL) {
            MaxHeapNode *temp = root->right;
            free(root);
            return temp;
        }
        if (root->right == NULL) {
            MaxHeapNode *temp = root->left;
            free(root);
            return temp;
        }

        MaxHeapNode *temp = root->right;
        while (temp->left != NULL) {
            temp = temp->left;
        }

        root->val = temp->val;
        root->right = deleteMaxHeapNode(root->right, temp->val);
    }

    return root;
}

// Function to free the memory allocated for the max-heap BST
void freeMaxHeapNode(MaxHeapNode *root) {
    if (root == NULL) {
        return;
    }
    freeMaxHeapNode(root->left);
    freeMaxHeapNode(root->right);
    free(root);
}
