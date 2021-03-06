//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// CS??? removed icp-serviceid-apikey-secret from CommonSecretCheckNames, CommonSecretCheckDirs,
// CS???   and CommonSecretCheckVolumeMounts
// Linter doesn't like "Secret" in string var names so use "Zecret"

package resources

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const DefaultImageRegistry = "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom-amd64"
const DefaultDmImageName = "metering-data-manager"
const DefaultDmImageTag = "3.4.0"
const DefaultUIImageName = "metering-ui"
const DefaultUIImageTag = "3.4.0"
const DefaultMcmUIImageName = "metering-mcmui"
const DefaultMcmUIImageTag = "3.4.0"
const DefaultSenderImageName = "metering-data-manager"
const DefaultSenderImageTag = "3.4.0"

var TrueVar = true
var FalseVar = false
var Replica1 int32 = 1
var Seconds60 int64 = 60
var ManagementNodeSelector = map[string]string{"management": "true"}

var cpu100 = resource.NewMilliQuantity(100, resource.DecimalSI)          // 100m
var cpu500 = resource.NewMilliQuantity(500, resource.DecimalSI)          // 500m
var cpu1000 = resource.NewMilliQuantity(1000, resource.DecimalSI)        // 1000m
var memory100 = resource.NewQuantity(100*1024*1024, resource.BinarySI)   // 100Mi
var memory128 = resource.NewQuantity(128*1024*1024, resource.BinarySI)   // 128Mi
var memory256 = resource.NewQuantity(256*1024*1024, resource.BinarySI)   // 256Mi
var memory512 = resource.NewQuantity(512*1024*1024, resource.BinarySI)   // 512Mi
var memory2560 = resource.NewQuantity(2560*1024*1024, resource.BinarySI) // 2560Mi

const ClusterNameVar = "CLUSTER_NAME"
const HCClusterNameVar = "HC_CLUSTER_NAME"
const DefaultClusterName = "mycluster"

var CommonEnvVars = []corev1.EnvVar{
	{
		Name:  "NODE_TLS_REJECT_UNAUTHORIZED",
		Value: "0",
	},
}

var IAMEnvVars = []corev1.EnvVar{
	{
		Name:  "DEFAULT_IAM_TOKEN_SERVICE_PORT",
		Value: "10443",
	},
	{
		Name:  "DEFAULT_IAM_PAP_SERVICE_PORT",
		Value: "39001",
	},
}

var secretCheckCmd = `set -- $SECRET_LIST; ` +
	`for secretDirName in $SECRET_DIR_LIST; do` +
	`  while true; do` +
	`    echo ` + "`date`" + `: Checking for secret $1;` +
	`    ls /sec/$secretDirName/* && break;` +
	`    echo ` + "`date`" + `: Required secret $1 not found ... try again in 30s;` +
	`    sleep 30;` +
	`  done;` +
	`  echo ` + "`date`" + `: Secret $1 found;` +
	`  shift; ` +
	`done; ` +
	`echo ` + "`date`" + `: All required secrets exist`

var SenderSecretCheckCmd = secretCheckCmd + ";" +
	`echo ` + "`date`" + `: Further, checking for kubeConfig secret...;` +
	`node /datamanager/lib/metering_init.js kubeconfig_secretcheck `

var CommonZecretCheckNames = "icp-mongodb-admin icp-mongodb-admin cluster-ca-cert icp-mongodb-client-cert"

var CommonZecretCheckDirs = "muser-icp-mongodb-admin mpass-icp-mongodb-admin cluster-ca-cert icp-mongodb-client-cert"

var CommonSecretCheckVolumeMounts = []corev1.VolumeMount{
	{
		Name:      "mongodb-ca-cert",
		MountPath: "/sec/cluster-ca-cert",
	},
	{
		Name:      "mongodb-client-cert",
		MountPath: "/sec/icp-mongodb-client-cert",
	},
	{
		Name:      "muser-icp-mongodb-admin",
		MountPath: "/sec/muser-icp-mongodb-admin",
	},
	{
		Name:      "mpass-icp-mongodb-admin",
		MountPath: "/sec/mpass-icp-mongodb-admin",
	},
}

