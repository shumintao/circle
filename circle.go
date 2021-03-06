/*
Copyright 2021 The Xiadat Authors.

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

package circle

import (
	"context"
)

const (
	Offset  = 10
	Version = "1.0.4"
)

// Error is a domain error encountered while processing circle requests
type Error string

func (e Error) Error() string {
	return string(e)
}

// Source is a user request source.
type Source struct {
	URL                string
	Token              string
	Account            string
	Password           string
	Tuisongclientid    string
	InsecureSkipVerify bool
}

// UserInfo return user detail information.
type UserInfo struct {
	ID         int    `json:"id,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	Token      string `json:"token,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	Createtime int    `json:"createtime,omitempty"`
	Expiretime int    `json:"expiretime,omitempty"`
	ExpiresIn  int    `json:"expires_in,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
}

type Tasks []Task

type Task struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	EndTime int    `json:"end_time,omitempty"`
	// The current number of tasks and the specified number of completed tasks.
	// containing user is the statistics of the number of tasks completed by the user prefix.
	ReadCount       int `json:"read_count,omitempty"`        // task readings
	UserReadCount   int `json:"user_read_count,omitempty"`   // User readings
	FinishScore     int `json:"finish_score,omitempty"`      // task credits
	UserScore       int `json:"user_score,omitempty"`        // User's current points
	GroupCountRw    int `json:"group_count_rw,omitempty"`    // The number of task WeChat groups
	UserGroupCount  int `json:"user_group_count,omitempty"`  // User WeChat groups
	CircleCountRw   int `json:"circle_count_rw,omitempty"`   // mission circle of friends
	UserCircleCount int `json:"user_circle_count,omitempty"` // User Moments
	// Statistics on the number of user sharing tasks.
	// after adding the number of WeChat groups to the number of Moments.
	UserTaskCount int `json:"user_task_count,omitempty"` // User WeChat group + Moments sharing times
	// WeChat share pictures and connections.
	MicroURL string `json:"micro_url,omitempty"` // WeChat share link
}

// Fitter is processing parameters.
type Fitter struct {
	Page    string
	Version string
}

type WechatShares []WechatShare

// WechatShare define WeChat sharing parameters.
type WechatShare struct {
	Microgrid string // task id.
	Type      string // task type.
}

type WechatType string

const (
	Friends = "1" // WeChat Moments
	Group   = "2" // WeChat group
)

// Client is an interface for login.
type Client interface {
	Connect(ctx context.Context, src Source) error
	Login(ctx context.Context, u Source) (*UserInfo, error)
	// Fetch is an interface for fetch task list.
	Fetch(ctx context.Context, query Fitter) (Tasks, error)
	Detail(ctx context.Context, microgrid string) (*Task, error)
	// Wechat is an interface for share articles to wechat groups and friends.
	Wechat(ctx context.Context, share WechatShare) error
}

// Response Serialize returns json response data.
type Response interface {
	MarshalJSON() ([]byte, error)
}

type Service interface {
	Get(ctx context.Context, source Source) (Tasks, string, error)
	List(ctx context.Context, source Source) (WechatShares, string, error)
	UnfinishedWechatShares(tasks Tasks) (WechatShares, error)
	Do(ctx context.Context, wc WechatShare, token string) error
}

type Assignment interface {
	Pub(source Source) error
	Close() error
}