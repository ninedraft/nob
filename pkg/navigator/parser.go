package navigator

import (
	"archive/zip"

	"github.com/pkg/errors"
)

func ParseZIPArchive(filePath string) (*Root, error) {
	var reader, errOpenZIP = zip.OpenReader(filePath)
	if errOpenZIP != nil {
		return nil, errors.Wrap(errOpenZIP, "unable to open notion backup file")
	}
	defer reader.Close()
	return nil, nil
}
