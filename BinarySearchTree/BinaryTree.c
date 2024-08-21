#include <stdio.h>
#include <stdlib.h>
#include "BinaryTree.h"

// Function to create a new node
TreeNode* createTreeNode(int val) {
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

// Function to insert a value into the binary search tree
// This function is recursive
void insertTreeNode(TreeNode **root, const int val) {
    if (*root == NULL) {
        *root = createTreeNode(val);
        printf("Inserted %d as a new node.\n", val);
        return;
    }

    printf("At node %d: trying to insert %d\n", (*root)->val, val);

    if (val < (*root)->val) {
        // If the value is less than the root, insert it into the left subtree
        printf("Moving left to insert %d under %d\n", val, (*root)->val);
        insertTreeNode(&((*root)->left), val);
    } else {
        // If the value is greater than or equal to the root, insert it into the right subtree
        printf("Moving right to insert %d under %d\n", val, (*root)->val);
        insertTreeNode(&((*root)->right), val);
    }
}
// Function to search for a value in the binary search tree
// This function is recursive; it calls itself till it finds the value or reaches a leaf node
TreeNode* searchTreeNode(TreeNode *root, int val) {
  // base case
    if (root == NULL || root->val == val) {
        return root;
    }
    if (val < root->val) {
        return searchTreeNode(root->left, val);
    }
    return searchTreeNode(root->right, val);
}
// Function to delete a value from the binary search tree
// This function is recursive
TreeNode* deleteTreeNode(TreeNode *root, int val) {
    // Base case: If the root is NULL, the tree is empty or we have reached a leaf node
    if (root == NULL) {
        return root;
    }

    // If the value to be deleted is less than the root's value, it must be in the left subtree
    if (val < root->val) {
        root->left = deleteTreeNode(root->left, val);
    }
    // If the value to be deleted is greater than the root's value, it must be in the right subtree
    else if (val > root->val) {
        root->right = deleteTreeNode(root->right, val);
    }
    // If the value matches the root's value, this is the node to be deleted
    else {
        // Case 1: Node with only one child or no child
        if (root->left == NULL) {
            // If there's no left child, replace this node with its right child
            TreeNode *temp = root->right;
            free(root); // Free memory of the node to be deleted
            return temp; // Return the new subtree root
        }
        if (root->right == NULL) {
            // If there's no right child, replace this node with its left child
            TreeNode *temp = root->left;
            free(root); // Free memory of the node to be deleted
            return temp; // Return the new subtree root
        }

        // Case 2: Node with two children
        // Find the inorder successor (smallest node in the right subtree)
        TreeNode *temp = root->right;
        while (temp->left != NULL) {
            temp = temp->left;
        }

        // Copy the inorder successor's value to this node
        root->val = temp->val;

        // Delete the inorder successor
        root->right = deleteTreeNode(root->right, temp->val);
    }

    // Return the (possibly updated) root of the subtree
    return root;
}

// Function to free the memory allocated for the binary search tree
void freeTreeNode(TreeNode *root) {
    if (root == NULL) {
        return;
    }
    freeTreeNode(root->left);
    freeTreeNode(root->right);
    free(root);
}