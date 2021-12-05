package main

import (
	"flag"
	"fmt"

	"github.com/ThomVett/aoc2021/src/day_2"
	"github.com/ThomVett/aoc2021/src/day_3"
	"github.com/ThomVett/aoc2021/src/day_4"
	"github.com/ThomVett/aoc2021/version"
)

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	fileFlag := flag.String("input-file", "", "FilePath")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		return
	}

	day_3.NoOP()

	if len(*fileFlag) > 0 {
		fmt.Println(day_4.Part1(*fileFlag))
	} else {
		fmt.Println(version.BuildDate)
		fmt.Println(day_2.ComputeSolutionWithAim("data/day_2/day_2_test.txt"))
	}
}
