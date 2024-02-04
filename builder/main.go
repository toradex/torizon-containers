package main

import (
	"bufio"
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// get every Dockerfile file path
	relativeContextPaths, err := WalkDirAndReturnPathsThatMatch("../debian-docker-images/upstream", "Dockerfile")
	var relativeDockerfilePaths []string
	var relativeBuildContexts []string

	for _, path := range relativeContextPaths {
		relativeDockerfilePaths = append(relativeDockerfilePaths, path+"/Dockerfile")
		relativeBuildContexts = append(relativeBuildContexts, path)
	}
	if err != nil {
		log.Println("error: error while walking directory")
	}

	// for every Dockerfile found, parse it to get the labels containing the build metadata
	var labels []map[string]string
	log.Println("Parsed Dockerfiles:")
	for _, path := range relativeBuildContexts {
		l, err := parseDockerfileLabels(path + "/Dockerfile")
		if err != nil {
			log.Println("Error:", err)
			log.Println("Aborting")
			os.Exit(1)
		}
		log.Println(path)
		l["path"] = path
		labels = append(labels, l)
	}

	//log.Println(labels)

	dependencyMap := make(map[string][]string)

	log.Println("Generating a dependency map...")
	log.Println("Dependency Map:")
	for _, label := range labels {
		child := label["org.toradex.image.namespace"] + "/" + label["org.toradex.image.name"]
		parent := label["org.toradex.image.base.namespace"] + "/" + label["org.toradex.image.base.name"]

		// Check if the child already exists in the dependency map for the parent
		exists := false
		for _, existingChild := range dependencyMap[parent] {
			if existingChild == child {
				exists = true
				break
			}
		}

		if !exists {
			dependencyMap[parent] = append(dependencyMap[parent], child)
		}
	}

	startKey := "library/debian"

	buildWorld := os.Getenv("BUILD_WORLD")
	if buildWorld == "1" {
		var buildOrder []string

		traverseMap(dependencyMap, startKey, make(map[string]bool), &buildOrder)
		log.Println("World Build Order is:")
		log.Print(buildOrder)

		log.Println("Building...")
		tasks := matchOrderWithLabels(labels, buildOrder)
		log.Println(tasks)
		for _, task := range tasks {
			log.Printf("Building %s\n", task["IMAGE"])

			// can't pass the map directly, so a copy is needed
			baseImageRegistry := task["BASE_IMAGE_REGISTRY"]
			baseImageNamespace := task["BASE_IMAGE_NAMESPACE"]
			baseImageName := task["BASE_IMAGE_NAME"]
			baseImageTag := task["BASE_IMAGE_TAG"]
			contextPath := task["DOCKERFILE_PATH"]

			buildArgs := map[string]*string{
				"BASE_IMAGE_REGISTRY":  &baseImageRegistry,
				"BASE_IMAGE_NAMESPACE": &baseImageNamespace,
				"BASE_IMAGE_NAME":      &baseImageName,
				"BASE_IMAGE_TAG":       &baseImageTag,
			}

			err = buildImage(contextPath, buildArgs)
			if err != nil {
				log.Printf("Error building image: %v\n", err)
				os.Exit(1)
			}
		}
	}
}

func traverseMap(dependencyMap map[string][]string, currentKey string, visited map[string]bool, resultSlice *[]string) {
	if visited[currentKey] {
		return
	}

	visited[currentKey] = true
	*resultSlice = append(*resultSlice, currentKey)

	for _, child := range dependencyMap[currentKey] {
		traverseMap(dependencyMap, child, visited, resultSlice)
	}
}

func matchOrderWithLabels(labels []map[string]string, buildOrder []string) []map[string]string {
	result := make([]map[string]string, 0)

	for _, member := range buildOrder {
		for _, label := range labels {
			ImageNamespace := label["org.toradex.image.namespace"]
			ImageName := label["org.toradex.image.name"]
			ImageTag := label["org.toradex.image.tag.major"]

			BaseImageNamespace := label["org.toradex.image.base.namespace"]
			BaseImageName := label["org.toradex.image.base.name"]

			var BaseImageTag string
			if BaseImageNamespace == "library" && BaseImageName == "debian" {
				// debian follows the MAJOR.MINOR-VARIANT scheme instead of full semantic versioning, and will be hosted in the main "library" namespace
				BaseImageTag = label["org.toradex.image.base.tag.major"] + "." + label["org.toradex.image.base.tag.minor"] + "-" + label["org.toradex.image.base.variant"]
			} else if BaseImageNamespace == "torizon" && BaseImageName == "debian" {
				// the torizon debian image follows a MAJOR-VARIANT
				BaseImageTag = label["org.toradex.image.base.tag.major"] + "-" + label["org.toradex.image.base.variant"]
			} else {
				// for all other images the major suffices
				BaseImageTag = label["org.toradex.image.base.tag.major"]
			}

			if ImageNamespace+"/"+ImageName == member {
				result = append(result, map[string]string{
					"IMAGE":                member,
					"IMAGE_REGISTRY":       label["org.toradex.image.registry"],
					"IMAGE_NAMESPACE":      ImageNamespace,
					"IMAGE_NAME":           ImageName,
					"IMAGE_TAG":            ImageTag,
					"IMAGE_ARCH":           label["org.toradex.image.arch"],
					"BASE_IMAGE_REGISTRY":  label["org.toradex.image.base.registry"],
					"BASE_IMAGE_NAMESPACE": BaseImageNamespace,
					"BASE_IMAGE_NAME":      BaseImageName,
					"BASE_IMAGE_TAG":       BaseImageTag,
					"DOCKERFILE_PATH":      label["path"],
				})
				// Break the inner loop once a match is found to avoid unnecessary iterations
				break
			}
		}
	}
	return result
}

