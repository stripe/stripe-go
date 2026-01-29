//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
)

// v2CoreBatchJobService is used to invoke batchjob related APIs.
type v2CoreBatchJobService struct {
	B   Backend
	Key string
}

// Creates a new batch job.
func (c v2CoreBatchJobService) Create(ctx context.Context, params *V2CoreBatchJobCreateParams) (*V2CoreBatchJob, error) {
	if params == nil {
		params = &V2CoreBatchJobCreateParams{}
	}
	params.Context = ctx
	batchjob := &V2CoreBatchJob{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/batch_jobs", c.Key, params, batchjob)
	return batchjob, err
}

// Serializes a Customer update request into a batch job JSONL line.
func (c v2CoreBatchJobService) SerializeV1CustomerUpdate(ctx context.Context, id string, params *CustomerUpdateParams) (string, error) {
	// Generate a UUID v4 using crypto/rand
	var uuidBytes [16]byte
	if _, err := rand.Read(uuidBytes[:]); err != nil {
		return "", err
	}
	uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40 // version 4
	uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80 // variant bits
	itemID := fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:16])

	item := BatchItemParams{
		ID:            itemID,
		PathParams:    map[string]string{"customer": id},
		Params:        params,
		StripeVersion: APIVersion, // Default to global API version
	}
	// Override StripeVersion and Context from the params if specified
	if container, ok := interface{}(params).(batchItemParamsContainer); ok {
		if v := container.GetStripeVersion(); v != nil {
			item.StripeVersion = *v
		}
		if c := container.GetStripeContext(); c != nil {
			item.Context = *c
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Serializes a SubscriptionSchedule create request into a batch job JSONL line.
func (c v2CoreBatchJobService) SerializeV1SubscriptionScheduleCreate(ctx context.Context, params *SubscriptionScheduleCreateParams) (string, error) {
	// Generate a UUID v4 using crypto/rand
	var uuidBytes [16]byte
	if _, err := rand.Read(uuidBytes[:]); err != nil {
		return "", err
	}
	uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40 // version 4
	uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80 // variant bits
	itemID := fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:16])

	item := BatchItemParams{
		ID:            itemID,
		PathParams:    nil,
		Params:        params,
		StripeVersion: APIVersion, // Default to global API version
	}
	// Override StripeVersion and Context from the params if specified
	if container, ok := interface{}(params).(batchItemParamsContainer); ok {
		if v := container.GetStripeVersion(); v != nil {
			item.StripeVersion = *v
		}
		if c := container.GetStripeContext(); c != nil {
			item.Context = *c
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Serializes a Subscription update request into a batch job JSONL line.
func (c v2CoreBatchJobService) SerializeV1SubscriptionUpdate(ctx context.Context, id string, params *SubscriptionUpdateParams) (string, error) {
	// Generate a UUID v4 using crypto/rand
	var uuidBytes [16]byte
	if _, err := rand.Read(uuidBytes[:]); err != nil {
		return "", err
	}
	uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40 // version 4
	uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80 // variant bits
	itemID := fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:16])

	item := BatchItemParams{
		ID:            itemID,
		PathParams:    map[string]string{"subscription_exposed_id": id},
		Params:        params,
		StripeVersion: APIVersion, // Default to global API version
	}
	// Override StripeVersion and Context from the params if specified
	if container, ok := interface{}(params).(batchItemParamsContainer); ok {
		if v := container.GetStripeVersion(); v != nil {
			item.StripeVersion = *v
		}
		if c := container.GetStripeContext(); c != nil {
			item.Context = *c
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
