package u

// IsImageFile checks if the given file is an image file based on its MIME type.
func IsImageFile(filePath string) (bool, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Get the file's MIME type
	fileHeader := make([]byte, 512) // We only need the first 512 bytes to determine the MIME type
	_, err = file.Read(fileHeader)
	if err != nil {
		return false, err
	}

	// Detect the MIME type
	fileType := http.DetectContentType(fileHeader)

	// Check if the MIME type indicates an image
	return strings.HasPrefix(fileType, "image/"), nil
}