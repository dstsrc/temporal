package iplocate_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
	"temporal-ip-geolocation/iplocate"
	"testing"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	activities := &iplocate.IPActivities{}

	// Mock activity implementation
	env.OnActivity(activities.GetIP, mock.Anything).Return("1.1.1.1", nil)
	env.OnActivity(activities.GetLocationInfo, mock.Anything, "1.1.1.1").Return("Planet Earth", nil)

	env.ExecuteWorkflow(iplocate.GetAddressFromIP, "Temporal")

	var result string
	assert.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, "Hello, Temporal. Your IP is 1.1.1.1 and your location is Planet Earth", result)
}
