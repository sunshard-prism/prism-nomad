// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"prism/pkg"
	"regexp"
)

func artifact(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"options", "headers"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "options":
			artifactOptions(&block.Block[index], &blockChanges)
		case "headers":
			artifactHeaders(&block.Block[index], &blockChanges)
		}
	}
}

func artifactOptions(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func artifactHeaders(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func affinity(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func changeScript(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func check(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"header", "check_restart"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "header":
			checkHeader(&block.Block[index], &blockChanges)
		case "check_restart":
			checkRestart(&block.Block[index], &blockChanges)
		}
	}
}

func checkHeader(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func checkRestart(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func connect(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"sidecar_service", "sidecar_task", "gateway"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "sidecar_service":
			sidecarService(&block.Block[index], &blockChanges)
		case "sidecar_task":
			sidecarTask(&block.Block[index], &blockChanges)
		case "gateway":
			gateway(&block.Block[index], &blockChanges)
		}
	}
}

func consul(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)

	for index, item := range block.Parameter {
		for k := range item {
			if k == "namespace" {
				block.Parameter[index][k] = changes.Namespace
			}
		}
	}
}

func constraint(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func csiPlugin(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

// Common.
func config(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func device(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"affinity", "constraint"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	if changes.Release != "" {
		block.Label = fmt.Sprintf("%s-%s", block.Label, changes.Release)
	}

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "affinity":
			affinity(&block.Block[index], &blockChanges)
		case "constraint":
			constraint(&block.Block[index], &blockChanges)
		}
	}
}

func dispatchPayload(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

// Common.
func env(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func ephemeralDisk(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func expose(block *model.TemplateBlock, changes *model.BlockChanges) {
	unnamedDublicateBlock := map[string]string{"path": "path"}

	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, unnamed, unnamedDublicateBlock,
		)

		if item.Type == "path" {
			exposePath(&block.Block[index], &blockChanges)
		}
	}
}

func exposePath(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func gateway(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"proxy", "ingress", "terminating"}

	checkSingleBlocks(block, &changes.File, singleBlock)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "proxy":
			gatewayProxy(&block.Block[index], &blockChanges)
		case "ingress":
			gatewayIngress(&block.Block[index], &blockChanges)
		case "terminating":
			gatewayTerminating(&block.Block[index], &blockChanges)
		}
	}
}

func gatewayProxy(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"config"}
	namedDuplicateBlock := []string{"envoy_gateway_bind_address"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkNamedDublicateBlocks(block, &changes.File, namedDuplicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {

		switch item.Type {
		case "config":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			config(&block.Block[index], &blockChanges)
		case "envoy_gateway_bind_address":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, named,
			)
			gatewayProxyAddress(&block.Block[index], &blockChanges)
		}
	}
}

func gatewayProxyAddress(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func gatewayIngress(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"tls"}
	unnamedDublicateBlock := map[string]string{"listener": "port"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)

	for index, item := range block.Block {

		switch item.Type {
		case "tls":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			gatewayIngressTLS(&block.Block[index], &blockChanges)
		case "listener":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, unnamed, unnamedDublicateBlock,
			)
			gatewayIngressListener(&block.Block[index], &blockChanges)
		}
	}
}

func gatewayIngressTLS(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func gatewayIngressListener(block *model.TemplateBlock, changes *model.BlockChanges) {
	unnamedDublicateBlock := map[string]string{"service": "name"}

	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, unnamed, unnamedDublicateBlock,
		)

		if item.Type == "service" {
			gatewayIngressListenerService(&block.Block[index], &blockChanges)
		}
	}
}

func gatewayIngressListenerService(
	block *model.TemplateBlock,
	changes *model.BlockChanges,
) {
	setFileChanges(block, &changes.File)
}

func gatewayTerminating(block *model.TemplateBlock, changes *model.BlockChanges) {
	unnamedDublicateBlock := map[string]string{"service": "name"}

	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, unnamed, unnamedDublicateBlock,
		)

		if item.Type == "service" {
			gatewayTerminatingService(&block.Block[index], &blockChanges)
		}
	}
}

func gatewayTerminatingService(
	block *model.TemplateBlock,
	changes *model.BlockChanges,
) {
	setFileChanges(block, &changes.File)
}

