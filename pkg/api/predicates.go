package api

import (
	"encoding/json"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

func predicates(r *restful.Request, w *restful.Response) {
	logrus.Infof("predicates begin select route path: %s", r.SelectedRoutePath())
	var args schedulerapi.ExtenderArgs

	if err := r.ReadEntity(&args); err != nil {
		logrus.Errorf("predicate read entity error: %v", err)
		w.WriteError(http.StatusInternalServerError, err)
		return
	}

	if args.Pod == nil {
		msg := "predicate extender args pod is nil"
		logrus.Errorf("%s", msg)
		w.WriteErrorString(http.StatusInternalServerError, msg)
		return
	}

	extenderFilterResult := handleFilter(&args)

	logrus.Debugf("predicate pod: %v", args.Pod.Name)

	body, err := json.Marshal(extenderFilterResult)
	if err != nil {

	}
	w.Write(body)
}

func handleFilter(args *schedulerapi.ExtenderArgs) *schedulerapi.ExtenderFilterResult {
	pod := args.Pod
	canSchedule := make([]v1.Node, 0, len(args.Nodes.Items))
	canNotSchedule := make(map[string]string)

	for _, node := range args.Nodes.Items {
		result, err := predicateHandler(*pod, node)
		if err != nil {
			canNotSchedule[node.Name] = err.Error()
		} else if result {
			canSchedule = append(canSchedule, node)
		}
	}
	return &schedulerapi.ExtenderFilterResult{
		Nodes: &v1.NodeList{
			Items: canSchedule,
		},
		FailedNodes: canNotSchedule,
		Error:       "",
	}
}

func predicateHandler(pod v1.Pod, node v1.Node) (bool, error) {
	for _, vol := range pod.Spec.Volumes {
		logrus.Debugf("predicateHandler pod: %s, vol: %s", pod.Name, vol.Name)
		if vol.PersistentVolumeClaim == nil {
			continue
		}

		// pvc, err :=
	}

	return true, nil
}
