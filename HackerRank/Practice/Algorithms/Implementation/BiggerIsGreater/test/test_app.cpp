/**
 * Name: test_app.cpp
 * Description: Contains unit tests for the challenge's answer.
 * Created on: 02.09.2019
 * Last modified: 02.09.2019
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <catch2/catch.hpp>
#include "../src/impl.hpp"

TEST_CASE("The program gives correct answer to all the queries from the sample input file(s)")
{
  SECTION("sample input 1")
  {
    CHECK(getNextLexicographicPermutation("lmno") == "lmon");
    CHECK(getNextLexicographicPermutation("dcba") == "no answer");
    CHECK(getNextLexicographicPermutation("dcbb") == "no answer");
    CHECK(getNextLexicographicPermutation("abdc") == "acbd");
    CHECK(getNextLexicographicPermutation("abcd") == "abdc");
    CHECK(getNextLexicographicPermutation("fedcabcd") == "fedcabdc");
  }

  SECTION("Empty string and one-letter string")
  {
    CHECK(getNextLexicographicPermutation("") == "no answer");
    CHECK(getNextLexicographicPermutation("a") == "no answer");
  }
}
