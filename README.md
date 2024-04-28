# generating:

use `make ps-api` or go generate line in `ps_service.go`

# issue

1. Duplicate constant values,
   check: `get_organizations_organization_databases_database_deploy_requests_number_responses.go`
2. The newest spec define `operationId` with a long name and use http verb,
   like: `PostOrganizationsOrganizationDatabasesDatabaseDeployRequests`
3. the `Number` from **Deploy Request** is float64, should be uint64, or a string that is supported for all language.
4. `InitialRestoreID` from **Get Branch** defined as string but the api return number, decoding failed.