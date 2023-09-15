#include <vector>
#include <iostream>


struct ListNode {
  ListNode(): val(0), next(nullptr) {}
  ListNode(int x): val(x), next(nullptr) {}
  ListNode(int x, ListNode* next): val(x), next(next) {}

  int val;
  ListNode *next;
};

ListNode* make_list(std::vector<int> contents) {
  if (contents.empty()) {
    return nullptr;
  }

  ListNode* head = new ListNode(contents[0]);
  ListNode* prev = head;
  ListNode* curr;
  for (int i = 1; i < contents.size(); i++) {
    curr = new ListNode(contents[i]);
    prev->next = curr;
    prev = curr;
  }

  return head;
}

void delete_list(ListNode* head) {
  if (head == nullptr) {
    return;
  }
  ListNode* prev = head;
  ListNode* curr = head->next;
  while (curr != nullptr) {
    delete prev;
    prev = curr;
    curr = curr->next;
  }
  delete prev;
}

void print_list(ListNode* node, int limit=20) {
  std::cout << "[";
  while (limit > 0 && node != nullptr) {
    std::cout << node->val << ", ";
    limit -= 1;
    node = node->next;
  }
  if (node != nullptr) {
    // If after hitting a limit we're still not at list tail, mark it:
    std::cout << "...";
  }
  std::cout << "]\n";
}

class Solution {
public:
  ListNode* reverseBetween(ListNode* head, int left, int right) {
    // NOTE: left and right are indexes, the problem uses 1-indexing instead of 0-indexing system.
    if (left == right) {
      return head;
    }

    // Store all nodes in a vector:
    std::vector<ListNode*> nodes;
    ListNode* curr = head;
    while (curr != nullptr) {
      nodes.push_back(curr);
      curr = curr->next;
    }
    
    // Convert left & right indices to 0-indexed:
    left -= 1;
    right -= 1;

    // Reverse nodes between left & right, inclusive:
    for (int i = left; i < right; i++) {
      nodes[i+1]->next = nodes[i];
    }
    if (left == 0) {
      head = nodes[right];
    } else {
      nodes[left-1]->next = nodes[right];
    }

    if (right + 1 == nodes.size()) {
      nodes[left]->next = nullptr;
    } else {
      nodes[left]->next = nodes[right+1];
    }

    return head;
  };
};


int main() {
  Solution s;

  // testcase 1: reverse [1, 2, 3, 4, 5] into [1, 4, 3, 2, 5].
  std::cout << "Example 1:\n";
  ListNode* list1 = make_list({1, 2, 3, 4, 5});
  print_list(list1);
  auto ans1 = s.reverseBetween(list1, 2, 4);
  print_list(ans1);
  std::cout << "\n";

  // testcase 2: reverse [1, 2, 3, 4, 5] into [5, 4, 3, 2, 1].
  std::cout << "Example 1:\n";
  ListNode* list2 = make_list({1, 2, 3, 4, 5});
  print_list(list2);
  auto ans2 = s.reverseBetween(list1, 1, 5);
  print_list(ans2);
  std::cout << "\n";
}
