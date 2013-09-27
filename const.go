/*
 * const.go: Go bindings for lxc
 *
 * Copyright © 2013, S.Çağlar Onur
 *
 * Authors:
 * S.Çağlar Onur <caglar@10ur.org>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.

 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.

 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
 */

package lxc

const (
	// WaitForever timeout
	WaitForever int = iota - 1
	// DontWait timeout
	DontWait
)

const (
	// CloneKeepName means don't edit the rootfs to change the hostname.
	CloneKeepName int = 1 << iota
	// CloneCopyHooks means copy all hooks into the container directory.
	CloneCopyHooks
	// CloneKeepMACAddr means don't change the mac address on network interfaces.
	CloneKeepMACAddr
	// CloneSnapshot means snapshot the original filesystem(s).
	CloneSnapshot
)

const (
	errAlreadyDefined                = "container %q already defined"
	errAlreadyFrozen                 = "container %q is already frozen"
	errAlreadyRunning                = "container %q is already running"
	errAttachFailed                  = "attaching to the container %q failed"
	errClearingCgroupItemFailed      = "clearing cgroup item for the container %q failed (key: %s)"
	errCloneFailed                   = "cloning the container %q failed"
	errCloseAllFdsFailed             = "setting close_all_fds flag for container %q failed"
	errCreateFailed                  = "creating the container %q failed"
	errDaemonizeFailed               = "setting daemonize flag for container %q failed"
	errDestroyFailed                 = "destroying the container %q failed"
	errFreezeFailed                  = "freezing the container %q failed"
	errInsufficientNumberOfArguments = "insufficient number of arguments were supplied"
	errInterfaces                    = "getting interface names for the container %q failed"
	errIPAddresses                   = "getting IP addresses of the container %q failed"
	errIPAddress                     = "getting IP address on the interface %s of the container %q failed"
	errIPv4Addresses                 = "getting IPv4 addresses of the container %q failed"
	errIPv6Addresses                 = "getting IPv6 addresses of the container %q failed"
	errLoadConfigFailed              = "loading config file for the container %q failed (path: %s)"
	errNotDefined                    = "there is no container named %q"
	errNotFrozen                     = "container %q is not frozen"
	errNotRunning                    = "container %q is not running"
	errRebootFailed                  = "rebooting the container %q failed"
	errSaveConfigFailed              = "saving config file for the container %q failed (path: %s)"
	errSettingCgroupItemFailed       = "setting cgroup item for the container %q failed (key: %s, value: %s)"
	errSettingConfigItemFailed       = "setting config item for the container %q failed (key: %s, value: %s)"
	errSettingConfigPathFailed       = "setting config file for the container %q failed (path: %s)"
	errShutdownFailed                = "shutting down the container %q failed"
	errStartFailed                   = "starting the container %q failed"
	errStopFailed                    = "stopping the container %q failed"
	errUnfreezeFailed                = "unfreezing the container %q failed"
)
