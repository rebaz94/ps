ps-api:
	@swagger generate client -f ps_api/ps_api_spec.json -A PlanetScale -c api -t ./ps_api