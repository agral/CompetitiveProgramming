#include <cassert>
#include <iostream>
#include <vector>
#include <random>

class Node {
public:
  Node(int x): val(x), next(nullptr), random(nullptr) {}

  int val;
  Node *next;
  Node *random;
};

Node* make_list(std::vector<int> contents) {
  if (contents.empty()) {
    return nullptr;
  }

  std::vector<Node*> nodes; // store pointers to all nodes for setting up random links
  nodes.resize(contents.size());

  Node* head = new Node(contents[0]);
  nodes[0] = head;
  Node* prev = head;
  Node* curr;
  for (int i = 1; i < contents.size(); i++) {
    curr = new Node(contents[i]);
    nodes[i] = curr;
    prev->next = curr;
    prev = curr;
  }

  // Set up random links:
  std::random_device rand_dev;
  std::mt19937 rng(rand_dev());
  std::uniform_int_distribution<std::mt19937::result_type> dist_nodes(0, nodes.size());
  for (int i = 0; i < nodes.size(); i++) {
    int random_node_index = dist_nodes(rng);
    nodes[i]->random = nodes[random_node_index];
  }

  return head;
}

void delete_list(Node* head) {
  if (head == nullptr) {
    return;
  }
  Node* prev = head;
  Node* curr = head->next;
  while (curr != nullptr) {
    delete prev;
    prev = curr;
    curr = curr->next;
  }
  delete prev;
}

void print(Node* node) {
  std::cout << "[";
  while (node != nullptr) {
    std::cout << "(" << node->val << ", r:"
      << (node->random == nullptr ? 0 : node->random->val) << "), ";
    node = node->next;
  }
  std::cout << "]\n";
}

class Solution {
public:
  Node* copyRandomList(Node* head) {
    if (head == nullptr) {
      // Nothing to be done for empty lists.
      return nullptr;
    }

    // 1st pass: create a copy of list with val and next pointer set, random pointer unset.
    // Store all old and copied (new) nodes in respective vectors.
    std::vector<Node*> new_nodes;
    std::vector<Node*> old_nodes;
    Node* old_curr = head;
    old_nodes.push_back(head);
    Node *new_curr = new Node(old_curr->val);
    new_nodes.push_back(new_curr);
    Node* new_prev = new_curr;
    old_curr = old_curr->next;
    while (old_curr != nullptr) {
      old_nodes.push_back(old_curr);
      new_curr = new Node(old_curr->val);
      new_nodes.push_back(new_curr);
      new_prev->next = new_curr;
      new_prev = new_curr;
      old_curr = old_curr->next;
    }

    // 2nd pass: for each node from old vector: find out to which element it points to in the old list
    //   (note: for some reason, random pointer can be set to nullptr - handle it too).
    // then use this index to link random pointers in the new list accordingly.
    for (int i = 0; i < new_nodes.size(); i++) {
      if (old_nodes[i]->random == nullptr) {
        new_nodes[i]->random = nullptr;
        // We could do nothing - Node constructor is supposed to set random ptr to nullptr.
        // But who knows if that's true, it's leetcode. Set it up to be sure.
      } else {
        // Find the index the random pointer is pointing to in the old list:
        int idx = 0;
        while (idx < old_nodes.size() && old_nodes[i]->random != old_nodes[idx]) {
          idx++;
        }
        new_nodes[i]->random = new_nodes[idx];
      }
    }
    return new_nodes[0];
  }
};

int main() {
  Solution s;
  Node* src0 = make_list({41, 42, 43, 44, 45, 46, 47, 48});
  Node* ans0 = s.copyRandomList(src0);
  print(src0);
  print(ans0);

  Node* ans_null = s.copyRandomList(nullptr);
  print(ans_null);
}