const APICertName = "icp-metering-api-ca-cert"
const APICertCommonName = "metering-server"
const APICertZecretName = "icp-metering-api-secret"
const APICertVolumeName = "icp-metering-api-certs"

var APICertVolumeMount = corev1.VolumeMount{
	Name:      APICertVolumeName,
	MountPath: "/sec/" + APICertZecretName,
}
var APICertVolume = corev1.Volume{
	Name: APICertVolumeName,
	VolumeSource: corev1.VolumeSource{
		Secret: &corev1.SecretVolumeSource{
			SecretName:  APICertZecretName,
			DefaultMode: &DefaultMode,
			Optional:    &TrueVar,
		},
	},
}

const ReceiverCertName = "icp-metering-receiver-ca-cert"
const ReceiverCertCommonName = "metering-receiver"
const ReceiverCertZecretName = "icp-metering-receiver-secret"
const ReceiverCertVolumeName = "icp-metering-receiver-certs"

var ReceiverCertVolumeMount = corev1.VolumeMount{
	Name:      ReceiverCertVolumeName,
	MountPath: "/sec/" + ReceiverCertZecretName,
}
var ReceiverCertVolume = corev1.Volume{
	Name: ReceiverCertVolumeName,
	VolumeSource: corev1.VolumeSource{
		Secret: &corev1.SecretVolumeSource{
			SecretName:  ReceiverCertZecretName,
			DefaultMode: &DefaultMode,
			Optional:    &TrueVar,
		},
	},
}

const PlatformOidcZecretName = "platform-oidc-credentials"
const PlatformOidcVolumeName = "platform-oidc-credentials"

var PlatformOidcVolumeMount = corev1.VolumeMount{
	Name:      PlatformOidcVolumeName,
	MountPath: "/sec/" + PlatformOidcZecretName,
}
var PlatformOidcVolume = corev1.Volume{
	Name: PlatformOidcVolumeName,
	VolumeSource: corev1.VolumeSource{
		Secret: &corev1.SecretVolumeSource{
			SecretName:  PlatformOidcZecretName,
			DefaultMode: &DefaultMode,
			Optional:    &TrueVar,
		},
	},
}

const APIKeyZecretName = "icp-serviceid-apikey-secret"
const APIKeyVolumeName = "icp-serviceid-apikey-secret"

var APIKeyVolumeMount = corev1.VolumeMount{
	Name:      APIKeyVolumeName,
	MountPath: "/sec/" + APIKeyZecretName,
}
var APIKeyVolume = corev1.Volume{
	Name: APIKeyVolumeName,
	VolumeSource: corev1.VolumeSource{
		Secret: &corev1.SecretVolumeSource{
			SecretName:  APIKeyZecretName,
			DefaultMode: &DefaultMode,
			Optional:    &TrueVar,
		},
	},
}

var commonInitVolumeMounts = []corev1.VolumeMount{
	{
		Name:      "mongodb-ca-cert",
		MountPath: "/certs/mongodb-ca",
	},
	{
		Name:      "mongodb-client-cert",
		MountPath: "/certs/mongodb-client",
	},
}

var commonInitResources = corev1.ResourceRequirements{
	Limits: map[corev1.ResourceName]resource.Quantity{
		corev1.ResourceCPU:    *cpu100,
		corev1.ResourceMemory: *memory100},
	Requests: map[corev1.ResourceName]resource.Quantity{
		corev1.ResourceCPU:    *cpu100,
		corev1.ResourceMemory: *memory100},
}

var commonSecurityContext = corev1.SecurityContext{
	AllowPrivilegeEscalation: &FalseVar,
	Privileged:               &FalseVar,
	ReadOnlyRootFilesystem:   &TrueVar,
	RunAsNonRoot:             &TrueVar,
	Capabilities: &corev1.Capabilities{
		Drop: []corev1.Capability{
			"ALL",
		},
	},
}

var CommonMainVolumeMounts = []corev1.VolumeMount{
	{
		Name:      "mongodb-ca-cert",
		MountPath: "/certs/mongodb-ca",
	},
	{
		Name:      "mongodb-client-cert",
		MountPath: "/certs/mongodb-client",
	},
}

