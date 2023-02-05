package main

import (
	"app/dao"
	"app/domain"
	_ "app/utils"
	"app/view"
)

func main() {
	for {
		menu := view.GetMenu()
		if menu == 4 {
			view.PrintToExit()
			break
		}
		navigate(menu)
	}
}

func navigate(menu int64) {
	switch menu {
	case 1:
		doInsert()
	case 2:
		doGetIdList()
	case 3:
		doGetData()

	default:
		view.PrintWrongInput()
	}

}

func doInsert() {
	// fileName := view.GetFileInput()
	test_string := `{"controller": {"name": "controller", "image": {"repository": "quay.io/kubernetes-ingress-controller/nginx-ingress-controller", "tag": "0.32.0", "pullPolicy": "IfNotPresent", "runAsUser": 101, "allowPrivilegeEscalation": true}, "useComponentLabel": false, "containerPort": {"http": 80, "https": 443}, "config": {}, "maxmindLicenseKey": "", "proxySetHeaders": {}, "addHeaders": {}, "hostNetwork": false, "dnsConfig": {}, "dnsPolicy": "ClusterFirst", "reportNodeInternalIp": false, "daemonset": {"useHostPort": false, "hostPorts": {"http": 80, "https": 443}}, "defaultBackendService": "", "electionID": "ingress-controller-leader", "ingressClass": "nginx", "deploymentLabels": {}, "podLabels": {}, "podSecurityContext": {}, "publishService": {"enabled": false, "pathOverride": ""}, "scope": {"enabled": false, "namespace": ""}, "configMapNamespace": "", "tcp": {"configMapNamespace": ""}, "udp": {"configMapNamespace": ""}, "extraArgs": {}, "extraEnvs": [], "kind": "Deployment", "deploymentAnnotations": {}, "updateStrategy": {}, "minReadySeconds": 0, "tolerations": [], "affinity": {}, "terminationGracePeriodSeconds": 60, "nodeSelector": {}, "livenessProbe": {"failureThreshold": 3, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 1, "port": 10254}, "readinessProbe": {"failureThreshold": 3, "initialDelaySeconds": 10, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 1, "port": 10254}, "podAnnotations": {}, "replicaCount": 1, "minAvailable": 1, "resources": {}, "autoscaling": {"enabled": false, "minReplicas": 2, "maxReplicas": 11, "targetCPUUtilizationPercentage": 50, "targetMemoryUtilizationPercentage": 50}, "customTemplate": {"configMapName": "", "configMapKey": ""}, "service": {"enabled": true, "annotations": {}, "labels": {}, "omitClusterIP": false, "externalIPs": [], "loadBalancerIP": "", "loadBalancerSourceRanges": [], "enableHttp": true, "enableHttps": true, "externalTrafficPolicy": "", "sessionAffinity": "", "healthCheckNodePort": 0, "ports": {"http": 80, "https": 443}, "targetPorts": {"http": "http", "https": "https"}, "type": "LoadBalancer", "nodePorts": {"http": "", "https": "", "tcp": {}, "udp": {}}}, "extraContainers": [], "extraVolumeMounts": [], "extraVolumes": [], "extraInitContainers": [], "admissionWebhooks": {"enabled": false, "failurePolicy": "Fail", "port": 8443, "service": {"annotations": {}, "omitClusterIP": false, "externalIPs": [], "loadBalancerIP": "", "loadBalancerSourceRanges": [], "servicePort": 443, "type": "ClusterIP"}, "patch": {"enabled": true, "image": {"repository": "jettech/kube-webhook-certgen", "tag": "v1.0.0", "pullPolicy": "IfNotPresent"}, "priorityClassName": "", "podAnnotations": {}, "nodeSelector": {}}}, "metrics": {"port": 10254, "enabled": false, "service": {"annotations": {}, "omitClusterIP": false, "externalIPs": [], "loadBalancerIP": "", "loadBalancerSourceRanges": [], "servicePort": 9913, "type": "ClusterIP"}, "serviceMonitor": {"enabled": false, "additionalLabels": {}, "namespace": "", "namespaceSelector": {}, "scrapeInterval": "30s"}, "prometheusRule": {"enabled": false, "additionalLabels": {}, "namespace": "", "rules": []}}, "lifecycle": {}, "priorityClassName": ""}, "revisionHistoryLimit": 10, "defaultBackend": {"enabled": true, "name": "default-backend", "image": {"repository": "k8s.gcr.io/defaultbackend-amd64", "tag": "1.5", "pullPolicy": "IfNotPresent", "runAsUser": 65534}, "useComponentLabel": false, "extraArgs": {}, "serviceAccount": {"create": true, "name": null}, "extraEnvs": [], "port": 8080, "livenessProbe": {"failureThreshold": 3, "initialDelaySeconds": 30, "periodSeconds": 10, "successThreshold": 1, "timeoutSeconds": 5}, "readinessProbe": {"failureThreshold": 6, "initialDelaySeconds": 0, "periodSeconds": 5, "successThreshold": 1, "timeoutSeconds": 5}, "tolerations": [], "affinity": {}, "podSecurityContext": {}, "deploymentLabels": {}, "podLabels": {}, "nodeSelector": {}, "podAnnotations": {}, "replicaCount": 1, "minAvailable": 1, "resources": {}, "service": {"annotations": {}, "omitClusterIP": false, "externalIPs": [], "loadBalancerIP": "", "loadBalancerSourceRanges": [], "servicePort": 80, "type": "ClusterIP"}, "priorityClassName": ""}, "releaseLabelOverride": "", "rbac": {"create": true, "scope": false}, "podSecurityPolicy": {"enabled": false}, "serviceAccount": {"create": true, "name": null, "annotations": {}}, "imagePullSecrets": [], "tcp": {}, "udp": {}}`
	data := domain.Data{}
	data.MapDataToMongo(test_string)
	dao.InsertData(data)
}

func doGetIdList() {

}

func doGetData() {

}
