#include <stdlib.h>
#include <stdio.h>
#include "ListNode.h"

ListNode* addTwoNumbers(const ListNode* l1, const ListNode* l2) {
    if (l1 == NULL || l2 == NULL) {
        fprintf(stderr, "One of the linked lists is NULL\n");
        exit(EXIT_FAILURE);
    }

    ListNode* dummyHead = malloc(sizeof(ListNode));
    if (dummyHead == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        exit(EXIT_FAILURE);
    }
    dummyHead->val = NULL;
    dummyHead->next = NULL;

    ListNode* current = dummyHead;
    int carry = 0;

    while (l1 != NULL || l2 != NULL || carry != 0) {
        int val1 = (l1 != NULL && l1->val != NULL) ? *(int *)l1->val : 0;
        int val2 = (l2 != NULL && l2->val != NULL) ? *(int *)l2->val : 0;
        int sum = val1 + val2 + carry;

        carry = sum / 10;
        sum = sum % 10;

        ListNode* newNode = malloc(sizeof(ListNode));
        if (newNode == NULL) {
            fprintf(stderr, "Memory allocation failed\n");
            exit(EXIT_FAILURE);
        }
        newNode->val = malloc(sizeof(int));
        if (newNode->val == NULL) {
            fprintf(stderr, "Memory allocation failed\n");
            exit(EXIT_FAILURE);
        }
        *(int *)(newNode->val) = sum;
        newNode->next = NULL;

        current->next = newNode;
        current = current->next;

        if (l1 != NULL) l1 = l1->next;
        if (l2 != NULL) l2 = l2->next;
    }

    ListNode* result = dummyHead->next;
    free(dummyHead);

    return result;
}
