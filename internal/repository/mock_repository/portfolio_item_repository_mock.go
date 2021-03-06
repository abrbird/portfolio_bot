package mock_repository

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/abrbird/portfolio_bot/internal/repository.PortfolioItemRepository -o ./portfolio_item_repository_mock_test.go -n PortfolioItemRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/abrbird/portfolio_bot/internal/domain"
)

// PortfolioItemRepositoryMock implements PortfolioItemRepository
type PortfolioItemRepositoryMock struct {
	t minimock.Tester

	funcDelete          func(ctx context.Context, portfolioItemId int64) (err error)
	inspectFuncDelete   func(ctx context.Context, portfolioItemId int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mPortfolioItemRepositoryMockDelete

	funcRetrieveOrCreate          func(ctx context.Context, portfolioData domain.PortfolioItemCreate) (p1 domain.PortfolioItemRetrieve)
	inspectFuncRetrieveOrCreate   func(ctx context.Context, portfolioData domain.PortfolioItemCreate)
	afterRetrieveOrCreateCounter  uint64
	beforeRetrieveOrCreateCounter uint64
	RetrieveOrCreateMock          mPortfolioItemRepositoryMockRetrieveOrCreate

	funcRetrievePortfolioItems          func(ctx context.Context, portfolioId int64) (pp1 *domain.PortfolioItemsRetrieve)
	inspectFuncRetrievePortfolioItems   func(ctx context.Context, portfolioId int64)
	afterRetrievePortfolioItemsCounter  uint64
	beforeRetrievePortfolioItemsCounter uint64
	RetrievePortfolioItemsMock          mPortfolioItemRepositoryMockRetrievePortfolioItems
}

// NewPortfolioItemRepositoryMock returns a mock for PortfolioItemRepository
func NewPortfolioItemRepositoryMock(t minimock.Tester) *PortfolioItemRepositoryMock {
	m := &PortfolioItemRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mPortfolioItemRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*PortfolioItemRepositoryMockDeleteParams{}

	m.RetrieveOrCreateMock = mPortfolioItemRepositoryMockRetrieveOrCreate{mock: m}
	m.RetrieveOrCreateMock.callArgs = []*PortfolioItemRepositoryMockRetrieveOrCreateParams{}

	m.RetrievePortfolioItemsMock = mPortfolioItemRepositoryMockRetrievePortfolioItems{mock: m}
	m.RetrievePortfolioItemsMock.callArgs = []*PortfolioItemRepositoryMockRetrievePortfolioItemsParams{}

	return m
}

type mPortfolioItemRepositoryMockDelete struct {
	mock               *PortfolioItemRepositoryMock
	defaultExpectation *PortfolioItemRepositoryMockDeleteExpectation
	expectations       []*PortfolioItemRepositoryMockDeleteExpectation

	callArgs []*PortfolioItemRepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// PortfolioItemRepositoryMockDeleteExpectation specifies expectation struct of the PortfolioItemRepository.Delete
type PortfolioItemRepositoryMockDeleteExpectation struct {
	mock    *PortfolioItemRepositoryMock
	params  *PortfolioItemRepositoryMockDeleteParams
	results *PortfolioItemRepositoryMockDeleteResults
	Counter uint64
}

// PortfolioItemRepositoryMockDeleteParams contains parameters of the PortfolioItemRepository.Delete
type PortfolioItemRepositoryMockDeleteParams struct {
	ctx             context.Context
	portfolioItemId int64
}

// PortfolioItemRepositoryMockDeleteResults contains results of the PortfolioItemRepository.Delete
type PortfolioItemRepositoryMockDeleteResults struct {
	err error
}

// Expect sets up expected params for PortfolioItemRepository.Delete
func (mmDelete *mPortfolioItemRepositoryMockDelete) Expect(ctx context.Context, portfolioItemId int64) *mPortfolioItemRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("PortfolioItemRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &PortfolioItemRepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &PortfolioItemRepositoryMockDeleteParams{ctx, portfolioItemId}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the PortfolioItemRepository.Delete
func (mmDelete *mPortfolioItemRepositoryMockDelete) Inspect(f func(ctx context.Context, portfolioItemId int64)) *mPortfolioItemRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for PortfolioItemRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by PortfolioItemRepository.Delete
func (mmDelete *mPortfolioItemRepositoryMockDelete) Return(err error) *PortfolioItemRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("PortfolioItemRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &PortfolioItemRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &PortfolioItemRepositoryMockDeleteResults{err}
	return mmDelete.mock
}

//Set uses given function f to mock the PortfolioItemRepository.Delete method
func (mmDelete *mPortfolioItemRepositoryMockDelete) Set(f func(ctx context.Context, portfolioItemId int64) (err error)) *PortfolioItemRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the PortfolioItemRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the PortfolioItemRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the PortfolioItemRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mPortfolioItemRepositoryMockDelete) When(ctx context.Context, portfolioItemId int64) *PortfolioItemRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("PortfolioItemRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &PortfolioItemRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &PortfolioItemRepositoryMockDeleteParams{ctx, portfolioItemId},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up PortfolioItemRepository.Delete return parameters for the expectation previously defined by the When method
func (e *PortfolioItemRepositoryMockDeleteExpectation) Then(err error) *PortfolioItemRepositoryMock {
	e.results = &PortfolioItemRepositoryMockDeleteResults{err}
	return e.mock
}

// Delete implements PortfolioItemRepository
func (mmDelete *PortfolioItemRepositoryMock) Delete(ctx context.Context, portfolioItemId int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, portfolioItemId)
	}

	mm_params := &PortfolioItemRepositoryMockDeleteParams{ctx, portfolioItemId}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := PortfolioItemRepositoryMockDeleteParams{ctx, portfolioItemId}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("PortfolioItemRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the PortfolioItemRepositoryMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, portfolioItemId)
	}
	mmDelete.t.Fatalf("Unexpected call to PortfolioItemRepositoryMock.Delete. %v %v", ctx, portfolioItemId)
	return
}

// DeleteAfterCounter returns a count of finished PortfolioItemRepositoryMock.Delete invocations
func (mmDelete *PortfolioItemRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of PortfolioItemRepositoryMock.Delete invocations
func (mmDelete *PortfolioItemRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to PortfolioItemRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mPortfolioItemRepositoryMockDelete) Calls() []*PortfolioItemRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*PortfolioItemRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *PortfolioItemRepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *PortfolioItemRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PortfolioItemRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PortfolioItemRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to PortfolioItemRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to PortfolioItemRepositoryMock.Delete")
	}
}

