package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type serviceConfig struct {
	YamlFile  string
	Port      string
	AdminPort string
}

func main() {
	// Define your service configurations.
	services := []serviceConfig{
		{
			YamlFile:  "app-local-service1.yaml",
			Port:      "8888",
			AdminPort: "8000",
		},
		{
			YamlFile:  "app-local-service2.yaml",
			Port:      "8889",
			AdminPort: "8001",
		},
	}

	// cache directory for the project running by appserver.py:
	baseStorage := "/Users/mathewmozaffari/speer/api-admin-cache/demo-cache"

	if err := os.MkdirAll(baseStorage, 0755); err != nil {
		fmt.Printf("Error creating storage path: %v\n", err)
		return
	}

	// Path to the dev_appserver.py script.
	devAppServerPath := "/Users/mathewmozaffari/google-cloud-sdk/bin/dev_appserver.py"

	var wg sync.WaitGroup

	yamlFiles := []string{}
	for _, svc := range services {
		yamlFiles = append(yamlFiles, svc.YamlFile)
	}

	args := []string{
		"--max_module_instances=1",
		"--storage_path=" + baseStorage,
	}

	args = append(args, yamlFiles...)

	cmd := exec.Command("python3", append([]string{devAppServerPath}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	wg.Add(1)
	go func() {
		fmt.Println("Starting all services with a shared storage path...")
		if err := cmd.Run(); err != nil {
			fmt.Printf("Dev appserver exited with error: %v\n", err)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Dev appserver has exited.")
}
