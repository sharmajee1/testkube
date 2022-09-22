package client

import (
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor/output"
)

// Client is the Testkube API client abstraction
type Client interface {
	TestAPI
	ExecutionAPI
	TestSuiteAPI
	TestSuiteExecutionAPI
	ExecutorAPI
	WebhookAPI
	ServiceAPI
	ConfigAPI
	TestSourceAPI
}

// TestAPI describes test api methods
type TestAPI interface {
	GetTest(id string) (test testkube.Test, err error)
	GetTestWithExecution(id string) (test testkube.TestWithExecution, err error)
	CreateTest(options UpsertTestOptions) (test testkube.Test, err error)
	UpdateTest(options UpsertTestOptions) (test testkube.Test, err error)
	DeleteTest(name string) error
	DeleteTests(selector string) error
	ListTests(selector string) (tests testkube.Tests, err error)
	ListTestWithExecutions(selector string) (tests testkube.TestWithExecutions, err error)
	ExecuteTest(id, executionName string, options ExecuteTestOptions) (executions testkube.Execution, err error)
	ExecuteTests(selector string, concurrencyLevel int, options ExecuteTestOptions) (executions []testkube.Execution, err error)
	Logs(id string) (logs chan output.Output, err error)
}

// ExecutionAPI describes execution api methods
type ExecutionAPI interface {
	GetExecution(executionID string) (execution testkube.Execution, err error)
	ListExecutions(id string, limit int, selector string) (executions testkube.ExecutionsResult, err error)
	AbortExecution(test string, id string) error
	GetExecutionArtifacts(executionID string) (artifacts testkube.Artifacts, err error)
	DownloadFile(executionID, fileName, destination string) (artifact string, err error)
}

// TestSuiteAPI describes test suite api methods
type TestSuiteAPI interface {
	CreateTestSuite(options UpsertTestSuiteOptions) (testSuite testkube.TestSuite, err error)
	UpdateTestSuite(options UpsertTestSuiteOptions) (testSuite testkube.TestSuite, err error)
	GetTestSuite(id string) (testSuite testkube.TestSuite, err error)
	GetTestSuiteWithExecution(id string) (testSuite testkube.TestSuiteWithExecution, err error)
	ListTestSuites(selector string) (testSuites testkube.TestSuites, err error)
	ListTestSuiteWithExecutions(selector string) (testSuitesWithExecutions testkube.TestSuiteWithExecutions, err error)
	DeleteTestSuite(name string) error
	DeleteTestSuites(selector string) error
	ExecuteTestSuite(id, executionName string, options ExecuteTestSuiteOptions) (executions testkube.TestSuiteExecution, err error)
	ExecuteTestSuites(selector string, concurrencyLevel int, options ExecuteTestSuiteOptions) (executions []testkube.TestSuiteExecution, err error)
}

// TestSuiteExecutionAPI describes test suite execution api methods
type TestSuiteExecutionAPI interface {
	GetTestSuiteExecution(executionID string) (execution testkube.TestSuiteExecution, err error)
	ListTestSuiteExecutions(test string, limit int, selector string) (executions testkube.TestSuiteExecutionsResult, err error)
	WatchTestSuiteExecution(executionID string) (execution chan testkube.TestSuiteExecution, err error)
}

// ExecutorAPI describes executor api methods
type ExecutorAPI interface {
	CreateExecutor(options CreateExecutorOptions) (executor testkube.ExecutorDetails, err error)
	GetExecutor(name string) (executor testkube.ExecutorDetails, err error)
	ListExecutors(selector string) (executors testkube.ExecutorsDetails, err error)
	DeleteExecutor(name string) (err error)
	DeleteExecutors(selector string) (err error)
}

// WebhookAPI describes webhook api methods
type WebhookAPI interface {
	CreateWebhook(options CreateWebhookOptions) (webhook testkube.Webhook, err error)
	GetWebhook(name string) (webhook testkube.Webhook, err error)
	ListWebhooks(selector string) (webhooks testkube.Webhooks, err error)
	DeleteWebhook(name string) (err error)
	DeleteWebhooks(selector string) (err error)
}

