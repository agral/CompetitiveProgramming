#include <iostream>
#include <map>
#include <memory>

struct LRUNode {
    int key;
    int value;
    std::shared_ptr<LRUNode> prev;
    std::shared_ptr<LRUNode> next;

    LRUNode(int k, int v): key(k), value(v), prev(nullptr), next(nullptr) {}
};

// Subjective level: medium/hard
// Solved on: 2026-01-12
class LRUCache {
public:
    // Runtime complexity: O(1)
    // Auxiliary space complexity: O(capacity+2) at max cache occupancy; initially O(2).
    LRUCache(int capacity) {
        m_head = std::make_shared<LRUNode>(-1, -1);
        m_tail = std::make_shared<LRUNode>(-1, -1);
        link(m_head, m_tail);
        m_maxCapacity = capacity;
    }

    // Runtime complexity: O(1)
    // Auxiliary space complexity: O(1)
    int get(int key) {
        const auto it = m_keyToLRUNode.find(key);
        if (it == m_keyToLRUNode.cend()) {
            return -1;
        }
        // otherwise, key exists in cache. Move it to cache's head and return its value.
        auto node = it->second;
        remove(node);
        moveToFront(node);
        return node->value;
    }

    // Runtime complexity: O(1) (amortized)
    // Auxiliary space complexity: O(1)
    void put(int key, int value) {
        // If the key already exists in the cache, just update its value and put it at front, then return.
        const auto it = m_keyToLRUNode.find(key);
        if (it != m_keyToLRUNode.cend()) {
            auto node = it->second;
            node->value = value;
            remove(node);
            moveToFront(node);
            return;
        }

        // Otherwise, the key has to be added to cache.
        // If cache is already full, evict the LRU entry:
        if (m_keyToLRUNode.size() == m_maxCapacity) {
            auto lruNode = m_tail->prev;
            m_keyToLRUNode.erase(lruNode->key);
            remove(lruNode);
        }
        // Finally, add the entry to cache and put it at front.
        auto newNode = std::make_shared<LRUNode>(key, value);
        moveToFront(newNode);
        m_keyToLRUNode[key] = newNode;
    }

private:
    std::shared_ptr<LRUNode> m_head;
    std::shared_ptr<LRUNode> m_tail;
    std::map<int, std::shared_ptr<LRUNode>> m_keyToLRUNode;
    int m_maxCapacity;

    // Runtime complexity: O(1).
    // Auxiliary space complexity: O(1).
    void remove(std::shared_ptr<LRUNode> node) {
        link(node->prev, node->next);
    }

    // Runtime complexity: O(1).
    // Auxiliary space complexity: O(1).
    void moveToFront(std::shared_ptr<LRUNode> node) {
        link(node, m_head->next);
        link(m_head, node);
    }

    // Runtime complexity: O(1).
    // Auxiliary space complexity: O(1).
    void link(std::shared_ptr<LRUNode> preceding, std::shared_ptr<LRUNode> following) {
        preceding->next = following;
        following->prev = preceding;
    }
};

int main() {
    struct Testcase {
        int input;
        LRUCache expected;
    };

    int numBad = 0, rv = 0;
    LRUCache cache(2);
    cache.put(1, 1); // cache should be {1:1}
    cache.put(2, 2); // cache should be {2:2, 1:1}

    rv = cache.get(1); // cache should still be {2:2, 1:1}, so expect that value of 1 from key 1 is returned.
    if (rv != 1) {
        std::cout << "cache.get(1) answered wrongly. Expect: 1, got: " << rv << "\n";
        numBad++;
    }
    // note: after this operation, key 1 becomes the most recently used. So cache is now: {1:1, 2:2}.

    cache.put(3, 3); // LRU key was 2. This evicts key 2, and cache should be: {3:3, 1:1}.

    rv = cache.get(2); // should return -1, as key 2 was evicted in the step above.
    if (rv != -1) {
        std::cout << "cache.get(2) answered wrongly. Expect: -1, got: " << rv << "\n";
        numBad++;
    }

    cache.put(4, 4); // LRU key was 1, this evicts key 1, with resulting cache: {4:4, 3:3}.

    rv = cache.get(1); // should return -1, as key 1 was evicted when key 4 was added.
    if (rv != -1) {
        std::cout << "cache.get(1) answered wrongly. Expect: -1, got: " << rv << "\n";
        numBad++;
    }

    rv = cache.get(3); // should return 3, as key 3 is still in cache. Cache is now: {3:3, 4:4}.
    if (rv != 3) {
        std::cout << "cache.get(3) answered wrongly. Expect: 3, got: " << rv << "\n";
        numBad++;
    }

    rv = cache.get(4); // should return 4, as key 4 is still in cache. Cache is now: {4:4, 3:3}.
    if (rv != 4) {
        std::cout << "cache.get(4) answered wrongly. Expect: 4, got: " << rv << "\n";
        numBad++;
    }

    if (numBad == 0) {
        std::cout << "[OK] All tests passed successfully.\n";
    } else {
        std::cout << "[FAIL] " << numBad << " failures in total.\n";
    }
}
