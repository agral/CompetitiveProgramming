#include <set>

class SmallestInfiniteSet {
private:
    int m_nextNum = 1;
    std::set<int> m_added;

public:
    int popSmallest() {
        if (m_added.empty()) {
            return m_nextNum++;
        } else {
            int ans = *m_added.begin();
            m_added.erase(m_added.begin());
            return ans;
        }
    }

    void addBack(int num) {
        if (num < m_nextNum) {
            m_added.insert(num);
        }
    }
};
