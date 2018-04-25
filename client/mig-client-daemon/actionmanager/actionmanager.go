// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Contributor: Zack Mullaly zmullaly@mozilla.com [:zack]

package actionmanager

import (
	//"strings"
	"time"

	"mig.ninja/mig"
	//migclient "mig.ninja/mig/client"
	"mig.ninja/mig/client/mig-client-daemon/ident"
	"mig.ninja/mig/client/mig-client-daemon/modules"
	"mig.ninja/mig/client/mig-client-daemon/targeting"
)

// ActionCatalog maintains information about actions that have been
// created, dispatched, etc.
type ActionCatalog struct {
	actions map[ident.Identifier]mig.Action
}

// NewActionCatalog creates a new `ActionCatalog`.
func NewActionCatalog() ActionCatalog {
	return ActionCatalog{
		actions: make(map[ident.Identifier]mig.Action),
	}
}

// CreateAction attempts to create a new action with the supplied information
// and register it to the catalog, returning the identifier of the
// newly-created action as a string.
func (catalog *ActionCatalog) CreateAction(
	module modules.Module,
	agentTargetSpecifiers []targeting.Query,
	expireAfter time.Duration,
) (ident.Identifier, error) {
	id := ident.GenerateUniqueID(3, 250*time.Millisecond, func(id ident.Identifier) bool {
		_, alreadyTaken := catalog.actions[id]
		return !alreadyTaken
	})

	queryStrings := []string{}
	for _, query := range agentTargetSpecifiers {
		whereClause, err := query.ToSQLWhereClause()
		if err != nil {
			return "", err
		}
		queryStrings = append(queryStrings, whereClause)
	}
	// target := strings.Join(queryStrings, " AND ")

	return id, nil
}
