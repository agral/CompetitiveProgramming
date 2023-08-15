#include <iostream>
#include <utility>
#include <vector>

struct ListNode {
  int val;
  ListNode* next;
  ListNode() : val(0), next(nullptr) {};
  ListNode(int x) : val(x), next(nullptr) {};
  ListNode(int x, ListNode* next) : val(x), next(next) {};
};

class Solution {
public:
  ListNode* partition(ListNode* head, int x) {
    if (head == nullptr) {
      return nullptr;
    }

    ListNode* lhs = nullptr;
    ListNode* rhs = nullptr;
    ListNode* lhs_tail = nullptr;
    ListNode* rhs_tail = nullptr;
    while (head != nullptr) {
      ListNode* node = new ListNode(head->val);
      if (node->val < x) {
        if (lhs == nullptr) {
          lhs = node;
          lhs_tail = node;
        }
        else {
          lhs_tail->next = node;
          lhs_tail = node;
        }
      }
      else {
        if (rhs == nullptr) {
          rhs = node;
          rhs_tail = node;
        }
        else {
          rhs_tail->next = node;
          rhs_tail = node;
        }
      }
      head = head->next;
    }
    
    // Now join lhs+rhs together in one list:
    if (lhs == nullptr) {
      return rhs;
    }
    if (rhs == nullptr) {
      return lhs;
    }
    lhs_tail->next = rhs;
    return lhs;
  }
};

ListNode* make_list(std::vector<int> values) {
  if (values.empty()) {
    return nullptr;
  }
  ListNode* head = new ListNode(values[0]);
  ListNode* last = head;
  for (auto i = 1; i < values.size(); i++) {
    ListNode* curr = new ListNode(values[i]);
    last->next = curr;
    last = curr;
  }
  return head;
}

void delete_list(ListNode* head) {
  if (head == nullptr) {
    return;
  }
  while (head->next != nullptr) {
    ListNode* next = head->next;
    delete head;
    head = next;
  }
  delete head;
}

void print_list(ListNode* head) {
  while (head != nullptr) {
    std::cout << head->val << ' ';
    head = head->next;
  }
  std::cout << '\n';
}

int main() {
  Solution s{};
  ListNode* testcase1 = make_list({1, 4, 3, 2, 5, 2});
  ListNode* ans1 = s.partition(testcase1, 3);
  print_list(ans1);
  delete_list(ans1);

  // degenerate case: this one results with empty lhs in partition():
  ListNode* ans1b = s.partition(testcase1, -5);
  print_list(ans1b);
  delete_list(ans1b);

  // degenerate case: this one results with empty rhs in partition():
  ListNode* ans1c = s.partition(testcase1, 100);
  print_list(ans1c);
  delete_list(ans1c);

  delete_list(testcase1);


  ListNode* testcase2 = make_list({2, 1});
  ListNode* ans2 = s.partition(testcase2, 2);
  print_list(ans2);
  delete_list(ans2);
  delete_list(testcase2);

  // Make sure that list with 0 elements is handled correctly too:
  ListNode* ans3 = s.partition(nullptr, 10);
  print_list(ans3);
  delete_list(ans3);
}
