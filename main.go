package main

import (
	"fmt"
	"os"
	"regexp"
)

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

func fileGrep(f string, regExpr regexp.Regexp) {

	fmt.Printf("\tFile: %s\n", f)

	result := regExpr.FindAllString(f, -1)
	if result != nil {
		fmt.Printf("File %s matched regex\n", f)
	}
}

func main() {
	args := os.Args[1:]

	reg, err := regexp.Compile(args[1])
	if err != nil {
		fmt.Printf("ERROR with regular expression: %s", err.Error())
		return
	}

	var work = make([]string, 0)

	// push second arg, which should be dir/file, to work queue
	work = append(work, args[1])

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
				work = append(work, job+"/"+f.Name())
			}
			continue
		}

		// job is file
		fileGrep(job, *reg)
	}
}
