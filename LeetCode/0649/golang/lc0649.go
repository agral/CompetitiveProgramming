package lc0649

import "fmt"

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("Dequeue() from an empty Queue.")
	}
	val := (*q)[0]
	(*q) = (*q)[1:]
	return val, nil
}

func predictPartyVictory(senate string) string {
	sz := len(senate)
	qD := make(Queue, 0)
	qR := make(Queue, 0)
	for i, ch := range senate {
		if ch == 'D' {
			qD.Enqueue(i)
		} else {
			qR.Enqueue(i)
		}
	}
	var numSenatorD, numSenatorR int
	for len(qD) > 0 && len(qR) > 0 {
		numSenatorD, _ = qD.Dequeue()
		numSenatorR, _ = qR.Dequeue()
		// The senator from a party that comes first (lower index)
		// votes the lowest senator from the opposing party out of the senate building.
		if numSenatorD < numSenatorR {
			// in effect both senators are destroyed; but the one voting returns with the index
			// greater than all currently existing indices.
			// Using their own index + initial senate size does the job.
			qD.Enqueue(numSenatorD + sz)
		} else {
			qR.Enqueue(numSenatorR + sz)
		}
	}

	if len(qD) == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
