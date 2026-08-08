package webRepository

import (
	"fmt"
	"os"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/errors"
)

func (p *WebRepository) GetFile(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found %w", core_errors.ErrNotFound)
		}
		return nil, err
	}
	return file, nil
}
