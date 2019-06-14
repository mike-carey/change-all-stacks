//

package cf

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"io"
	"net/http"
	"net/url"
)

// CFClient ...
type CFClient interface {
	UpdateApp(guid string, aur cfclient.AppUpdateResource) (cfclient.UpdateResponse, error)
	// cfclient.ListAppUsageEventsByQuery lists all events matching the provided query.
	ListAppUsageEventsByQuery(query url.Values) ([]cfclient.AppUsageEvent, error)
	// cfclient.ListAppUsageEvents lists all unfiltered events.
	ListAppUsageEvents() ([]cfclient.AppUsageEvent, error)
	// cfclient.ListAppEvents returns all app events based on eventType
	ListAppEvents(eventType string) ([]cfclient.AppEventEntity, error)
	// cfclient.ListAppEventsByQuery returns all app events based on eventType and queries
	ListAppEventsByQuery(eventType string, queries []cfclient.AppEventQuery) ([]cfclient.AppEventEntity, error)
	// cfclient.ListAppsByQueryWithLimits queries totalPages app info. cfclient.When totalPages is
	// less and equal than 0, it queries all app info
	// cfclient.When there are no more than totalPages apps on server side, all apps info will be returned
	ListAppsByQueryWithLimits(query url.Values, totalPages int) ([]cfclient.App, error)
	ListAppsByQuery(query url.Values) ([]cfclient.App, error)
	// cfclient.GetAppByGuidNoInlineCall will fetch app info including space and orgs information
	// cfclient.Without using inline-relations-depth=2 call
	GetAppByGuidNoInlineCall(guid string) (cfclient.App, error)
	ListApps() ([]cfclient.App, error)
	ListAppsByRoute(routeGuid string) ([]cfclient.App, error)
	GetAppInstances(guid string) (map[string]cfclient.AppInstance, error)
	GetAppEnv(guid string) (cfclient.AppEnv, error)
	GetAppRoutes(guid string) ([]cfclient.Route, error)
	GetAppStats(guid string) (map[string]cfclient.AppStats, error)
	KillAppInstance(guid string, index string) error
	GetAppByGuid(guid string) (cfclient.App, error)
	AppByGuid(guid string) (cfclient.App, error)
	//AppByName takes an appName, and cfclient.GUIDs for a space and org, and performs
	// the cfclient.API lookup with those query parameters set to return you the desired
	// cfclient.App object.
	AppByName(appName, spaceGuid, orgGuid string) (app cfclient.App, err error)
	// cfclient.UploadAppBits uploads the application's contents
	UploadAppBits(file io.Reader, appGUID string) error
	// cfclient.GetAppBits downloads the application's bits as a tar file
	GetAppBits(guid string) (io.ReadCloser, error)
	// cfclient.CreateApp creates a new empty application that still needs it's
	// app bit uploaded and to be started
	CreateApp(req cfclient.AppCreateRequest) (cfclient.App, error)
	StartApp(guid string) error
	StopApp(guid string) error
	DeleteApp(guid string) error
	CreateBuildpack(bpr *cfclient.BuildpackRequest) (*cfclient.Buildpack, error)
	ListBuildpacks() ([]cfclient.Buildpack, error)
	DeleteBuildpack(guid string, async bool) error
	GetBuildpackByGuid(buildpackGUID string) (cfclient.Buildpack, error)
	// cfclient.NewRequest is used to create a new cfclient.Request
	NewRequest(method, path string) *cfclient.Request
	// cfclient.NewRequestWithBody is used to create a new request with
	// arbigtrary body io.Reader.
	NewRequestWithBody(method, path string, body io.Reader) *cfclient.Request
	// cfclient.DoRequest runs a request with our client
	DoRequest(r *cfclient.Request) (*http.Response, error)
	// cfclient.DoRequestWithoutRedirects executes the request without following redirects
	DoRequestWithoutRedirects(r *cfclient.Request) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
	GetToken() (string, error)
	ListDomainsByQuery(query url.Values) ([]cfclient.Domain, error)
	ListDomains() ([]cfclient.Domain, error)
	ListSharedDomainsByQuery(query url.Values) ([]cfclient.SharedDomain, error)
	ListSharedDomains() ([]cfclient.SharedDomain, error)
	GetSharedDomainByGuid(guid string) (cfclient.SharedDomain, error)
	CreateSharedDomain(name string, internal bool, router_group_guid string) (*cfclient.SharedDomain, error)
	DeleteSharedDomain(guid string, async bool) error
	GetDomainByName(name string) (cfclient.Domain, error)
	GetSharedDomainByName(name string) (cfclient.SharedDomain, error)
	CreateDomain(name, orgGuid string) (*cfclient.Domain, error)
	DeleteDomain(guid string) error
	GetRunningEnvironmentVariableGroup() (cfclient.EnvironmentVariableGroup, error)
	GetStagingEnvironmentVariableGroup() (cfclient.EnvironmentVariableGroup, error)
	SetRunningEnvironmentVariableGroup(evg cfclient.EnvironmentVariableGroup) error
	SetStagingEnvironmentVariableGroup(evg cfclient.EnvironmentVariableGroup) error
	// cfclient.ListEventsByQuery lists all events matching the provided query.
	ListEventsByQuery(query url.Values) ([]cfclient.Event, error)
	// cfclient.ListEvents lists all unfiltered events.
	ListEvents() ([]cfclient.Event, error)
	// cfclient.TotalEventsByQuery returns the number of events matching the provided query.
	TotalEventsByQuery(query url.Values) (int, error)
	// cfclient.TotalEvents returns the number of unfiltered events.
	TotalEvents() (int, error)
	// cfclient.GetInfo retrieves cfclient.Info from the cfclient.Cloud cfclient.Controller cfclient.API
	GetInfo() (*cfclient.Info, error)
	CreateIsolationSegment(name string) (*cfclient.IsolationSegment, error)
	GetIsolationSegmentByGUID(guid string) (*cfclient.IsolationSegment, error)
	ListIsolationSegmentsByQuery(query url.Values) ([]cfclient.IsolationSegment, error)
	ListIsolationSegments() ([]cfclient.IsolationSegment, error)
	DeleteIsolationSegmentByGUID(guid string) error
	AddIsolationSegmentToOrg(isolationSegmentGUID, orgGUID string) error
	RemoveIsolationSegmentFromOrg(isolationSegmentGUID, orgGUID string) error
	AddIsolationSegmentToSpace(isolationSegmentGUID, spaceGUID string) error
	RemoveIsolationSegmentFromSpace(isolationSegmentGUID, spaceGUID string) error
	ListOrgQuotasByQuery(query url.Values) ([]cfclient.OrgQuota, error)
	ListOrgQuotas() ([]cfclient.OrgQuota, error)
	GetOrgQuotaByName(name string) (cfclient.OrgQuota, error)
	CreateOrgQuota(orgQuote cfclient.OrgQuotaRequest) (*cfclient.OrgQuota, error)
	UpdateOrgQuota(orgQuotaGUID string, orgQuota cfclient.OrgQuotaRequest) (*cfclient.OrgQuota, error)
	DeleteOrgQuota(guid string, async bool) error
	ListOrgsByQuery(query url.Values) ([]cfclient.Org, error)
	ListOrgs() ([]cfclient.Org, error)
	GetOrgByName(name string) (cfclient.Org, error)
	GetOrgByGuid(guid string) (cfclient.Org, error)
	OrgSpaces(guid string) ([]cfclient.Space, error)
	ListOrgUsersByQuery(orgGUID string, query url.Values) ([]cfclient.User, error)
	ListOrgUsers(orgGUID string) ([]cfclient.User, error)
	ListOrgManagersByQuery(orgGUID string, query url.Values) ([]cfclient.User, error)
	ListOrgManagers(orgGUID string) ([]cfclient.User, error)
	ListOrgAuditorsByQuery(orgGUID string, query url.Values) ([]cfclient.User, error)
	ListOrgAuditors(orgGUID string) ([]cfclient.User, error)
	ListOrgBillingManagersByQuery(orgGUID string, query url.Values) ([]cfclient.User, error)
	ListOrgBillingManagers(orgGUID string) ([]cfclient.User, error)
	AssociateOrgManager(orgGUID, userGUID string) (cfclient.Org, error)
	AssociateOrgManagerByUsername(orgGUID, name string) (cfclient.Org, error)
	AssociateOrgManagerByUsernameAndOrigin(orgGUID, name, origin string) (cfclient.Org, error)
	AssociateOrgUser(orgGUID, userGUID string) (cfclient.Org, error)
	AssociateOrgAuditor(orgGUID, userGUID string) (cfclient.Org, error)
	AssociateOrgUserByUsername(orgGUID, name string) (cfclient.Org, error)
	AssociateOrgUserByUsernameAndOrigin(orgGUID, name, origin string) (cfclient.Org, error)
	AssociateOrgAuditorByUsername(orgGUID, name string) (cfclient.Org, error)
	AssociateOrgAuditorByUsernameAndOrigin(orgGUID, name, origin string) (cfclient.Org, error)
	AssociateOrgBillingManager(orgGUID, userGUID string) (cfclient.Org, error)
	AssociateOrgBillingManagerByUsername(orgGUID, name string) (cfclient.Org, error)
	AssociateOrgBillingManagerByUsernameAndOrigin(orgGUID, name, origin string) (cfclient.Org, error)
	RemoveOrgManager(orgGUID, userGUID string) error
	RemoveOrgManagerByUsername(orgGUID, name string) error
	RemoveOrgManagerByUsernameAndOrigin(orgGUID, name, origin string) error
	RemoveOrgUser(orgGUID, userGUID string) error
	RemoveOrgAuditor(orgGUID, userGUID string) error
	RemoveOrgUserByUsername(orgGUID, name string) error
	RemoveOrgUserByUsernameAndOrigin(orgGUID, name, origin string) error
	RemoveOrgAuditorByUsername(orgGUID, name string) error
	RemoveOrgAuditorByUsernameAndOrigin(orgGUID, name, origin string) error
	RemoveOrgBillingManager(orgGUID, userGUID string) error
	RemoveOrgBillingManagerByUsername(orgGUID, name string) error
	RemoveOrgBillingManagerByUsernameAndOrigin(orgGUID, name, origin string) error
	ListOrgSpaceQuotas(orgGUID string) ([]cfclient.SpaceQuota, error)
	ListOrgPrivateDomains(orgGUID string) ([]cfclient.Domain, error)
	ShareOrgPrivateDomain(orgGUID, privateDomainGUID string) (*cfclient.Domain, error)
	UnshareOrgPrivateDomain(orgGUID, privateDomainGUID string) error
	CreateOrg(req cfclient.OrgRequest) (cfclient.Org, error)
	UpdateOrg(orgGUID string, orgRequest cfclient.OrgRequest) (cfclient.Org, error)
	DeleteOrg(guid string, recursive, async bool) error
	DefaultIsolationSegmentForOrg(orgGUID, isolationSegmentGUID string) error
	ResetDefaultIsolationSegmentForOrg(orgGUID string) error
	// cfclient.ListAllProcesses will call the v3 processes api
	ListAllProcesses() ([]cfclient.Process, error)
	// cfclient.ListAllProcessesByQuery will call the v3 processes api
	ListAllProcessesByQuery(query url.Values) ([]cfclient.Process, error)
	MappingAppAndRoute(req cfclient.RouteMappingRequest) (*cfclient.RouteMapping, error)
	ListRouteMappings() ([]*cfclient.RouteMapping, error)
	ListRouteMappingsByQuery(query url.Values) ([]*cfclient.RouteMapping, error)
	GetRouteMappingByGuid(guid string) (*cfclient.RouteMapping, error)
	DeleteRouteMapping(guid string) error
	// cfclient.CreateRoute creates a regular http route
	CreateRoute(routeRequest cfclient.RouteRequest) (cfclient.Route, error)
	// cfclient.CreateTcpRoute creates a cfclient.TCP route
	CreateTcpRoute(routeRequest cfclient.RouteRequest) (cfclient.Route, error)
	// cfclient.BindRoute associates the specified route with the application
	BindRoute(routeGUID, appGUID string) error
	ListRoutesByQuery(query url.Values) ([]cfclient.Route, error)
	ListRoutes() ([]cfclient.Route, error)
	DeleteRoute(guid string) error
	ListSecGroups() (secGroups []cfclient.SecGroup, err error)
	ListRunningSecGroups() ([]cfclient.SecGroup, error)
	ListStagingSecGroups() ([]cfclient.SecGroup, error)
	GetSecGroupByName(name string) (secGroup cfclient.SecGroup, err error)
	/*
	   cfclient.CreateSecGroup contacts the cfclient.CF endpoint for creating a new security group.
	   name: the name to give to the created security group
	   rules: cfclient.A slice of rule objects that describe the rules that this security group enforces.
	   	This can technically be nil or an empty slice - we won't judge you
	   spaceGuids: cfclient.The security group will be associated with the spaces specified by the contents of this slice.
	   	If nil, the security group will not be associated with any spaces initially.
	*/
	CreateSecGroup(name string, rules []cfclient.SecGroupRule, spaceGuids []string) (*cfclient.SecGroup, error)
	/*
	   cfclient.UpdateSecGroup contacts the cfclient.CF endpoint to update an existing security group.
	   guid: identifies the security group that you would like to update.
	   name: the new name to give to the security group
	   rules: cfclient.A slice of rule objects that describe the rules that this security group enforces.
	   	If this is left nil, the rules will not be changed.
	   spaceGuids: cfclient.The security group will be associated with the spaces specified by the contents of this slice.
	   	If nil, the space associations will not be changed.
	*/
	UpdateSecGroup(guid, name string, rules []cfclient.SecGroupRule, spaceGuids []string) (*cfclient.SecGroup, error)
	/*
	   cfclient.DeleteSecGroup contacts the cfclient.CF endpoint to delete an existing security group.
	   guid: cfclient.Indentifies the security group to be deleted.
	*/
	DeleteSecGroup(guid string) error
	/*
	   cfclient.GetSecGroup contacts the cfclient.CF endpoint for fetching the info for a particular security group.
	   guid: cfclient.Identifies the security group to fetch information from
	*/
	GetSecGroup(guid string) (*cfclient.SecGroup, error)
	/*
	   cfclient.BindSecGroup contacts the cfclient.CF endpoint to associate a space with a security group
	   secGUID: identifies the security group to add a space to
	   spaceGUID: identifies the space to associate
	*/
	BindSecGroup(secGUID, spaceGUID string) error
	/*
	   cfclient.BindSpaceStagingSecGroup contacts the cfclient.CF endpoint to associate a space with a security group for staging functions only
	   secGUID: identifies the security group to add a space to
	   spaceGUID: identifies the space to associate
	*/
	BindStagingSecGroupToSpace(secGUID, spaceGUID string) error
	/*
	   cfclient.BindRunningSecGroup contacts the cfclient.CF endpoint to associate  a security group
	   secGUID: identifies the security group to add a space to
	*/
	BindRunningSecGroup(secGUID string) error
	/*
	   cfclient.UnbindRunningSecGroup contacts the cfclient.CF endpoint to dis-associate  a security group
	   secGUID: identifies the security group to add a space to
	*/
	UnbindRunningSecGroup(secGUID string) error
	/*
	   cfclient.BindStagingSecGroup contacts the cfclient.CF endpoint to associate a space with a security group
	   secGUID: identifies the security group to add a space to
	*/
	BindStagingSecGroup(secGUID string) error
	/*
	   cfclient.UnbindStagingSecGroup contacts the cfclient.CF endpoint to dis-associate a space with a security group
	   secGUID: identifies the security group to add a space to
	*/
	UnbindStagingSecGroup(secGUID string) error
	/*
	   cfclient.UnbindSecGroup contacts the cfclient.CF endpoint to dissociate a space from a security group
	   secGUID: identifies the security group to remove a space from
	   spaceGUID: identifies the space to dissociate from the security group
	*/
	UnbindSecGroup(secGUID, spaceGUID string) error
	ListServiceBindingsByQuery(query url.Values) ([]cfclient.ServiceBinding, error)
	ListServiceBindings() ([]cfclient.ServiceBinding, error)
	GetServiceBindingByGuid(guid string) (cfclient.ServiceBinding, error)
	ServiceBindingByGuid(guid string) (cfclient.ServiceBinding, error)
	DeleteServiceBinding(guid string) error
	CreateServiceBinding(appGUID, serviceInstanceGUID string) (*cfclient.ServiceBinding, error)
	CreateRouteServiceBinding(routeGUID, serviceInstanceGUID string) error
	DeleteRouteServiceBinding(routeGUID, serviceInstanceGUID string) error
	DeleteServiceBroker(guid string) error
	UpdateServiceBroker(guid string, usb cfclient.UpdateServiceBrokerRequest) (cfclient.ServiceBroker, error)
	CreateServiceBroker(csb cfclient.CreateServiceBrokerRequest) (cfclient.ServiceBroker, error)
	ListServiceBrokersByQuery(query url.Values) ([]cfclient.ServiceBroker, error)
	ListServiceBrokers() ([]cfclient.ServiceBroker, error)
	GetServiceBrokerByGuid(guid string) (cfclient.ServiceBroker, error)
	GetServiceBrokerByName(name string) (cfclient.ServiceBroker, error)
	ListServiceInstancesByQuery(query url.Values) ([]cfclient.ServiceInstance, error)
	ListServiceInstances() ([]cfclient.ServiceInstance, error)
	GetServiceInstanceByGuid(guid string) (cfclient.ServiceInstance, error)
	ServiceInstanceByGuid(guid string) (cfclient.ServiceInstance, error)
	CreateServiceInstance(req cfclient.ServiceInstanceRequest) (cfclient.ServiceInstance, error)
	UpdateServiceInstance(serviceInstanceGuid string, updatedConfiguration io.Reader, async bool) error
	DeleteServiceInstance(guid string, recursive, async bool) error
	ListServiceKeysByQuery(query url.Values) ([]cfclient.ServiceKey, error)
	ListServiceKeys() ([]cfclient.ServiceKey, error)
	GetServiceKeyByName(name string) (cfclient.ServiceKey, error)
	// cfclient.GetServiceKeyByInstanceGuid is deprecated in favor of cfclient.GetServiceKeysByInstanceGuid
	GetServiceKeyByInstanceGuid(guid string) (cfclient.ServiceKey, error)
	// cfclient.GetServiceKeysByInstanceGuid returns the service keys for a service instance.
	// cfclient.If none are found, it returns an error.
	GetServiceKeysByInstanceGuid(guid string) ([]cfclient.ServiceKey, error)
	// cfclient.CreateServiceKey creates a service key from the request. cfclient.If a service key
	// exists already, it returns an error containing `CF-ServiceKeyNameTaken`
	CreateServiceKey(csr cfclient.CreateServiceKeyRequest) (cfclient.ServiceKey, error)
	// cfclient.DeleteServiceKey removes a service key instance
	DeleteServiceKey(guid string) error
	ListServicePlanVisibilitiesByQuery(query url.Values) ([]cfclient.ServicePlanVisibility, error)
	ListServicePlanVisibilities() ([]cfclient.ServicePlanVisibility, error)
	GetServicePlanVisibilityByGuid(guid string) (cfclient.ServicePlanVisibility, error)
	//a uniqueID is the id of the service in the catalog and not in cf internal db
	CreateServicePlanVisibilityByUniqueId(uniqueId string, organizationGuid string) (cfclient.ServicePlanVisibility, error)
	CreateServicePlanVisibility(servicePlanGuid string, organizationGuid string) (cfclient.ServicePlanVisibility, error)
	DeleteServicePlanVisibilityByPlanAndOrg(servicePlanGuid string, organizationGuid string, async bool) error
	DeleteServicePlanVisibility(guid string, async bool) error
	UpdateServicePlanVisibility(guid string, servicePlanGuid string, organizationGuid string) (cfclient.ServicePlanVisibility, error)
	ListServicePlansByQuery(query url.Values) ([]cfclient.ServicePlan, error)
	ListServicePlans() ([]cfclient.ServicePlan, error)
	GetServicePlanByGUID(guid string) (*cfclient.ServicePlan, error)
	MakeServicePlanPublic(servicePlanGUID string) error
	MakeServicePlanPrivate(servicePlanGUID string) error
	// cfclient.ListServiceUsageEventsByQuery lists all events matching the provided query.
	ListServiceUsageEventsByQuery(query url.Values) ([]cfclient.ServiceUsageEvent, error)
	// cfclient.ListServiceUsageEvents lists all unfiltered events.
	ListServiceUsageEvents() ([]cfclient.ServiceUsageEvent, error)
	GetServiceByGuid(guid string) (cfclient.Service, error)
	ListServicesByQuery(query url.Values) ([]cfclient.Service, error)
	ListServices() ([]cfclient.Service, error)
	ListSpaceQuotasByQuery(query url.Values) ([]cfclient.SpaceQuota, error)
	ListSpaceQuotas() ([]cfclient.SpaceQuota, error)
	GetSpaceQuotaByName(name string) (cfclient.SpaceQuota, error)
	AssignSpaceQuota(quotaGUID, spaceGUID string) error
	CreateSpaceQuota(spaceQuote cfclient.SpaceQuotaRequest) (*cfclient.SpaceQuota, error)
	UpdateSpaceQuota(spaceQuotaGUID string, spaceQuote cfclient.SpaceQuotaRequest) (*cfclient.SpaceQuota, error)
	CreateSpace(req cfclient.SpaceRequest) (cfclient.Space, error)
	UpdateSpace(spaceGUID string, req cfclient.SpaceRequest) (cfclient.Space, error)
	DeleteSpace(guid string, recursive, async bool) error
	ListSpaceManagersByQuery(spaceGUID string, query url.Values) ([]cfclient.User, error)
	ListSpaceManagers(spaceGUID string) ([]cfclient.User, error)
	ListSpaceAuditorsByQuery(spaceGUID string, query url.Values) ([]cfclient.User, error)
	ListSpaceAuditors(spaceGUID string) ([]cfclient.User, error)
	ListSpaceDevelopersByQuery(spaceGUID string, query url.Values) ([]cfclient.User, error)
	ListSpaceDevelopers(spaceGUID string) ([]cfclient.User, error)
	AssociateSpaceDeveloper(spaceGUID, userGUID string) (cfclient.Space, error)
	AssociateSpaceDeveloperByUsername(spaceGUID, name string) (cfclient.Space, error)
	AssociateSpaceDeveloperByUsernameAndOrigin(spaceGUID, name, origin string) (cfclient.Space, error)
	RemoveSpaceDeveloper(spaceGUID, userGUID string) error
	RemoveSpaceDeveloperByUsername(spaceGUID, name string) error
	RemoveSpaceDeveloperByUsernameAndOrigin(spaceGUID, name, origin string) error
	AssociateSpaceAuditor(spaceGUID, userGUID string) (cfclient.Space, error)
	AssociateSpaceAuditorByUsername(spaceGUID, name string) (cfclient.Space, error)
	AssociateSpaceAuditorByUsernameAndOrigin(spaceGUID, name, origin string) (cfclient.Space, error)
	RemoveSpaceAuditor(spaceGUID, userGUID string) error
	RemoveSpaceAuditorByUsername(spaceGUID, name string) error
	RemoveSpaceAuditorByUsernameAndOrigin(spaceGUID, name, origin string) error
	AssociateSpaceManager(spaceGUID, userGUID string) (cfclient.Space, error)
	AssociateSpaceManagerByUsername(spaceGUID, name string) (cfclient.Space, error)
	AssociateSpaceManagerByUsernameAndOrigin(spaceGUID, name, origin string) (cfclient.Space, error)
	RemoveSpaceManager(spaceGUID, userGUID string) error
	RemoveSpaceManagerByUsername(spaceGUID, name string) error
	RemoveSpaceManagerByUsernameAndOrigin(spaceGUID, name, origin string) error
	ListSpaceSecGroups(spaceGUID string) (secGroups []cfclient.SecGroup, err error)
	ListSpacesByQuery(query url.Values) ([]cfclient.Space, error)
	ListSpaces() ([]cfclient.Space, error)
	GetSpaceByName(spaceName string, orgGuid string) (space cfclient.Space, err error)
	GetSpaceByGuid(spaceGUID string) (cfclient.Space, error)
	IsolationSegmentForSpace(spaceGUID, isolationSegmentGUID string) error
	ResetIsolationSegmentForSpace(spaceGUID string) error
	ListStacksByQuery(query url.Values) ([]cfclient.Stack, error)
	ListStacks() ([]cfclient.Stack, error)
	// cfclient.ListTasks returns all tasks the user has access to.
	// cfclient.See http://v3-apidocs.cloudfoundry.org/version/3.12.0/index.html#list-tasks
	ListTasks() ([]cfclient.Task, error)
	// cfclient.ListTasksByQuery returns all tasks the user has access to, with query parameters.
	// cfclient.See http://v3-apidocs.cloudfoundry.org/version/3.12.0/index.html#list-tasks
	ListTasksByQuery(query url.Values) ([]cfclient.Task, error)
	// cfclient.TasksByApp returns task structures which aligned to an app identified by the given guid.
	// cfclient.See: http://v3-apidocs.cloudfoundry.org/version/3.12.0/index.html#list-tasks-for-an-app
	TasksByApp(guid string) ([]cfclient.Task, error)
	// cfclient.TasksByAppByQuery returns task structures which aligned to an app identified by the given guid
	// and filtered by the given query parameters.
	// cfclient.See: http://v3-apidocs.cloudfoundry.org/version/3.12.0/index.html#list-tasks-for-an-app
	TasksByAppByQuery(guid string, query url.Values) ([]cfclient.Task, error)
	// cfclient.CreateTask creates a new task in cfclient.CF system and returns its structure.
	CreateTask(tr cfclient.TaskRequest) (task cfclient.Task, err error)
	// cfclient.GetTaskByGuid returns a task structure by requesting it with the tasks cfclient.GUID.
	GetTaskByGuid(guid string) (task cfclient.Task, err error)
	TaskByGuid(guid string) (task cfclient.Task, err error)
	// cfclient.TerminateTask cancels a task identified by its cfclient.GUID.
	TerminateTask(guid string) error
	ListUserProvidedServiceInstancesByQuery(query url.Values) ([]cfclient.UserProvidedServiceInstance, error)
	ListUserProvidedServiceInstances() ([]cfclient.UserProvidedServiceInstance, error)
	GetUserProvidedServiceInstanceByGuid(guid string) (cfclient.UserProvidedServiceInstance, error)
	UserProvidedServiceInstanceByGuid(guid string) (cfclient.UserProvidedServiceInstance, error)
	CreateUserProvidedServiceInstance(req cfclient.UserProvidedServiceInstanceRequest) (*cfclient.UserProvidedServiceInstance, error)
	DeleteUserProvidedServiceInstance(guid string) error
	UpdateUserProvidedServiceInstance(guid string, req cfclient.UserProvidedServiceInstanceRequest) (*cfclient.UserProvidedServiceInstance, error)
	// cfclient.GetUserByGUID retrieves the user with the provided guid.
	GetUserByGUID(guid string) (cfclient.User, error)
	ListUsersByQuery(query url.Values) (cfclient.Users, error)
	ListUsers() (cfclient.Users, error)
	ListUserSpaces(userGuid string) ([]cfclient.Space, error)
	ListUserAuditedSpaces(userGuid string) ([]cfclient.Space, error)
	ListUserManagedSpaces(userGuid string) ([]cfclient.Space, error)
	ListUserOrgs(userGuid string) ([]cfclient.Org, error)
	ListUserManagedOrgs(userGuid string) ([]cfclient.Org, error)
	ListUserAuditedOrgs(userGuid string) ([]cfclient.Org, error)
	ListUserBillingManagedOrgs(userGuid string) ([]cfclient.Org, error)
	CreateUser(req cfclient.UserRequest) (cfclient.User, error)
	DeleteUser(userGuid string) error
}
