type Task struct {
	id int
	data int
}

type Result struct {
	taskID int
	output int
}

type Worker struct {
	id int
}

func (w Worker) execute(task Task) Result {
	time.Sleep(time.Millisecond * time.Duration(50+w.id*5))
	return Result{taskID: task.id, output: task.data * 2}
}

func workerPool(id int, tasks < chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	w := Worker{id: id}
	for task := range tasks {
		res := w.execute(task)
		results <- res
	}
}

func generateTasks(n int) []Task {
	tasks := make([]Task, n)
	for i := 0; i < n; i++ {
		tasks[i] = Task{id: i + 1, data: (i + 1) * 10}
	}
	return tasks
}

func collectResults(results <-chan Result, done chan<- bool) {
	for res := range results {
		fmt.Println("Task", res.taskID, "Output", res.output)
	}
	done <- true
}

func main() {
	numWorkers := 6
	numTasks := 12

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	var wg sync.WaitGroup
	done := make(chan bool)

	go collectResults(results, done)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerPool(i, tasks, results, &wg)
	}

	taskList := generateTasks(numTasks)

	for _, t := range taskList {
		tasks <- t
	}
	close(tasks)

	wg.Wait()
	close(results)

	<-done
} 
