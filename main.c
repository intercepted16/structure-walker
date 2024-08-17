//
// Created by ab on 17/08/24.
//

#include <stdio.h>
#include <string.h>
#include "linkedList/addTwoNums.c"
#include "linkedList/ListNode.c"
#include "linkedList/ListNode.h"

int addTwoNumsExample(void) {
    // cast int to void pointer
    void *val = malloc(sizeof(int));
    *(int *)val = 2;
    ListNode *l1 = createNode(val, 2);
    ListNode *l2 = createNode(val, 2);
    printf("Node 1: %d\n", *(int *)l1->val);
    printf("Node 2: %d\n", *(int *)l2->val);
    ListNode *result = addTwoNumbers(l1, l2);
    printf("Result: %d\n", *(int *)result->val);
    freeList(l1);
    freeList(l2);
    freeList(result);
    return 0;
}

int main(void) {
    // inputExample();
    addTwoNumsExample();
    return 0;
}

