#include <cassert>
#include <set> // orders elements automatically
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

class FoodRatings {
public:
  FoodRatings(std::vector<std::string>& foods,
              std::vector<std::string>& cuisines,
              std::vector<int>&         ratings)
  {
    for (int i = 0; i < foods.size(); ++i) {
      m_foodToRating[foods[i]] = ratings[i];
      m_foodToCuisine[foods[i]] = cuisines[i];
      m_cuisineToData[cuisines[i]].insert({-ratings[i], foods[i]});
    }
  }

  void changeRating(std::string food, int newRating) {
    const int oldRating = m_foodToRating[food];
    const std::string cuisine = m_foodToCuisine[food];
    auto& fullCuisineData = m_cuisineToData[cuisine];

    fullCuisineData.erase({-oldRating, food});
    fullCuisineData.insert({-newRating, food});
    m_foodToRating[food] = newRating;
  }

  std::string highestRated(std::string cuisine) {
    //return m_cuisineToData[cuisine].rbegin()->second; //rbegin() points to the last element
    // Note: rbegin() approach would get ties wrong. Switched to using negative ratings and begin().
    return m_cuisineToData[cuisine].begin()->second;
  }

private:
  std::unordered_map<std::string, int> m_foodToRating;
  std::unordered_map<std::string, std::string> m_foodToCuisine;
  // note: has to be <rating, name> instead of <name, rating> so that automatic sorting works as required.
  std::unordered_map<std::string, std::set<std::pair<int, std::string>>> m_cuisineToData;
};

int main() {
  std::vector<std::string>    foods = {"kimchi",     "miso",    "sushi",  "moussaka",    "ramen", "bulgogi"};
  std::vector<std::string> cuisines = {"korean", "japanese", "japanese",     "greek", "japanese",  "korean"};
  std::vector<int>          ratings = {       9,         12,          8,          15,         14,         7};

  FoodRatings fr(foods, cuisines, ratings);
  assert(fr.highestRated("korean") == "kimchi");
  assert(fr.highestRated("japanese") == "ramen");
  fr.changeRating("sushi", 16);
  assert(fr.highestRated("japanese") == "sushi");
  fr.changeRating("ramen", 16);
  assert(fr.highestRated("japanese") == "ramen"); // lexicographically smaller than "sushi".
}
