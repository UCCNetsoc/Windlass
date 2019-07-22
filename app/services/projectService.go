package services

import (
	"context"
	"net/http"

	"github.com/UCCNetworkingSociety/Windlass-worker/app/models/project"
	repo "github.com/UCCNetworkingSociety/Windlass/app/repositories"
	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
)

type ProjectService struct {
	consulRepo *repo.ConsulRepository
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		consulRepo: repo.NewConsulRepository(),
	}
}

func (p *ProjectService) CreateProject(ctx context.Context, project project.Project) error {
	workerAddr, err := p.consulRepo.SelectWorker(ctx)
	if err != nil {
		return err
	}

	log.Info("address: %s", workerAddr)

	resp, err := http.Get("http://" + workerAddr + "/health")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Info("%d", resp.StatusCode)
	return nil
}
