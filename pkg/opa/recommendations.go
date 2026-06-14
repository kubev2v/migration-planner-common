package opa

// This file contains OS upgrade recommendation functionality.
// This is a short-term solution. The long-term solution should be implemented as part of ECOPROJECT-3571.
// See: https://issues.redhat.com/browse/ECOPROJECT-3571

import (
	"fmt"
	"strings"

	"github.com/kubev2v/migration-planner-common/pkg/duckdb_parser/models"
)

// OSUpgradeMap maps unsupported operating systems to their recommended upgrade targets
var OSUpgradeMap = map[string]string{
	"red hat enterprise linux 6": "Red Hat Enterprise Linux 7",
	"centos 7":                   "Red Hat Enterprise Linux 7",
	"centos 8":                   "Red Hat Enterprise Linux 8",
	"centos 9":                   "Red Hat Enterprise Linux 9",
	"amazon linux 2":             "Red Hat Enterprise Linux 8",
}

// GetOSUpgradeConcern returns an OS upgrade recommendation concern for models.VM if applicable.
// Returns nil if the OS is not in the upgrade map.
func GetOSUpgradeConcern(osName string) *models.Concern {
	osNameLower := strings.ToLower(strings.TrimSpace(osName))
	upgradeTarget, found := OSUpgradeMap[osNameLower]
	if !found {
		for k, v := range OSUpgradeMap {
			if strings.HasPrefix(osNameLower, k) {
				upgradeTarget = v
				break
			}
		}
	}

	if upgradeTarget == "" {
		return nil
	}

	return &models.Concern{
		Id:         "vmware.os.upgrade.recommendation",
		Category:   "Information",
		Label:      "OS Upgrade Recommendation",
		Assessment: fmt.Sprintf("The guest operating system: %s is not currently supported. The operating system can be upgraded to %s", osName, upgradeTarget),
	}
}
