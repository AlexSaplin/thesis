package k8s

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	tesseract "tesseract/pkg/service/pb"
)

type K8SClient struct {
	clientSet *kubernetes.Clientset
	yaml      []byte
}

func NewK8SClient(configPath string) (*K8SClient, error) {

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	f, err := os.Open("service.yaml")
	if err != nil {
		return nil, err
	}
	yaml, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}

	return &K8SClient{
		clientSet: clientSet,
		yaml:      yaml,
	}, nil
}

func (client *K8SClient) checkPods(ctx context.Context, namespace string) (string, error) {
	label := fmt.Sprintf("appid=%s", namespace)
	pods, err := client.clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return "", err
	}
	for _, pod := range pods.Items {
		for _, status := range pod.Status.ContainerStatuses {
			if status.State.Waiting != nil {
				return status.State.Waiting.Reason, nil
			}
		}
		if pod.Status.Phase == corev1.PodPending {
			for _, condition := range pod.Status.Conditions {
				if condition.Type == corev1.PodScheduled && condition.Status == corev1.ConditionFalse {
					return condition.Reason, nil
				}
			}
		}
		// port check
		/*
		if pod.Status.Phase == corev1.PodRunning {
			for _, condition := range pod.Status.Conditions {
				if condition.Type == corev1.ContainersReady && condition.Status == corev1.ConditionFalse {
					return condition.Reason, nil
				}
			}
		}*/
	}
	return "MODIFYING", nil
}

func (client *K8SClient) Get(ctx context.Context, namespace string) (string, error) {

	label := fmt.Sprintf("appid=%s", namespace)
	result, err := client.clientSet.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return "", err
	}
	if len(result.Items) != 1 {
		return "", errors.New("not 1 deployment with correct label")
	}

	deployment := result.Items[0]

	for _, item := range deployment.Status.Conditions {
		if item.Type == v1.DeploymentProgressing && item.Reason != "NewReplicaSetAvailable" ||
			item.Type == v1.DeploymentAvailable && item.Reason == "MinimumReplicasUnavailable" {
			return client.checkPods(ctx, namespace)
		}
	}

	return "OK", nil
}

func (client *K8SClient) Down(ctx context.Context, namespace string) error {
	return client.clientSet.CoreV1().Namespaces().Delete(ctx, namespace, metav1.DeleteOptions{})
}

type TemplateArgs struct {
	Namespace string
	Name      string
	DNS       string
	Image     string
	Port      uint
	Scale     uint
	CPU       uint
	MemoryMB  uint
	GPU       string
	Env       []*tesseract.KV
	Auth      string
}

func (client *K8SClient) Apply(ctx context.Context, args TemplateArgs) error {
	// https://github.com/kubernetes/enhancements/issues/555

	t, err := template.New("yaml").Parse(string(client.yaml))
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	err = t.Execute(&buf, args)
	if err != nil {
		return err
	}
	yamlFilled := buf.Bytes()
	fmt.Println("args %+v", args)
	fmt.Println(string(yamlFilled))

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	_, err = tf.Write(yamlFilled)
	if err != nil {
		return err
	}
	err = tf.Close()
	if err != nil {
		return err
	}
	fname := tf.Name()
	cmd := exec.CommandContext(ctx, "kubectl", "apply", "-f", fname)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	// err = cmd.Run()
	if err != nil {
		return err
	}
	return os.Remove(fname)
}
