/*
* Copyright 2024 Google LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     https://www.apache.org/licenses/LICENSE-2.0
*
*     Unless required by applicable law or agreed to in writing, software
*     distributed under the License is distributed on an "AS IS" BASIS,
*     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*     See the License for the specific language governing permissions and
*     limitations under the License.
*/

package v1

// PolicyConditionType is a type of condition for a service policy.
type PolicyConditionType string

// PolicyConditionReason is a reason for a service policy condition.
type PolicyConditionReason string

const (
	// PolicyConditionAttached indicates whether the policy has been accepted or rejected
	// by a targeted resource, and why.
	//
	// Possible reasons for this condition to be true are:
	//
	// * "Attached"
	//
	// Possible reasons for this condition to be False are:
	//
	// * "Conflicted"
	//
	PolicyConditionAttached PolicyConditionType = "Attached"

	// PolicyReasonAttached is used with the "Attached" condition when the policy has been
	// accepted by the targeted resource.
	PolicyReasonAttached PolicyConditionReason = "Attached"

	// PolicyReasonConflicted is used with the "Attached" condition when the policy has not
	// been accepted by a targeted resource because there is another policy that targets the same
	// resource and has higher precedence.
	PolicyReasonConflicted PolicyConditionReason = "Conflicted"

	// PolicyReasonInvalid is used with the "Attached" condition when the policy is syntactically
	// or semantically invalid.
	PolicyReasonInvalid PolicyConditionReason = "Invalid"

	// PolicyReasonTargetNotFound is used with the "Attached" condition when the policy is attached to
	// an invalid target resource
	PolicyReasonTargetNotFound PolicyConditionReason = "TargetNotFound"
)
