"""
Name:          ans_iterative.py
Description:   Calculates in how many ways can a given value be realized
               with the use of an infinite number of coins of given values.
Options:       None
Created on:    27.12.2017
Last modified: 23.10.2019
Author:        Adam Graliński (adam@gralin.ski)
License:       MIT
"""


def count_ways(target_value, coins):
    """Returns number of ways in which a given `target_value` may be realized
    with the use of an infinite number of `available_coins` of given values."""
    dp_table = [[0 for column in range(len(coins))]
                for row in range(target_value+1)]

    # Fills the entries for the base case [target_value == 0]:
    for column in range(len(coins)):
        dp_table[0][column] = 1

    # Fills the rest of the table:
    for row in range(1, target_value+1):
        for c, _ in enumerate(coins):
            # Counts the number of distinct solutions including c'th coin:
            inclusive = dp_table[row-coins[c]][c] if row-coins[c] >= 0 else 0

            # Counts the number of distinct solutions excluding c'th coin:
            exclusive = dp_table[row][c-1] if c > 0 else 0

            dp_table[row][c] = inclusive + exclusive

    return dp_table[target_value][len(coins)-1]


def main():
    """A driver function to answer the challenge."""
    target_value, _ = [int(k) for k in input().split(" ")]
    available_coins = [int(v) for v in input().split(" ")]
    print(count_ways(target_value, available_coins))


if __name__ == "__main__":
    main()
