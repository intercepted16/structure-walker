//
// Created by ab on 25/08/24.
//

#ifndef QUEUE_H
#define QUEUE_H
typedef struct Node {
    int val;
    struct Node *next;
} Node;

typedef struct Queue {
    Node *head;
    Node *tail;
} Queue;

#endif //QUEUE_H
