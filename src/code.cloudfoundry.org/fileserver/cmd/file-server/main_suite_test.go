package main_test

import (
	"code.cloudfoundry.org/bbs/test_helpers"
	"code.cloudfoundry.org/diego-logging-client/testhelpers"
	"code.cloudfoundry.org/go-loggregator/v9/rpc/loggregator_v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	fileServerBinary   string
	testIngressServer  *testhelpers.TestIngressServer
	metronIngressSetup *test_helpers.MetronIngressSetup

	testMetricsChan   chan *loggregator_v2.Envelope
	signalMetricsChan chan struct{}
)

func TestFileServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Server Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	fileServerPath, err := gexec.Build("code.cloudfoundry.org/fileserver/cmd/file-server")
	Expect(err).NotTo(HaveOccurred())
	return []byte(fileServerPath)
}, func(fileServerPath []byte) {
	fileServerBinary = string(fileServerPath)

})

var _ = BeforeEach(func() {
	var err error
	metronIngressSetup, err = test_helpers.StartMetronIngress("fixtures")
	Expect(err).NotTo(HaveOccurred())
	testIngressServer = metronIngressSetup.Server
	signalMetricsChan = metronIngressSetup.SignalMetricsChan
	testMetricsChan = metronIngressSetup.TestMetricsChan

})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
})

var _ = AfterEach(func() {
	testIngressServer.Stop()
	close(signalMetricsChan)

})
