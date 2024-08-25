#ifndef STACK_H
#define STACK_H

typedef struct Node {
    int val;
    struct Node *next;
} Node;

typedef struct Stack {
    Node *top;
    Node *bottom;
} Struct;

#endif //STACK_H
