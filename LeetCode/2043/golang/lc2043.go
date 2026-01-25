package lc2043

// Subjective level: easy+
// Solved on: 2026-01-25

type Bank struct {
	Accounts map[int]int64
}

func Constructor(balances []int64) Bank {
	ans := Bank{}
	ans.Accounts = make(map[int]int64)
	for i, balance := range balances {
		// Note: balances is 0-indexed, but initially bank uses 1-indexed account numbers,
		// so if there are n entries, these have to be numbered from 1 to n.
		ans.Accounts[i+1] = balance
	}
	return ans
}

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func (this *Bank) Transfer(fromAccount int, toAccount int, amount int64) bool {
	_, fromAccountExists := this.Accounts[fromAccount]
	_, toAccountExists := this.Accounts[toAccount]
	if !(fromAccountExists && toAccountExists) {
		return false
	}
	if this.Accounts[fromAccount] < amount {
		return false
	}

	this.Accounts[fromAccount] -= amount
	this.Accounts[toAccount] += amount
	return true
}

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func (this *Bank) Deposit(account int, amount int64) bool {
	_, accountExists := this.Accounts[account]
	if !accountExists {
		return false
	}
	this.Accounts[account] += amount
	return true
}

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func (this *Bank) Withdraw(account int, amount int64) bool {
	_, accountExists := this.Accounts[account]
	if !accountExists {
		return false
	}
	if this.Accounts[account] < amount {
		return false
	}
	this.Accounts[account] -= amount
	return true
}
