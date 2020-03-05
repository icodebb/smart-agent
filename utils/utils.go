/**
 * Utility methods.
 *
 */

package utils

import (
	"fmt"
	"math/rand"
	"os"

	log "github.com/sirupsen/logrus"
)

// ModuleDict is used in logging.
var ModuleDict = map[int]string{
	0: "<MAIN>",
	1: "<NATS>",
	2: "<gRPC>",
	3: "<HTTP>",
}

const (
	// MAIN and following are all for logging.
	MAIN = 0
	// NATS is for NATS logging.
	NATS = 1
	// GRPC server
	GRPC = 2
	// HTTP server
	HTTP = 3
	// INFO the Info level
	INFO = 3
	// WARNING the warning level
	WARNING = 2
	// ERROR the error level
	ERROR = 1
	// FATAL the fatal level
	FATAL = 0
)

const (
	// KeyNatsServer key for NATS server
	KeyNatsServer = "CLUSTER_MANAGER_NATS_SERVER"
	// KeyNatsClusterID key for cluster ID
	KeyNatsClusterID = "CLUSTER_MANAGER_NATS_CLUSTER_ID"
	// KeyNatsSubjec key for subject
	KeyNatsSubjec = "CLUSTER_MANAGER_NATS_CLUSTER_SUBJECT"
	// KeyHTTPPort key for HTTP server port number
	KeyHTTPPort = "HTTP_SERVER_PORT"
)

// GetEnv gets variable from env.
func GetEnv(key string, defaultVal string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal, nil
	}
	return v, nil
}

// MyLog is a wrapper for logrus.
func MyLog(module int, level int, msg ...interface{}) {
	line := fmt.Sprintf("%v", msg...)
	line = ModuleDict[module] + " " + line

	switch level {
	case INFO:
		log.Info(line)

	case WARNING:
		log.Warning(line)

	case ERROR:
		log.Error(line)

	case FATAL:
		log.Fatal(line)

	default:
		log.Info(line)
	}
}

// MyLogf is a wrapper for logrus with format.
func MyLogf(module int, level int, format string, msg ...interface{}) {
	line := fmt.Sprintf(format, msg...)
	line = ModuleDict[module] + " " + line

	switch level {
	case INFO:
		log.Info(line)

	case WARNING:
		log.Warning(line)

	case ERROR:
		log.Error(line)

	case FATAL:
		log.Fatal(line)

	default:
		log.Info(line)
	}
}

// RandomInt returns an int >= min, < max
func RandomInt(min, max int) int {
	// rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// RandomString generates a random string of A-Z chars with len = l
func RandomString(len int) string {
	// rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}
