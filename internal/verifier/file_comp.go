package verifier

import (
	"fmt"
	"os"
)

import (
	"a-puzzle-a-day/internal/solver"
)

func CheckUniqueness(month string, day string) {
	solution_dir := solver.GetSolutionDir(month, day)
	stat, err := os.Stat(solution_dir)
	if err != nil {
		fmt.Println(err)
		return
	} else if !stat.IsDir() {
		fmt.Println(solution_dir, "is not a dir")
	}

	if are_all_files_unique(solution_dir) {
		fmt.Println("All files in", solution_dir, "are unique")
	} else {
		fmt.Println("Duplicate files found in", solution_dir)
	}
}

func are_all_files_unique(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read directory:", err)
		return false
	}

	unique_files := make(map[string]struct{})
	for _, file := range files {
		filename := fmt.Sprintf("%s/%s", dir, file.Name())
		contents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Failed to read file", filename, ":", err)
			return false
		}
		unique_files[string(contents)] = struct{}{}
	}

	return len(unique_files) == len(files)
}
