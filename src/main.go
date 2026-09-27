package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	m := map[string]string{
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

	if len(os.Args) <= 1 {
		log.Fatal("Directory not specified")
	}
	nameDir := os.Args[1]

	openDir, errReadDir := os.ReadDir(nameDir)
	if errReadDir != nil {
		log.Fatal(errReadDir)
	}

	for _, dirEntry := range openDir {

		if dirEntry.IsDir() == false {

			nameFile := dirEntry.Name()
			extFile, ok := m[filepath.Ext(nameFile)]

			if !ok {
				extFile = m[""]
			}

			fmt.Println(nameFile, " -> ", extFile)

			p := filepath.Join(nameDir, extFile)
			errMkdir := os.MkdirAll(p, 0755)
			if errMkdir != nil {
				log.Println(errMkdir)
				continue
			}

			oldPath := filepath.Join(nameDir, nameFile)
			newPath := filepath.Join(nameDir, extFile, nameFile)

			errRename := os.Rename(oldPath, newPath)
			if errRename != nil {
				log.Println(errRename)
			}
		}
	}

}
