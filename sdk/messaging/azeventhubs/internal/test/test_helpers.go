// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	azlog "github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/exported"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

// CaptureLogsForTest adds a logging listener which captures messages to an
// internal channel.
// Returns a function that ends log capturing and returns any captured messages.
// It's safe to call endCapture() multiple times, so a simple call pattern is:
//
//	endCapture := CaptureLogsForTest()
//	defer endCapture()  // ensure cleanup in case of test assert failures
//
//	/* some test code */
//
//	messages := endCapture()
//	/* do inspection of log messages */
func CaptureLogsForTest() func() []string {
	messagesCh := make(chan string, 10000)
	return CaptureLogsForTestWithChannel(messagesCh)
}

func CaptureLogsForTestWithChannel(messagesCh chan string) func() []string {
	setAzLogListener(func(e azlog.Event, s string) {
		messagesCh <- fmt.Sprintf("[%s] %s", e, s)
	})

	return func() []string {
		if messagesCh == nil {
			// already been closed, probably manually.
			return nil
		}

		setAzLogListener(nil)
		close(messagesCh)

		var messages []string

		for msg := range messagesCh {
			messages = append(messages, msg)
		}

		messagesCh = nil
		return messages
	}
}

// EnableStdoutLogging turns on logging to stdout for diagnostics.
func EnableStdoutLogging() {
	azlog.SetEvents(exported.EventAuth, exported.EventConn, exported.EventConsumer, exported.EventProducer)
	setAzLogListener(func(e azlog.Event, s string) {
		log.Printf("%s %s", e, s)
	})
}

var logMu sync.Mutex

func setAzLogListener(listener func(e azlog.Event, s string)) {
	logMu.Lock()
	defer logMu.Unlock()
	azlog.SetListener(listener)
}

// RandomString generates a random string with prefix
func RandomString(prefix string, length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return prefix + hex.EncodeToString(b)
}

type ConnectionParamsForTest struct {
	ConnectionString        string
	EventHubName            string
	EventHubNamespace       string
	ResourceGroup           string
	StorageConnectionString string
	SubscriptionID          string
}

func GetConnectionParamsForTest(t *testing.T) ConnectionParamsForTest {
	_ = godotenv.Load()

	envVars := mustGetEnvironmentVars(t, []string{
		"AZURE_SUBSCRIPTION_ID",
		"CHECKPOINTSTORE_STORAGE_CONNECTION_STRING",
		"EVENTHUB_CONNECTION_STRING",
		"EVENTHUB_NAME",
		"RESOURCE_GROUP",
	})

	connProps, err := exported.ParseConnectionString(envVars["EVENTHUB_CONNECTION_STRING"])
	require.NoError(t, err)

	return ConnectionParamsForTest{
		ConnectionString:        envVars["EVENTHUB_CONNECTION_STRING"],
		EventHubName:            envVars["EVENTHUB_NAME"],
		EventHubNamespace:       connProps.FullyQualifiedNamespace,
		ResourceGroup:           envVars["RESOURCE_GROUP"],
		StorageConnectionString: envVars["CHECKPOINTSTORE_STORAGE_CONNECTION_STRING"],
		SubscriptionID:          envVars["AZURE_SUBSCRIPTION_ID"],
	}
}

func mustGetEnvironmentVars(t *testing.T, names []string) map[string]string {
	var missingVars []string
	envVars := map[string]string{}

	for _, name := range names {
		val := os.Getenv(name)

		if val == "" {
			missingVars = append(missingVars, name)
			continue
		}

		envVars[name] = val
	}

	if len(missingVars) > 0 {
		t.Skipf("Missing env vars for live test: %s. Skipping...", strings.Join(missingVars, ","))
		return nil
	}

	return envVars
}

func RequireClose(t *testing.T, closeable interface {
	Close(ctx context.Context) error
}) {
	require.NoError(t, closeable.Close(context.Background()))
}

// RequireContextHasDefaultTimeout checks that the context has a deadline set, and that it's
// using the right timeout.
// NOTE: There's some wiggle room since some time will expire before this is called.
func RequireContextHasDefaultTimeout(t *testing.T, ctx context.Context, timeout time.Duration) {
	tm, hasDeadline := ctx.Deadline()

	require.True(t, hasDeadline, "deadline must exist, we always set an operation timeout")
	duration := time.Until(tm)

	require.Greater(t, duration, time.Duration(0))
	require.LessOrEqual(t, duration, timeout)
}
