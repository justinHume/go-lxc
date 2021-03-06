/*
 * create.go
 *
 * Copyright © 2013, 2014, S.Çağlar Onur
 *
 * Authors:
 * S.Çağlar Onur <caglar@10ur.org>
 *
 * This library is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 2, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package main

import (
	"flag"
	"log"
	"os"

	"github.com/lxc/go-lxc"
)

var (
	lxcpath  string
	template string
	distro   string
	release  string
	arch     string
	name     string
)

func init() {
	flag.StringVar(&lxcpath, "lxcpath", lxc.DefaultConfigPath(), "Use specified container path")
	flag.StringVar(&template, "template", "busybox", "Template to use")
	flag.StringVar(&distro, "distro", "ubuntu", "Template to use")
	flag.StringVar(&release, "release", "trusty", "Template to use")
	flag.StringVar(&arch, "arch", "amd64", "Template to use")
	flag.StringVar(&name, "name", "rubik", "Name of the container")
	flag.Parse()
}

func main() {
	c, err := lxc.NewContainer(name, lxcpath)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
	defer lxc.PutContainer(c)

	log.Printf("Creating container...\n")
	c.SetVerbosity(lxc.Verbose)

	if os.Geteuid() != 0 {
		if err := c.CreateAsUser(distro, release, arch); err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		}
	} else {
		if err := c.Create(template, "-a", arch, "-r", release); err != nil {
			log.Printf("ERROR: %s\n", err.Error())
		}
	}
}
