package repo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ReadRepoContents(repo string) []string {
	files, mdErr := getAllMdFilesInPath(repo)
	if mdErr != nil {
		log.Fatalln("Error during markdown files parsing.")
	}

	contentFiles := make([]string, 0)
	for _, f := range files {
		readFile, _ := os.ReadFile(f)
		contentFiles = append(contentFiles, string(readFile))
	}
	return contentFiles
}

func GetGitRepository(repoUrl string, overwrite bool) {
	projectPath := viper.GetString("targetPath") + "/" + getSlug(repoUrl)
	fmt.Printf("Target repository clone path: %s\n", projectPath)
	pathExists, pathErr := pathExists(projectPath)
	cobra.CheckErr(pathErr)

	if pathExists {
		if overwrite {
			err := os.RemoveAll(projectPath)
			cobra.CheckErr(err)
			syncRepo(repoUrl, projectPath)
		} else {
			fmt.Printf("Repository already exists.")
		}
	} else {
		syncRepo(repoUrl, projectPath)
	}
}

func syncRepo(repoUrl string, projectPath string) {
	cloneErr := CloneToFilesystem(projectPath, validateRepoUrl(repoUrl))
	if cloneErr != nil {
		log.Fatalf("Error during git clone. Path: %+x\n", projectPath)
	}
}

func validateRepoUrl(repoUrl string) string {
	if !strings.HasPrefix(repoUrl, "https://") {
		repoUrl = "https://" + repoUrl
	}
	fmt.Printf("Target repository url: %s\n", repoUrl)
	return repoUrl
}

func CloneToFilesystem(path, url string) error {
	if path == "" {
		path = "./temp"
	}
	if url == "" {
		url = "https://github.com/am8850/acli.git"
	}
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if path == "./temp" {
		os.RemoveAll(path)
	}
	return err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getAllMdFilesInPath(projectName string) ([]string, error) {
	absoluteContentPath := viper.GetString("targetPath") + "/" + projectName + viper.GetString("relativePath")
	return walkMatch(absoluteContentPath, "*.md")
}

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func getSlug(s string) string {
	slices := strings.Split(s, "/")
	return slices[len(slices)-1]
}
