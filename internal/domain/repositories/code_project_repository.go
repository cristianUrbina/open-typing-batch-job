package repositories

import (
	"cristianUrbina/open-typing-batch-job/internal/domain"
)

type CodeRepository interface {
  Create(CodeProject *domain.Repository) error
}
