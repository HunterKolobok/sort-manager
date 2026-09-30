package sortDir

import(
	"fmt"
	"os"
	"path/filepath"
	"log"
)

var m = map[string]string{
		".jpg": "images", ".png": "images", ".jpeg": "images",
		".webp": "images", ".bmp": "images", ".gif": "images",
		".jif": "images", ".jfif": "images", ".jpe": "images",
		".heic": "images",

		".mp4": "videos", ".avi": "videos", ".mkv": "videos",
		".mov": "videos", ".webm": "videos",

		".mp3": "music", ".aac": "music", ".ogg": "music",
		".wma": "music",

		".docx": "documents", ".doc": "documents", ".odt": "documents",
		".pdf": "documents", ".xlsx": "documents", ".xls": "documents",
		".txt": "documents", ".md": "documents", ".csv": "documents",
		".ppt": "documents", ".pptx": "documents", ".odp": "documents",
		".epub": "documents", ".fb2": "documents", ".mobi": "documents",
		".djvu": "documents", ".azw3": "documents",

		".gz": "archive", ".zip": "archive", ".rar": "archive",
		".7z": "archive", ".tar": "archive", ".bin": "archive",
		".iso": "archive", ".xz": "archive",

		".gitignore": "prog", ".c": "prog", ".h": "prog",
		".cpp": "prog", ".py": "prog", ".go": "prog",
		".ipynb": "prog", ".sh": "prog", ".json": "prog",
		".yaml": "prog", ".sql": "prog",

		"": "other",
	}

func SortDirectory(nameDir string) error{

	openDir, errReadDir := os.ReadDir(nameDir)
	if errReadDir != nil {
		return errReadDir
	}

		for _, dirEntry := range openDir {

			if dirEntry.IsDir() == false {

				catFile, ok := getCategory(dirEntry)

				if !ok {
					catFile = m[""]
				}

				nameFile := dirEntry.Name()
				fmt.Println(nameFile, " -> ", catFile)

				if errCreateDir := createDirectory(nameDir, catFile); errCreateDir != nil {
					return errCreateDir
				}

				if errMove := moveObject(nameDir, nameFile, catFile); errMove != nil {
					log.Println(errMove)
				}

			}
		}

	return nil
}

func getCategory(dirEntry os.DirEntry) (string, bool){
	nameFile := dirEntry.Name()
	catFile, ok := m[filepath.Ext(nameFile)]
	return catFile, ok
}

func createDirectory(nameDir, catFile string) error {
	p := filepath.Join(nameDir, catFile)
	errMkdir := os.MkdirAll(p, 0755)

	return errMkdir
}

func moveObject(nameDir, nameFile, catFile string) error {
	oldPath := filepath.Join(nameDir, nameFile)
	newPath := filepath.Join(nameDir, catFile, nameFile)

	errRename := os.Rename(oldPath, newPath)

	return errRename
}
