package main

func canCompleteCircuit(gas []int, cost []int) int {
    i := 0
    n := len(gas)
    for i < n {
        gasAmount, costAmount, current := 0, 0, 0
        for current < n {
            j := (i + current) % n
            gasAmount += gas[j]
            costAmount += cost[j]
            if gasAmount < costAmount {
                break;
            }
            current ++
        }
        if current == n {
            return i
        } else {
            i += current + 1
        }
    }
    return -1
}