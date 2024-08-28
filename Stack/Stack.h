#ifndef STACK_H
#define STACK_H

typedef struct MinHeapNode {
    int val;
    struct MinHeapNode *next;
} Node;

typedef struct Stack {
    Node *top;
    Node *bottom;
} Struct;

#endif //STACK_H
