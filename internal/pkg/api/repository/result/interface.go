package result

import (
	"context"
	"time"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

const PageDefaultLimit int = 100

type Filter interface {
	TestName() string
	TestNameDefined() bool
	LastNDays() int
	LastNDaysDefined() bool
	StartDate() time.Time
	StartDateDefined() bool
	EndDate() time.Time
	EndDateDefined() bool
	Statuses() testkube.ExecutionStatuses
	StatusesDefined() bool
	Page() int
	PageSize() int
	TextSearchDefined() bool
	TextSearch() string
	Selector() string
	TypeDefined() bool
	Type() string
}

type Repository interface {
	Sequences
	// Get gets execution result by id
	Get(ctx context.Context, id string) (testkube.Execution, error)
	// GetByName gets execution result by name
	GetByName(ctx context.Context, id string) (testkube.Execution, error)
	// GetByNameAndTest gets execution result by name and test name
	GetByNameAndTest(ctx context.Context, name, testName string) (testkube.Execution, error)
	// GetLatestByTest gets latest execution result by test
	GetLatestByTest(ctx context.Context, testName, sortField string) (testkube.Execution, error)
	// GetLatestByTests gets latest execution results by test names
	GetLatestByTests(ctx context.Context, testNames []string, sortField string) (executions []testkube.Execution, err error)
	// GetExecutions gets executions using a filter, use filter with no data for all
	GetExecutions(ctx context.Context, filter Filter) ([]testkube.Execution, error)
	// GetExecutionTotals gets the statistics on number of executions using a filter, but without paging
	GetExecutionTotals(ctx context.Context, paging bool, filter ...Filter) (result testkube.ExecutionsTotals, err error)
	// Insert inserts new execution result
	Insert(ctx context.Context, result testkube.Execution) error
	// Update updates execution result
	Update(ctx context.Context, result testkube.Execution) error
	// UpdateExecution updates result in execution
	UpdateResult(ctx context.Context, id string, execution testkube.ExecutionResult) error
	// StartExecution updates execution start time
	StartExecution(ctx context.Context, id string, startTime time.Time) error
	// EndExecution updates execution end time
	EndExecution(ctx context.Context, id string, endTime time.Time, duration time.Duration) error
	// GetLabels get all available labels
	GetLabels(ctx context.Context) (labels map[string][]string, err error)
	// DeleteByTest deletes execution results by test
	DeleteByTest(ctx context.Context, testName string) error
	// DeleteByTestSuite deletes execution results by test suite
	DeleteByTestSuite(ctx context.Context, testSuiteName string) error
	// DeleteAll deletes all execution results
	DeleteAll(ctx context.Context) error
	// DeleteByTests deletes execution results by tests
	DeleteByTests(ctx context.Context, testNames []string) (err error)
	// DeleteByTestSuites deletes execution results by test suites
	DeleteByTestSuites(ctx context.Context, testSuiteNames []string) (err error)
	// DeleteForAllTestSuites deletes execution results for all test suites
	DeleteForAllTestSuites(ctx context.Context) (err error)

	GetTestMetrics(ctx context.Context, name string, limit, last int) (metrics testkube.ExecutionsMetrics, err error)
}

type Sequences interface {
	// GetNextExecutionNumber gets next execution number by test name
	GetNextExecutionNumber(ctx context.Context, testName string) (number int32, err error)
}
