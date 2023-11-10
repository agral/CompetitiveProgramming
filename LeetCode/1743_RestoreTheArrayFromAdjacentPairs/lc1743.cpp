#include <cassert>
#include <map>
#include <vector>

#include <algorithm> // just for testing purposes, not used in Solution itself.
#include <iostream>

class Node {
public:
  Node() : val(0), neighbors{nullptr, nullptr}, numNeighbors(0) {}
  Node(int val) : val(val), neighbors{nullptr, nullptr}, numNeighbors(0) {}

  int val;
  Node* neighbors[2];
  int numNeighbors;
};

class Solution {
public:
  std::vector<int> restoreArray(std::vector<std::vector<int>>& adjacentPairs) {
    std::map<int, Node*> existing_nodes;
    for (const std::vector<int>& pair: adjacentPairs) {
      // Find or create appropriate Nodes:
      Node *n0 = nullptr, *n1 = nullptr;
      auto it0 = existing_nodes.find(pair[0]), it1 = existing_nodes.find(pair[1]);
      if (it0 == existing_nodes.end()) {
        n0 = new Node(pair[0]);
        existing_nodes.insert({pair[0], n0});
      } else {
        n0 = it0->second;
      }
      if (it1 == existing_nodes.end()) {
        n1 = new Node(pair[1]);
        existing_nodes.insert({pair[1], n1});
      } else {
        n1 = it1->second;
      }

      // Create a linkage between these Nodes:
      n0->neighbors[n0->numNeighbors] = n1;
      n0->numNeighbors++;
      n1->neighbors[n1->numNeighbors] = n0;
      n1->numNeighbors++;
    }

    // Find a Node that has exactly one neighbor:
    Node* starter = nullptr;
    for (auto it = existing_nodes.begin(); (starter == nullptr) && (it != existing_nodes.end()); it++) {
      if (it->second->numNeighbors == 1) {
        starter = it->second;
      }
    }

    // Reconstruct the array starting from the found starter Node:
    std::vector<int> ans;
    ans.push_back(starter->val);
    Node* current = starter->neighbors[0];
    Node* previous = starter;
    while (current->numNeighbors == 2) {
      ans.push_back(current->val);
      Node* next = current->neighbors[0];
      // Ensure that no steps back are taken:
      if (next == previous) {
        next = current->neighbors[1];
      }

      previous = current;
      current = next;
    }

    // Include the last Node too:
    ans.push_back(current->val);

    return ans;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<int>> pairs = {{2, 1}, {3, 4}, {3, 2}};
    std::vector<int> expected_answer = {1, 2, 3, 4};
    auto reversed_expected_answer = expected_answer;
    std::reverse(expected_answer.begin(), expected_answer.end());
    std::vector<int> answer = s.restoreArray(pairs);
    assert(answer == expected_answer || answer == reversed_expected_answer);
  }
  {
    std::vector<std::vector<int>> pairs = {{4, -2}, {1, 4}, {-3, 1}};
    std::vector<int> expected_answer = {-2, 4, 1, -3};
    auto reversed_expected_answer = expected_answer;
    std::reverse(expected_answer.begin(), expected_answer.end());
    std::vector<int> answer = s.restoreArray(pairs);
    assert(answer == expected_answer || answer == reversed_expected_answer);
  }
  {
    std::vector<std::vector<int>> pairs = {{100000, -100000}};
    std::vector<int> expected_answer = {100000, -100000};
    auto reversed_expected_answer = expected_answer;
    std::reverse(expected_answer.begin(), expected_answer.end());
    std::vector<int> answer = s.restoreArray(pairs);
    assert(answer == expected_answer || answer == reversed_expected_answer);
  }
}