func matchOrderWithLabels1(labels []map[string]string, buildOrder []string) map[string]map[string]string {
	result := make(map[string]map[string]string)

	for _, member := range buildOrder {
		for _, label := range labels {
			ImageNamespace := label["org.toradex.image.namespace"]
			ImageName := label["org.toradex.image.name"]
			ImageTag := label["org.toradex.image.tag.major"]

			BaseImageNamespace := label["org.toradex.image.base.namespace"]
			BaseImageName := label["org.toradex.image.base.name"]

			var BaseImageTag string
			if BaseImageNamespace == "library" && BaseImageName == "debian" {
				// debian follows the MAJOR.MINOR-VARIANT scheme instead of full semantic versioning, and will be hosted in the main "library" namespace
				BaseImageTag = label["org.toradex.image.base.tag.major"] + "." + label["org.toradex.image.base.tag.minor"] + "-" + label["org.toradex.image.base.variant"]
			} else if BaseImageNamespace == "torizon" && BaseImageName == "debian" {
				// the torizon debian image follows a MAJOR-VARIANT
				BaseImageTag = label["org.toradex.image.base.tag.major"] + "-" + label["org.toradex.image.base.variant"]
			} else {
				// for all other images the major suffices
				BaseImageTag = label["org.toradex.image.base.tag.major"]
			}

			if ImageNamespace+"/"+ImageName == member {
				result[member] = map[string]string{
					"IMAGE_REGISTRY":       label["org.toradex.image.registry"],
					"IMAGE_NAMESPACE":      ImageNamespace,
					"IMAGE_NAME":           ImageName,
					"IMAGE_TAG":            ImageTag,
					"IMAGE_ARCH":           label["org.toradex.image.arch"],
					"BASE_IMAGE_REGISTRY":  label["org.toradex.image.base.registry"],
					"BASE_IMAGE_NAMESPACE": BaseImageNamespace,
					"BASE_IMAGE_NAME":      BaseImageName,
					"BASE_IMAGE_TAG":       BaseImageTag,
					"DOCKERFILE_PATH":      label["path"],
				}
				break
			}
		}
	}
	return result
}

func WalkDirAndReturnPathsThatMatch(root string, name string) ([]string, error) {
	var relativeContextPaths []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if strings.Contains(path, name) {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			completeRelativePath := root + "/" + relPath
			relDir := strings.TrimSuffix(completeRelativePath, "/Dockerfile")
			relativeContextPaths = append(relativeContextPaths, relDir)
		}
		return nil
	})
	return relativeContextPaths, err
}

func buildImage(contextPath string, buildArgs map[string]*string) error {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return err
	}

	// Always assume that contextPath is also the build context
	buildContext, err := archive.TarWithOptions(contextPath, &archive.TarOptions{})
	if err != nil {
		return err
	}

	buildOptions := types.ImageBuildOptions{
		BuildArgs:  buildArgs,
		Dockerfile: "Dockerfile",
		Context:    buildContext,
		Tags:       []string{"testimage:tag"},
	}

	buildResponse, err := cli.ImageBuild(context.Background(), buildContext, buildOptions)
	if err != nil {
		return err
	}

	// Print build logs
	defer buildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		return err
	}

	log.Println("Image built successfully.")

	return nil
}

func parseDockerfileLabels(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var labelLine string
	var foundLastLine bool
	var foundDuplicatedEntries bool
	var foundFirstLine bool

	labels := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// ignore lines that are not LABEL
		if !strings.HasPrefix(line, "LABEL") && !foundFirstLine {
			continue
		}
		if strings.HasPrefix(line, "LABEL") && (strings.HasSuffix(line, `\`)) {
			foundFirstLine = true
			labelLine = strings.TrimSpace(strings.TrimPrefix(line, "LABEL"))
		} else if strings.HasPrefix(line, "LABEL") && !(strings.HasSuffix(line, `\`)) {
			foundLastLine = true
			foundFirstLine = true
		} else if !strings.HasPrefix(line, "LABEL") && strings.HasSuffix(line, `\`) {
			labelLine += strings.TrimSpace(strings.TrimSuffix(line, `\`))
		} else if !strings.HasPrefix(line, "LABEL") && !(strings.HasSuffix(line, `\`)) {
			foundLastLine = true
			labelLine = line
		}

		if labelLine != "" {
			labelPairs := strings.Split(labelLine, " ")
			for _, labelPair := range labelPairs {
				labelKeyValue := strings.SplitN(labelPair, "=", 2)
				if len(labelKeyValue) == 2 {
					key := strings.TrimSpace(labelKeyValue[0])
					value := strings.Trim(labelKeyValue[1], "\"")
					if _, exists := labels[key]; exists {
						log.Printf("Key '%s' already exists with value %s in %s\n", key, labels[key], filePath)
						foundDuplicatedEntries = true
					} else {
						labels[key] = value
					}
				}
			}
			labelLine = ""
		}
		// if we reached the last line, drop out of the scanner loop
		if foundLastLine {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if foundDuplicatedEntries {
		return nil, errors.New("cannot continue with duplicated LABEL statements")
	}

	return labels, nil
}
