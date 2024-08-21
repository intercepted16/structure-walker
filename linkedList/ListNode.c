#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "ListNode.h"


ListNode *createNode(const void *val, const size_t val_size) {
    ListNode *node = malloc(sizeof(ListNode));
    if (node == NULL) {
        // Failed to allocate memory
        fprintf(stderr, "Failed to allocate memory\n");
        exit(EXIT_FAILURE);
    }
    if (val != NULL) {
        // create a copy of the value
        node->val = malloc(val_size);
        if (node->val == NULL) {
            // Failed to allocate memory for the value
            free(node);
            fprintf(stderr, "Failed to allocate memory for value\n");
            exit(EXIT_FAILURE);
        }
        memcpy(node->val, val, val_size);
    } else {
        node->val = NULL;
    }
    node->next = NULL;
    return node;
}

void freeList(ListNode *head) {
    ListNode *current = head;
    ListNode *next = NULL;
    while (current != NULL) {
        next = current->next;
        free(current->val);
        free(current);
        current = next;
    }
}