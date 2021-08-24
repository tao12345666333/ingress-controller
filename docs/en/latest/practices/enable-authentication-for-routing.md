---
title: Enable authentication for routing 
---

<!--
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
-->

In this practice, we will explain how to enable authentication for APISIX routing. 

## Prerequisites

* Prepare an available Kubernetes cluster in your workstation, we recommend you to use [KIND](https://kind.sigs.k8s.io/docs/user/quick-start/) to create a local Kubernetes cluster.
* Install Apache APISIX in Kubernetes by [Helm Chart](https://github.com/apache/apisix-helm-chart).
* Install [apisix-ingress-controller](https://github.com/apache/apisix-ingress-controller/blob/master/install.md).

## Deploy httpbin service

We use [kennethreitz/httpbin](https://hub.docker.com/r/kennethreitz/httpbin/) as the service image, see its overview page for details.

Now, try to deploy it to your Kubernetes cluster:

```shell
kubectl run httpbin --image kennethreitz/httpbin --port 80
kubectl expose pod httpbin --port 80
```

