//
// Created by ab on 25/08/24.
//
#include <stdlib.h>
#include "Queue.h"

// Allocate memory for a new queue
Queue* createQueue() {
    Queue *q = malloc(sizeof(Queue));
    if (q == NULL) {
        return NULL;
    }
    q->head = NULL;
    q->tail = NULL;
    return q;
}

// Add a new element to the queue
int enqueue(Queue *q, const int val) {
    Node *newNode = malloc(sizeof(Node));
    if (newNode == NULL) {
        return -1;
    }
    newNode->val = val;
    newNode->next = NULL;
    // If there are no elements in the queue, set the head and tail to the new node
    if (q->head == NULL) {
        q->head = newNode;
        q->tail = newNode;
    } else {
        // If there are elements in the queue, set the tail's next to the new node and update the tail
        q->tail->next = newNode;
        q->tail = newNode;
    }
    // Return 0 to indicate success
    return 0;
}

// Removes the first element in a queue and returns an integer indicating success
int dequeue(Queue *q) {
    // The first element in a queue is it's head
    Node *temp = q->head;
    // If the queue is empty, return -1 to indicate failure
    if (temp == NULL) {
        return -1;
    }
    // Set the head to the next element in the queue
    q->head = temp->next;
    // Free the memory allocated for the first element
    free(temp);
    // Return 0 to indicate success
    return 0;
}

// Frees the memory allocated for a queue
void freeQueue(Queue *q) {
    Node *current = q->head;
    Node *next = NULL;
    while (current != NULL) {
        next = current->next;
        free(current);
        current = next;
    }
    free(q);
}