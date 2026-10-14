package main

import "fmt"

func main() {
	logs := []Ticket{
		{"Dwight", "close", 2}, {"Dwight", "open", 2},
		{"Drew", "open", 32}, {"Drew", "close", 32},
		{"Drew", "open", 32}, {"Drew", "close", 32},
		{"Susa", "open", 7}, {"Jo", "close", 7},
		{"Susa", "open", 33}, {"Jo", "open", 8},
		{"Jo", "open", 36}, {"Jo", "close", 8},
		{"Susa", "close", 33},
	}
	anomalies := actionLogAnomalies2(logs)
	fmt.Println(anomalies)
}

type Ticket struct {
	agent  string
	action string
	number int
}

type Action struct {
	name   string
	by     string
	ticket int
}

// the logs is sorted chronologically
// ticket doesn't have anomalies if:
// - it is opened and closed once, in that order
// - the opening and closing agent is the same
// - the agent didn't do any action for a different ticket between opening and closing the current ticket
func actionLogAnomalies2(logs []Ticket) []int {
	lines := map[int]int{}
	agentToPrevLine := map[string]Ticket{}
	openClose := map[int]bool{}
	for _, log := range logs {
		lines[log.number]++
		if log.action == "close" {
			openClose[log.number] = agentToPrevLine[log.agent] == Ticket{log.agent, "open", log.number}
		}
		agentToPrevLine[log.agent] = log
	}

	anomaliesArr := []int{}
	for ticketId, n := range lines {
		if n != 2 || !openClose[ticketId] {
			anomaliesArr = append(anomaliesArr, ticketId)
		}
	}

	return anomaliesArr
}

func actionLogAnomalies(logs []Ticket) []int {
	ticketToActions := map[int][]Action{}
	agentToActions := map[string][]Action{}
	anomalies := map[int]struct{}{}
	for _, log := range logs {
		ticketToActions[log.number] = append(ticketToActions[log.number],
			Action{
				name: log.action,
				by:   log.agent})

		agentToActions[log.agent] = append(agentToActions[log.agent], Action{
			name:   log.action,
			ticket: log.number,
		})
	}

	for ticket, actions := range ticketToActions {
		if len(actions) != 2 {
			anomalies[ticket] = struct{}{}
			continue
		}

		if actions[0].name != "open" || actions[1].name != "close" {
			anomalies[ticket] = struct{}{}
			continue
		}

		if actions[0].by != actions[1].by {
			anomalies[ticket] = struct{}{}
			continue
		}
	}

	// fmt.Println(ticketToActions)
	fmt.Println(agentToActions)

	for _, actions := range agentToActions {
		for idx, action := range actions {
			if action.name == "close" && idx != 0 {
				if actions[idx-1].name != "open" || actions[idx-1].ticket != actions[idx].ticket {
					anomalies[action.ticket] = struct{}{}
				}
			}

		}
		fmt.Println()
	}

	anomaliesArr := []int{}
	for anomalie := range anomalies {
		anomaliesArr = append(anomaliesArr, anomalie)
	}

	// fmt.Println(agentToActions)

	return anomaliesArr
}
