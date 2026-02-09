#include <cassert>
#include <stack>

class MyQueue {
public:
  void push(int x) {
    m_input.push(x);
  }

  int pop() {
    peek();
    auto value = m_output.top();
    m_output.pop();
    return value;
  }

  int peek() {
    if (m_output.empty()) {
      while (!m_input.empty()) {
        auto value = m_input.top();
        m_input.pop();
        m_output.push(value);
      }
    }
    return m_output.top();
  }

  bool empty() {
    return m_input.empty() && m_output.empty();
  }

private:
  std::stack<int> m_input;
  std::stack<int> m_output;
};

int main() {
  MyQueue q;
  assert(q.empty() == true);

  q.push(1);
  assert(q.empty() == false);
  assert(q.peek() == 1);

  q.push(2);
  assert(q.peek() == 1); // This is a FIFO queue, 1 is still on top.

  q.pop();
  assert(q.empty() == false);
  assert(q.peek() == 2);

  q.pop();
  assert(q.empty() == true);
}
