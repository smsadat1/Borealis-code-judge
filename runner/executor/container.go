// prepares container with all OCI specs
package executor

import (
	"context"
	"fmt"
	"local/runner/utils"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/oci"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func buildSpecOpts(
	image containerd.Image,
	command string,
	rules utils.ExecRules,
) []oci.SpecOpts {

	memoryBytes := uint64(rules.MemoryLimitMB * 1024 * 1024)
	quota := int64(rules.CpuCores * 100000)
	period := uint64(100000)
	// ociEnvs := parseToEnv(rules.Env)

	opts := []oci.SpecOpts{
		// image
		oci.WithImageConfig(image),

		// resource limits
		oci.WithMemoryLimit(memoryBytes),
		oci.WithPidsLimit(rules.PidLimit),
		oci.WithCPUShares(rules.CpuShares),
		oci.WithCPUCFS(quota, period),

		// env
		// oci.WithEnv(ociEnvs),
	}

	// file mount
	if rules.HostSrcpath != "" && rules.ContainerDestPath != "" {
		opts = append(opts, oci.WithMounts([]specs.Mount{
			{
				Type:        "bind",
				Source:      rules.HostSrcpath,
				Destination: rules.ContainerDestPath,
				/* "ro" makes it read-only for security, "rw" makes writable
				* "rbind" ensures sub-mounts are included, "nodev"/"nosuid" are standard sandbox protections
				 */
				Options: []string{"rbind", "ro", "nodev", "nosuid", "create=dir"},
			},
		}))
	}

	// opts = []oci.SpecOpts{
	// 	oci.WithCapabilities(stageCapabilities(rules.Stage)),
	// }

	if rules.ReadOnlyRootfs {
		opts = append(opts, oci.WithRootFSReadonly())
	}

	// evaluated last to guarantee execution parameters survive
	if command != "" {
		processArgs := []string{"/bin/sh", "-c", command}
		opts = append(opts, oci.WithProcessArgs(processArgs...))
	}

	return opts
}

func enforceContainerLimits(
	ctx context.Context,
	client *containerd.Client,
	image containerd.Image,
	rules utils.ExecRules,
) (containerd.Container, string, error) {

	if image == nil {
		return nil, "", fmt.Errorf("cannot enforce limits: provided image object is nil")
	}

	fmt.Printf("Command: %v\n", rules.Command)

	snapshotID := rules.ContainerID + "-snapshot"

	opts := buildSpecOpts(image, rules.Command, rules)

	container, err := client.NewContainer(
		ctx,
		rules.ContainerID,
		containerd.WithNewSnapshot(snapshotID, image),
		containerd.WithNewSpec(opts...),
		// containerd.WithRuntime("runsc", nil), // gVisor interception
	)

	if err != nil {
		return nil, "", err
	}

	return container, snapshotID, nil
}
