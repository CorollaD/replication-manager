// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017-2021 SIGNAL18 CLOUD SAS
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.
// Redistribution/Reuse of this code is permitted under the GNU v3 license, as
// an additional term, ALL code must carry the original Author(s) credit in comment form.
// See LICENSE in this directory for the integral text.
package cluster

import (
	"os"
	"strings"
)

func (proxy *Proxy) IsFilterInTags(filter string) bool {
	tags := proxy.ClusterGroup.GetProxyTags()
	for _, tag := range tags {
		if strings.Contains(filter, "."+tag) {
			//	fmt.Println(server.ClusterGroup.Conf.ProvTags + " vs tag: " + tag + "  against " + filter)
			return true
		}
	}
	return false
}

func (proxy *Proxy) HasProvisionCookie() bool {
	if _, err := os.Stat(proxy.Datadir + "/@cookie_prov"); os.IsNotExist(err) {
		return false
	}
	return true
}

func (proxy *Proxy) HasWaitStartCookie() bool {
	if _, err := os.Stat(proxy.Datadir + "/@cookie_waitstart"); os.IsNotExist(err) {
		return false
	}
	return true
}

func (proxy *Proxy) HasWaitStopCookie() bool {
	if _, err := os.Stat(proxy.Datadir + "/@cookie_waitstop"); os.IsNotExist(err) {
		return false
	}
	return true
}

func (proxy *Proxy) HasRestartCookie() bool {
	if _, err := os.Stat(proxy.Datadir + "/@cookie_restart"); os.IsNotExist(err) {
		return false
	}
	return true
}

func (proxy *Proxy) HasReprovCookie() bool {
	if _, err := os.Stat(proxy.Datadir + "/@cookie_reprov"); os.IsNotExist(err) {
		return false
	}
	return true
}

func (proxy *Proxy) IsRunning() bool {
	return !proxy.IsDown()
}

func (proxy *Proxy) IsDown() bool {
	if proxy.State == stateFailed || proxy.State == stateSuspect || proxy.State == stateErrorAuth {
		return true
	}
	return false
}
