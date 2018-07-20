/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"github.com/spf13/cobra"
	"github.com/projectriff/riff-cli/pkg/core"
)

func System() *cobra.Command {
	return &cobra.Command{
		Use:   "system",
		Short: "Manage system related resources",
	}
}

func SystemInstall(kc *core.KubectlClient) *cobra.Command {
	options := core.SystemInstallOptions{}

	command := &cobra.Command{
		Use:   "install",
		Short: "install the riff and Knative system components",
		Example: `  riff system install`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := (*kc).SystemInstall(options)
			if err != nil {
				return err
			}
			return nil
		},
	}

	command.Flags().BoolVarP(&options.NodePort, "node-port", "", false, "use NodePort instead of LoadBalancer for ingress gateways")

	return command
}
