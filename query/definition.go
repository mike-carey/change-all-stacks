package query

//go:generate ifacemaker -f ../vendor/github.com/cloudfoundry-community/go-cfclient/app_update.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/app_usage_events.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/appevents.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/apps.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/buildpacks.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/cf_error.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/client.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/domains.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/environmentvariablegroups.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/error.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/events.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/gen_error.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/info.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/isolationsegments.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/org_quotas.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/orgs.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/processes.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/route_mappings.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/routes.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/secgroups.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_bindings.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_brokers.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_instances.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_keys.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_plan_visibilities.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_plans.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_usage_events.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/services.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/space_quotas.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/spaces.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/stacks.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/tasks.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/types.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/user_provided_service_instances.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/users.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/v3types.go -s Client -i CFClient -p query -o cf-client.go
//go:generate ./patch.sh cf-client.go

//go:generate ./inquisitor.sh -pkg query generic/inquisitor.go.erb
//go:generate ./inquisitor.sh -pkg query_test generic/inquisitor_test.go.erb

//go:generate genny -in generic/generic-setup_test.go -out app-setup_test.go -pkg query_test gen Item=App
//go:generate genny -in generic/generic-service.go -out app-service.go -pkg query gen Item=App
//go:generate genny -in generic/generic-service_test.go -out app-service_test.go -pkg query_test gen Item=App
//go:generate genny -in generic/generic-filter-by.go -out app-filter-by.go -pkg query gen Item=App
//go:generate genny -in generic/generic-filter-by_test.go -out app-filter-by_test.go -pkg query_test gen Item=App
//go:generate genny -in generic/generic-group-by.go -out app-group-by.go -pkg query gen Item=App
//go:generate genny -in generic/generic-group-by_test.go -out app-group-by_test.go -pkg query_test gen Item=App

//go:generate genny -in generic/generic-setup_test.go -out space-setup_test.go -pkg query_test gen Item=Space
//go:generate genny -in generic/generic-service.go -out space-service.go -pkg query gen Item=Space
//go:generate genny -in generic/generic-service_test.go -out space-service_test.go -pkg query_test gen Item=Space
//go:generate genny -in generic/generic-filter-by.go -out space-filter-by.go -pkg query gen Item=Space
//go:generate genny -in generic/generic-filter-by_test.go -out space-filter-by_test.go -pkg query_test gen Item=Space
//go:generate genny -in generic/generic-group-by.go -out space-group-by.go -pkg query gen Item=Space
//go:generate genny -in generic/generic-group-by_test.go -out space-group-by_test.go -pkg query_test gen Item=Space

//go:generate genny -in generic/generic-setup_test.go -out org-setup_test.go -pkg query_test gen Item=Org
//go:generate genny -in generic/generic-service.go -out org-service.go -pkg query gen Item=Org
//go:generate genny -in generic/generic-service_test.go -out org-service_test.go -pkg query_test gen Item=Org
//go:generate genny -in generic/generic-filter-by.go -out org-filter-by.go -pkg query gen Item=Org
//go:generate genny -in generic/generic-filter-by_test.go -out org-filter-by_test.go -pkg query_test gen Item=Org
//go:generate genny -in generic/generic-group-by.go -out org-group-by.go -pkg query gen Item=Org
//go:generate genny -in generic/generic-group-by_test.go -out org-group-by_test.go -pkg query_test gen Item=Org

//go:generate genny -in generic/generic-setup_test.go -out buildpack-setup_test.go -pkg query_test gen Item=Buildpack
//go:generate genny -in generic/generic-service.go -out buildpack-service.go -pkg query gen Item=Buildpack
//go:generate genny -in generic/generic-service_test.go -out buildpack-service_test.go -pkg query_test gen Item=Buildpack
//go:generate genny -in generic/generic-filter-by.go -out buildpack-filter-by.go -pkg query gen Item=Buildpack
//go:generate genny -in generic/generic-filter-by_test.go -out buildpack-filter-by_test.go -pkg query_test gen Item=Buildpack
//go:generate genny -in generic/generic-group-by.go -out buildpack-group-by.go -pkg query gen Item=Buildpack
//go:generate genny -in generic/generic-group-by_test.go -out buildpack-group-by_test.go -pkg query_test gen Item=Buildpack

//go:generate genny -in generic/generic-setup_test.go -out stack-setup_test.go -pkg query_test gen Item=Stack
//go:generate genny -in generic/generic-service.go -out stack-service.go -pkg query gen Item=Stack
//go:generate genny -in generic/generic-service_test.go -out stack-service_test.go -pkg query_test gen Item=Stack
//go:generate genny -in generic/generic-filter-by.go -out stack-filter-by.go -pkg query gen Item=Stack
//go:generate genny -in generic/generic-filter-by_test.go -out stack-filter-by_test.go -pkg query_test gen Item=Stack
//go:generate genny -in generic/generic-group-by.go -out stack-group-by.go -pkg query gen Item=Stack
//go:generate genny -in generic/generic-group-by_test.go -out stack-group-by_test.go -pkg query_test gen Item=Stack

//go:generate ./patch.sh --ignore data_test.go --ignore inquisitor_test.go --ignore query_suite_test.go --ignore *-setup_test.go -- *_test.go

// TODO: Remove this workaround once merged: https://github.com/cloudfoundry-community/go-cfclient/pull/234
//go:generate ./workaround-stack.sh stack-service.go

//go:generate counterfeiter -o fakes/fake_cf_client.go cf-client.go CFClient
