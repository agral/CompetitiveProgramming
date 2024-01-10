#include <cassert>
#include <deque>

class RecentCounter {
public:
  int ping(int t) {
    q.push_back(t);
    while (q.front() < t - 3000) {
      q.pop_front();
    }
    return q.size();
  }

private:
  std::deque<int> q;
};

int main() {
  RecentCounter rc;
  assert(rc.ping(1) == 1);
  assert(rc.ping(100) == 2);
  assert(rc.ping(3001) == 3);
  assert(rc.ping(3002) == 3);
}
