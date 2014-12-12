package stripe

import (
	"math/rand"
	"time"
)

// A function that returns an error upon failure, and nil upon success.
type Operation func() (retryable, unretryable error)

type UnsuccessfulResponseStrategy interface {
	Submit(op Operation) error
}

type UnsuccessfulResponseHandler interface {
	// Returning Inf implies that we should just exit now.
	Advice() (delay time.Duration, stop bool)
}

func (strategy *ExponentialBackoffStrategy) Submit(op Operation) error {
	var attempt float64
	var retryable, unretryable error
	var sleepDuration time.Duration
	var stop bool

	handler := strategy.NewHandler()

	for attempt = 0; ; attempt++ {
		retryable, unretryable = op()

		// Stop if we encounter an unretryable error
		if unretryable != nil {
			return unretryable
		}

		// Stop if there was no error that we should retry
		if retryable == nil {
			break
		}

		if sleepDuration, stop = handler.Advice(); stop {
			break
		}

		time.Sleep(sleepDuration)
	}
	return unretryable
}

type ExponentialBackoffStrategy struct {
	MaxElapsedTime time.Duration

	MaxRetryInterval     time.Duration
	InitialRetryInterval time.Duration
	IntervalMultiplier   float64

	RandomizationFactor float64
}

type ExponentialBackoffHandler struct {
	MaxElapsedTime time.Duration

	MaxRetryInterval     time.Duration
	InitialRetryInterval time.Duration
	IntervalMultiplier   float64

	RandomizationFactor float64

	CurrentInterval time.Duration

	// Implementation-specific data
	Started time.Time
}

// TODO: document what numbers do.
// TODO: revisit these values.
const (
	// The default initial time that a delay will occur before a retry occurs
	DEFAULT_INITIAL_RETRY_INTERVAL_MS = 500 * time.Millisecond

	// The default randomization factor. 0.5 results in a number varying between 50% and 150% the initial number.
	DEFAULT_RANDOMIZATION_FACTOR = 0.5

	// The default multiplier. 1.5 is a 50% increase per back-off
	DEFAULT_INTERVAL_MULTIPLIER = 1.5

	// The maximum interval that will ever occur, after randomization.
	DEFAULT_MAX_INDIVIDUAL_INTERVAL = 1 * time.Minute

	// The maximum total elapsed time that will occur before a request fails.
	DEFAULT_MAX_ELAPSED_TIME = 15 * time.Minute
)

func (strategy *ExponentialBackoffStrategy) NewHandler() UnsuccessfulResponseHandler {
	MaxElapsedTime := strategy.MaxElapsedTime
	if MaxElapsedTime == 0 {
		MaxElapsedTime = DEFAULT_MAX_ELAPSED_TIME
	}
	MaxRetryInterval := strategy.MaxRetryInterval
	if MaxRetryInterval == 0 {
		// TODO: rename MaxRetryInterval -> MaxIndividualInterval
		MaxRetryInterval = DEFAULT_MAX_INDIVIDUAL_INTERVAL
	}
	InitialRetryInterval := strategy.InitialRetryInterval
	if InitialRetryInterval == 0 {
		InitialRetryInterval = DEFAULT_INITIAL_RETRY_INTERVAL_MS
	}
	IntervalMultiplier := strategy.IntervalMultiplier
	if IntervalMultiplier == 0 {
		IntervalMultiplier = DEFAULT_INTERVAL_MULTIPLIER
	}
	RandomizationFactor := strategy.RandomizationFactor
	if RandomizationFactor == 0 {
		RandomizationFactor = DEFAULT_RANDOMIZATION_FACTOR
	}
	return &ExponentialBackoffHandler{
		MaxElapsedTime:       MaxElapsedTime,
		MaxRetryInterval:     MaxRetryInterval,
		InitialRetryInterval: InitialRetryInterval,
		IntervalMultiplier:   IntervalMultiplier,
		RandomizationFactor:  RandomizationFactor,
		CurrentInterval:      InitialRetryInterval,
		Started:              time.Now(),
	}
}

func (handler *ExponentialBackoffHandler) Advice() (recommendedDelay time.Duration, stop bool) {

	// Don't run again if we've already passed the deadline.
	if time.Now().Sub(handler.Started) > handler.MaxElapsedTime {
		return 0, true
	}

	// TODO: use fewer steps to calculate the range

	// Choose a number in the range [1- handler.RandomizationFactor, 1 + handler.RandomizationFactor]
	// [0, 1) -> [-1, 1)
	intermediateRandom1 := (rand.Float64() - 0.5) * 2
	// [-1, 1) -> [-handler.RandomizationFactor, handler.RandomizationFactor)
	intermediateRandom2 := intermediateRandom1 * handler.RandomizationFactor
	// [-handler.RandomizationFactor, handler.RandomizationFactor) -> [1 - handler.RandomizationFactor, 1 + handler.RandomizationFactor)
	intermediateRandom3 := 1 + intermediateRandom2

	// TODO: check for overflow.
	// TODO/DOING: weeping about type coercion
	recommendedDelay = time.Duration(int64(float64(handler.CurrentInterval.Nanoseconds()) * intermediateRandom3))

	if recommendedDelay > handler.MaxRetryInterval {
		recommendedDelay = handler.MaxRetryInterval
	}

	handler.updateCurrentInterval()

	// Don't re-run if delaying would cause us to go past the MaxElapsedTime
	if time.Now().Add(recommendedDelay).Sub(handler.Started) > handler.MaxElapsedTime {
		return 0, true
	}

	return recommendedDelay, false
}

func (handler *ExponentialBackoffHandler) updateCurrentInterval() {
	proposedInterval := time.Duration(float64(handler.CurrentInterval) * handler.IntervalMultiplier)
	// Check for overflow
	if proposedInterval < handler.CurrentInterval {
		handler.CurrentInterval = handler.MaxRetryInterval
		// Check for an interval larger than the MaxRetryInterval
	} else if proposedInterval > handler.MaxRetryInterval {
		handler.CurrentInterval = handler.MaxRetryInterval
	} else {
		handler.CurrentInterval = proposedInterval
	}
}
