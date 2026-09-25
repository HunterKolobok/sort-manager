package main

import(
	"fmt"
	"path/filepath"
	"os"
)

func main() {

	m := map[string]string{
			".jpg": "images",
			".png": "images",
			".jpeg": "images",
			".webp": "images",

			".mp4": "videos",
			".avi": "videos",
			".mkv": "videos",
			".mov": "videos",

			".mp3": "music",
			".aac": "music",
			".ogg": "music",
			".wma": "music",

			".docx": "documents",
			".doc": "documents",
			".odt": "documents",
			".pdf": "documents",
			".xlsx": "documents",
			".xls": "documents",
			".txt": "documents",

			".gz": "archive",
			".zip": "archive",
			".rar": "archive",

			".gitignore": "prog",

			"": "other",
			}

	structFile, _ := os.ReadDir("./test")

	for _, structik := range structFile {
		if structik.IsDir() == false {
			nameFile := structik.Name()
			extFile := m[filepath.Ext(nameFile)]
			fmt.Println(nameFile, " -> ", extFile )
		}
	}
}
