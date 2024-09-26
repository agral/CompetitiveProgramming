#include <cassert>
#include <map>
#include <vector>

// Brute force approach || O(n^2) complexity
/*
class MyCalendar {
public:
    MyCalendar() {}
    bool book(int start, int end) {
        for (const auto& [s, e]: m_calendar) {
            if (std::max(s, start) < std::min(e, end)) {
                return false;
            }
        }
        m_calendar.emplace_back(start, end);
        return true;
    }

private:
    std::vector<std::pair<int, int>> m_calendar;
};
*/

// Ordered map approach || O(nlogn) complexity
class MyCalendar {
public:
    MyCalendar() {}
    bool book(int start, int end) {
        // lower_bound returns an iterator to the first element not less (>=) than a key.
        auto lbound = m_calendar.lower_bound(end);
        if (lbound == m_calendar.begin()  // case 1) no elements or every element is >= key
        || (--lbound)->second <= start) { // case 2) previous event ends before (or *right* before) this one starts
            m_calendar[start] = end;
            return true;
        }
        return false;
    }
private:
    std::map<int, int> m_calendar; // Maps calendar events' start times to their end times.
                                   // Keeps them sorted by start times.
};

int main() {
    MyCalendar c;
    assert(c.book(10, 20) == true);
    assert(c.book(15, 25) == false);
    assert(c.book(20, 30) == true);
}
