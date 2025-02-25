package configmap

import (
	"context"

	"github.com/kubeshop/testkube/pkg/k8sclient"
	"github.com/kubeshop/testkube/pkg/log"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
)

// Client provide methods to manage configmaps
type Client struct {
	ClientSet *kubernetes.Clientset
	Log       *zap.SugaredLogger
	Namespace string
}

// NewClient is a method to create new configmap client
func NewClient(namespace string) (*Client, error) {
	clientSet, err := k8sclient.ConnectToK8s()
	if err != nil {
		return nil, err
	}

	return &Client{
		ClientSet: clientSet,
		Log:       log.DefaultLogger,
		Namespace: namespace,
	}, nil
}

// Create is a method to create new configmap
func (c *Client) Create(id string, stringData map[string]string) error {
	configMapsClient := c.ClientSet.CoreV1().ConfigMaps(c.Namespace)
	ctx := context.Background()

	configMapSpec := NewSpec(id, c.Namespace, stringData)
	if _, err := configMapsClient.Create(ctx, configMapSpec, metav1.CreateOptions{}); err != nil {
		return err
	}

	return nil
}

// Get is a method to retrieve an existing configmap
func (c *Client) Get(id string) (map[string]string, error) {
	configMapsClient := c.ClientSet.CoreV1().ConfigMaps(c.Namespace)
	ctx := context.Background()

	configMapSpec, err := configMapsClient.Get(ctx, id, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	stringData := map[string]string{}
	for key, value := range configMapSpec.Data {
		stringData[key] = value
	}

	return stringData, nil
}

// Update is a method to update an existing configmap
func (c *Client) Update(id string, stringData map[string]string) error {
	configMapsClient := c.ClientSet.CoreV1().ConfigMaps(c.Namespace)
	ctx := context.Background()

	configMapSpec := NewSpec(id, c.Namespace, stringData)
	if _, err := configMapsClient.Update(ctx, configMapSpec, metav1.UpdateOptions{}); err != nil {
		return err
	}

	return nil
}

// Apply is a method to create or update a configmap
func (c *Client) Apply(id string, stringData map[string]string) error {
	configMapsClient := c.ClientSet.CoreV1().ConfigMaps(c.Namespace)
	ctx := context.Background()

	configMapSpec := NewApplySpec(id, c.Namespace, stringData)
	if _, err := configMapsClient.Apply(ctx, configMapSpec, metav1.ApplyOptions{
		FieldManager: "application/apply-patch"}); err != nil {
		return err
	}

	return nil
}

// NewSpec is a method to return configmap spec
func NewSpec(id, namespace string, stringData map[string]string) *v1.ConfigMap {
	configuration := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      id,
			Namespace: namespace,
		},
		Data: stringData,
	}

	return configuration
}

// NewApplySpec is a method to return configmap apply spec
func NewApplySpec(id, namespace string, stringData map[string]string) *corev1.ConfigMapApplyConfiguration {
	configuration := corev1.ConfigMap(id, namespace).
		WithData(stringData)

	return configuration
}
