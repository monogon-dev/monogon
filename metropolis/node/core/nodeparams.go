package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/metropolis/proto/api"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/osbase/supervisor"
)

func nodeParamsFWCFG(ctx context.Context) (*apb.NodeParameters, error) {
	bytes, err := os.ReadFile("/sys/firmware/qemu_fw_cfg/by_name/dev.monogon.metropolis/parameters.pb/raw")
	if err != nil {
		return nil, fmt.Errorf("could not read firmware enrolment file: %w", err)
	}

	var config apb.NodeParameters
	err = proto.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return &config, nil
}

// nodeParamsGCPMetadata attempts to retrieve the node parameters from the
// GCP metadata service. Returns nil if the metadata service is available,
// but no node parameters are specified.
func nodeParamsGCPMetadata(ctx context.Context) (*apb.NodeParameters, error) {
	const metadataURL = "http://169.254.169.254/computeMetadata/v1/instance/attributes/metropolis-node-params"
	req, err := http.NewRequestWithContext(ctx, "GET", metadataURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}
	decoded, err := io.ReadAll(base64.NewDecoder(base64.StdEncoding, resp.Body))
	if err != nil {
		return nil, fmt.Errorf("cannot decode base64: %w", err)
	}
	var config apb.NodeParameters
	err = proto.Unmarshal(decoded, &config)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshalling NodeParameters: %w", err)
	}
	return &config, nil
}

func getDMIBoardName() (string, error) {
	b, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		return "", fmt.Errorf("could not read board name: %w", err)
	}
	return strings.TrimRight(string(b), "\n"), nil
}

func isGCPInstance(boardName string) bool {
	return boardName == "Google Compute Engine"
}

func getLocalNodeParams(ctx context.Context, storage *localstorage.Root) (*apb.NodeParameters, error) {
	// Retrieve node parameters from qemu's fwcfg interface or ESP.
	// TODO(q3k): probably abstract this away and implement per platform/build/...
	paramsFWCFG, err := nodeParamsFWCFG(ctx)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			supervisor.Logger(ctx).Infof("No qemu fwcfg params.")
		} else {
			supervisor.Logger(ctx).Warningf("Could not retrieve node parameters from qemu fwcfg: %v", err)
		}
		paramsFWCFG = nil
	} else {
		supervisor.Logger(ctx).Infof("Retrieved node parameters from qemu fwcfg")
	}
	paramsESP, err := storage.ESP.Metropolis.NodeParameters.Unmarshal()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			supervisor.Logger(ctx).Infof("No ESP node parameters.")
		} else {
			supervisor.Logger(ctx).Warningf("Could not retrieve node parameters from ESP: %v", err)
		}
		paramsESP = nil
	} else {
		supervisor.Logger(ctx).Infof("Retrieved node parameters from ESP")
	}
	if paramsFWCFG == nil && paramsESP == nil {
		return nil, fmt.Errorf("could not find node parameters in ESP or qemu fwcfg")
	}
	if paramsFWCFG != nil && paramsESP != nil {
		supervisor.Logger(ctx).Warningf("Node parameters found both in both ESP and qemu fwcfg, using the latter")
		return paramsFWCFG, nil
	} else if paramsFWCFG != nil {
		return paramsFWCFG, nil
	} else {
		return paramsESP, nil
	}
}

func getNodeParams(ctx context.Context, storage *localstorage.Root) (*apb.NodeParameters, error) {
	boardName, err := getDMIBoardName()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			supervisor.Logger(ctx).Infof("Board name: UNKNOWN")
		} else {
			supervisor.Logger(ctx).Warningf("Could not get board name, cannot detect platform: %v", err)
		}
	} else {
		supervisor.Logger(ctx).Infof("Board name: %q", boardName)
	}

	// When running on GCP, attempt to retrieve the node parameters from the
	// metadata server first. Retry until we get a response, since we need to
	// wait for the network service to assign an IP address first.
	if isGCPInstance(boardName) {
		var params *apb.NodeParameters
		op := func() error {
			supervisor.Logger(ctx).Info("Running on GCP, attempting to retrieve node parameters from metadata server")
			params, err = nodeParamsGCPMetadata(ctx)
			return err
		}
		err := backoff.Retry(op, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
		if err != nil {
			supervisor.Logger(ctx).Errorf("Failed to retrieve node parameters: %v", err)
		}
		if params != nil {
			supervisor.Logger(ctx).Info("Retrieved parameters from GCP metadata server")
			return params, nil
		}
		supervisor.Logger(ctx).Infof("\"metropolis-node-params\" metadata not found")
	}

	return getLocalNodeParams(ctx, storage)
}