var LoglevelVolumeMount = corev1.VolumeMount{
	Name:      "loglevel",
	MountPath: "/etc/config",
}

var Log4jsVolumeMount = corev1.VolumeMount{
	Name:      "log4js",
	MountPath: "/etc/config",
}

const SecretListVarNdx = 0
const SecretDirVarNdx = 1
const SecretCheckCmdNdx = 2

var BaseSecretCheckContainer = corev1.Container{
	Image:           "metering-data-manager",
	Name:            "metering-secret-check",
	ImagePullPolicy: corev1.PullAlways,
	Command: []string{
		"sh",
		"-c",
		secretCheckCmd,
	},
	Env: []corev1.EnvVar{
		{
			Name: "SECRET_LIST",
			// CommonZecretCheckNames will be added by the controller
			Value: "",
		},
		{
			// CommonZecretCheckDirs will be added by the controller
			Name:  "SECRET_DIR_LIST",
			Value: "",
		},
	},
	// CommonSecretCheckVolumeMounts will be added by the controller
	VolumeMounts:    []corev1.VolumeMount{},
	Resources:       commonInitResources,
	SecurityContext: &commonSecurityContext,
}

var BaseInitContainer = corev1.Container{
	Image:           "metering-data-manager",
	Name:            "metering-init",
	ImagePullPolicy: corev1.PullAlways,
	Command: []string{
		"node",
		"/datamanager/lib/metering_init.js",
		"verifyOnlyMongo",
	},
	// CommonEnvVars and mongoDBEnvVars will be added by the controller
	Env:             []corev1.EnvVar{},
	VolumeMounts:    commonInitVolumeMounts,
	Resources:       commonInitResources,
	SecurityContext: &commonSecurityContext,
}

var DmMainContainer = corev1.Container{
	Image: "metering-data-manager",
	//CS??? Image: "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom-amd64/metering-data-manager:3.3.1",
	Name:            "metering-dm",
	ImagePullPolicy: corev1.PullAlways,
	// CommonMainVolumeMounts will be added by the controller
	VolumeMounts: []corev1.VolumeMount{
		{
			Name:      ReceiverCertVolumeName,
			MountPath: "/certs/metering-receiver",
		},
		LoglevelVolumeMount,
	},
	// CommonEnvVars, IAMEnvVars and mongoDBEnvVars will be added by the controller
	Env: []corev1.EnvVar{
		{
			Name:  "METERING_API_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_USE_HTTPS",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCM_RECEIVER_ENABLED",
			Value: "false",
			//CS??? Value: "true",
		},
		{
			Name:  "HC_DM_MCM_SENDER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_STORAGEREADER_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_REPORTER2_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_PURGER2_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_PREAGGREGATOR_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_METRICS_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_SELFMETER_PURGER_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_IS_ICP",
			Value: "true",
		},
		{
			Name:  "HC_RECEIVER_SSL_CA",
			Value: "/certs/metering-receiver/ca.crt",
		},
		{
			Name:  "HC_RECEIVER_SSL_CERT",
			Value: "/certs/metering-receiver/tls.crt",
		},
		{
			Name:  "HC_RECEIVER_SSL_KEY",
			Value: "/certs/metering-receiver/tls.key",
		},
		{
			Name:  "HC_DM_ALLOW_TEST",
			Value: "false",
		},
	},
	Ports: []corev1.ContainerPort{
		{ContainerPort: 3000},
		{ContainerPort: 5000},
	},
	LivenessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/livenessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3000,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 305,
		TimeoutSeconds:      5,
		PeriodSeconds:       300,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	ReadinessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/readinessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3000,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 15,
		TimeoutSeconds:      15,
		PeriodSeconds:       30,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	Resources: corev1.ResourceRequirements{
		Limits: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu1000,
			corev1.ResourceMemory: *memory2560},
		Requests: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu100,
			corev1.ResourceMemory: *memory256},
	},
	SecurityContext: &commonSecurityContext,
}