func group(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{
		"affinity",
		"consul",
		"constraint",
		"ephemeral_disk",
		"migrate",
		"network",
		"reschedule",
		"scaling",
		"spread",
		"update",
		"meta",
		"restart",
		"vault",
	}

	unnamedDublicateBlock := map[string]string{"service": "name"}
	namedDublicateBlock := []string{"task", "volume"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)
	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)

	setFileChanges(block, &changes.File)

	if changes.Release != "" {
		block.Label = fmt.Sprintf("%s-%s", block.Label, changes.Release)
	}

	for index, item := range block.Block {
		blockChanges := checkFileChanges(&block.Block[index], changes, single)

		switch item.Type {
		case "affinity":
			affinity(&block.Block[index], &blockChanges)
		case "consul":
			consul(&block.Block[index], &blockChanges)
		case "constraint":
			constraint(&block.Block[index], &blockChanges)
		case "ephemeral_disk":
			ephemeralDisk(&block.Block[index], &blockChanges)
		case "network":
			network(&block.Block[index], &blockChanges)
		case "migrate":
			migrate(&block.Block[index], &blockChanges)
		case "reschedule":
			reschedule(&block.Block[index], &blockChanges)
		case "spread":
			spread(&block.Block[index], &blockChanges)
		case "update":
			update(&block.Block[index], &blockChanges)
		case "task":
			blockChanges := checkFileChanges(&block.Block[index], changes, named)
			task(&block.Block[index], &blockChanges)
		case "scaling":
			scaling(&block.Block[index], &blockChanges)
		case "service":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, unnamed, unnamedDublicateBlock,
			)
			service(&block.Block[index], &blockChanges)
		case "meta":
			meta(&block.Block[index], &blockChanges)
		case "restart":
			restart(&block.Block[index], &blockChanges)
		case "volume":
			blockChanges := checkFileChanges(&block.Block[index], changes, named)
			volume(&block.Block[index], &blockChanges)
		case "vault":
			vault(&block.Block[index], &blockChanges)
		}
	}
}

func identity(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func job(block *model.TemplateBlock, changes *model.BlockChanges) {
	var (
		haveType      bool
		haveNamespace bool
		haveMeta      bool
	)

	// Get type from pack file and set namespace.
	for index, item := range block.Parameter {
		for key := range item {
			switch key {
			case "type":
				haveType = true
				block.Parameter[index][key] = changes.Pack.Type
			case "namespace":
				haveNamespace = true
				block.Parameter[index][key] = changes.Namespace
			}
		}
	}

	for _, i := range block.Block {
		if i.Type == "meta" {
			haveMeta = true
		}
	}

	if !haveType {
		schedulerType := make(map[string]interface{})
		schedulerType["type"] = changes.Pack.Type
		block.Parameter = append(block.Parameter, schedulerType)
	}

	if !haveNamespace {
		namespaceParameter := map[string]interface{}{"namespace": changes.Namespace}
		block.Parameter = append(block.Parameter, namespaceParameter)
	}

	if !haveMeta {
		meta := model.TemplateBlock{
			Type: "meta",
		}

		deployVersion := map[string]interface{}{"run_uuid": changes.Pack.DeployVersion}
		meta.Parameter = append(meta.Parameter, deployVersion)

		block.Block = append(block.Block, meta)
	}

	// Checking for blocks.
	// If the block is not in the configuration, it will be added.
	singleBlock := []string{
		"affinity",
		"constraint",
		"meta",
		"multiregion",
		"parameterized",
		"periodic",
		"migrate",
		"reschedule",
		"spread",
		"update",
		"vault",
	}

	namedDublicateBlock := []string{"group"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)

	// Set changes.
	setFileChanges(block, &changes.File)

	// Adding the release name to the job name.
	if changes.Release != "" {
		block.Label = fmt.Sprintf("%s-%s", block.Label, changes.Release)
	}

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "affinity":
			affinity(&block.Block[index], &blockChanges)
		case "constraint":
			constraint(&block.Block[index], &blockChanges)
		case "group":
			blockChanges = checkFileChanges(
				&block.Block[index], changes, named,
			)
			group(&block.Block[index], &blockChanges)
		case "meta":
			jobMeta(&block.Block[index], &blockChanges)
		case "multiregion":
			multiregion(&block.Block[index], &blockChanges)
		case "parameterized":
			parameterized(&block.Block[index], &blockChanges)
		case "periodic":
			periodic(&block.Block[index], &blockChanges)
		case "migrate":
			migrate(&block.Block[index], &blockChanges)
		case "reschedule":
			reschedule(&block.Block[index], &blockChanges)
		case "spread":
			spread(&block.Block[index], &blockChanges)
		case "update":
			update(&block.Block[index], &blockChanges)
		case "vault":
			vault(&block.Block[index], &blockChanges)
		}
	}
}

