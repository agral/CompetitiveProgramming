#include <iostream>
#include <list>
#include <unordered_map>

struct LFUNode {
    std::list<int>::const_iterator it;
    int key;
    int value;
    int frequency;
};

// Subjective level: medium/hard
// Solved on: 2026-02-01
class LFUCache {
private:
    int m_capacity;
    int m_minFrequency;
    std::unordered_map<int, LFUNode> m_keyToNode;
    std::unordered_map<int, std::list<int>> m_frequencyToKeys;

public:
    // Runtime complexity: O(1)
    // Auxiliary space complexity: O(1)
    LFUCache(int capacity) 
    : m_capacity(capacity)
    , m_minFrequency(0)
    {}

    // Runtime complexity: O(1) (amortized); up to O(n) in worst case.
    // Auxiliary space complexity: O(1)
    int get(int key) {
        const auto it = m_keyToNode.find(key);
        if (it != m_keyToNode.cend()) {
            LFUNode& node = it->second;
            tap(node);
            return node.value;
        }
        return -1;
    }

    // Runtime complexity: O(1) amortized; up to O(n) in worst case.
    // Auxiliary space complexity: O(n)
    void put(int key, int value) {
        // Check if the key is already present in cache; if so - just update & tap it.
        const auto it = m_keyToNode.find(key);
        if (it != m_keyToNode.cend()) {
            LFUNode& node = it->second;
            node.value = value;
            tap(node);
        } else {
            // Otherwise, check if the capacity is already at max:
            if (m_capacity == m_keyToNode.size()) {
                // if so, the LFU key has to be deleted:
                int keyToRemove = m_frequencyToKeys[m_minFrequency].back();
                m_frequencyToKeys[m_minFrequency].pop_back();
                m_keyToNode.erase(keyToRemove);
            }
            // Regardless of whether the above deletion took place or not,
            // put the new node in the front of the cache.
            m_minFrequency = 1; // the new entry has frequency 1, it is a minimum (maybe a new minimum).
            m_frequencyToKeys[1].push_front(key);
            m_keyToNode[key] = LFUNode{m_frequencyToKeys[1].cbegin(), key, value, 1};
        }
    }

private:
    void tap(LFUNode& node) {
        const int formerFreq = node.frequency;
        node.frequency += 1;

        m_frequencyToKeys[formerFreq].erase(node.it);
        if (m_frequencyToKeys[formerFreq].empty()) {
            m_frequencyToKeys.erase(formerFreq);
            if (m_minFrequency == formerFreq) {
                m_minFrequency += 1;
            }
        }

        // The new key always gets put in the front of the cache's list:
        m_frequencyToKeys[node.frequency].push_front(node.key);
        node.it = m_frequencyToKeys[node.frequency].cbegin();
    }
};

int main() {
    int numBad = 0, rv = 0;
    LFUCache lfu(2);

    lfu.put(1, 1); // lfu: {1, _}. cnt[1] = 1
    lfu.put(2, 2); // lfu: {2, 1}. cnt[2]=1, cnt[1]=1
    rv = lfu.get(1);
    if (rv != 1) {
        std::cout << "#1 lfu.get(1) answered wrongly. Expect: 1, got: " << rv << "\n";
        numBad++;
    }
    // After this operation: lfu={1,2}. cnt[1]=2, cnt[2]=1.

    lfu.put(3, 3);
    // After this operation: LFU entry was 2, so it got deleted. lfu={3,1}. cnt[3]=1, cnt[1]=2.
    rv = lfu.get(2);
    if (rv != -1) {
        std::cout << "#2 lfu.get(2) answered wrongly. Expect: -1, got: " << rv << "\n";
        numBad++;
    }
    rv = lfu.get(3);
    if (rv != 3) {
        std::cout << "#3 lfu.get(3) answered wrongly. Expect: 3, got: " << rv << "\n";
        numBad++;
    }
    // After this operation: lfu={3,1}. cnt[3]=2, cnt[1]=2.

    lfu.put(4, 4);
    // After this operation: lfu={4,3}. cnt[4]=1, cnt[3]=2.
    // Entries 1 and 3 had same use frequency, tiebreaking by least recently used (1) takes place.

    rv = lfu.get(1);
    if (rv != -1) {
        std::cout << "#4 lfu.get(1) answered wrongly. Expect: -1, got: " << rv << "\n";
        numBad++;
    }
    rv = lfu.get(3);
    if (rv != 3) {
        std::cout << "#5 lfu.get(3) answered wrongly. Expect: 3, got: " << rv << "\n";
        numBad++;
    }
    // After this operation: lfu={3,4}. cnt[3]=3, cnt[4]=1.

    rv = lfu.get(4);
    if (rv != 4) {
        std::cout << "#6 lfu.get(4) answered wrongly. Expect: 4, got: " << rv << "\n";
        numBad++;
    }
    // After this operation: lfu={4,3}. cnt[3]=3, cnt[4]=2.

    // End of testing; display the summary.
    if (numBad == 0) {
        std::cout << "[OK] All tests passed successfully.\n";
    } else {
        std::cout << "[FAIL] " << numBad << " failures in total.\n";
    }
}
