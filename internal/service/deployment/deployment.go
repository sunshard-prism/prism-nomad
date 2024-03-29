// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package deployment

import (
	"fmt"
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/parser"
	"slices"
	"time"

	"github.com/hashicorp/nomad/api"
)

type Deployment struct {
	parser  parser.Parser
	builder builder.StructureBuilder
	changes builder.Changes
}

func NewDeployment(
	parser parser.Parser,
	builder builder.StructureBuilder,
	changes builder.Changes,
) *Deployment {
	return &Deployment{
		parser:  parser,
		builder: builder,
		changes: changes,
	}
}

// Checks whether the namespace exists in the cluster.
// If the --create-namespace flag is specified and
// the specified namespace does not exist, then it will be created.
func (s *Deployment) CheckNamespace(namespace model.CheckNamespace) error {
	var namespaces []string
	client := namespace.Client

	availableNamespaces, _, err := client.Namespaces().List(&api.QueryOptions{})
	if err != nil {
		return err
	}

	for _, n := range availableNamespaces {
		namespaces = append(namespaces, n.Name)
	}

	if !slices.Contains(namespaces, namespace.Namespace) {
		if namespace.CreateNamespace {
			newNamespace := &api.Namespace{
				Name: namespace.Namespace,
			}

			_, err := client.Namespaces().Register(
				newNamespace,
				&api.WriteOptions{},
			)

			if err != nil {
				return fmt.Errorf("error create namespace %s", err)
			}

			fmt.Printf(
				"Namespace \"%s\" was successfully created.\n",
				namespace.Namespace,
			)

			return nil
		}

		return fmt.Errorf(
			"%s %s",
			"The specified namespace does not exist.",
			"Use the --create-namespace flag to create a new namespace.",
		)
	}

	return nil
}

// Job configuration deployment in the nomad cluster.
func (s *Deployment) Deployment(d model.Deployment) (string, error) {
	jobConfig, err := d.Client.Jobs().ParseHCL(d.Config, true)
	if err != nil {
		return d.JobName, fmt.Errorf("failed to parse hcl: %s", err)
	}

	writeOptions := &api.WriteOptions{
		Namespace: d.Namespace,
	}

	_, _, err = d.Client.Jobs().Register(jobConfig, writeOptions)
	if err != nil {
		return *jobConfig.ID, fmt.Errorf("job registration error, %s", err)
	}

	fmt.Printf("Running of job \"%s\" deployment.\n", *jobConfig.ID)

	timeNow := time.Now().UTC()
	deployment := &api.Deployment{}

	queryOptions := &api.QueryOptions{
		Namespace: d.Namespace,
	}

	for {
		var (
			jobName          string
			jobStatus        string
			deploymentStatus string
			allocationStatus string

			startTime = fmt.Sprint(
				time.Duration(time.Since(timeNow).Seconds()) * time.Second,
			)
		)

		timeIsOver := time.Duration(
			time.Since(timeNow).Seconds(),
		)*time.Second == time.Duration(d.WaitTime)*time.Second

		if timeIsOver {
			return *jobConfig.ID, fmt.Errorf("job deployment time out has expired")
		}

		job, _, err := d.Client.Jobs().Info(*jobConfig.ID, queryOptions)
		if err != nil {
			return *jobConfig.ID, fmt.Errorf("failed to get job status: %s", err)
		}

		if job != nil {
			jobName = *job.Name
			jobStatus = *job.Status

			deployment, _, err = d.Client.Jobs().LatestDeployment(*job.ID, queryOptions)
			if err != nil {
				return *job.Name, fmt.Errorf(
					"failed to get deployment status: %s", err,
				)
			}

			if deployment != nil {
				deploymentStatus = deployment.Status

				allocation, _, err := d.Client.Jobs().Allocations(
					deployment.JobID, false, queryOptions,
				)

				if err != nil {
					return *job.Name, fmt.Errorf(
						"failed to get allocation status: %s", err,
					)
				}

				if len(allocation) > 0 {
					allocationStatus = allocation[0].ClientStatus

					if allocationStatus == "failed" {
						printDeploymentStatus(
							jobName, startTime, jobStatus,
							deploymentStatus, allocationStatus,
						)

						return jobName, fmt.Errorf("allocation status \"failed\"")
					}

					if allocationStatus == "running" {
						if jobStatus == "dead" {
							printDeploymentStatus(
								jobName, startTime, jobStatus,
								deploymentStatus, allocationStatus,
							)

							return jobName, fmt.Errorf("job status \"dead\"")
						}

						if jobStatus == "running" {
							if deploymentStatus == "successful" {
								printDeploymentStatus(
									jobName, startTime, jobStatus,
									deploymentStatus, allocationStatus,
								)

								return jobName, nil
							}
						}
					}

					printDeploymentStatus(
						jobName, startTime, jobStatus,
						deploymentStatus, allocationStatus,
					)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func printDeploymentStatus(
	jobName, startTime, jobStatus, deploymentStatus, allocationStatus string,
) {
	fmt.Printf("Job deployment \"%s\" started %s ago\n", jobName, startTime)
	fmt.Printf("Job status: \t\t%+v\n", jobStatus)
	fmt.Printf("Deployment status: \t%+v\n", deploymentStatus)
	fmt.Printf("Allocation status: \t%+v\n\n", allocationStatus)
}
