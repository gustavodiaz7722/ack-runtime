// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package types

import (
	ctrlrt "sigs.k8s.io/controller-runtime"
)

// StatusReconciler is responsible for reconciling an adopted resource conditions
// It acts a wrapper over Status controller implemented in github.com/awslabs/operatorpkg/status
type StatusReconciler interface {
	// BindControllerManager sets up the StatusReconciler with an
	// instance of an upstream controller-runtime.Manager
	BindControllerManager(ctrlrt.Manager) error
}
