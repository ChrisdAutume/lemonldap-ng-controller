/*
Copyright 2018 Mathieu Parent <math.parent@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"time"

	clientset "k8s.io/client-go/kubernetes"

	"github.com/lemonldap-ng-controller/lemonldap-ng-controller/internal/filesystem"
)

// Configuration stores lemonldap-ng-controller configuration
type Configuration struct {
	APIServerHost  string
	KubeConfigFile string
	Client         clientset.Interface

	ResyncPeriod time.Duration

	ConfigMapName string

	Namespace string

	ForceNamespaceIsolation bool

	FS                              filesystem.FileSystem
	LemonLDAPConfigurationDirectory string
}
