package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func(string) string

type TaskResult struct {
	TaskName    string
	Result      string
	Duration    time.Duration
	CompletedAt time.Time
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks []struct {
		name string
		task Task
		arg  string
	}
}


func (s *Scheduler) AddTask(name string, task Task, arg string) {
	s.tasks = append(s.tasks, struct {
		name string
		task Task
		arg  string
	}{name, task, arg})
}

func (s *Scheduler) Run() []TaskResult {
	var wg sync.WaitGroup
	results := make([]TaskResult, len(s.tasks))
	
	// 为每个任务启动一个协程
	for i, t := range s.tasks {
		wg.Add(1)
		fmt.Println("for ------------------------start", i)
		go func(index int, name string, task Task, arg string) {
			fmt.Println("go ------------------------start", i)
			defer wg.Done()
						
			startTime := time.Now()						
			result := task(arg)						
			duration := time.Since(startTime)
						
			results[index] = TaskResult{
				TaskName:    name,
				Result:      result,
				Duration:    duration,
				CompletedAt: time.Now(),
			}
			fmt.Println("go ------------------------end", i)
		}(i, t.name, t.task, t.arg)
		fmt.Println("for ------------------------end", i)
	}
		
	wg.Wait()
	return results
}


func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]struct {
			name string
			task Task
			arg  string
		}, 0),
	}
}


func main() {
	
	scheduler := NewScheduler()
		
	scheduler.AddTask("数据处理", processData, "dataset.csv")	
	scheduler.AddTask("计算任务", computeResult, "1000000")
		
	fmt.Println("开始执行所有任务...")
	start := time.Now()
	results := scheduler.Run()
	totalDuration := time.Since(start)
	
	// 输出结果
	fmt.Printf("\n所有任务完成，总耗时: %v\n", totalDuration)
	fmt.Println("------------------------")
	for _, res := range results {
		fmt.Printf("任务: %s\n", res.TaskName)
		fmt.Printf("  结果: %s\n", res.Result)
		fmt.Printf("  耗时: %v\n", res.Duration)
		fmt.Printf("  完成时间: %s\n", res.CompletedAt.Format("15:04:05.000"))
		fmt.Println("------------------------")
	}
}

func processData(filename string) string {
	time.Sleep(1500 * time.Millisecond)
	return fmt.Sprintf("已处理文件: %s，记录数: 10000", filename)
}


// 示例任务3：计算任务
func computeResult(iterations string) string {
	time.Sleep(1000 * time.Millisecond)
	return fmt.Sprintf("计算完成，迭代次数: %s，结果: 42", iterations)
}
