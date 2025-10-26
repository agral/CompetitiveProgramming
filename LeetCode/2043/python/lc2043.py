from typing import List

# Subjective level: easy
class Bank:
    # Runtime complexity: O(n)
    # Auxiliary space complexity: O(n)
    def __init__(self, balance: List[int]):
        self.accounts = dict()
        for i in range(len(balance)):
            self.accounts[i+1] = balance[i]
        print(self.accounts)

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def transfer(self, fromAccount: int, toAccount: int, amount: int):
        if fromAccount in self.accounts.keys() and toAccount in self.accounts.keys():
            if self.accounts[fromAccount] >= amount:
                self.accounts[fromAccount] -= amount
                self.accounts[toAccount] += amount
                return True
        return False

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def deposit(self, account: int, amount: int):
        if account in self.accounts.keys():
            self.accounts[account] += amount
            return True
        return False

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def withdraw(self, account: int, amount: int):
        if account in self.accounts.keys():
            if self.accounts[account] >= amount:
                self.accounts[account] -= amount
                return True
        return False
