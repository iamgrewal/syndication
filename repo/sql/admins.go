/*
  Copyright (C) 2017 Jorge Martinez Hernandez

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// Package repo provides routines to operate on Syndications SQL repo
// using models defined in the models package to map data in said repo.
package sql

import (
	"github.com/jmartinezhern/syndication/models"
	"github.com/jmartinezhern/syndication/repo"
)

type (
	Admins struct {
		db *DB
	}
)

// NewAdmin returns an instance of an SQL implementation of repo.Admins
func NewAdmins(db *DB) Admins {
	return Admins{
		db,
	}
}

// Create a new admin
func (u Admins) Create(admin *models.Admin) {
	u.db.db.Create(admin)
}

// Update an admin
func (u Admins) Update(id string, admin *models.Admin) error {
	dbAdmin, found := u.AdminWithID(id)
	if !found {
		return repo.ErrModelNotFound
	}

	u.db.db.Model(&dbAdmin).Updates(admin).RecordNotFound()
	return nil
}

// AdminWithID returns an admin with id
func (u Admins) AdminWithID(id string) (admin models.Admin, found bool) {
	found = !u.db.db.First(&admin, "api_id = ?", id).RecordNotFound()
	return
}

func (u Admins) InitialUser() (admin models.Admin, found bool) {
	found = !u.db.db.First(&admin, "id = ?", 1).RecordNotFound()
	return
}

// Delete admin with id
func (u Admins) Delete(id string) error {
	admin, found := u.AdminWithID(id)
	if !found {
		return repo.ErrModelNotFound
	}

	u.db.db.Delete(&admin)
	return nil
}

// AdminWithName returns a admin with username
func (u Admins) AdminWithName(name string) (admin models.Admin, found bool) {
	found = !u.db.db.First(&admin, "username = ?", name).RecordNotFound()
	return
}

// OwnsKey returns true if the given APIKey is owned by admin
func (u Admins) OwnsKey(key *models.APIKey, admin *models.Admin) bool {
	return !u.db.db.Model(admin).Where("key = ?", key.Key).Related(key).RecordNotFound()
}

// AddAPIKey associates an API key with admin
func (u Admins) AddAPIKey(key *models.APIKey, admin *models.Admin) {
	u.db.db.Model(admin).Association("APIKeys").Append(key)
}