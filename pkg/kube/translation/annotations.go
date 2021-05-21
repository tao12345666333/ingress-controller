// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package translation

import (
	"fmt"

	configv1 "github.com/apache/apisix-ingress-controller/pkg/kube/apisix/apis/config/v1"
	"github.com/apache/apisix-ingress-controller/pkg/kube/translation/annotations"
	"github.com/apache/apisix-ingress-controller/pkg/log"
	apisix "github.com/apache/apisix-ingress-controller/pkg/types/apisix/v1"
	"go.uber.org/zap"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
)

var (
	_handlers = []annotations.Handler{
		annotations.NewCorsHandler(),
		annotations.NewIPRestrictionHandler(),
	}
)

func (t *translator) translateAnnotations(ing interface{}) apisix.Plugins {
	var extractor annotations.Extractor
	switch v := ing.(type) {
	case *networkingv1.Ingress:
		extractor = annotations.NewExtractor(ing.(*networkingv1.Ingress).Annotations, ing.(*networkingv1.Ingress).Namespace, ing.(*networkingv1.Ingress).Name)
		fmt.Printf("%T", v)
	case *networkingv1beta1.Ingress:
		extractor = annotations.NewExtractor(ing.(*networkingv1beta1.Ingress).Annotations, ing.(*networkingv1beta1.Ingress).Namespace, ing.(*networkingv1beta1.Ingress).Name)
		fmt.Printf("%T", v)
	case *extensionsv1beta1.Ingress:
		extractor = annotations.NewExtractor(ing.(*extensionsv1beta1.Ingress).Annotations, ing.(*extensionsv1beta1.Ingress).Namespace, ing.(*extensionsv1beta1.Ingress).Name)
		fmt.Printf("%T", v)
	case *configv1.ApisixRoute:
		extractor = annotations.NewExtractor(ing.(*configv1.ApisixRoute).Annotations, ing.(*configv1.ApisixRoute).Namespace, ing.(*configv1.ApisixRoute).Name)
		fmt.Printf("%T", v)
	default:
		log.Error("unknown type")
		return nil
	}

	plugins := make(apisix.Plugins)
	for _, handler := range _handlers {
		out, err := handler.Handle(extractor)
		if err != nil {
			log.Warnw("failed to handle annotations",
				zap.Error(err),
			)
			continue
		}
		if out != nil {
			plugins[handler.PluginName()] = out
		}
	}
	return plugins
}