func jobMeta(block *model.TemplateBlock, changes *model.BlockChanges) {
	var haveUUID bool

	for index, p := range block.Parameter {
		for key := range p {
			if key == "run_uuid" {
				haveUUID = true
				block.Parameter[index][key] = changes.Pack.DeployVersion
			}
		}
	}

	if !haveUUID {
		deployVersion := map[string]interface{}{"run_uuid": changes.Pack.DeployVersion}
		block.Parameter = append(block.Parameter, deployVersion)
	}

	setFileChanges(block, &changes.File)
}

func lifecycle(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func logs(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

// Common.
func meta(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func migrate(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func multiregion(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"strategy"}
	namedDublicateBlock := []string{"region"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)

	for index, item := range block.Block {

		switch item.Type {
		case "strategy":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			multiregionStrategy(&block.Block[index], &blockChanges)
		case "region":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, named,
			)
			multiregionRegion(&block.Block[index], &blockChanges)
		}
	}
}

func multiregionStrategy(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func multiregionRegion(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"meta"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		if item.Type == "meta" {
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			meta(&block.Block[index], &blockChanges)
		}
	}
}

func network(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"dns"}
	namedDublicateBlock := []string{"port"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {

		switch item.Type {
		case "dns":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			networkDNS(&block.Block[index], &blockChanges)
		case "port":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, named,
			)
			networkPort(&block.Block[index], &blockChanges)
		}
	}
}

func networkPort(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func networkDNS(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func numa(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func parameterized(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func periodic(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func proxy(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"config", "upstream", "expose"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "config":
			config(&block.Block[index], &blockChanges)
		case "upstreams":
			upstreams(&block.Block[index], &blockChanges)
		case "expose":
			expose(&block.Block[index], &blockChanges)
		}
	}
}

func reschedule(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func resources(block *model.TemplateBlock, changes *model.BlockChanges) {
	namedDublicateBlock := []string{"device"}

	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, named,
		)

		switch item.Type {
		case "device":
			device(&block.Block[index], &blockChanges)
		case "numa":
			numa(&block.Block[index], &blockChanges)
		}
	}
}

func restart(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func scaling(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"policy"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		if item.Type == "policy" {
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			scalingPolicy(&block.Block[index], &blockChanges)
		}
	}
}

func scalingPolicy(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func service(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{
		"tagged_addresses",
		"check_restart",
		"meta",
		"canary_meta",
		"connect",
	}

	unnamedDublicateBlock := map[string]string{"check": "name"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(&block.Block[index], changes, single)

		switch item.Type {
		case "tagged_addresses":
			serviceTaggedAddresses(&block.Block[index], &blockChanges)
		case "check":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, unnamed, unnamedDublicateBlock,
			)
			check(&block.Block[index], &blockChanges)
		case "check_restart":
			checkRestart(&block.Block[index], &blockChanges)
		case "meta":
			meta(&block.Block[index], &blockChanges)
		case "canary_meta":
			serviceCanaryMeta(&block.Block[index], &blockChanges)
		case "connect":
			connect(&block.Block[index], &blockChanges)
		}
	}
}

func serviceCanaryMeta(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func serviceTaggedAddresses(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func sidecarService(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"meta", "proxy"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(&block.Block[index], changes, single)

		switch item.Type {
		case "meta":
			meta(&block.Block[index], &blockChanges)
		case "proxy":
			proxy(&block.Block[index], &blockChanges)
		}
	}
}

func sidecarTask(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{
		"config",
		"env",
		"meta",
		"logs",
		"resources",
	}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "config":
			config(&block.Block[index], &blockChanges)
		case "env":
			env(&block.Block[index], &blockChanges)
		case "meta":
			meta(&block.Block[index], &blockChanges)
		case "logs":
			logs(&block.Block[index], &blockChanges)
		case "resources":
			resources(&block.Block[index], &blockChanges)
		}
	}
}

func spread(block *model.TemplateBlock, changes *model.BlockChanges) {
	namedDublicateBlock := []string{"target"}

	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		if item.Type == "target" {
			blockChanges := checkFileChanges(
				&block.Block[index], changes, named,
			)
			spreadTarget(&block.Block[index], &blockChanges)
		}
	}
}

func spreadTarget(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func task(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{
		"artifact",
		"affinity",
		"config",
		"constraint",
		"env",
		"meta",
		"resources",
		"csi_plugin",
		"dispatch_payload",
		"identity",
		"lifecycle",
		"logs",
		"restart",
		"vault",
	}

	unnamedDublicateBlock := map[string]string{
		"service":      "name",
		"template":     "name",
		"volume_mount": "volume",
	}

	namedDublicateBlock := []string{"scaling"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	checkNamedDublicateBlocks(block, &changes.File, namedDublicateBlock)
	checkUnnamedDublicateBlocks(block, &changes.File, unnamedDublicateBlock)

	setFileChanges(block, &changes.File)

	if changes.Release != "" {
		block.Label = fmt.Sprintf("%s-%s", block.Label, changes.Release)
	}

	for index, item := range block.Block {
		blockChanges := checkFileChanges(&block.Block[index], changes, single)

		switch item.Type {
		case "artifact":
			artifact(&block.Block[index], &blockChanges)
		case "affinity":
			affinity(&block.Block[index], &blockChanges)
		case "consul":
			consul(&block.Block[index], &blockChanges)
		case "config":
			config(&block.Block[index], &blockChanges)
		case "constraint":
			constraint(&block.Block[index], &blockChanges)
		case "env":
			env(&block.Block[index], &blockChanges)
		case "meta":
			meta(&block.Block[index], &blockChanges)
		case "template":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, unnamed, unnamedDublicateBlock,
			)
			template(&block.Block[index], &blockChanges)
		case "csi_plugin":
			csiPlugin(&block.Block[index], &blockChanges)
		case "resources":
			resources(&block.Block[index], &blockChanges)
		case "dispatch_payload":
			dispatchPayload(&block.Block[index], &blockChanges)
		case "identity":
			identity(&block.Block[index], &blockChanges)
		case "lifecycle":
			lifecycle(&block.Block[index], &blockChanges)
		case "logs":
			logs(&block.Block[index], &blockChanges)
		case "restart":
			restart(&block.Block[index], &blockChanges)
		case "scaling":
			blockChanges := checkFileChanges(&block.Block[index], changes, named)
			scaling(&block.Block[index], &blockChanges)
		case "service":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, unnamed, unnamedDublicateBlock,
			)
			service(&block.Block[index], &blockChanges)
		case "vault":
			vault(&block.Block[index], &blockChanges)
		case "volume_mount":
			blockChanges := checkFileChanges(
				&block.Block[index], changes, unnamed, unnamedDublicateBlock,
			)
			volumeMount(&block.Block[index], &blockChanges)
		}
	}
}

func template(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"wait", "change_script"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for _, item := range block.Parameter {
		for k, v := range item {
			if k == "file" {
				var fileFullPath string

				// Check the full file path or file name.
				separatorFormat, err := regexp.Compile(`\\|\/`)
				if err != nil {
					fmt.Printf("failed check OS separator in file path, %s", err)
					os.Exit(1)
				}

				findSeparator := separatorFormat.FindStringSubmatch(v.(string))

				if len(findSeparator) > 0 {
					fileFullPath = v.(string)
				} else {
					fileFullPath = filepath.Join(changes.FilesDirPath, v.(string))
				}

				// Read the file and add data to the "data" parameter.
				file, err := os.ReadFile(fileFullPath)
				if err != nil {
					fmt.Printf("error read template files - %v\n", err)
					os.Exit(1)
				}

				i := make(map[string]interface{})
				i["data"] = string(file)
				block.Parameter = append(block.Parameter, i)

				pkg.RemoveParameter(block, "file")
			}
		}
	}

	pkg.RemoveParameter(block, "name")

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "wait":
			templateWait(&block.Block[index], &blockChanges)
		case "change_script":
			changeScript(&block.Block[index], &blockChanges)
		}
	}
}

