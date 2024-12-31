package usecases

import (
	"github.com/DamiaoCanndido/document-api/entities"
	"github.com/DamiaoCanndido/document-api/repository"
)

type DocumentUseCase struct {
	repository repository.DocumentRepository
}

func NewDocumentUseCase(repo repository.DocumentRepository) DocumentUseCase {
	return DocumentUseCase{
		repository: repo,
	}
}

func (doc *DocumentUseCase) GetDocuments(docType string) ([]entities.Document, error) {
	return doc.repository.GetDocuments(docType)
}

func (doc *DocumentUseCase) CreateDocuments(document entities.Document) (entities.Document, error) {
	document, err := doc.repository.CreateDocuments(document)
	if (err != nil) {
		return entities.Document{}, err
	}
	return document, nil
}