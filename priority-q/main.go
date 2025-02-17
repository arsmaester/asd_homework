package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	priority int
	duration int
	index    int
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}
func (pq *PriorityQueue) Push(x interface{}) {
	task := x.(*Task)
	task.index = len(*pq)
	*pq = append(*pq, task)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[0 : n-1]
	return task
}

func worker(pq *PriorityQueue) {
	for pq.Len() > 0 {
		task := heap.Pop(pq).(*Task)
		fmt.Printf("Processing task with priority %d for %d seconds\n", task.priority, task.duration)
		time.Sleep(time.Duration(task.duration) * time.Second)
	}
	fmt.Println("All tasks completed.")
}

func readTasks(filename string, pq *PriorityQueue) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			continue
		}
		priority, _ := strconv.Atoi(parts[0])
		duration, _ := strconv.Atoi(parts[1])
		heap.Push(pq, &Task{priority: priority, duration: duration})
	}
}

func main() {
	pq := &PriorityQueue{}
	heap.Init(pq)
	readTasks("tasks.txt", pq)
	worker(pq)
}
