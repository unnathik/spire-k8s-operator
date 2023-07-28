package controller

import (
	"context"
	"testing"

	// "sigs.k8s.io/controller-runtime/pkg/client/fake"

	spirev1 "github.com/glcp/spire-k8s-operator/api/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MockClient struct {
	client.Client
	CreateFn func(context.Context, client.Object, ...client.CreateOption) error
}

func (m *MockClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, obj, opts...)
	}
	return nil
}

func createReconciler() *SpireServerReconciler {
	reconciler := &SpireServerReconciler{
		Client: &MockClient{
			CreateFn: func(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
				// Handle the create logic in the mock client
				return nil
			},
		},
		Scheme: scheme.Scheme,
	}
	return reconciler
}
func TestSpireserverController(t *testing.T) {
	// Create the objects needed for the test
	reconciler := createReconciler()

	spireserver := &spirev1.SpireServer{}
	spireServiceNamespace := "test-namespace"
	spireServiceAccount := reconciler.createServiceAccount(spireServiceNamespace)
	bundle := reconciler.spireBundleDeployment(spireServiceNamespace)
	roles := reconciler.spireRoleDeployment(spireServiceNamespace)
	roleBinding := reconciler.spireRoleBindingDeployment(spireServiceNamespace)
	clusterRoles := reconciler.spireClusterRoleDeployment(spireServiceNamespace)
	clusterRoleBinding := reconciler.spireClusterRoleBindingDeployment(spireServiceNamespace)
	serverConfigMap := reconciler.spireConfigMapDeployment(spireserver, spireServiceNamespace)
	spireStatefulSet := reconciler.spireStatefulSetDeployment(2, spireServiceNamespace)
	spireService := reconciler.spireServiceDeployment(8081, spireServiceNamespace)

	// Call the method you want to test
	// Assert the expected behavior

	if spireServiceAccount.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, spireServiceAccount.Namespace)
	}
	if bundle.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, bundle.Namespace)
	}
	if roles.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, roles.Namespace)
	}
	if roleBinding.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, roleBinding.Namespace)
	}
	if clusterRoles.Namespace != "" {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, clusterRoles.Namespace)
	}
	if clusterRoleBinding.Namespace != "" {
		t.Errorf("Expected namespace \"\", got %s", clusterRoleBinding.Namespace)
	}
	if serverConfigMap.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, serverConfigMap.Namespace)
	}
	if spireStatefulSet.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, spireStatefulSet.Namespace)
	}
	if spireService.Namespace != spireServiceNamespace {
		t.Errorf("Expected namespace %s, got %s", spireServiceNamespace, spireService.Namespace)
	}
}
func TestValidNameSpaceRoles(t *testing.T) {
	// Create the objects needed for the test
	reconcilerForRoles := createReconciler()
	roles := reconcilerForRoles.spireRoleDeployment("default")
	assert.Equal(t, roles.Namespace, "default")
}

func TestInvalidNameSpaceRoles(t *testing.T) {
	reconcilerForRoles := createReconciler()
	roles := reconcilerForRoles.spireRoleDeployment("default1")
	assert.NotEqual(t, roles.Namespace, "default2")
}

func TestEmptyNameSpaceRoles(t *testing.T) {
	reconcilerForRoles := createReconciler()
	roles := reconcilerForRoles.spireRoleDeployment("")
	assert.Equal(t, roles.Namespace, "")
}

func TestValidNameSpaceRoleBinding(t *testing.T) {
	reconcilerForRoleBinding := createReconciler()
	roleBinding := reconcilerForRoleBinding.spireRoleDeployment("default")
	assert.Equal(t, roleBinding.Namespace, "default")
}

func TestInvalidNameSpaceRoleBinding(t *testing.T) {
	reconcilerForRoleBinding := createReconciler()
	roleBinding := reconcilerForRoleBinding.spireRoleDeployment("default1")
	assert.NotEqual(t, roleBinding.Namespace, "default2")
}

func TestEmptyNameSpaceRoleBinding(t *testing.T) {
	reconcilerForRoleBinding := createReconciler()
	roleBinding := reconcilerForRoleBinding.spireRoleDeployment("")
	assert.Equal(t, roleBinding.Namespace, "")
}

func TestValidNameSpaceClusterRoles(t *testing.T) {
	reconcilerForClusterRoles := createReconciler()
	clusterRoles := reconcilerForClusterRoles.spireClusterRoleDeployment("default")
	assert.Equal(t, clusterRoles.Namespace, "")
}

func TestInvalidNameSpaceClusterRoles(t *testing.T) {
	reconcilerForClusterRoles := createReconciler()
	clusterRoles := reconcilerForClusterRoles.spireClusterRoleDeployment("default1")
	assert.Equal(t, clusterRoles.Namespace, "")
}

func TestEmptyNameSpaceClusterRoles(t *testing.T) {
	reconcilerForClusterRoles := createReconciler()
	clusterRoles := reconcilerForClusterRoles.spireClusterRoleDeployment("")
	assert.Equal(t, clusterRoles.Namespace, "")
}

func TestValidNameSpaceClusterRoleBinding(t *testing.T) {
	reconcilerForClusterRoleBinding := createReconciler()
	clusterRoleBinding := reconcilerForClusterRoleBinding.spireClusterRoleDeployment("default")
	assert.Equal(t, clusterRoleBinding.Namespace, "")
}

func TestInvalidNameSpaceClusterRoleBinding(t *testing.T) {
	reconcilerForClusterRoleBinding := createReconciler()
	clusterRoleBinding := reconcilerForClusterRoleBinding.spireClusterRoleDeployment("default1")
	assert.Equal(t, clusterRoleBinding.Namespace, "")
}

func TestEmptyNameSpaceClusterRoleBinding(t *testing.T) {
	reconcilerForClusterRoleBinding := createReconciler()
	clusterRoleBinding := reconcilerForClusterRoleBinding.spireClusterRoleDeployment("")
	assert.Equal(t, clusterRoleBinding.Namespace, "")
}
