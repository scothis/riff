/*
 * Copyright 2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"context"
	"fmt"

	"github.com/projectriff/riff/pkg/cli"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RequestProcessorListOptions struct {
	cli.ListOptions
}

func (opts *RequestProcessorListOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := &cli.FieldError{}

	errs = errs.Also(opts.ListOptions.Validate(ctx))

	return errs
}

func NewRequestProcessorListCommand(c *cli.Config) *cobra.Command {
	opts := &RequestProcessorListOptions{}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "<todo>",
		Example: "<todo>",
		Args:    cli.Args(),
		PreRunE: cli.ValidateOptions(opts),
		RunE: func(cmd *cobra.Command, args []string) error {
			requestprocessors, err := c.Request().RequestProcessors(opts.Namespace).List(metav1.ListOptions{})
			if err != nil {
				return err
			}

			if len(requestprocessors.Items) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), "No request processors found.")
			}
			for _, requestprocessor := range requestprocessors.Items {
				// TODO pick a generic table formatter
				fmt.Fprintln(cmd.OutOrStdout(), requestprocessor.Name)
			}

			return nil
		},
	}

	cli.AllNamespacesFlag(cmd, c, &opts.Namespace, &opts.AllNamespaces)

	return cmd
}