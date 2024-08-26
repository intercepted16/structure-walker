#include "DoublyLinkedList.h"

#include <stdlib.h>
#include <stdio.h>

// Create a new doubly linked node
DoubleNode *createDoubleNode() {
    DoubleNode *newNode = malloc(sizeof(DoubleNode));
    if (newNode == NULL) {
        return NULL;
    }
    newNode->val = NULL;
    newNode->next = NULL;
    newNode->prev = NULL;
    return newNode;
}

// Insert a new node into a doubly linked node
int insertDoubleNode(DoubleNode *head, void *val) {
    DoubleNode *newNode = createDoubleNode();
    if (newNode == NULL) {
        return -1;
    }
    newNode->val = val;
    newNode->next = head->next;
    newNode->prev = head;
    head->next = newNode;
    // If the new node is not the last item in the doubly linked list
    // set the previous pointer of the next node to the new node
    // as to insert itself between the head and the next node
    if (newNode->next != NULL) {
        newNode->next->prev = newNode;
    }
    // If it is the last node, which is more likely, return.
    return 0;
}

int appendDoubleNode(DoubleNode **head, void *val) {
    DoubleNode *newNode = createDoubleNode();
    if (newNode == NULL) {
        return -1; // Failed to create a new node
    }

    newNode->val = val;
    newNode->next = NULL;

    // If the list is empty or if it's the first node being inserted, set the head's value to the new node
    if (*head == NULL || (*head)->val == NULL) {
        newNode->prev = NULL;
        *head = newNode;
        return 0;
    }
    // Traverse to the end of the list
    DoubleNode *current = *head;
    while (current->next != NULL) {
        current = current->next;
    }
    // Append the new node at the end
    current->next = newNode;
    newNode->prev = current;

    return 0; // Success
}


// Delete a node from a doubly linked list
int deleteDoubleNode(const DoubleNode *head, const void *val) {
    DoubleNode *current = head->next;
    // Recursively traverse the doubly linked list from the head
    // to find the node with the value to be deleted
    // and delete it by updating the pointers of the previous and next nodes
    while (current != NULL) {
        if (current->val == val) {
            // If the node to be deleted is not the last node
            // update the next node's previous pointer to the previous node
            // as to remove itself in between the previous and next node
            // and keep the doubly linked list intact
            current->prev->next = current->next;
            // The same goes for the next node's previous pointer
            if (current->next != NULL) {
                current->next->prev = current->prev;
            }
            // Now we can safely free the memory allocated for the node
            free(current);
            return 0;
        }
        // If the node is not found, move to the next node
        current = current->next;
    }
    // If the node is not found, return -1 to indicate failure
    return -1;
}

// Free the memory allocated for a doubly linked list
void freeDoubleNode(DoubleNode *head) {
    DoubleNode *current = head->next;
    DoubleNode *next = NULL;
    // Traverse the doubly linked list from the head
    // and free the memory allocated for each node
    while (current != NULL) {
        next = current->next;
        free(current);
        current = next;
    }
    free(head);
}