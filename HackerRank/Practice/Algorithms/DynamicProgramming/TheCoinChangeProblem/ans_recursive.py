"""
Name:          ans_recursive.py
Description:   Calculates in how many ways can a given value be realized
               with the use of an infinite number of coins of given values.
Options:       None
Created on:    27.12.2017
Last modified: 23.10.2019
Author:        Adam Graliński (adam@gralin.ski)
License:       MIT
"""


def count_ways(target_value, available_coins, currently_considered=0):
    """Returns number of ways in which a given `target_value` may be realized
    with the use of an infinite number of `available_coins` of given values."""
    # If target == 0, then there is only one valid solution
    # (use 0 of each of the available coins):
    if target_value == 0:
        return 1

    # There is no way to obtain a negative sum, either:
    if target_value < 0:
        return 0

    # If there are no more coins to use, but the target value is not zero,
    # there is no way to obtain the target value:
    if currently_considered >= len(available_coins) and target_value != 0:
        return 0

    # The number of the ways is a sum of two simpler cases:
    # (1): solutions that do not use the first coin at all
    # (2): solutions that use the first coin at least once.
    return (count_ways(target_value, available_coins, currently_considered+1) +
            count_ways(target_value-available_coins[currently_considered],
                       available_coins, currently_considered))


def main():
    """A driver function to answer the challenge."""
    target, _ = [int(k) for k in input().split(' ')]
    print(target)
    denominations = [int(k) for k in input().split(' ')]
    print(count_ways(target, denominations))


if __name__ == "__main__":
    main()
