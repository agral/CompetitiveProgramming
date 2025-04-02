struct Node {
    int val;
    Node *left;
    Node *right;
    Node *next;
};

class Solution {
public:
    // Runtime complexity: O(n)
    // Auxiliary space complexity: O(h), where h is the height of the binary tree.
    Node* connect(Node* root) {
        if (root != nullptr) {
            doConnect(root->left, root->right);
        }
        return root;
    }

private:
    void doConnect(Node* lhs, Node* rhs) {
        if (lhs == nullptr) {
            return;
        }
        // Note: The tree is guaranteed to be perfect, i.e. a full one. So if lhs != nullptr,
        // then rhs is also guaranteed to not be a nullptr.

        // connect lhs to rhs:
        lhs->next = rhs;

        // recursively connect lhs and rhs descendants:
        doConnect(lhs->left, lhs->right);
        doConnect(lhs->right, rhs->left);
        doConnect(rhs->left, rhs->right);
    }
};
