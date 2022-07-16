package mysql

import (
	"fmt"
	log "go-database/src/util/log"
	stringUtil "go-database/src/util/string"
	"strconv"
	"sync"
	"time"
)

/*
	JobContainId -- 任务组编号
	Work -- Mysql任务
		Model -- 任务执行模式、执行参数:
			FULL_GOROUTINE 满协程模式：即所有任务 都各自使用一个协程，在任务未进入 执行队列 前堵塞，进行执行队列 后开始运行。
			REDUCE_GOROUTINE 节省协程模式：只有任务进入了 执行队列 才分配一个协程。主线程会 尝试向执行队列 中插入任务（任务选择是通过轮训），若队列满则会
	Jobs -- 任务组内所有任务集合
	IsAlive -- 任务正在执行队列
	IsEnd -- 任务已经结束队列
	waitGroup -- 等待被执行的任务组（这个同步量的主要目的在于，协调 协程与主线程之间的同步关系）
	Result -- 各任务获取结果集合队列


*/
type JobContain struct {
	JobContainId string
	Work         MysqlWork
	Jobs         []*Job
	IsAlive      chan string
	IsEnd        chan string
	waitGroup    sync.WaitGroup
	Result       interface{}
}

func NewJobContain() *JobContain {
	jobContain := new(JobContain)

	// 初始化 任务组 管道
	jobContain.Work = *Work
	// jobContain.IsAlive = make(chan string, jobContain.Work.ThreadNumber)
	// jobContain.IsEnd = make(chan string, len(jobContain.Jobs))
	jobContain.JobContainId = "x"
	jobContain.Jobs = make([]*Job, 20)
	jobContain.waitGroup.Add(len(jobContain.Jobs))
	return jobContain
}

type Job struct {
	jobId     string
	startTime string
	endTime   string
	Result    interface{}
	isEnd     bool
}

var jobContains []JobContain = make([]JobContain, 10)

func check(jobContain *JobContain) bool {
	log.PrintInfoLog("开始对配置文件中 work 参数进行校验")
	if !stringUtil.IsNotEmpty(jobContain.Work.SQL) {
		log.PrintErrorLog("配置文件中 work 参数中 缺少 SQL")
		return false
	}
	if !(jobContain.Work.ThreadNumber > 0 && jobContain.Work.ThreadNumber <= 20) {
		log.PrintErrorLog("配置文件中 work 参数中 ThreadNumber 数值 不在预定范围内")
		return false
	}
	return true
}

// 进行工作参数 的校验工作
func prepare(jobContain *JobContain) {

	for i := 0; i < 20; i++ {
		jobx := new(Job)
		jobx.jobId = strconv.Itoa(i)
		jobContain.Jobs[i] = jobx
	}

}

func start(job *Job, jobContain *JobContain) {

	if jobContain.Work.Model == "FULL_GOROUTINE" {
		jobContain.IsAlive <- job.jobId
	}

	fmt.Println("任务：" + job.jobId + " 开始执行")
	time.Sleep(20 * time.Second)

	jobContain.waitGroup.Done()

	// 任务完成，退出 正在执行任务 队列
	<-jobContain.IsAlive
	fmt.Println("任务：" + job.jobId + " 执行完毕")
}

func Exe() {
	jobContain := NewJobContain()
	prepare(jobContain)

	if jobContain.Work.Model == "REDUCE_GOROUTINE" {
		for i := 0; i != -1; i++ {
			if i == len(jobContain.Jobs) {
				i = 0
			}
			if len(jobContain.IsEnd) == len(jobContain.Jobs) {
				log.PrintInfoLog("任务组：" + jobContain.JobContainId + ",完成所有任务")
				return
			} else if jobContain.Jobs[i].isEnd == false {
				jobContain.IsAlive <- jobContain.Jobs[i].jobId
				go start(jobContain.Jobs[i], jobContain)
			}
		}
	} else if jobContain.Work.Model == "FULL_GOROUTINE" {
		for i := 0; i < len(jobContain.Jobs); i++ {
			go start(jobContain.Jobs[i], jobContain)
		}
		jobContain.waitGroup.Wait()
	}

}
