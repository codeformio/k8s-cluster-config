package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates/v1beta1"
	"sigs.k8s.io/yaml"
)

// Merge all <name>.rego files into <name>.yaml constraint templates.
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	regoFiles, err := filepath.Glob("*.rego")
	if err != nil {
		return fmt.Errorf("listing rego files: %w", err)
	}

	for _, regoFile := range regoFiles {
		if strings.HasSuffix(regoFile, "_test.rego") {
			continue
		}

		if err := build(regoFile); err != nil {
			return fmt.Errorf("building constraint: %w", err)
		}
	}

	return nil
}

func build(regoFile string) error {
	regoSrc, err := ioutil.ReadFile(regoFile)
	if err != nil {
		return fmt.Errorf("reading rego file: %w", err)
	}

	k8sFile := strings.TrimSuffix(regoFile, ".rego") + ".yaml"

	k8sYaml, err := ioutil.ReadFile(k8sFile)
	if err != nil {
		return fmt.Errorf("reading gatekeeper template file: %w", err)
	}

	var template v1beta1.ConstraintTemplate
	if err := yaml.Unmarshal(k8sYaml, &template); err != nil {
		return fmt.Errorf("unmarshaling gatekeeper template: %w", err)
	}

	template.Spec.Targets = []v1beta1.Target{
		{
			Target: "admission.k8s.gatekeeper.sh",
			Rego:   string(regoSrc),
		},
	}

	newK8sYaml, err := yaml.Marshal(template)
	if err != nil {
		return fmt.Errorf("marshalling gatekeeper template: %w", err)
	}

	fmt.Printf("Building %v into %v\n", regoFile, k8sFile)
	if err := ioutil.WriteFile(k8sFile, newK8sYaml, 0644); err != nil {
		return fmt.Errorf("writing gatekeeper template file: %w", err)
	}

	return nil
}
