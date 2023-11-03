#include <cassert>
#include <memory>
#include <string>
#include <unordered_map>

struct TrieNode {
  TrieNode(): isLeafNode(false), next() {}
  bool isLeafNode;
  std::unordered_map<char, std::unique_ptr<TrieNode>> next;
};

class Trie {
public:
  Trie() {
    m_root = std::make_unique<TrieNode>();
  }

  void insert(std::string word) {
    TrieNode* tn = m_root.get();
    for (const auto& ch: word) {
      if (tn->next.find(ch) == tn->next.end()) {
        // Insert new node to the trie:
        std::unique_ptr<TrieNode> new_node = std::make_unique<TrieNode>();
        tn->next.insert({ch, std::move(new_node)});
      }

      tn = tn->next[ch].get();
    }
    tn->isLeafNode = true;
  }

  bool search(std::string word) {
    TrieNode* tn = m_root.get();
    for (const auto& ch: word) {
      if (tn->next.find(ch) == tn->next.end()) {
        return false;
      }
      tn = tn->next[ch].get();
    }
    return tn->isLeafNode;
  }

  // Implementation same as search, except for the last case - don't care whether it's leaf node.
  bool startsWith(std::string prefix) {
    TrieNode* tn = m_root.get();
    for (const auto& ch: prefix) {
      if (tn->next.find(ch) == tn->next.end()) {
        return false;
      }
      tn = tn->next[ch].get();
    }
    return true;
  }

private:
  std::unique_ptr<TrieNode> m_root;
};

int main() {
  Trie trie;
  trie.insert("apple");
  assert(trie.search("apple") == true);
  assert(trie.search("app") == false);
  assert(trie.startsWith("app") == true);
  trie.insert("app");
  assert(trie.search("app") == true);
}
