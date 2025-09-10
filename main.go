package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Image Merger CLI started. Press Ctrl+C to exit.")

	for {
		fmt.Println("Merging images from upload/...")

		err := MergeImages("upload", "output/file.png")
		if err != nil {
			log.Println("Error:", err)
		} else {
			fmt.Println("✅ Done! Saved to output/file.png")
		}

		fmt.Println("Press Enter to re-run or Ctrl+C to exit.")
		var input string
		fmt.Scanln(&input)
	}
}
