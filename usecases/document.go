package usecases

import (
	"go-api/entities"
	"go-api/repository"
)

type DocumentUseCase struct {
	repository repository.DocumentRepository
}

func NewDocumentUseCase(repo repository.DocumentRepository) DocumentUseCase {
	return DocumentUseCase{
		repository: repo,
	}
}

func (doc *DocumentUseCase) GetDocuments() ([]entities.Document, error) {
	return doc.repository.GetDocuments()
}

func (doc *DocumentUseCase) CreateDocuments(document entities.Document) (entities.Document, error) {
	document, err := doc.repository.CreateDocuments(document)
	if (err != nil) {
		return entities.Document{}, err
	}
	return document, nil
}