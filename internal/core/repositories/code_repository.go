package repositories

import "cristianUrbina/open-typing-batch-job/internal/core/entities"

type CodeRepository interface {
  Create(CodeProject *entitites.Code) error
}
