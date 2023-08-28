/**
 * Note: this is a bit nonsense challenge which asks to implement a LIFO stack
 * using two classic queues (i.e. using only push_back, peek_front/pop_front, size and is_empty).
 * The resulting LIFO stack has push() and pop()/top() operations of O(n) complexity - i.e. really bad.
 * Even approaching this using C-arrays would result in better performance.
 */

#include <queue>
#include <cassert>

class MyStack {
public:
  void push(int x) {
    int size_before_push = m_q.size();
    m_q.push(x); // push x to back of queue
    for (int k = 0; k < size_before_push; k++) {
      m_q.push(m_q.front());
      m_q.pop();
    }
  }

  int pop() {
    int ans = m_q.front();
    m_q.pop();
    return ans;
  }

  int top() {
    return m_q.front();
  }

  bool empty() {
    return m_q.empty();
  }

private:
  std::queue<int> m_q;
};

int main() {
  {
    MyStack s{};
    s.push(1);
    s.push(2);
    assert(s.top() == 2);
    assert(s.pop() == 2);
    assert(s.empty() == false);
  }

  {
    MyStack s{};
    s.push(1);
    s.push(2);
    s.push(3);
    assert(s.pop() == 3);
    s.push(4);
    assert(s.top() == 4);
    assert(s.pop() == 4);
    assert(s.pop() == 2);
    assert(s.empty() == false);
    assert(s.pop() == 1);
    assert(s.empty() == true);
  }
}
