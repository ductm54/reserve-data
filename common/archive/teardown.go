package archive

import (
	"log"
	"os"
)

func TearDown(tmpDir string) {
	rErr := os.RemoveAll(tmpDir)
	if rErr != nil {
		log.Printf("cannot remove temp dir %s, err: %s. Please remove it manually", tmpDir, rErr)
	}
}