type mPortfolioItemRepositoryMockRetrieveOrCreate struct {
	mock               *PortfolioItemRepositoryMock
	defaultExpectation *PortfolioItemRepositoryMockRetrieveOrCreateExpectation
	expectations       []*PortfolioItemRepositoryMockRetrieveOrCreateExpectation

	callArgs []*PortfolioItemRepositoryMockRetrieveOrCreateParams
	mutex    sync.RWMutex
}

// PortfolioItemRepositoryMockRetrieveOrCreateExpectation specifies expectation struct of the PortfolioItemRepository.RetrieveOrCreate
type PortfolioItemRepositoryMockRetrieveOrCreateExpectation struct {
	mock    *PortfolioItemRepositoryMock
	params  *PortfolioItemRepositoryMockRetrieveOrCreateParams
	results *PortfolioItemRepositoryMockRetrieveOrCreateResults
	Counter uint64
}

// PortfolioItemRepositoryMockRetrieveOrCreateParams contains parameters of the PortfolioItemRepository.RetrieveOrCreate
type PortfolioItemRepositoryMockRetrieveOrCreateParams struct {
	ctx           context.Context
	portfolioData domain.PortfolioItemCreate
}

// PortfolioItemRepositoryMockRetrieveOrCreateResults contains results of the PortfolioItemRepository.RetrieveOrCreate
type PortfolioItemRepositoryMockRetrieveOrCreateResults struct {
	p1 domain.PortfolioItemRetrieve
}

// Expect sets up expected params for PortfolioItemRepository.RetrieveOrCreate
func (mmRetrieveOrCreate *mPortfolioItemRepositoryMockRetrieveOrCreate) Expect(ctx context.Context, portfolioData domain.PortfolioItemCreate) *mPortfolioItemRepositoryMockRetrieveOrCreate {
	if mmRetrieveOrCreate.mock.funcRetrieveOrCreate != nil {
		mmRetrieveOrCreate.mock.t.Fatalf("PortfolioItemRepositoryMock.RetrieveOrCreate mock is already set by Set")
	}

	if mmRetrieveOrCreate.defaultExpectation == nil {
		mmRetrieveOrCreate.defaultExpectation = &PortfolioItemRepositoryMockRetrieveOrCreateExpectation{}
	}

	mmRetrieveOrCreate.defaultExpectation.params = &PortfolioItemRepositoryMockRetrieveOrCreateParams{ctx, portfolioData}
	for _, e := range mmRetrieveOrCreate.expectations {
		if minimock.Equal(e.params, mmRetrieveOrCreate.defaultExpectation.params) {
			mmRetrieveOrCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRetrieveOrCreate.defaultExpectation.params)
		}
	}

	return mmRetrieveOrCreate
}

