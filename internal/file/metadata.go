package file

import (
	"bytes"
	"image"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"fmt"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	pb "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnQueued func(bool)
type OnProgress func(data []byte)
type OnError func(err error, method string)

// Struct to Implement Node Callback Methods
type FileCallback struct {
	Queued   OnQueued
	Progress OnProgress
	Error    OnError
}

// ^ File that safely sets metadata and thumbnail in routine ^ //
type SafeFile struct {
	Path     string
	Call     FileCallback
	mutex    sync.Mutex
	metadata pb.Metadata
}

// ^ Create generates file metadata ^ //
func (sf *SafeFile) Create() {
	// ** Lock ** //
	sf.mutex.Lock()

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(sf.Path)
	if err != nil {
		fmt.Println("Error opening File", err)
		sf.Call.Error(err, "AddFile")
		sf.Call.Queued(false)
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		sf.Call.Error(err, "AddFile")
		sf.Call.Queued(false)
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
		sf.Call.Error(err, "AddFile")
		sf.Call.Queued(false)
	}
	file.Close()

	// @ 2. Create Thumbnail
	thumbBuffer := new(bytes.Buffer)
	if filetype.IsImage(head) {
		// New File for ThumbNail
		file, err := os.Open(sf.Path)
		if err != nil {
			fmt.Println(err)
			sf.Call.Error(err, "AddFile")
			sf.Call.Queued(false)
		}
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
			sf.Call.Error(err, "AddFile")
			sf.Call.Queued(false)
		}

		// Find Image Bounds
		b := img.Bounds()
		width := b.Max.X
		height := b.Max.Y
		fmt.Println("width = ", width, " height = ", height)

		// Calculate Fit
		newWidth, newHeight := calculateRatioFit(width, height)
		fmt.Println("w = ", newWidth, " h = ", newHeight)

		// Call the resize library for image scaling
		scaledImage := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

		// Encode as Jpeg into buffer
		err = jpeg.Encode(thumbBuffer, scaledImage, nil)
		if err != nil {
			fmt.Println(err)
			sf.Call.Error(err, "AddFile")
			sf.Call.Queued(false)
		}
		fmt.Println("Thumbnail created")
	}

	// @ 3. Set Metadata Protobuf Values
	sf.metadata = pb.Metadata{
		FileId:    uuid.New().String(),
		Name:      fileName(sf.Path),
		Path:      sf.Path,
		Size:      info.Size(),
		Kind:      kind.MIME.Type,
		Thumbnail: thumbBuffer.Bytes(),
	}
	// ** Unlock ** //
	sf.mutex.Unlock()
	sf.Call.Queued(true)
}

// ^ Safely returns metadata depending on lock ^ //
func (sf *SafeFile) Metadata() *pb.Metadata {
	// ** Lock File wait for access ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// @ 2. Return Value
	return &sf.metadata
}

// ** Resize Constants ** //
const MAX_WIDTH float64 = 320
const MAX_HEIGHT float64 = 240

// @ Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(MAX_WIDTH/float64(srcWidth), MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// @ Get FileName without Extension
func fileName(path string) string {
	fileBase := filepath.Base(path)
	return strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
}
