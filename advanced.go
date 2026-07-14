type Task struct {
	id int
	load int
}

type Node struct {
	id int
}

type Result struct {
	taskID int
	nodeID int
	timeTaken int
}

func processTask(n Node, t Task) Result {
	start := time.Now()
	time.Sleep(time.Millisecond * time.Duration(40+n.id*3))
	elapsed := time.Since(start).Milliseconds()
	return Result{taskID: t.id, nodeID: n.id, timeTaken: int(elapsed)}
}

func dispatcher(tasks []Task, nodes []Node, results chan<- Result, wg *sync.WaitGroup) {
	for _, t := range tasks {
		for _, n := range nodes {
			wg.Add(1)
			go func(node Node, task Task) {
				defer wg.Done()
				res := processTask(node, task)
				results <- res
			}(n, t)
		}
	}
}

func aggregate(results <-chan Result, done chan<- bool) {
	count := 0
	total := 0
	for r := range results {
		fmt.Println("Task", r.taskID, "Node", r.nodeID, "Time", r.timeTaken)
		total += r.timeTaken
		count++
	}
	if count > 0 {
		fmt.Println("Average Time:", total/count)
	}
	done <- true
}

func createTasks(n int) []Task {
	list := make([]Task, n)
	for i := 0; i < n; i++ {
		list[i] = Task{id: i + 1, load: (i + 1) * 5}
	}
	return list
}

func createNodes(n int) []Node {
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{id: i + 1}
	}
	return nodes
}

func main() {
	numNodes := 5
	numTasks := 6

	tasks := createTasks(numTasks)
	nodes := createNodes(numNodes)

	results := make(chan Result, numTasks*numNodes)
	done := make(chan bool)

	var wg sync.WaitGroup

	go aggregate(results, done)

	dispatcher(tasks, nodes, results, &wg)

	wg.Wait()
	close(results)

	<-done
}
