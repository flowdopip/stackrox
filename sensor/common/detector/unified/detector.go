package unified

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/detection"
	"github.com/stackrox/rox/pkg/detection/deploytime"
	"github.com/stackrox/rox/pkg/detection/runtime"
	"github.com/stackrox/rox/pkg/logging"
)

var (
	log = logging.LoggerForModule()
)

// Detector is a thin layer atop the other detectors that provides a unified interface.
type Detector interface {
	ReconcilePolicies(newList []*storage.Policy)
	DetectDeployment(ctx deploytime.DetectionContext, deployment *storage.Deployment, images []*storage.Image) []*storage.Alert
	DetectProcess(deployment *storage.Deployment, images []*storage.Image, processIndicator *storage.ProcessIndicator, processOutsideWhitelist bool) []*storage.Alert
}

// NewDetector returns a new detector.
func NewDetector() Detector {
	return &detectorImpl{
		deploytimeDetector: deploytime.NewDetector(detection.NewPolicySet()),
		runtimeDetector:    runtime.NewDetector(detection.NewPolicySet()),
	}
}
