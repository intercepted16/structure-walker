#include "Stack.h"
#include <stdlib.h>

// Allocate memory for a new stack
struct Stack* createStack() {
    struct Stack *s = malloc(sizeof(struct Stack));
    if (s == NULL) {
        return NULL;
    }
    s->top = NULL;
    return s;
}

// Push a new element onto the stack
int push(struct Stack *s, const int val) {
    Node *newNode = malloc(sizeof(struct MinHeapNode));
    if (newNode == NULL) {
        return -1;  // Return error code if memory allocation fails
    }
    newNode->val = val;
    newNode->next = s->top;
    s->top = newNode;
    return 0;  // Return success code
}

// Pop the top element from the stack
int pop(struct Stack *s) {
    Node *temp = s->top;
    if (temp == NULL) {
        return -1;  // Return error code if stack is empty
    }
    s->top = temp->next;
    free(temp);
    return 0;  // Return success code
}

// Frees the memory allocated for a stack
void freeStack(struct Stack *s) {
    Node *current = s->top;
    Node *next = NULL;
    while (current != NULL) {
        next = current->next;
        free(current);
        current = next;
    }
    free(s);  // Free the stack structure itself
}
