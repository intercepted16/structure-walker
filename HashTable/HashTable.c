#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "HashTable.h"
#include "murmurhash.h"

// Function to create a new hash entry
HashEntry* createHashEntry(const char *key, const void *val, size_t val_size) {
    HashEntry *entry = malloc(sizeof(HashEntry));
    if (entry == NULL) {
        fprintf(stderr, "Failed to allocate memory for HashEntry\n");
        return NULL;
    }
    entry->key = strdup(key); // Copy the key
    if (entry->key == NULL) {
        fprintf(stderr, "Failed to allocate memory for key\n");
        free(entry);
        return NULL;
    }
    entry->val = malloc(val_size); // Copy the value
    if (entry->val == NULL) {
        fprintf(stderr, "Failed to allocate memory for value\n");
        free(entry->key);
        free(entry);
        return NULL;
    }
    memcpy(entry->val, val, val_size);
    entry->val_size = val_size;
    entry->next = NULL;
    return entry;
}

// Function to create a hash table
HashTable* createHashTable() {
    HashTable *ht = malloc(sizeof(HashTable));
    if (ht == NULL) {
        fprintf(stderr, "Failed to allocate memory for HashTable\n");
        return NULL;
    }
    memset(ht->buckets, 0, sizeof(ht->buckets)); // Initialize all buckets to NULL
    return ht;
}

// Function to append an entry to the hash table
int appendToHashTable(HashTable *ht, const char *key, const void *val, size_t val_size) {
    uint32_t hash = murmurhash(key, (uint32_t)strlen(key), 0);
    int index = hash % TABLE_SIZE;

    HashEntry *new_entry = createHashEntry(key, val, val_size);
    if (new_entry == NULL) {
        return -1; // Memory allocation failed
    }

    new_entry->next = ht->buckets[index];
    ht->buckets[index] = new_entry;
    return 0;
}

// Function to get a value from the hash table by key
void* getFromHashTable(HashTable *ht, const char *key) {
    uint32_t hash = murmurhash(key, (uint32_t)strlen(key), 0);
    int index = hash % TABLE_SIZE;

    HashEntry *current = ht->buckets[index];
    while (current != NULL) {
        if (strcmp(current->key, key) == 0) {
            return current->val; // Found the key
        }
        current = current->next;
    }
    return NULL; // Key not found
}

// Function to free a hash entry
void freeHashEntry(HashEntry *entry) {
    if (entry) {
        free(entry->key);
        free(entry->val);
        free(entry);
    }
}

// Function to free the entire hash table
void freeHashTable(HashTable *ht) {
    if (ht) {
        for (int i = 0; i < TABLE_SIZE; i++) {
            HashEntry *current = ht->buckets[i];
            while (current != NULL) {
                HashEntry *to_free = current;
                current = current->next;
                freeHashEntry(to_free);
            }
        }
        free(ht);
    }
}