var RdrMainContainer = corev1.Container{
	Image: "metering-data-manager",
	//CS??? Image: "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom-amd64/metering-data-manager:3.3.1",
	Name:            "metering-reader",
	ImagePullPolicy: corev1.PullAlways,
	// CommonMainVolumeMounts will be added by the controller
	VolumeMounts: []corev1.VolumeMount{
		{
			Name:      APICertVolumeName,
			MountPath: "/certs/metering-api",
		},
		LoglevelVolumeMount,
	},
	// CommonEnvVars, IAMEnvVars and mongoDBEnvVars will be added by the controller
	Env: []corev1.EnvVar{
		{
			Name:  "METERING_API_ENABLED",
			Value: "true",
		},
		{
			Name:  "METERING_INTERNALAPI_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_USE_HTTPS",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCM_RECEIVER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCM_SENDER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCMREADER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_READER_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_STORAGEREADER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_READER_APIENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_REPORTER2_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_PURGER2_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_PREAGGREGATOR_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_SELFMETER_PURGER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_IS_ICP",
			Value: "true",
		},
		{
			Name:  "HC_DM_API_PORT",
			Value: "4000",
		},
		{
			Name:  "HC_DM_INTERNAL_API_PORT",
			Value: "4002",
		},
		{
			Name:  "HC_API_SSL_CA",
			Value: "/certs/metering-api/ca.crt",
		},
		{
			Name:  "HC_API_SSL_CERT",
			Value: "/certs/metering-api/tls.crt",
		},
		{
			Name:  "HC_API_SSL_KEY",
			Value: "/certs/metering-api/tls.key",
		},
		{
			Name: "MY_NODE_NAME",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "spec.nodeName",
				},
			},
		},
	},
	Ports: []corev1.ContainerPort{
		{ContainerPort: 3000},
		{ContainerPort: 4000},
		{ContainerPort: 4002},
	},
	LivenessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/livenessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3000,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 305,
		TimeoutSeconds:      5,
		PeriodSeconds:       300,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	ReadinessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/readinessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3000,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 15,
		TimeoutSeconds:      15,
		PeriodSeconds:       30,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	Resources: corev1.ResourceRequirements{
		Limits: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu500,
			corev1.ResourceMemory: *memory512},
		Requests: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu100,
			corev1.ResourceMemory: *memory128},
	},
	SecurityContext: &commonSecurityContext,
}

var SenderMainContainer = corev1.Container{
	Image:           "metering-data-manager",
	Name:            "metering-sender",
	ImagePullPolicy: corev1.PullAlways,
	// CommonMainVolumeMounts will be added by the controller
	VolumeMounts: []corev1.VolumeMount{
		LoglevelVolumeMount,
	},
	// CommonEnvVars, IAMEnvVars and mongoDBEnvVars will be added by the controller
	Env: []corev1.EnvVar{
		{
			Name:  "METERING_API_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_SELFMETER_PURGER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_REPORTER2_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_PURGER2_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_PREAGGREGATOR_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_METRICS_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_READER_APIENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCM_RECEIVER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCMREADER_ENABLED",
			Value: "false",
		},
		{
			Name:  "HC_DM_MCM_SENDER_ENABLED",
			Value: "true",
		},
		{
			Name:  "HC_DM_IS_ICP",
			Value: "true",
		},
		{
			Name:  "HC_DM_ALLOW_TEST",
			Value: "false",
		},
	},
	LivenessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/livenessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3000,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 305,
		TimeoutSeconds:      5,
		PeriodSeconds:       300,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	ReadinessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/readinessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3000,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 15,
		TimeoutSeconds:      15,
		PeriodSeconds:       30,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	Resources: corev1.ResourceRequirements{
		Limits: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu500,
			corev1.ResourceMemory: *memory512},
		Requests: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu100,
			corev1.ResourceMemory: *memory128},
	},
	SecurityContext: &commonSecurityContext,
}