// Inspect accepts an inspector function that has same arguments as the PortfolioItemRepository.RetrieveOrCreate
func (mmRetrieveOrCreate *mPortfolioItemRepositoryMockRetrieveOrCreate) Inspect(f func(ctx context.Context, portfolioData domain.PortfolioItemCreate)) *mPortfolioItemRepositoryMockRetrieveOrCreate {
	if mmRetrieveOrCreate.mock.inspectFuncRetrieveOrCreate != nil {
		mmRetrieveOrCreate.mock.t.Fatalf("Inspect function is already set for PortfolioItemRepositoryMock.RetrieveOrCreate")
	}

	mmRetrieveOrCreate.mock.inspectFuncRetrieveOrCreate = f

	return mmRetrieveOrCreate
}

// Return sets up results that will be returned by PortfolioItemRepository.RetrieveOrCreate
func (mmRetrieveOrCreate *mPortfolioItemRepositoryMockRetrieveOrCreate) Return(p1 domain.PortfolioItemRetrieve) *PortfolioItemRepositoryMock {
	if mmRetrieveOrCreate.mock.funcRetrieveOrCreate != nil {
		mmRetrieveOrCreate.mock.t.Fatalf("PortfolioItemRepositoryMock.RetrieveOrCreate mock is already set by Set")
	}

	if mmRetrieveOrCreate.defaultExpectation == nil {
		mmRetrieveOrCreate.defaultExpectation = &PortfolioItemRepositoryMockRetrieveOrCreateExpectation{mock: mmRetrieveOrCreate.mock}
	}
	mmRetrieveOrCreate.defaultExpectation.results = &PortfolioItemRepositoryMockRetrieveOrCreateResults{p1}
	return mmRetrieveOrCreate.mock
}

//Set uses given function f to mock the PortfolioItemRepository.RetrieveOrCreate method
func (mmRetrieveOrCreate *mPortfolioItemRepositoryMockRetrieveOrCreate) Set(f func(ctx context.Context, portfolioData domain.PortfolioItemCreate) (p1 domain.PortfolioItemRetrieve)) *PortfolioItemRepositoryMock {
	if mmRetrieveOrCreate.defaultExpectation != nil {
		mmRetrieveOrCreate.mock.t.Fatalf("Default expectation is already set for the PortfolioItemRepository.RetrieveOrCreate method")
	}

	if len(mmRetrieveOrCreate.expectations) > 0 {
		mmRetrieveOrCreate.mock.t.Fatalf("Some expectations are already set for the PortfolioItemRepository.RetrieveOrCreate method")
	}

	mmRetrieveOrCreate.mock.funcRetrieveOrCreate = f
	return mmRetrieveOrCreate.mock
}

// When sets expectation for the PortfolioItemRepository.RetrieveOrCreate which will trigger the result defined by the following
// Then helper
func (mmRetrieveOrCreate *mPortfolioItemRepositoryMockRetrieveOrCreate) When(ctx context.Context, portfolioData domain.PortfolioItemCreate) *PortfolioItemRepositoryMockRetrieveOrCreateExpectation {
	if mmRetrieveOrCreate.mock.funcRetrieveOrCreate != nil {
		mmRetrieveOrCreate.mock.t.Fatalf("PortfolioItemRepositoryMock.RetrieveOrCreate mock is already set by Set")
	}

	expectation := &PortfolioItemRepositoryMockRetrieveOrCreateExpectation{
		mock:   mmRetrieveOrCreate.mock,
		params: &PortfolioItemRepositoryMockRetrieveOrCreateParams{ctx, portfolioData},
	}
	mmRetrieveOrCreate.expectations = append(mmRetrieveOrCreate.expectations, expectation)
	return expectation
}

