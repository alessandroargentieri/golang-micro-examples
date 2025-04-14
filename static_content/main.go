package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define the port to listen on
	port := ":8080"

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	// Define the path to the static files directory
	staticDir := filepath.Join(cwd, "static")

	// Check if the static directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		// Create the directory structure if it doesn't exist
		err = os.MkdirAll(staticDir, 0755)
		if err != nil {
			log.Fatal("Error creating static directory:", err)
		}

		// Create a default index.html file with a message
		indexPath := filepath.Join(staticDir, "index.html")
		defaultHTML := `<!DOCTYPE html>
<html>
<head>
    <title>Go File Server</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .alert { background-color: #f8d7da; border: 1px solid #f5c6cb; color: #721c24; padding: 10px; border-radius: 4px; }
    </style>
</head>
<body>
    <h1>Go File Server</h1>
    <div class="alert">
        <p><strong>Note:</strong> This is a default page. To display your custom HTML, place your file explorer HTML in the "static" folder as "index.html".</p>
    </div>
</body>
</html>`

		err = os.WriteFile(indexPath, []byte(defaultHTML), 0644)
		if err != nil {
			log.Fatal("Error creating default index.html:", err)
		}

		fmt.Println("Created static directory and default index.html")
	}

	// Create a file server handler for the static directory
	fs := http.FileServer(http.Dir(staticDir))

	// Register the file server handler for the root path
	http.Handle("/", fs)

	// Start the server
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Printf("Serving files from: %s\n", staticDir)
	log.Fatal(http.ListenAndServe(port, nil))
}