var UIEnvVars = []corev1.EnvVar{
	{
		Name:  "IS_PRIVATECLOUD",
		Value: "true",
	},
	{
		Name: "ICP_API_KEY",
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: APIKeyZecretName,
				},
				Key:      "ICP_API_KEY",
				Optional: &TrueVar,
			},
		},
	},
	{
		Name: "WLP_CLIENT_ID",
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: PlatformOidcZecretName,
				},
				Key:      "WLP_CLIENT_ID",
				Optional: &TrueVar,
			},
		},
	},
	{
		Name: "WLP_CLIENT_SECRET",
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: PlatformOidcZecretName,
				},
				Key:      "WLP_CLIENT_SECRET",
				Optional: &TrueVar,
			},
		},
	},
	{
		Name:  "USE_PRIVATECLOUD_SECURITY",
		Value: "true",
	},
	{
		Name:  "DEFAULT_PLATFORM_IDENTITY_MANAGEMENT_SERVICE_PORT",
		Value: "4500",
	},
	{
		Name:  "DEFAULT_PLATFORM_HEADER_SERVICE_PORT",
		Value: "3000",
	},
	{
		Name:  "HC_DM_ALLOW_TEST",
		Value: "false",
	},
}

var UIMainContainer = corev1.Container{
	Image: "metering-ui",
	//CS??? Image: "hyc-cloud-private-edge-docker-local.artifactory.swg-devops.com/ibmcom-amd64/metering-ui:3.3.1",
	Name:            "metering-ui",
	ImagePullPolicy: corev1.PullAlways,
	// CommonMainVolumeMounts will be added by the controller
	VolumeMounts: []corev1.VolumeMount{
		LoglevelVolumeMount,
	},
	// CommonEnvVars, IAMEnvVars, UIEnvVars and mongoDBEnvVars will be added by the controller. removed HC_HCAPI_URI
	Env: []corev1.EnvVar{
		{
			Name:  "ICP_DEFAULT_DASHBOARD",
			Value: "cpi.icp.main",
		},
		{
			Name:  "PORT",
			Value: "3130",
		},
		{
			Name:  "PROXY_URI",
			Value: "metering",
		},
	},
	Ports: []corev1.ContainerPort{
		{ContainerPort: 3130},
	},
	LivenessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/unsecure/c2c/status/livenessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3130,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 305,
		TimeoutSeconds:      5,
		PeriodSeconds:       300,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	ReadinessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/unsecure/c2c/status/readinessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3130,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 15,
		TimeoutSeconds:      15,
		PeriodSeconds:       30,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	Resources: corev1.ResourceRequirements{
		Limits: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu500,
			corev1.ResourceMemory: *memory512},
		Requests: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu100,
			corev1.ResourceMemory: *memory128},
	},
	SecurityContext: &commonSecurityContext,
}

var McmUIMainContainer = corev1.Container{
	Image:           "metering-mcmui",
	Name:            "metering-mcmui",
	ImagePullPolicy: corev1.PullAlways,
	// CommonMainVolumeMounts will be added by the controller
	VolumeMounts: []corev1.VolumeMount{
		Log4jsVolumeMount,
	},
	// CommonEnvVars, IAMEnvVars, UIEnvVars and mongoDBEnvVars will be added by the controller
	Env: []corev1.EnvVar{
		{
			Name:  "PORT",
			Value: "3001",
		},
		{
			Name:  "PROXY_URI",
			Value: "metering-mcm",
		},
	},
	Ports: []corev1.ContainerPort{
		{ContainerPort: 3001},
	},
	LivenessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/unsecure/livenessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3001,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 305,
		TimeoutSeconds:      5,
		PeriodSeconds:       300,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	ReadinessProbe: &corev1.Probe{
		Handler: corev1.Handler{
			HTTPGet: &corev1.HTTPGetAction{
				Path: "/unsecure/readinessProbe",
				Port: intstr.IntOrString{
					Type:   intstr.Int,
					IntVal: 3001,
				},
				Scheme: "",
			},
		},
		InitialDelaySeconds: 15,
		TimeoutSeconds:      5,
		PeriodSeconds:       15,
		SuccessThreshold:    1,
		FailureThreshold:    3,
	},
	Resources: corev1.ResourceRequirements{
		Limits: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu500,
			corev1.ResourceMemory: *memory256},
		Requests: map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    *cpu100,
			corev1.ResourceMemory: *memory128},
	},
	SecurityContext: &commonSecurityContext,
}