// Then sets up PortfolioItemRepository.RetrieveOrCreate return parameters for the expectation previously defined by the When method
func (e *PortfolioItemRepositoryMockRetrieveOrCreateExpectation) Then(p1 domain.PortfolioItemRetrieve) *PortfolioItemRepositoryMock {
	e.results = &PortfolioItemRepositoryMockRetrieveOrCreateResults{p1}
	return e.mock
}

// RetrieveOrCreate implements PortfolioItemRepository
func (mmRetrieveOrCreate *PortfolioItemRepositoryMock) RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioItemCreate) (p1 domain.PortfolioItemRetrieve) {
	mm_atomic.AddUint64(&mmRetrieveOrCreate.beforeRetrieveOrCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmRetrieveOrCreate.afterRetrieveOrCreateCounter, 1)

	if mmRetrieveOrCreate.inspectFuncRetrieveOrCreate != nil {
		mmRetrieveOrCreate.inspectFuncRetrieveOrCreate(ctx, portfolioData)
	}

	mm_params := &PortfolioItemRepositoryMockRetrieveOrCreateParams{ctx, portfolioData}

	// Record call args
	mmRetrieveOrCreate.RetrieveOrCreateMock.mutex.Lock()
	mmRetrieveOrCreate.RetrieveOrCreateMock.callArgs = append(mmRetrieveOrCreate.RetrieveOrCreateMock.callArgs, mm_params)
	mmRetrieveOrCreate.RetrieveOrCreateMock.mutex.Unlock()

	for _, e := range mmRetrieveOrCreate.RetrieveOrCreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1
		}
	}

	if mmRetrieveOrCreate.RetrieveOrCreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRetrieveOrCreate.RetrieveOrCreateMock.defaultExpectation.Counter, 1)
		mm_want := mmRetrieveOrCreate.RetrieveOrCreateMock.defaultExpectation.params
		mm_got := PortfolioItemRepositoryMockRetrieveOrCreateParams{ctx, portfolioData}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRetrieveOrCreate.t.Errorf("PortfolioItemRepositoryMock.RetrieveOrCreate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRetrieveOrCreate.RetrieveOrCreateMock.defaultExpectation.results
		if mm_results == nil {
			mmRetrieveOrCreate.t.Fatal("No results are set for the PortfolioItemRepositoryMock.RetrieveOrCreate")
		}
		return (*mm_results).p1
	}
	if mmRetrieveOrCreate.funcRetrieveOrCreate != nil {
		return mmRetrieveOrCreate.funcRetrieveOrCreate(ctx, portfolioData)
	}
	mmRetrieveOrCreate.t.Fatalf("Unexpected call to PortfolioItemRepositoryMock.RetrieveOrCreate. %v %v", ctx, portfolioData)
	return
}

// RetrieveOrCreateAfterCounter returns a count of finished PortfolioItemRepositoryMock.RetrieveOrCreate invocations
func (mmRetrieveOrCreate *PortfolioItemRepositoryMock) RetrieveOrCreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRetrieveOrCreate.afterRetrieveOrCreateCounter)
}

// RetrieveOrCreateBeforeCounter returns a count of PortfolioItemRepositoryMock.RetrieveOrCreate invocations
func (mmRetrieveOrCreate *PortfolioItemRepositoryMock) RetrieveOrCreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRetrieveOrCreate.beforeRetrieveOrCreateCounter)
}

// Calls returns a list of arguments used in each call to PortfolioItemRepositoryMock.RetrieveOrCreate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRetrieveOrCreate *mPortfolioItemRepositoryMockRetrieveOrCreate) Calls() []*PortfolioItemRepositoryMockRetrieveOrCreateParams {
	mmRetrieveOrCreate.mutex.RLock()

	argCopy := make([]*PortfolioItemRepositoryMockRetrieveOrCreateParams, len(mmRetrieveOrCreate.callArgs))
	copy(argCopy, mmRetrieveOrCreate.callArgs)

	mmRetrieveOrCreate.mutex.RUnlock()

	return argCopy
}

