package repo

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Temp_folder_name = "./temp_acli"
	Default_git_repo = "https://github.com/am8850/acli.git"
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
		path = Temp_folder_name
	}
	if url == "" {
		url = Default_git_repo
	}
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
		//Progress: os.Stdout,
	})
	// if path == temp_folder_name {
	// 	os.RemoveAll(path)
	// }
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

func CopyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func CheckDir(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true // Directory exists
	}
	if os.IsNotExist(err) {
		return false // Directory does not exist
	}
	return false // Error occurred while checking directory existence
}

func CopyFolder(project, template string) error {
	src := Temp_folder_name + "/templates/" + template + "/"
	dst := "./" + project + "/"

	if !CheckDir(src) || !CheckDir(dst) {
		fmt.Println("Source: ", src)
		fmt.Println("Destination: ", dst)
		log.Fatalln("The target or the source directory does not exist.")
	}

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			err := os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func RemoveTempFolder() {
	// Always delete the temp folder
	os.RemoveAll(Temp_folder_name)
}
