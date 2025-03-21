#include <iostream>
#include <queue>
#include <sstream>
#include <unordered_map>
#include <unordered_set>
#include <vector>

// Runtime complexity: O( |V| + |E| )
// Auxiliary space complexity: O( |V| + |E| )
class Solution {
public:
    std::vector<std::string> findAllRecipes(std::vector<std::string>& recipes,
                                            std::vector<std::vector<std::string>>& ingredients,
                                            std::vector<std::string>& supplies) {

        std::unordered_map<std::string, std::vector<std::string>> graph;
        std::unordered_map<std::string, int> indegrees;
        std::unordered_set<std::string> setSupplies(supplies.begin(), supplies.end());

        for (int i = 0; i < recipes.size(); i++) {
            for (const std::string& ing: ingredients[i]) {
                if (!setSupplies.contains(ing)) {
                    graph[ing].push_back(recipes[i]);
                    indegrees[recipes[i]] += 1;
                }
            }
        }

        std::vector<std::string> ans = {};
        std::queue<std::string> q;
        // Topological sort:
        for (const std::string& recipe: recipes) {
            if (!indegrees.contains(recipe)) {
                q.push(recipe);
            }
        }
        while (!q.empty()) {
            std::string f = q.front();
            q.pop();
            ans.push_back(f);
            if (graph.contains(f)) {
                for (std::string& v: graph[f]) {
                    indegrees[v] -= 1;
                    if (indegrees[v] == 0) {
                        q.push(v);
                    }
                }
            }
        }

        return ans;
    }
};

template<typename T>
std::string vectorToString(std::vector<T> vec) {
    std::string separator = "";
    std::stringstream ss;
    ss << "[";
    for (const auto& elem: vec) {
        ss << separator << elem;
        separator = ", ";
    }
    ss << "]";
    return ss.str();
}

int main() {
    struct Testcase {
        std::vector<std::string> recipes;
        std::vector<std::vector<std::string>> ingredients;
        std::vector<std::string> supplies;
        std::vector<std::string> expected;
    };
    Testcase testcases[] = {
        {
            /*recipes=*/{"bread"},
            /*ingredients=*/{{"yeast", "flour"}},
            /*supplies=*/{"yeast", "flour", "corn"},
            /*expected=*/{"bread"},
        },
        {
            /*recipes=*/{"bread", "sandwich"},
            /*ingredients=*/{{"yeast", "flour"}, {"bread", "meat"}},
            /*supplies=*/{"yeast", "flour", "meat"},
            /*expected=*/{"bread", "sandwich"},
        },
        {
            /*recipes=*/{"bread", "sandwich", "burger"},
            /*ingredients=*/{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
            /*supplies=*/{"yeast", "flour", "meat"},
            /*expected=*/{"bread", "sandwich", "burger"},
        },
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.findAllRecipes(tc.recipes, tc.ingredients, tc.supplies);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << vectorToString(actual)
                << ", want: " << vectorToString(tc.expected) << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
