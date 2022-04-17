package main

import (
	"fmt"
	"github.com/google/uuid"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	new_notes_dir := "/home/abdullah/projects/notes_v2"
	old_notes_dir := "/home/abdullah/projects/notes"

	exclude := map[string]bool{"home": true, "abdullah": true, "projects": true, "notes": true, "note": true, "": true}

	filepath.WalkDir(old_notes_dir, func(path string, d fs.DirEntry, err error) error {

		fmt.Println("processing path: ", path)

		if d.IsDir() {
			if strings.Contains(path, ".git") {
				fmt.Println("path contains .git, skipping this directory")
				return fs.SkipDir
			}

			return nil
		}


		tags := strings.Split(path, "/")
		filteredTags := filter(tags, func(s string) bool {
			return exclude[s] == false
		})
		readFile, readErr := os.ReadFile(path)
		if readErr != nil {
			fmt.Println("Could not open file")
			return nil
		}
		content := string(readFile)
		author := os.Getenv("USER")

		marshalled := Marshal(author, filteredTags, content)
		newNoteId := uuid.New().String() + ".md"

		fullFilePath := filepath.Join(new_notes_dir, newNoteId)

		createdFile, createErr := os.Create(fullFilePath)
		if createErr != nil {
			fmt.Println("Could not create file: ", createErr)
			return nil
		}
		defer createdFile.Close()

		_, writeErr := createdFile.WriteString(marshalled)
		if writeErr != nil {
			fmt.Println("Could not write to file: ", writeErr)
			return nil
		}

		return nil

	})
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

const NoteTemplate = `<!-- Metadata start: do not remove, this section is managed by notes app -->
[_metadata_:author]:# "<author>"
[_metadata_:tags]:# "<tags>"
<!-- Metadata end -->

`

func Marshal(author string, tags []string, content string) string {
	withAuthor := strings.ReplaceAll(NoteTemplate, "<author>", author)
	sort.Strings(tags)
	withTags := strings.ReplaceAll(withAuthor, "<tags>", strings.Join(tags, ","))
	withContent := withTags + content

	return withContent
}
