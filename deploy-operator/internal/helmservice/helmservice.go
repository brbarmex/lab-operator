package helmservice

import (
	"context"
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
)

type ChartParameter struct {
	Namespace     string
	ReleaseName   string
	ReleaseVerion string
	ChartName     string
	DryRun        bool
}

func HelmInstall(ctx context.Context, chartParameter *ChartParameter) error {

	settings := cli.New()
	actionCfg := new(action.Configuration)
	settings.SetNamespace(chartParameter.Namespace)

	if err := actionCfg.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return err
	}

	actionCfg.KubeClient.(*kube.Client).Namespace = chartParameter.Namespace
	actionInstall := action.NewInstall(actionCfg)
	actionInstall.Namespace = chartParameter.Namespace
	actionInstall.ReleaseName = chartParameter.ReleaseName
	actionInstall.DryRun = chartParameter.DryRun
	actionInstall.ChartPathOptions.InsecureSkipTLSverify = true
	actionInstall.ChartPathOptions.Verify = false

	chartPath, err := actionInstall.ChartPathOptions.LocateChart(chartParameter.ChartName, settings)
	if err != nil {
		return err
	}

	chart, err := loader.Load(chartPath)
	if err != nil {
		return err
	}

	vals := map[string]interface{}{}
	releaseName, err := actionInstall.RunWithContext(ctx, chart, vals)
	if err != nil {
		return err
	}

	fmt.Sprintln(releaseName)

	return nil

}
