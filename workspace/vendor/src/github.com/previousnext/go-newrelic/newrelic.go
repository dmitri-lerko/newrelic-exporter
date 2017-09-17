package newrelic

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Client is used to interacting with New Relics API.
type Client struct {
	key string
}

// DeploymentInput is passed the REST client to register a deployment.
type DeploymentInput struct {
	Deployment Deployment `json:"deployment"`
}

// Deployment contains the parameters for registering a deployment.
type Deployment struct {
	Revision    string `json:"revision"`
	Changelog   string `json:"changelog"`
	Description string `json:"description"`
	User        string `json:"user"`
}

// Applications is a collection of New Relic applications.
type Applications struct {
	Applications []Application `json:"applications"`
}

// Application is a New Relic application.
type Application struct {
	ID                 int64              `json:"id"`
	Name               string             `json:"name"`
	Language           string             `json:"language"`
	HealthStatus       string             `json:"health_status"`
	ApplicationSummary ApplicationSummary `json:"application_summary"`
	EndUserSummary     EndUserSummary     `json:"end_user_summary"`
}

// ApplicationSummary is used to return a summary of the applications health.
type ApplicationSummary struct {
	ResponseTime  float64 `json:"response_time"`
	Throughput    float64 `json:"throughput"`
	ErrorRate     float64 `json:"error_rate"`
	ApdexTarget   float64 `json:"apdex_target"`
	ApdexScore    float64 `json:"apdex_score"`
	HostCount     float64 `json:"host_count"`
	InstanceCount float64 `json:"instance_count"`
}

// EndUserSummary is used to return the end users experience.
type EndUserSummary struct {
	ReponseTime float64 `json:"response_time"`
	Throughput  float64 `json:"throughput"`
	ApdexTarget float64 `json:"apdex_target"`
	ApdexScore  float64 `json:"apdex_score"`
}

// New returns a new New Relic client.
func New(key string) Client {
	return Client{
		key: key,
	}
}

// NameToApplicationID returns an App ID based on App Name.
func (n Client) NameToApplicationID(name string) (int64, error) {
	resp, err := n.ListApplications()
	if err != nil {
		return 0, err
	}

	for _, app := range resp.Applications {
		if app.Name == name {
			return app.ID, nil
		}
	}

	return 0, fmt.Errorf("Cannot find application with name: %s", name)
}

// Application returns an Application status.
func (n Client) Application(name string) (Application, error) {
	var app Application

	resp, err := n.ListApplications()
	if err != nil {
		return app, err
	}

	for _, app := range resp.Applications {
		if app.Name == name {
			return app, nil
		}
	}

	return app, fmt.Errorf("Cannot find application with name: %s", name)
}

// ListApplications returns a list of applications.
func (n Client) ListApplications() (Applications, error) {
	var apps Applications

	_, _, errs := gorequest.New().Get("https://api.newrelic.com/v2/applications.json").
		Set("X-Api-Key", n.key).
		Set("Content-Type", "application/json").
		EndStruct(&apps)

	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}

		return apps, errs[0]
	}

	return apps, nil
}

// Deployment sends a deployment tag to a New Relic application.
// https://docs.newrelic.com/docs/apm/new-relic-apm/maintenance/recording-deployments
func (n Client) Deployment(id int64, d DeploymentInput) error {
	_, body, errs := gorequest.New().Post(fmt.Sprintf("https://api.newrelic.com/v2/applications/%d/deployments.json", id)).
		Set("X-Api-Key", n.key).
		Set("Content-Type", "application/json").
		Send(d).
		End()

	if len(errs) > 0 {
		return fmt.Errorf(body)
	}

	return nil
}
