package lc0649

type Queue []int

func (q Queue) Enqueue(val int) Queue {
	return append(q, val)
}

func (q Queue) Dequeue() (Queue, int) {
	sz := len(q)
	if sz == 0 {
		panic("Just tried to Pop() from an empty Queue.")
	}
	return q[1:], q[0]
}

func predictPartyVictory(senate string) string {
	sz := len(senate)
	qD := make(Queue, 0)
	qR := make(Queue, 0)
	for i, ch := range senate {
		if ch == 'D' {
			qD = qD.Enqueue(i)
		} else {
			qR = qR.Enqueue(i)
		}
	}
	var numSenatorD, numSenatorR int
	for len(qD) > 0 && len(qR) > 0 {
		qD, numSenatorD = qD.Dequeue()
		qR, numSenatorR = qR.Dequeue()
		// The senator from a party that comes first (lower index)
		// votes the lowest senator from the opposing party out of the senate building.
		if numSenatorD < numSenatorR {
			// in effect both senators are destroyed; but the one voting returns with the index
			// greater than all currently existing indices.
			// Using their own index + initial senate size does the job.
			qD = qD.Enqueue(numSenatorD + sz)
		} else {
			qR = qR.Enqueue(numSenatorR + sz)
		}
	}

	if len(qD) == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
