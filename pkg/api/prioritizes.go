package api

import (
	"encoding/json"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

func priorities(r *restful.Request, w *restful.Response) {
	logrus.Infof("priorities begin ...")
	var extenderArgs schedulerapi.ExtenderArgs

	if err := r.ReadEntity(&extenderArgs); err != nil {
		logrus.Errorf("priorities read entity error: %v", err)
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	hostPriorityList := handlepriorities(extenderArgs)

	logrus.Debug("priorities pod: %v request: %#v", extenderArgs.Pod.Name, extenderArgs.NodeNames)

	body, err := json.Marshal(hostPriorityList)
	if err != nil {

	}
	w.Write(body)
}

func handlepriorities(args schedulerapi.ExtenderArgs) *schedulerapi.HostPriorityList {
	var priorityList schedulerapi.HostPriorityList
	priorityList = make([]schedulerapi.HostPriority, len(args.Nodes.Items))

	for i, node := range args.Nodes.Items {
		priorityList[i] = schedulerapi.HostPriority{
			Host:  node.Name,
			Score: 0,
		}
	}

	return &priorityList
}
