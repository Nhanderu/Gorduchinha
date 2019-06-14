package service

import (
	"net/http"

	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/Nhanderu/gorduchinha/src/infra/config"
	"github.com/Nhanderu/gorduchinha/src/infra/logger"
)

// Service holds the domain service repositories.
type Service struct {
	db         contract.DataManager
	cfg        config.Config
	log        logger.Logger
	cache      contract.CacheManager
	tracking   contract.TrackingManager
	messaging  contract.MessagingManager
	monitoring contract.MonitoringManager
	httpClient *http.Client

	Scraper scraperService
	Champ   champService
	Team    teamService
}

// New returns a new domain Service instance.
func New(db contract.DataManager, cache contract.CacheManager, cfg config.Config, log logger.Logger, tracking contract.TrackingManager, messaging contract.MessagingManager, monitoring contract.MonitoringManager) *Service {

	svc := new(Service)
	svc.db = db
	svc.cfg = cfg
	svc.log = log
	svc.cache = cache
	svc.tracking = tracking
	svc.messaging = messaging
	svc.monitoring = monitoring

	svc.httpClient = new(http.Client)
	svc.httpClient.Timeout = cfg.HTTPClient.Timeout

	svc.Scraper = newScraperService(svc)
	svc.Champ = newChampService(svc)
	svc.Team = newTeamService(svc)

	return svc
}
