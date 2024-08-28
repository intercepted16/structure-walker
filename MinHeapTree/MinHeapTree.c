#include <stdio.h>
#include <stdlib.h>
#include "MinHeapTree.h"

// Function to create a new node in the min-heap tree
TreeNode* createMinHeapNode(const int val) {
    TreeNode *node = malloc(sizeof(TreeNode));
    if (node == NULL) {
        fprintf(stderr, "Failed to allocate memory for node\n");
        return NULL;
    }
    node->val = val;
    node->left = NULL;
    node->right = NULL;
    return node;
}

// Function to insert a value into the min-heap BST
void insertMinHeapNode(TreeNode **root, const int val) {
    if (*root == NULL) {
        *root = createMinHeapNode(val);
        printf("Inserted %d as a new node.\n", val);
        return;
    }

    printf("At node %d: trying to insert %d\n", (*root)->val, val);

    if (val < (*root)->val) {
        // Swap values to maintain min-heap property
        int temp = (*root)->val;
        (*root)->val = val;
        insertMinHeapNode(&((*root)->left), temp);
    } else {
        insertMinHeapNode(&((*root)->right), val);
    }
}

// Function to search for a value in the min-heap BST
TreeNode* searchMinHeapNode(TreeNode *root, int val) {
    if (root == NULL || root->val == val) {
        return root;
    }
    if (val < root->val) {
        return searchMinHeapNode(root->left, val);
    }
    return searchMinHeapNode(root->right, val);
}

// Function to delete a value from the min-heap BST
TreeNode* deleteMinHeapNode(TreeNode *root, int val) {
    if (root == NULL) {
        return root;
    }

    if (val < root->val) {
        root->left = deleteMinHeapNode(root->left, val);
    } else if (val > root->val) {
        root->right = deleteMinHeapNode(root->right, val);
    } else {
        if (root->left == NULL) {
            TreeNode *temp = root->right;
            free(root);
            return temp;
        }
        if (root->right == NULL) {
            TreeNode *temp = root->left;
            free(root);
            return temp;
        }

        TreeNode *temp = root->right;
        while (temp->left != NULL) {
            temp = temp->left;
        }

        root->val = temp->val;
        root->right = deleteMinHeapNode(root->right, temp->val);
    }

    return root;
}

// Function to free the memory allocated for the min-heap BST
void freeMinHeapNode(TreeNode *root) {
    if (root == NULL) {
        return;
    }
    freeMinHeapNode(root->left);
    freeMinHeapNode(root->right);
    free(root);
}
