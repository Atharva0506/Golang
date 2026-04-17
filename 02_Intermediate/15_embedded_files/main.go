package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

// 1. Embed a single file into a []byte variable.
// The //go:embed directive MUST be on the line immediately before the variable declaration.
// The path is relative to the source file.
//
//go:embed static/greeting.txt
var greetingBytes []byte

// 2. Embed a single file into a string variable.
//
//go:embed static/greeting.txt
var greetingString string

// 3. Embed an entire directory into an embed.FS.
// embed.FS is a read-only, virtual file system — perfect for serving multiple static assets.
//
//go:embed static
var staticFiles embed.FS

func main() {
	// 1. Use the embedded []byte directly
	fmt.Println("--- As []byte ---")
	fmt.Printf("Length: %d bytes\n", len(greetingBytes))
	fmt.Println(string(greetingBytes))

	// 2. Use the embedded string
	fmt.Println("--- As string ---")
	fmt.Print(greetingString)

	// 3. Read a file from the embedded FS
	fmt.Println("\n--- From embed.FS ---")
	data, err := staticFiles.ReadFile("static/greeting.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes from FS\n", len(data))

	// 4. Walk the embedded directory to list all files
	fmt.Println("\n--- Walking embedded FS ---")
	fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Printf("[dir]  %s\n", path)
		} else {
			info, _ := d.Info()
			fmt.Printf("[file] %s (%d bytes)\n", path, info.Size())
		}
		return nil
	})

	// 5. Serve the embedded FS over HTTP — the most common production use case.
	// Use http.FS to convert an embed.FS into an http.FileSystem.
	fmt.Println("\n--- HTTP Server serving embedded files on :8081 ---")
	fmt.Println("Visit http://localhost:8081/static/greeting.txt")
	// Use fs.Sub to strip the "static" prefix so the URL path is /greeting.txt
	subFS, _ := fs.Sub(staticFiles, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(subFS))))

	// NOTE: In a real server you'd call http.ListenAndServe — skipping here to keep the demo short.
}
