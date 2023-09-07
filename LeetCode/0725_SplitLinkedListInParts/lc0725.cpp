#include <cassert>
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
  std::vector<ListNode*> splitListToParts(ListNode* head, int k) {
    std::vector<ListNode*> ans;

    // store all nodes in a vector.
    std::vector<ListNode*> nodes;
    ListNode* curr = head;
    while (curr != nullptr) {
      nodes.push_back(curr);
      curr = curr->next;
    }

    int base_length = nodes.size() / k;
    int extra_length = nodes.size() % k;
    ans.resize(k, nullptr);
    int head_idx = 0;
    for (int i = 0; i < k; i++) {
      int part_length = base_length;
      if (extra_length > 0) {
        extra_length -= 1;
        part_length += 1;
      }
      
      // Perform the cut, if applicable:
      int tail_idx = head_idx + part_length - 1;
      if (tail_idx < nodes.size() && nodes[tail_idx] != nullptr) {
        nodes[tail_idx]->next = nullptr;
      }

      // Store the head of ith part into answer array:
      if (head_idx < nodes.size()) {
        ans[i] = nodes[head_idx];
      } // otherwise ans[i] stays a nullptr assigned in the beginning.

      head_idx += part_length;
    }
    return ans;
  };
};

int main() {
  Solution s;

  // testcase 1: split [1, 2, 3] into 5 parts.
  std::cout << "Example 1:\n";
  ListNode* list1 = make_list({1, 2, 3});
  print_list(list1);
  auto ans1 = s.splitListToParts(list1, 5);
  for (const auto& elem: ans1) { print_list(elem); }
  std::cout << "\n";

  // testcase 2: split [1, 2, 3, 4, ... , 9, 10] into 3 parts.
  std::cout << "Example 2:\n";
  ListNode* list2 = make_list({1, 2, 3, 4, 5, 6, 7, 8, 9, 10});
  print_list(list2);
  auto ans2 = s.splitListToParts(list2, 3);
  for (const auto& elem: ans2) { print_list(elem); }
  std::cout << "\n";

  // testcase 3: split [] into 3 parts.
  std::cout << "Example 3:\n";
  print_list(nullptr);
  auto ans3 = s.splitListToParts(nullptr, 3);
  for (const auto& elem: ans3) { print_list(elem); }
  std::cout << "\n";

}
