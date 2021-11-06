/*
Copyright 2021 The SHUMIN Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"
	"x6t.io/circle"
)

type TaskService struct {
	userSvc  circle.User
	fetchSvc circle.Fetcher
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) StatisticsTask(ctx context.Context) (circle.Tasks,error)  {

	return nil,nil
}

func (s *TaskService) ProcessTask(ctx context.Context) error  {

	return nil
}