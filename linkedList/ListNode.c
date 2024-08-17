#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "ListNode.h"



ListNode *createNode(const char *val) {
    ListNode *node =  malloc(sizeof(ListNode));
    if (node == NULL) {
        // Failed to allocate memory
        fprintf(stderr, "Failed to allocate memory\n");
        exit(EXIT_FAILURE);
    }
    if (val != NULL) {
        node->val = strdup(val);

        if (node->val == NULL) {
            // Failed to allocate memory for the string
            free(node);
            fprintf(stderr, "Failed to allocate memory for value\n");
            exit(EXIT_FAILURE);
        }
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
    ListNode *head = createNode(NULL);
    ListNode *current = head;
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
        // append it to the linked list
        ListNode *newNode = createNode(input);
        // Make the next node the new node
        current->next = newNode;
        // Reassign the current node to the new node
        // ^^ this seems odd, but it's because we recursively loop through it
        current = newNode;
    }

    current = head;

    while (current != NULL) {
        if (current->val != NULL) {
            printf("%p\n", current->val);
        }
        current = current->next;
    }
    freeList(head);
    return 0;
}
