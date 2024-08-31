package helmservice

import (
	"context"
	"testing"
)

func TestHelmInstall(t *testing.T) {

	chart := &ChartParameter{
		ReleaseName: "sample-1",
		Namespace:   "sample-1",
		ChartName:   "brbarmex-tenant/tenants",
		DryRun:      false,
	}

	err := HelmInstall(context.TODO(), chart)
	if err != nil {
		t.Error(err)
	}

}
