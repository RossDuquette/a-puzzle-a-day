package verifier

import (
	"fmt"
	"os"
)

import (
	"a-puzzle-a-day/internal/solver"
)

func CheckUniqueness(month string, day string) {
	solutionDir := solver.GetSolutionDir(month, day)
	stat, err := os.Stat(solutionDir)
	if err != nil {
		fmt.Println(err)
		return
	} else if !stat.IsDir() {
		fmt.Println(solutionDir, "is not a dir")
	}

	if areAllFilesUnique(solutionDir) {
		fmt.Println("All files in", solutionDir, "are unique")
	} else {
		fmt.Println("Duplicate files found in", solutionDir)
	}
}

func areAllFilesUnique(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read directory:", err)
		return false
	}

	uniqueFiles := make(map[string]struct{})
	for _, file := range files {
		filename := fmt.Sprintf("%s/%s", dir, file.Name())
		contents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Failed to read file", filename, ":", err)
			return false
		}
		uniqueFiles[string(contents)] = struct{}{}
	}

	return len(uniqueFiles) == len(files)
}
