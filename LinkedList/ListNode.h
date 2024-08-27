//
// Created by ab on 17/08/24.
//

#ifndef LISTNODE_H
#define LISTNODE_H
typedef struct ListNode {
    void *val;
    struct ListNode *next;
} ListNode;

ListNode *createNode(const void *val, const size_t val_size);
void freeList(ListNode *head);


#endif //LISTNODE_H
