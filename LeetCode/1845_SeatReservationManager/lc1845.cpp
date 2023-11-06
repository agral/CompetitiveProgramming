#include <cassert>
#include <queue>

class SeatManager {
public:
  SeatManager(int n) {
    // Clear the queue:
    while (!m_queue.empty()) {
      m_queue.pop();
    }
    // Then populate the queue with numbers (seats) from 1 to n:
    for (int i = 1; i <= n; i++) {
      m_queue.push(i);
    }
  }

  int reserve() {
    int ans = m_queue.top();
    m_queue.pop();
    return ans;
  }

  void unreserve(int seatNumber) {
    m_queue.push(seatNumber);
  }

private:
  std::priority_queue<int, std::vector<int>, std::greater<int>> m_queue;
};

int main() {
  {
    SeatManager sm1{5};
    assert(sm1.reserve() == 1); // All seats available, so 1 - the lowest number - should be returned.
    assert(sm1.reserve() == 2);
    sm1.unreserve(2);
    assert(sm1.reserve() == 2);
    assert(sm1.reserve() == 3);
    assert(sm1.reserve() == 4);
    assert(sm1.reserve() == 5);
    sm1.unreserve(5);
    assert(sm1.reserve() == 5);
    sm1.unreserve(5);
    sm1.unreserve(3);
    assert(sm1.reserve() == 3);
    assert(sm1.reserve() == 5);
  }
}
