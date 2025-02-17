/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/apisix/manager-api/internal/conf"
	"github.com/apisix/manager-api/internal/utils"
)

func newStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "stop Apache APISIX Dashboard service/program",
		Run: func(cmd *cobra.Command, args []string) {
			pid, err := utils.ReadPID(conf.PIDPath)
			if err != nil {
				if syscall.ENOENT.Error() != err.Error() {
					fmt.Fprintf(os.Stderr, "failed to get manager-api pid: %s\n", err)
				} else {
					fmt.Fprintf(os.Stderr, "pid path %s not found, is manager-api running?\n", conf.PIDPath)
				}
				return
			}
			p, err := os.FindProcess(pid)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to find process manager-api: %s\n", err)
				return
			}
			if err := p.Signal(syscall.SIGINT); err != nil {
				fmt.Fprintf(os.Stderr, "failed to kill manager-api: %s", err)
			}
		},
	}
}