// ConfigAPI describes config api methods
type ConfigAPI interface {
	UpdateConfig(config testkube.Config) (outputConfig testkube.Config, err error)
	GetConfig() (config testkube.Config, err error)
}

// ServiceAPI describes service api methods
type ServiceAPI interface {
	GetServerInfo() (info testkube.ServerInfo, err error)
	GetDebugInfo() (info testkube.DebugInfo, err error)
}

// TestSourceAPI describes test source api methods
type TestSourceAPI interface {
	CreateTestSource(options UpsertTestSourceOptions) (testSource testkube.TestSource, err error)
	UpdateTestSource(options UpsertTestSourceOptions) (testSource testkube.TestSource, err error)
	GetTestSource(name string) (testSource testkube.TestSource, err error)
	ListTestSources(selector string) (testSources testkube.TestSources, err error)
	DeleteTestSource(name string) (err error)
	DeleteTestSources(selector string) (err error)
}

// TODO consider replacing below types by testkube.*

// UpsertTestSuiteOptions - mapping to OpenAPI schema for creating/changing testsuite
type UpsertTestSuiteOptions testkube.TestSuiteUpsertRequest

// UpsertTestOptions - is mapping for now to OpenAPI schema for creating/changing test
// if needed can be extended to custom struct
type UpsertTestOptions testkube.TestUpsertRequest

// CreateExecutorOptions - is mapping for now to OpenAPI schema for creating executor request
type CreateExecutorOptions testkube.ExecutorCreateRequest

// CreateWebhookOptions - is mapping for now to OpenAPI schema for creating/changing webhook
type CreateWebhookOptions testkube.WebhookCreateRequest

// UpsertTestSourceOptions - is mapping for now to OpenAPI schema for creating/changing test source
// if needed can be extended to custom struct
type UpsertTestSourceOptions testkube.TestSourceUpsertRequest

// TODO consider replacing it with testkube.ExecutionRequest - looks almost the samea and redundant
// ExecuteTestOptions contains test run options
type ExecuteTestOptions struct {
	ExecutionVariables            map[string]testkube.Variable
	ExecutionVariablesFileContent string
	ExecutionLabels               map[string]string
	Args                          []string
	Envs                          map[string]string
	SecretEnvs                    map[string]string
	HTTPProxy                     string
	HTTPSProxy                    string
	Image                         string
}

// ExecuteTestSuiteOptions contains test suite run options
type ExecuteTestSuiteOptions struct {
	ExecutionVariables map[string]testkube.Variable
	HTTPProxy          string
	HTTPSProxy         string
	ExecutionLabels    map[string]string
}

// Gettable is an interface of gettable objects
type Gettable interface {
	testkube.Test | testkube.TestSuite | testkube.ExecutorDetails |
		testkube.Webhook | testkube.TestWithExecution | testkube.TestSuiteWithExecution |
		testkube.Artifact | testkube.ServerInfo | testkube.Config | testkube.DebugInfo | testkube.TestSource
}

// Executable is an interface of executable objects
type Executable interface {
	testkube.Execution | testkube.TestSuiteExecution |
		testkube.ExecutionsResult | testkube.TestSuiteExecutionsResult
}

// All is an interface of all objects
type All interface {
	Gettable | Executable
}

// Transport provides methods to execute api calls
type Transport[A All] interface {
	Execute(method, uri string, body []byte, params map[string]string) (result A, err error)
	ExecuteMultiple(method, uri string, body []byte, params map[string]string) (result []A, err error)
	Delete(uri, selector string, isContentExpected bool) error
	GetURI(pathTemplate string, params ...interface{}) string
	GetLogs(uri string, logs chan output.Output) error
	GetFile(uri, fileName, destination string) (name string, err error)
}
