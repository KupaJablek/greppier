package main

import (
	"fmt"
	"os"
)

func enqueue(queue []string, elem string) []string {
	queue = append(queue, elem)
	return queue
}

func dequeue(queue []string) (string, []string) {
	if len(queue) == 0 {
		return "", []string{}
	}

	temp := queue[0]
	if len(queue) == 1 {
		return temp, []string{}
	}

	return temp, queue[1:]
}

func fileGrep(f string) {

	fmt.Printf("\tFile: %s\n", f)
}

func main() {
	args := os.Args[1:]

	for _, a := range args {
		fmt.Println(a)
	}

	var work = make([]string, 0)

	work = enqueue(work, args[0])

	for len(work) > 0 {
		var job string
		job, work = dequeue(work)

		stat, err := os.Stat(job)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if stat.IsDir() {
			fmt.Printf("Dir: %s\n", job)
			// iterate through and add items to work
			files, err := os.ReadDir(job)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			for _, f := range files {
				work = enqueue(work, job + "/" + f.Name())
			}
			continue
		}

		// job is file
		fileGrep(job)
	}
}
