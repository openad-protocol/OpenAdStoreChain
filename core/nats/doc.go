// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Package nats hold the implementation of the Publisher and PubSub
// interfaces for the NATS messaging system, the internal messaging
// broker of the Mainflux IoT platform. Due to the practical requirements
// implementation Publisher is created alongside PubSub. The reason for
// this is that AdNatsConnection implementation of NATS brings the burden of
// additional struct fields which are not used by Publisher. AdNatsConnection
// is not implemented separately because PubSub can be used where AdNatsConnection is needed.
package nats
