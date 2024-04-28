package main

import (
	"context"
	"ppp/ps_api/api/database_branches"
	"ppp/ps_api/api/deploy_requests"
)

func main() {
	// remove duplicated code
	// get_organizations_organization_databases_database_deploy_requests_number_responses

	ctx := context.Background()
	ps := NewPsService("org", "db", "tokenId", "token")

	// example 1
	deployRes, _ := ps.Ps.DeployRequests.PostOrganizationsOrganizationDatabasesDatabaseDeployRequests(
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

	// deploy number is float64
	_ = deployRes.Payload.Number

	// example 2
	branchRes, _ := ps.Ps.DatabaseBranches.GetOrganizationsOrganizationDatabasesDatabaseBranchesName(
		&database_branches.GetOrganizationsOrganizationDatabasesDatabaseBranchesNameParams{
			Database:     ps.DbName,
			Organization: ps.Organization,
			Context:      ctx,
			Name:         "dev",
		}, nil,
	)

	// defined as string but the api return number
	_ = branchRes.Payload.InitialRestoreID

}
