#include <cassert>
#include <iostream>
#include <vector>

struct ListNode {
  ListNode(int x): val(x), next(nullptr) {}

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

void limited_print(ListNode* node, int limit=20) {
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
  bool hasCycle(ListNode* head) {
    if (head == nullptr) {
      return false; // no cycles in an empty linked list.
    }
    ListNode* fast = head;
    ListNode* slow = head;
    while (true) {
      if (fast->next == nullptr || fast->next->next == nullptr) {
        // The list terminates. No cycle.
        return false;
      }
      fast = fast->next->next;
      slow = slow->next;
      if (slow == fast) {
        // Slow caught up with fast ==> there is a cycle.
        // This method will detect any cycles, no matter the loop's length.
        return true;
      }
    }
    return false;
  }
};

int main() {
  Solution s;
  ListNode* list0 = make_list({0, 1, 2, 3, 4, 5});
  limited_print(list0);
  assert(s.hasCycle(list0) == false);

  ListNode* list1 = make_list({3, 2, 0, -4});
  // Make a loop: -4 links back to 2 (3rd elem to 1st):
  (list1->next->next->next)->next = list1->next;
  limited_print(list1);
  assert(s.hasCycle(list1) == true);

  ListNode* list2 = make_list({1, 2});
  // Make a loop: 2 links back to 1 (1st elem to 0th):
  (list2->next)->next = list2;
  limited_print(list2);
  assert(s.hasCycle(list2) == true);

  ListNode* list3 = make_list({1});
  limited_print(list3);
  assert(s.hasCycle(list3) == false);
}
