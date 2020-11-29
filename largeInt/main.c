#include <stdio.h>
#include <stdlib.h>

struct Node {
	int val;
	int numNeighbors;
	struct Node** neighbors;
};
typedef struct Node* Node;

Node newNode(int val, int numNeighbors) {
	Node node = (Node) malloc(sizeof(struct Node));
	node->val = val;
	node->numNeighbors = numNeighbors;
	node->neighbors = (struct Node**) malloc(sizeof(struct Node*) * numNeighbors);
	for (int i = 0; i < numNeighbors; i++) {
		node->neighbors[i] = NULL;
	}
	return node;
}

int main() {
	int numNodes = 2;
	Node node = newNode(3, numNodes);
	for (int i = 0; i < numNodes; i++) {
		printf("%p\n", &node->neighbors[i]);
	}
	return 0;
}
