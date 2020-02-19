/* Copyright 2020 The Fedlearner Authors. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/bytedance/fedlearner/deploy/kubernetes_operator/pkg/apis/fedlearner.k8s.io/v1alpha1"
	"github.com/bytedance/fedlearner/deploy/kubernetes_operator/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type FedlearnerV1alpha1Interface interface {
	RESTClient() rest.Interface
	FLAppsGetter
}

// FedlearnerV1alpha1Client is used to interact with features provided by the fedlearner.k8s.io group.
type FedlearnerV1alpha1Client struct {
	restClient rest.Interface
}

func (c *FedlearnerV1alpha1Client) FLApps(namespace string) FLAppInterface {
	return newFLApps(c, namespace)
}

// NewForConfig creates a new FedlearnerV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*FedlearnerV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &FedlearnerV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new FedlearnerV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *FedlearnerV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new FedlearnerV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *FedlearnerV1alpha1Client {
	return &FedlearnerV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FedlearnerV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
