//go:build solution

// Package roundrobin shares tasks among threads by turns, like a scheduler.
package roundrobin

// RoundRobin deals the tasks among threads in turns and returns the list
// each thread gets.
func RoundRobin(tasks []string, threads int) [][]string {
	queues := make([][]string, threads)
	for i := range queues {
		queues[i] = []string{}
	}
	for i, task := range tasks {
		thread := i % threads
		queues[thread] = append(queues[thread], task)
	}
	return queues
}
