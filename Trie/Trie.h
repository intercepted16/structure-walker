#ifndef TRIE_H
#define TRIE_H
#include <stdbool.h>

typedef struct TrieNode {
    struct TrieNode *children[26];
    bool isEndOfWord;
    char *val;
} TrieNode;
#endif //TRIE_H
