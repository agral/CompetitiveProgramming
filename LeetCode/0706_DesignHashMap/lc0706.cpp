#include <cassert>
#include <utility>
#include <vector>

class MyHashMap {
public:
  const int BUCKET_SIZE = 1000;
  MyHashMap() {
  }

  void put(int key, int value) {
    auto h = hash(key);
    for (int i = 0; i < m_buckets[h].size(); i++) {
      if (m_buckets[h][i].first == key) {
        // key already present, update the found value and return:
        m_buckets[h][i].second = value;
        return;
      }
    }
    // key not present, push new key-value pair:
    m_buckets[h].push_back(std::make_pair(key, value));
  }

  int get(int key) {
    auto h = hash(key);
    for (int i = 0; i < m_buckets[h].size(); i++) {
      if (m_buckets[h][i].first == key) {
        return m_buckets[h][i].second;
      }
    }
    return -1;
  }

  void remove(int key) {
    auto h = hash(key);
    for (int i = 0; i < m_buckets[h].size(); i++) {
      if (m_buckets[h][i].first == key) {
        m_buckets[h].erase(m_buckets[h].begin() + i);
        return;
      }
    }
  }

private:
  int hash(int key) {
    return key % BUCKET_SIZE;
  }
  std::vector<std::pair<int, int>> m_buckets[1000] {};
};

int main() {
  MyHashMap hm;
  hm.put(1, 1);
  hm.put(2, 2);
  assert(hm.get(1) == 1); // the value stored at key 1 is 1.
  assert(hm.get(3) == -1); // no value stored under key 3, so -1 should be returned.
  assert(hm.get(2) == 2);
  hm.put(2, 1);
  assert(hm.get(2) == 1);
  hm.remove(2);
  assert(hm.get(2) == -1);
}
