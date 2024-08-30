#ifndef BTREE_H
#define BTREE_H
#define MAX_KEYS 3
#define MIN_DEGREE ((MAX_KEYS + 1) / 2)

typedef struct BNode {
    int keys[MAX_KEYS];
    struct BNode *children[MAX_KEYS + 1];
    int numKeys;
    int isLeaf;
} BNode;

#endif // BTREE_H