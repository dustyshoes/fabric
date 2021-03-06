/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/pkg/errors"

// This error is returned when trying to join or remove an application channel when the system channel exists.
var ErrSystemChannelExists = errors.New("system channel exists")

// This error is returned when trying to join a app channel that already exists (when the system channel does not
// exist), or when trying to join the system channel when it already exists.
var ErrChannelAlreadyExists = errors.New("channel already exists")

// This error is returned when trying to join a system channel (that does not exist) when application channels
// already exist.
var ErrAppChannelsAlreadyExists = errors.New("application channels already exist")

// This error is returned when trying to remove a channel that does not exist
var ErrChannelNotExist = errors.New("channel does not exist")
