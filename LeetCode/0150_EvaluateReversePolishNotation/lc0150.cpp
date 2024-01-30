#include <cassert>
#include <deque>
#include <functional>
#include <string>
#include <unordered_map>
#include <vector>

class Solution {
public:
  long evalRPN(std::vector<std::string>& tokens) {
    const std::unordered_map<std::string, std::function<long(long, long)>> token_to_operation {
      {"+", [](long lhs, long rhs){ return lhs + rhs; }},
      {"-", [](long lhs, long rhs){ return lhs - rhs; }},
      {"*", [](long lhs, long rhs){ return lhs * rhs; }},
      {"/", [](long lhs, long rhs){ return lhs / rhs; }},
    };
    std::deque<long> stack;

    for (const std::string& token: tokens) {
      // tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
      if (token_to_operation.count(token) == 0) {
        stack.push_back(std::stoi(token));
      } else {
        const long rhs = stack.back();
        stack.pop_back();
        const long lhs = stack.back();
        stack.pop_back();
        stack.push_back(token_to_operation.at(token)(lhs, rhs));
      }
    }
    return stack.back();
  }
};

int main() {
  Solution s;
  {
    std::vector<std::string> tokens = {"2", "1", "+", "3", "*"};
    assert(s.evalRPN(tokens) == 9);
  }
  {
    std::vector<std::string> tokens = {"4", "13", "5", "/", "+"};
    assert(s.evalRPN(tokens) == 6);
  }
  {
    std::vector<std::string> tokens =
      {"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"};
    assert(s.evalRPN(tokens) == 22);
  }
}
