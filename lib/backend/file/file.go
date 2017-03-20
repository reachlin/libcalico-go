// Copyright (c) 2017 reachlin@gmail.com. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
  goerrors "errors"
  "strings"

  "time"

  log "github.com/Sirupsen/logrus"
  "github.com/projectcalico/libcalico-go/lib/backend/api"
  "github.com/projectcalico/libcalico-go/lib/backend/model"
  "github.com/projectcalico/libcalico-go/lib/errors"
  "golang.org/x/net/context"
)

type BackendConfig struct {
  EtcdScheme     string `json:"etcdScheme" envconfig:"ETCD_SCHEME" default:"http"`
  EtcdAuthority  string `json:"etcdAuthority" envconfig:"ETCD_AUTHORITY" default:"127.0.0.1:2379"`
  EtcdEndpoints  string `json:"etcdEndpoints" envconfig:"ETCD_ENDPOINTS"`
  EtcdUsername   string `json:"etcdUsername" envconfig:"ETCD_USERNAME"`
  EtcdPassword   string `json:"etcdPassword" envconfig:"ETCD_PASSWORD"`
  EtcdKeyFile    string `json:"etcdKeyFile" envconfig:"ETCD_KEY_FILE"`
  EtcdCertFile   string `json:"etcdCertFile" envconfig:"ETCD_CERT_FILE"`
  EtcdCACertFile string `json:"etcdCACertFile" envconfig:"ETCD_CA_CERT_FILE"`
}

type FileClient struct {
  fileName     string `json:"filename" envconfig:"FILE_NAME" default:"/etc/calico/backend.conf"`
}

func NewConfigFile(config *BackendConfig) (*FileClient, error) {
  return &FileClient{fileName: client}, nil
}

// EnsureInitialized makes sure that the etcd data is initialized for use by
// Calico.
func (c *FileClient) EnsureInitialized() error {
  return nil
}

// EnsureCalicoNodeInitialized performs additional initialization required
// by the calico/node components [startup/ipip-allocation/confd].  This is a
// temporary requirement until the calico/node components are updated to not
// require special etcd setup, or until the global and per-node config is
// reworked to allow the node to perform the necessary updates.
func (c *FileClient) EnsureCalicoNodeInitialized(node string) error {
  return nil
}
