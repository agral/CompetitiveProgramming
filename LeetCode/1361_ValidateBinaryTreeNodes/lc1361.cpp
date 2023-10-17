#include <cassert>
#include <vector>
#include <unordered_set>
#include <deque>

class Solution {
public:
  bool validateBinaryTreeNodes(int n, std::vector<int>& leftChild, std::vector<int>& rightChild) {
    std::unordered_set<int> nodes_with_parent;
    for (const auto& num: leftChild) {
      if (num != -1) {
        nodes_with_parent.insert(num);
      }
    }
    for (const auto& num: rightChild) {
      if (num != -1) {
        nodes_with_parent.insert(num);
      }
    }

    // There have to be exactly n-1 nodes with a parent. The one node without parent is the root element.
    if (nodes_with_parent.size() != n - 1) {
      return false;
    }

    int root = -1;
    for (int i = 0; i < n; i++) {
      if (nodes_with_parent.find(i) == nodes_with_parent.end()) {
        root = i;
        break; // out of for loop
      }
    }

    std::unordered_set<int> visited_nodes;
    std::deque<int> nodes_to_visit;
    nodes_to_visit.push_back(root);
    while (!nodes_to_visit.empty()) {
      int node = nodes_to_visit.front();
      nodes_to_visit.pop_front();
      if (visited_nodes.find(node) != visited_nodes.end()) {
        // found an already visited node - this means either a cycle or a diamond pattern is present in that tree.
        return false;
      }
      visited_nodes.insert(node);
      int left_child = leftChild[node];
      if (left_child != -1) {
        nodes_to_visit.push_back(left_child);
      }
      int right_child = rightChild[node];
      if (right_child != -1) {
        nodes_to_visit.push_back(right_child);
      }
    }

    return (visited_nodes.size() == n);
  }
};

int main() {
  Solution s;

  std::vector<int> left1 = {1, -1, 3, -1};
  std::vector<int> right1 = {2, -1, -1, -1};
  assert(s.validateBinaryTreeNodes(4, left1, right1) == true);

  std::vector<int> left2 = {1, -1, 3, -1};
  std::vector<int> right2 = {2, 3, -1, -1};
  assert(s.validateBinaryTreeNodes(4, left2, right2) == false);

  std::vector<int> left3 = {1, 0};
  std::vector<int> right3 = {-1, -1};
  assert(s.validateBinaryTreeNodes(2, left3, right3) == false);
}