// MinimockRetrieveOrCreateDone returns true if the count of the RetrieveOrCreate invocations corresponds
// the number of defined expectations
func (m *PortfolioItemRepositoryMock) MinimockRetrieveOrCreateDone() bool {
	for _, e := range m.RetrieveOrCreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RetrieveOrCreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRetrieveOrCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRetrieveOrCreate != nil && mm_atomic.LoadUint64(&m.afterRetrieveOrCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockRetrieveOrCreateInspect logs each unmet expectation
func (m *PortfolioItemRepositoryMock) MinimockRetrieveOrCreateInspect() {
	for _, e := range m.RetrieveOrCreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PortfolioItemRepositoryMock.RetrieveOrCreate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RetrieveOrCreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRetrieveOrCreateCounter) < 1 {
		if m.RetrieveOrCreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PortfolioItemRepositoryMock.RetrieveOrCreate")
		} else {
			m.t.Errorf("Expected call to PortfolioItemRepositoryMock.RetrieveOrCreate with params: %#v", *m.RetrieveOrCreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRetrieveOrCreate != nil && mm_atomic.LoadUint64(&m.afterRetrieveOrCreateCounter) < 1 {
		m.t.Error("Expected call to PortfolioItemRepositoryMock.RetrieveOrCreate")
	}
}

type mPortfolioItemRepositoryMockRetrievePortfolioItems struct {
	mock               *PortfolioItemRepositoryMock
	defaultExpectation *PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation
	expectations       []*PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation

	callArgs []*PortfolioItemRepositoryMockRetrievePortfolioItemsParams
	mutex    sync.RWMutex
}

// PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation specifies expectation struct of the PortfolioItemRepository.RetrievePortfolioItems
type PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation struct {
	mock    *PortfolioItemRepositoryMock
	params  *PortfolioItemRepositoryMockRetrievePortfolioItemsParams
	results *PortfolioItemRepositoryMockRetrievePortfolioItemsResults
	Counter uint64
}

// PortfolioItemRepositoryMockRetrievePortfolioItemsParams contains parameters of the PortfolioItemRepository.RetrievePortfolioItems
type PortfolioItemRepositoryMockRetrievePortfolioItemsParams struct {
	ctx         context.Context
	portfolioId int64
}

// PortfolioItemRepositoryMockRetrievePortfolioItemsResults contains results of the PortfolioItemRepository.RetrievePortfolioItems
type PortfolioItemRepositoryMockRetrievePortfolioItemsResults struct {
	pp1 *domain.PortfolioItemsRetrieve
}

// Expect sets up expected params for PortfolioItemRepository.RetrievePortfolioItems
func (mmRetrievePortfolioItems *mPortfolioItemRepositoryMockRetrievePortfolioItems) Expect(ctx context.Context, portfolioId int64) *mPortfolioItemRepositoryMockRetrievePortfolioItems {
	if mmRetrievePortfolioItems.mock.funcRetrievePortfolioItems != nil {
		mmRetrievePortfolioItems.mock.t.Fatalf("PortfolioItemRepositoryMock.RetrievePortfolioItems mock is already set by Set")
	}

	if mmRetrievePortfolioItems.defaultExpectation == nil {
		mmRetrievePortfolioItems.defaultExpectation = &PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation{}
	}

	mmRetrievePortfolioItems.defaultExpectation.params = &PortfolioItemRepositoryMockRetrievePortfolioItemsParams{ctx, portfolioId}
	for _, e := range mmRetrievePortfolioItems.expectations {
		if minimock.Equal(e.params, mmRetrievePortfolioItems.defaultExpectation.params) {
			mmRetrievePortfolioItems.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRetrievePortfolioItems.defaultExpectation.params)
		}
	}

	return mmRetrievePortfolioItems
}

// Inspect accepts an inspector function that has same arguments as the PortfolioItemRepository.RetrievePortfolioItems
func (mmRetrievePortfolioItems *mPortfolioItemRepositoryMockRetrievePortfolioItems) Inspect(f func(ctx context.Context, portfolioId int64)) *mPortfolioItemRepositoryMockRetrievePortfolioItems {
	if mmRetrievePortfolioItems.mock.inspectFuncRetrievePortfolioItems != nil {
		mmRetrievePortfolioItems.mock.t.Fatalf("Inspect function is already set for PortfolioItemRepositoryMock.RetrievePortfolioItems")
	}

	mmRetrievePortfolioItems.mock.inspectFuncRetrievePortfolioItems = f

	return mmRetrievePortfolioItems
}

// Return sets up results that will be returned by PortfolioItemRepository.RetrievePortfolioItems
func (mmRetrievePortfolioItems *mPortfolioItemRepositoryMockRetrievePortfolioItems) Return(pp1 *domain.PortfolioItemsRetrieve) *PortfolioItemRepositoryMock {
	if mmRetrievePortfolioItems.mock.funcRetrievePortfolioItems != nil {
		mmRetrievePortfolioItems.mock.t.Fatalf("PortfolioItemRepositoryMock.RetrievePortfolioItems mock is already set by Set")
	}

	if mmRetrievePortfolioItems.defaultExpectation == nil {
		mmRetrievePortfolioItems.defaultExpectation = &PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation{mock: mmRetrievePortfolioItems.mock}
	}
	mmRetrievePortfolioItems.defaultExpectation.results = &PortfolioItemRepositoryMockRetrievePortfolioItemsResults{pp1}
	return mmRetrievePortfolioItems.mock
}

//Set uses given function f to mock the PortfolioItemRepository.RetrievePortfolioItems method
func (mmRetrievePortfolioItems *mPortfolioItemRepositoryMockRetrievePortfolioItems) Set(f func(ctx context.Context, portfolioId int64) (pp1 *domain.PortfolioItemsRetrieve)) *PortfolioItemRepositoryMock {
	if mmRetrievePortfolioItems.defaultExpectation != nil {
		mmRetrievePortfolioItems.mock.t.Fatalf("Default expectation is already set for the PortfolioItemRepository.RetrievePortfolioItems method")
	}

	if len(mmRetrievePortfolioItems.expectations) > 0 {
		mmRetrievePortfolioItems.mock.t.Fatalf("Some expectations are already set for the PortfolioItemRepository.RetrievePortfolioItems method")
	}

	mmRetrievePortfolioItems.mock.funcRetrievePortfolioItems = f
	return mmRetrievePortfolioItems.mock
}

// When sets expectation for the PortfolioItemRepository.RetrievePortfolioItems which will trigger the result defined by the following
// Then helper
func (mmRetrievePortfolioItems *mPortfolioItemRepositoryMockRetrievePortfolioItems) When(ctx context.Context, portfolioId int64) *PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation {
	if mmRetrievePortfolioItems.mock.funcRetrievePortfolioItems != nil {
		mmRetrievePortfolioItems.mock.t.Fatalf("PortfolioItemRepositoryMock.RetrievePortfolioItems mock is already set by Set")
	}

	expectation := &PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation{
		mock:   mmRetrievePortfolioItems.mock,
		params: &PortfolioItemRepositoryMockRetrievePortfolioItemsParams{ctx, portfolioId},
	}
	mmRetrievePortfolioItems.expectations = append(mmRetrievePortfolioItems.expectations, expectation)
	return expectation
}

// Then sets up PortfolioItemRepository.RetrievePortfolioItems return parameters for the expectation previously defined by the When method
func (e *PortfolioItemRepositoryMockRetrievePortfolioItemsExpectation) Then(pp1 *domain.PortfolioItemsRetrieve) *PortfolioItemRepositoryMock {
	e.results = &PortfolioItemRepositoryMockRetrievePortfolioItemsResults{pp1}
	return e.mock
}

// RetrievePortfolioItems implements PortfolioItemRepository
func (mmRetrievePortfolioItems *PortfolioItemRepositoryMock) RetrievePortfolioItems(ctx context.Context, portfolioId int64) (pp1 *domain.PortfolioItemsRetrieve) {
	mm_atomic.AddUint64(&mmRetrievePortfolioItems.beforeRetrievePortfolioItemsCounter, 1)
	defer mm_atomic.AddUint64(&mmRetrievePortfolioItems.afterRetrievePortfolioItemsCounter, 1)

	if mmRetrievePortfolioItems.inspectFuncRetrievePortfolioItems != nil {
		mmRetrievePortfolioItems.inspectFuncRetrievePortfolioItems(ctx, portfolioId)
	}

	mm_params := &PortfolioItemRepositoryMockRetrievePortfolioItemsParams{ctx, portfolioId}

	// Record call args
	mmRetrievePortfolioItems.RetrievePortfolioItemsMock.mutex.Lock()
	mmRetrievePortfolioItems.RetrievePortfolioItemsMock.callArgs = append(mmRetrievePortfolioItems.RetrievePortfolioItemsMock.callArgs, mm_params)
	mmRetrievePortfolioItems.RetrievePortfolioItemsMock.mutex.Unlock()

	for _, e := range mmRetrievePortfolioItems.RetrievePortfolioItemsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp1
		}
	}

	if mmRetrievePortfolioItems.RetrievePortfolioItemsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRetrievePortfolioItems.RetrievePortfolioItemsMock.defaultExpectation.Counter, 1)
		mm_want := mmRetrievePortfolioItems.RetrievePortfolioItemsMock.defaultExpectation.params
		mm_got := PortfolioItemRepositoryMockRetrievePortfolioItemsParams{ctx, portfolioId}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRetrievePortfolioItems.t.Errorf("PortfolioItemRepositoryMock.RetrievePortfolioItems got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRetrievePortfolioItems.RetrievePortfolioItemsMock.defaultExpectation.results
		if mm_results == nil {
			mmRetrievePortfolioItems.t.Fatal("No results are set for the PortfolioItemRepositoryMock.RetrievePortfolioItems")
		}
		return (*mm_results).pp1
	}
	if mmRetrievePortfolioItems.funcRetrievePortfolioItems != nil {
		return mmRetrievePortfolioItems.funcRetrievePortfolioItems(ctx, portfolioId)
	}
	mmRetrievePortfolioItems.t.Fatalf("Unexpected call to PortfolioItemRepositoryMock.RetrievePortfolioItems. %v %v", ctx, portfolioId)
	return
}

