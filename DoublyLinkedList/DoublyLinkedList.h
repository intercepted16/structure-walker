//
// Created by ab on 26/08/24.
//

#ifndef DOUBLYLINKEDLIST_H
#define DOUBLYLINKEDLIST_H

typedef struct DoubleNode {
    void *val;
    struct DoubleNode *next;
    struct DoubleNode *prev;
} DoubleNode;
#endif //DOUBLYLINKEDLIST_H
