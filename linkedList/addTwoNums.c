#include <stdlib.h>
#include <stdio.h>

 struct ListNode {
     int val;
     struct ListNode *next;
 };
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode* dummyHead = malloc(sizeof(struct ListNode));
    struct ListNode* current = dummyHead;
    // Loop over both nodes till exhausted, keeping track of carry(s) when adding
    int carry = 0;
    while (l1 != NULL || l2 != NULL || carry != 0) {
        // Add the place value
                int val1 = (l1 != NULL) ? l1->val : 0;
        int val2 = (l2 != NULL) ? l2->val : 0;
        int sum = val1 + val2 + carry;
        // Carry over the place value
        carry = sum / 10;
        // Get the remainding value
        sum = sum % 10;
        // update the tail and set the current val
        struct ListNode* newNode = malloc(sizeof(struct ListNode));
        current->next = newNode;
        current = current->next;
        current->val = sum;
        current->next = NULL;
        // set the linked lists to the next number
        if (l1 != NULL) l1 = l1->next;
        if (l2 != NULL) l2 = l2->next;
    }
    struct ListNode* result = dummyHead->next;
    return result;

}

int main(void) {
  // Create two linked lists
    struct ListNode* l1 = malloc(sizeof(struct ListNode));
    struct ListNode* l2 = malloc(sizeof(struct ListNode));
    // Set the values
    l1->val = 2;
    l1->next = NULL;
    l2->val = 5;
    l2->next = NULL;
    // Add the two numbers
    struct ListNode* result = addTwoNumbers(l1, l2);
    // Print the result
    while (result != NULL) {
        printf("%d\n", result->val);
        result = result->next;
    }
    return 0;
}