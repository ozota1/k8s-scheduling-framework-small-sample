package main

import (
	"context"
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type SampleScheduler struct {
	framework.PreFilterPlugin
	framework.PostBindPlugin
}

const Name = "SampleScheduler"

func (pl *SampleScheduler) Name() string {
	return Name
}

func (cs *SampleScheduler) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.Infof("pre filter called for pod %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}
func (cs *SampleScheduler) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}
func (cs *SampleScheduler) PostBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.Infof("pod %v is binded to %v", pod.Name, nodeName)
}
func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &SampleScheduler{}, nil
}
func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(Name, New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
