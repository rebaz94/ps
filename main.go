package main

import (
	"context"
	"ppp/ps_api/api/deploy_requests"
)

func main() {
	// remove duplicated code
	// get_organizations_organization_databases_database_deploy_requests_number_responses

	ctx := context.Background()
	ps := NewPsService("org", "db", "tokenId", "token")

	res, _ := ps.Ps.DeployRequests.PostOrganizationsOrganizationDatabasesDatabaseDeployRequests(
		&deploy_requests.PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsParams{
			Body: deploy_requests.PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsBody{
				Branch:     "",
				IntoBranch: "",
				Notes:      "",
			},
			Database:     ps.DbName,
			Organization: ps.Organization,
			Context:      ctx,
		},
		nil,
	)

	// deploy number
	_ = res.Payload.Number

}
