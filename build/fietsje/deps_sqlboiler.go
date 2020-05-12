// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func depsSQLBoiler(p *planner) {
	p.collect(
		"github.com/volatiletech/sqlboiler/v4", "v4.1.1",
	).use(
		"github.com/denisenkom/go-mssqldb",
		"github.com/ericlagergren/decimal",
		"github.com/friendsofgo/errors",
		"github.com/go-sql-driver/mysql",
		"github.com/golang-sql/civil",
		"github.com/hashicorp/hcl",
		"github.com/lib/pq",
		"github.com/magiconair/properties",
		"github.com/spf13/cast",
		"github.com/spf13/cobra",
		"github.com/spf13/jwalterweatherman",
		"github.com/spf13/viper",
		"github.com/subosito/gotenv",
		"github.com/volatiletech/inflect",
		"github.com/volatiletech/null/v8",
		"github.com/volatiletech/randomize",
		"github.com/volatiletech/strmangle",
		"gopkg.in/ini.v1",
	)
	// required by //build/sqlboiler autogeneration
	p.collect(
		"github.com/glerchundi/sqlboiler-crdb/v4", "d540ee52783ebbbfe010acc5d91a9043d88de3fd",
	).use(
		"github.com/gofrs/uuid",
	)
	p.collect(
		"github.com/rubenv/sql-migrate", "ae26b214fa431c314a5a9b986d5c90fb1719c68d",
	).use(
		"github.com/armon/go-radix",
		"github.com/mattn/go-sqlite3",
		"github.com/mitchellh/cli",
		"github.com/posener/complete",
		"github.com/joho/godotenv",
		"gopkg.in/gorp.v1",
	)
}