func templateWait(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func update(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func upstreams(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"config", "mesh_gateway"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Parameter {
		for k := range item {
			if k == "destination_namespace" {
				block.Parameter[index][k] = changes.Namespace
			}
		}
	}

	for index, item := range block.Block {
		blockChanges := checkFileChanges(
			&block.Block[index], changes, single,
		)

		switch item.Type {
		case "config":
			config(&block.Block[index], &blockChanges)
		case "mesh_gateway":
			upstreamMeshGateway(&block.Block[index], &blockChanges)
		}
	}
}

func upstreamMeshGateway(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func vault(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)

	for index, item := range block.Parameter {
		for k := range item {
			if k == "namespace" {
				block.Parameter[index][k] = changes.Namespace
			}
		}
	}
}

func volume(block *model.TemplateBlock, changes *model.BlockChanges) {
	singleBlock := []string{"mount_options"}

	checkSingleBlocks(block, &changes.File, singleBlock)
	setFileChanges(block, &changes.File)

	for index, item := range block.Block {
		if item.Type == "mount_options" {
			blockChanges := checkFileChanges(
				&block.Block[index], changes, single,
			)
			volumeMountOptions(&block.Block[index], &blockChanges)
		}
	}
}

func volumeMountOptions(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}

func volumeMount(block *model.TemplateBlock, changes *model.BlockChanges) {
	setFileChanges(block, &changes.File)
}
