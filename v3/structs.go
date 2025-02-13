/*
Copyright 2021 U. Cirello (cirello.io and github.com/cirello-io)

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

package dynamolock

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type acquireLockOptions struct {
	partitionKey                string
	data                        []byte
	replaceData                 bool
	deleteLockOnRelease         bool
	failIfLocked                bool
	refreshPeriod               time.Duration
	additionalTimeToWaitForLock time.Duration
	additionalAttributes        map[string]types.AttributeValue
	sessionMonitor              *sessionMonitor
}

type getLockOptions struct {
	partitionKey                      string
	deleteLockOnRelease               bool
	millisecondsToWait                time.Duration
	refreshPeriodDuration             time.Duration
	lockTryingToBeAcquired            *Lock
	alreadySleptOnceForOneLeasePeriod bool
	sessionMonitor                    *sessionMonitor
	start                             time.Time
	replaceData                       bool
	data                              []byte
	additionalAttributes              map[string]types.AttributeValue
	failIfLocked                      bool
}

type releaseLockOptions struct {
	lockItem   *Lock
	deleteLock bool
	data       []byte
}

type createDynamoDBTableOptions struct {
	billingMode           types.BillingMode
	provisionedThroughput *types.ProvisionedThroughput
	tags                  []types.Tag
}
