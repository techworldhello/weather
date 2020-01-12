package mock

import (
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

var testLog, _ = test.NewNullLogger()
var TestLogger = logrus.NewEntry(testLog)
