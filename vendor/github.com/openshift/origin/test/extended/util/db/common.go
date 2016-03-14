package db

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/openshift/origin/test/extended/util"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
)

// PodConfig holds configuration for a pod.
type PodConfig struct {
	Container string
	Env       map[string]string
}

func getPodConfig(c kclient.PodInterface, podName string) (conf *PodConfig, err error) {
	pod, err := c.Get(podName)
	if err != nil {
		return nil, err
	}
	env := make(map[string]string)
	for _, container := range pod.Spec.Containers {
		for _, e := range container.Env {
			env[e.Name] = e.Value
		}
	}
	return &PodConfig{pod.Spec.Containers[0].Name, env}, nil
}

func firstContainerName(c kclient.PodInterface, podName string) (string, error) {
	pod, err := c.Get(podName)
	if err != nil {
		return "", err
	}

	return pod.Spec.Containers[0].Name, nil
}

func isReady(oc *util.CLI, podName string, pingCommand, expectedOutput string) (bool, error) {
	out, err := executeShellCommand(oc, podName, pingCommand)
	ok := strings.Contains(out, expectedOutput)
	if !ok {
		err = fmt.Errorf("Expected output: %q but actual: %q", expectedOutput, out)
	}
	return ok, err
}

func executeShellCommand(oc *util.CLI, podName string, command string) (string, error) {
	out, err := oc.Run("exec").Args(podName, "--", "bash", "-c", command).Output()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return "", nil
		default:
			return "", err
		}
	}

	return out, nil
}
