#ifndef HASHTABLE_H
#define HASHTABLE_H

#include "../linkedList/ListNode.h"

#define TABLE_SIZE 100

// extends ListNode with a key
typedef struct HashEntry {
    void *val;
    struct HashEntry *next;
    size_t val_size;
    char *key;
} HashEntry;

typedef struct {
    HashEntry *buckets[TABLE_SIZE];
} HashTable;

#endif // HASHTABLE_H