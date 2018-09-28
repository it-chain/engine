/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ivm

type ContainerService interface {
	StartContainer(icode ICode) error
	StopContainer(id ID) error
	ExecuteRequest(request Request) (Result, error)
	GetRunningICodeList() []ICode
}

type GitService interface {
	//clone code from deploy info
	Clone(baseSavePath string, repositoryUrl string, sshPath string, password string) (ICode, error)
	CloneFromRawSsh(baseSavePath string, repositoryUrl string, rawSsh []byte, password string) (ICode, error)
}

type EventService interface {
	Publish(topic string, event interface{}) error
}
