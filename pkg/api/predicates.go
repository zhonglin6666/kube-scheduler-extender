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
	logrus.Infof("predicates begin")
	var extenderArgs schedulerapi.ExtenderArgs

	if err := r.ReadEntity(&extenderArgs); err != nil {
		logrus.Errorf("predicate read entity error: %v", err)
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	extenderFilterResult := handleFilter(extenderArgs)

	logrus.Debugf("predicate pod: %v request: %#v", extenderArgs.Pod.Name, extenderArgs.NodeNames)

	body, err := json.Marshal(extenderFilterResult)
	if err != nil {

	}
	w.Write(body)
}

func handleFilter(args schedulerapi.ExtenderArgs) *schedulerapi.ExtenderFilterResult {
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
	return true, nil
}
