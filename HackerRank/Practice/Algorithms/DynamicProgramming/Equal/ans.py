"""
Name:          ans.py
Description:   Solves the "Equal" challenge from HackerRank:
               https://www.hackerrank.com/challenges/equal/problem
Options:       None
Created on:    26.12.2017
Last modified: 24.10.2019
Author:        Adam Graliński (adam@gralin.ski)
License:       MIT
"""


T = int(input())
for t in range(T):
    N = int(input())
    colleagues = [int(x) for x in input().split(" ")]
    minimum = min(colleagues)
    for n in range(N):
        colleagues[n] -= minimum
    colleagues1 = colleagues[:]
    colleagues2 = colleagues[:]
    operations0 = 0
    operations1 = 0
    operations2 = 0
    for n in range(N):
        colleagues1[n] += 1
        colleagues2[n] += 2

        while(colleagues[n] != 0):
            if colleagues[n] >= 5:
                operations0 += colleagues[n] // 5
                colleagues[n] = colleagues[n] % 5
            elif colleagues[n] >= 2:
                colleagues[n] -= 2
                operations0 += 1
            else:
                colleagues[n] -= 1
                operations0 += 1

        while(colleagues1[n] != 0):
            if colleagues1[n] >= 5:
                operations1 += colleagues1[n] // 5
                colleagues1[n] = colleagues1[n] % 5
            elif colleagues1[n] >= 2:
                colleagues1[n] -= 2
                operations1 += 1
            else:
                colleagues1[n] -= 1
                operations1 += 1

        while(colleagues2[n] != 0):
            if colleagues2[n] >= 5:
                operations2 += colleagues2[n] // 5
                colleagues2[n] = colleagues2[n] % 5
            elif colleagues2[n] >= 2:
                colleagues2[n] -= 2
                operations2 += 1
            else:
                colleagues2[n] -= 1
                operations2 += 1

    print(min(operations0, operations1, operations2))