// RetrievePortfolioItemsAfterCounter returns a count of finished PortfolioItemRepositoryMock.RetrievePortfolioItems invocations
func (mmRetrievePortfolioItems *PortfolioItemRepositoryMock) RetrievePortfolioItemsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRetrievePortfolioItems.afterRetrievePortfolioItemsCounter)
}

// RetrievePortfolioItemsBeforeCounter returns a count of PortfolioItemRepositoryMock.RetrievePortfolioItems invocations
func (mmRetrievePortfolioItems *PortfolioItemRepositoryMock) RetrievePortfolioItemsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRetrievePortfolioItems.beforeRetrievePortfolioItemsCounter)
}

// Calls returns a list of arguments used in each call to PortfolioItemRepositoryMock.RetrievePortfolioItems.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRetrievePortfolioItems *mPortfolioItemRepositoryMockRetrievePortfolioItems) Calls() []*PortfolioItemRepositoryMockRetrievePortfolioItemsParams {
	mmRetrievePortfolioItems.mutex.RLock()

	argCopy := make([]*PortfolioItemRepositoryMockRetrievePortfolioItemsParams, len(mmRetrievePortfolioItems.callArgs))
	copy(argCopy, mmRetrievePortfolioItems.callArgs)

	mmRetrievePortfolioItems.mutex.RUnlock()

	return argCopy
}

