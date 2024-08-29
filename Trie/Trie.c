#include "Trie.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#define ALPHABET_SIZE 26

// Function to create a new Trie node
TrieNode *createTrieNode() {
    // Allocate memory for the new node
    TrieNode *node = malloc(sizeof(TrieNode));
    if (node == NULL) {
        fprintf(stderr, "Failed to allocate memory for node\n");
        return NULL;
    }
    // Initialize the node
    for (int i = 0; i < ALPHABET_SIZE; i++) {
        node->children[i] = NULL;
    }
    node->isEndOfWord = false;
    return node;
}

// Function to insert a word into the Trie
void insertTrieNode(TrieNode **root, const char *word) {
    if (*root == NULL) {
        *root = createTrieNode();
        if (*root == NULL) {
            return;  // Memory allocation failed
        }
    }
    TrieNode *current = *root;
    // Loop over the characters in the word
    for (int i = 0; word[i] != '\0'; i++) {
        // Validate the character
        if (word[i] < 'a' || word[i] > 'z') {
            fprintf(stderr, "Invalid character in word: %c\n", word[i]);
            return;
        }
        // Get its ASCII index
        int index = word[i] - 'a';
        // If the character's index on the current node's children is NULL, create a new node
        if (current->children[index] == NULL) {
            current->children[index] = createTrieNode();
            if (current->children[index] == NULL) {
                return;  // Memory allocation failed
            }
        }
        // Move to that node
        current = current->children[index];
    }
    // Mark the end of the word; it's been exhausted
    current->isEndOfWord = true;
}

// Function to search for a word in the Trie
bool searchWord(TrieNode *root, const char *word) {
    if (root == NULL) {
        return false;
    }
    TrieNode *current = root;
    // Loop over the characters in the word
    for (int i = 0; word[i] != '\0'; i++) {
        // Validate the character
        if (word[i] < 'a' || word[i] > 'z') {
            return false;  // Invalid character
        }
        // Get its ASCII index
        int index = word[i] - 'a';
        // If the character's index on the current node's children is NULL, the word doesn't exist
        if (current->children[index] == NULL) {
            return false;
        }
        // Move to that node
        current = current->children[index];
    }
    // If the end of the word is marked, the word exists
    return current->isEndOfWord;
}

// Function to free the Trie nodes
void freeTrieNode(TrieNode *node) {
    // Base case
    if (node == NULL) {
        return;
    }
    // Recursively free the children nodes
    for (int i = 0; i < ALPHABET_SIZE; i++) {
        freeTrieNode(node->children[i]);
    }
    // Free the current node
    free(node);
}