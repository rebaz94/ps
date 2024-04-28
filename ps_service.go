package main

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"ppp/ps_api/api"
)

// generate PlanetScale Notifier from swagger spec
//go:generate swagger generate client -f ps_api/ps_api_spec.json -A PlanetScale -c api -t ./ps_api

type PsService struct {
	// PlanetScale organization identifier
	Organization string

	// Name of the database
	DbName string
	Ps     *api.PlanetScale
}

func NewPsService(organizationName, dbName, tokenId, token string) *PsService {
	cfg := api.DefaultTransportConfig()
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", tokenId+":"+token)

	ps := &PsService{
		Organization: organizationName,
		DbName:       dbName,
		Ps:           api.New(transport, strfmt.Default),
	}

	return ps
}
