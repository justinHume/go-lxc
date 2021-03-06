/*
 * list.go
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

	"github.com/lxc/go-lxc"
)

var (
	lxcpath string
)

func init() {
	flag.StringVar(&lxcpath, "lxcpath", lxc.DefaultConfigPath(), "Use specified container path")
	flag.Parse()
}

func main() {
	log.Printf("Defined containers:\n")
	for _, v := range lxc.DefinedContainers(lxcpath) {
		log.Printf("%s (%s)\n", v.Name(), v.State())
	}

	log.Println()

	log.Printf("Active containers:\n")
	for _, v := range lxc.ActiveContainers(lxcpath) {
		log.Printf("%s (%s)\n", v.Name(), v.State())
	}

	log.Println()

	log.Printf("Active and Defined containers:\n")
	for _, v := range lxc.ActiveContainers(lxcpath) {
		log.Printf("%s (%s)\n", v.Name(), v.State())
	}
}
