// Solution to LeetCode problem #0002, "Two Numbers"
// https://leetcode.com/problems/add-two-numbers/
// by https://gralin.ski/

#include <initializer_list>
#include <iostream>

struct ListNode;

struct ListNode
{
  int val;
  ListNode* next;

  ListNode(int x = 0): val(x), next(nullptr) {};
};

ListNode* buildListNode(std::initializer_list<int> nums)
{
  if (nums.size() == 0) {
    nums = {0};
  }
  ListNode *previous = nullptr;
  ListNode *head = nullptr;
  for (auto& num: nums) {
    ListNode *current = new ListNode(num);
    if (previous == nullptr) {
      head = current;
    } else {
      previous->next = current;
    }
    previous = current;
  }
  return head;
}

void printListNode(ListNode* ptr) {
  std::cout << "[";
  while (ptr != nullptr) {
    std::cout << ptr->val << ", ";
    ptr = ptr->next;
  }
  std::cout << "]\n";
}

ListNode* add_two_lists(ListNode* lhs, ListNode* rhs) {
  auto carry = 0;
  ListNode *head = nullptr;
  ListNode *previous = nullptr;
  while(!(lhs == nullptr && rhs == nullptr)) {
    // Read the current digit if available; use value of 0 if the reversed number is already exhausted:
    auto lhsval = (lhs == nullptr) ? 0 : lhs->val;
    auto rhsval = (rhs == nullptr) ? 0 : rhs->val;

    auto sum = carry + lhsval + rhsval;
    carry = sum / 10;
    auto digit = sum % 10;
    ListNode* node = new ListNode(digit);
    if (previous == nullptr) {
      head = node;
    } else {
      previous->next = node;
    }
    previous = node;

    // Move each of the two pointers to the next digit, if available:
    // (*note*: lhs and rhs can be of different lengths, so need to check it separately)
    if (lhs != nullptr) {
      lhs = lhs->next;
    }
    if (rhs != nullptr) {
      rhs = rhs->next;
    }
  }
  if (carry > 0) {
    ListNode* node = new ListNode(carry);
    previous->next = node;
  }

  return head;
}

int main()
{
  {
    std::cout << "Example 1\n";
    ListNode* left_addend = buildListNode({2, 4, 3});
    ListNode* right_addend = buildListNode({5, 6, 4});
    ListNode* sum = add_two_lists(left_addend, right_addend);
    printListNode(sum);
  }
  {
    std::cout << "Example 2\n";
    ListNode* left_addend = buildListNode({0});
    ListNode* right_addend = buildListNode({0});
    ListNode* sum = add_two_lists(left_addend, right_addend);
    printListNode(sum);
  }
  {
    std::cout << "Example 3\n";
    ListNode* left_addend = buildListNode({9, 9, 9, 9, 9, 9, 9});
    ListNode* right_addend = buildListNode({9, 9, 9, 9});
    ListNode* sum = add_two_lists(left_addend, right_addend);
    printListNode(sum);
  }
  {
    std::cout << "Example that adds up to exactly 10000\n";
    ListNode* left_addend = buildListNode({6, 7, 8, 9});        // -> 9876, reverse notated
    ListNode* right_addend = buildListNode({4, 2, 1});          // ->  124, reverse notated
    ListNode* sum = add_two_lists(left_addend, right_addend);   // = 10000, when added.
    printListNode(sum); // Expect [0, 0, 0, 0, 1] in the output.
  }
}
