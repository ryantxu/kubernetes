package testing

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
)

var _ storage.Interface = &FakeStorageInterface{}

type FakeStorageInterface struct {
	ExpectedError          error
	ExpectedVersioner      storage.Versioner
	ExpectedWatchInterface watch.Interface

	FakeGetList func(ctx context.Context, key string, opts storage.ListOptions, listObj runtime.Object) error
}

// Count implements storage.Interface.
func (f *FakeStorageInterface) Count(key string) (int64, error) {
	return 0, f.ExpectedError
}

// Create implements storage.Interface.
func (f *FakeStorageInterface) Create(ctx context.Context, key string, obj runtime.Object, out runtime.Object, ttl uint64) error {
	return f.ExpectedError
}

// Delete implements storage.Interface.
func (f *FakeStorageInterface) Delete(ctx context.Context, key string, out runtime.Object, preconditions *storage.Preconditions, validateDeletion storage.ValidateObjectFunc, cachedExistingObject runtime.Object) error {
	return f.ExpectedError
}

// Get implements storage.Interface.
func (f *FakeStorageInterface) Get(ctx context.Context, key string, opts storage.GetOptions, objPtr runtime.Object) error {
	return f.ExpectedError
}

// GetList implements storage.Interface.
func (f *FakeStorageInterface) GetList(ctx context.Context, key string, opts storage.ListOptions, listObj runtime.Object) error {
	if f.FakeGetList != nil {
		return f.FakeGetList(ctx, key, opts, listObj)
	}
	return f.ExpectedError
}

// GuaranteedUpdate implements storage.Interface.
func (f *FakeStorageInterface) GuaranteedUpdate(ctx context.Context, key string, destination runtime.Object, ignoreNotFound bool, preconditions *storage.Preconditions, tryUpdate storage.UpdateFunc, cachedExistingObject runtime.Object) error {
	return f.ExpectedError
}

// ReadinessCheck implements storage.Interface.
func (f *FakeStorageInterface) ReadinessCheck() error {
	return f.ExpectedError
}

// RequestWatchProgress implements storage.Interface.
func (f *FakeStorageInterface) RequestWatchProgress(ctx context.Context) error {
	return f.ExpectedError
}

// Versioner implements storage.Interface.
func (f *FakeStorageInterface) Versioner() storage.Versioner {
	return f.ExpectedVersioner
}

// Watch implements storage.Interface.
func (f *FakeStorageInterface) Watch(ctx context.Context, key string, opts storage.ListOptions) (watch.Interface, error) {
	return f.ExpectedWatchInterface, f.ExpectedError
}
