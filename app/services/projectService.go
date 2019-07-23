package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"

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

	var buf *bytes.Buffer
	json.NewEncoder(buf).Encode(project)

	req, err := http.NewRequest(http.MethodPost, "http://"+workerAddr+"/projects", buf)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	req.Header.Set("X-Auth-Token", viper.GetString("windlass.secret"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Info("%d", resp.StatusCode)
	return nil
}
