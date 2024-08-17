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


int extern exampleLinkedList(void) {
    ListNode *l1 = createNode(NULL, 0);
    ListNode *l2 = createNode(NULL, 0);
    ListNode *current = l1;
    char input[30];
    while (1) {
        fgets(input, sizeof(input), stdin);
        // Strip the newline character
        input[strcspn(input, "\n")] = 0;
        printf("You entered: %s\n", input);
        // if input is "q", exit
        if (strcmp(input, "q") == 0) {
            break;
        }
        // create a new node
        ListNode *newNode = createNode(input, strlen(input) + 1);
        current->next = newNode;
        current = current->next;
    }
    freeList(l1);
    freeList(l2);
    return 0;
}