// MinimockRetrievePortfolioItemsDone returns true if the count of the RetrievePortfolioItems invocations corresponds
// the number of defined expectations
func (m *PortfolioItemRepositoryMock) MinimockRetrievePortfolioItemsDone() bool {
	for _, e := range m.RetrievePortfolioItemsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RetrievePortfolioItemsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRetrievePortfolioItemsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRetrievePortfolioItems != nil && mm_atomic.LoadUint64(&m.afterRetrievePortfolioItemsCounter) < 1 {
		return false
	}
	return true
}

// MinimockRetrievePortfolioItemsInspect logs each unmet expectation
func (m *PortfolioItemRepositoryMock) MinimockRetrievePortfolioItemsInspect() {
	for _, e := range m.RetrievePortfolioItemsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PortfolioItemRepositoryMock.RetrievePortfolioItems with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RetrievePortfolioItemsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRetrievePortfolioItemsCounter) < 1 {
		if m.RetrievePortfolioItemsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PortfolioItemRepositoryMock.RetrievePortfolioItems")
		} else {
			m.t.Errorf("Expected call to PortfolioItemRepositoryMock.RetrievePortfolioItems with params: %#v", *m.RetrievePortfolioItemsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRetrievePortfolioItems != nil && mm_atomic.LoadUint64(&m.afterRetrievePortfolioItemsCounter) < 1 {
		m.t.Error("Expected call to PortfolioItemRepositoryMock.RetrievePortfolioItems")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PortfolioItemRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteInspect()

		m.MinimockRetrieveOrCreateInspect()

		m.MinimockRetrievePortfolioItemsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PortfolioItemRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *PortfolioItemRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockRetrieveOrCreateDone() &&
		m.MinimockRetrievePortfolioItemsDone()
}
