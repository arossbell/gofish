// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// FailureDomainScope shall represent a geographic scope in which all components
// within that scope have similar vulnerabilities.
type FailureDomainScope string

const (
	// ServerFailureDomainScope are components of a CPU/memory complex that share
	// the same infrastructure.
	ServerFailureDomainScope FailureDomainScope = "Server"
	// RackFailureDomainScope is a container within a datacenter that provides
	// communication, power, or cooling to a set of components.
	RackFailureDomainScope FailureDomainScope = "Rack"
	// RackGroupFailureDomainScope is a set of racks that may share common
	// communication, power, or cooling.
	RackGroupFailureDomainScope FailureDomainScope = "RackGroup"
	// RowFailureDomainScope is a set of adjacent racks or rackgroups that may
	// share common communication, power, or cooling.
	RowFailureDomainScope FailureDomainScope = "Row"
	// DatacenterFailureDomainScope is a facility that provides communication,
	// power, or cooling infrastructure to a co-located set of servers,
	// networking and storage.
	DatacenterFailureDomainScope FailureDomainScope = "Datacenter"
	// RegionFailureDomainScope is a set of resources that are required to be
	// either geographically or politically isolated from resources not in
	// the resources.
	RegionFailureDomainScope FailureDomainScope = "Region"
)

// RecoveryAccessScope shall represent the relative time required to make a
// replica available as a source.
type RecoveryAccessScope string

const (
	// OnlineActiveRecoveryAccessScope shall be instantaneous.
	OnlineActiveRecoveryAccessScope RecoveryAccessScope = "OnlineActive"
	// OnlinePassiveRecoveryAccessScope shall be consistent with switching
	// access to a different path the same front-end interconnect. A restore
	// step shall not be required.
	OnlinePassiveRecoveryAccessScope RecoveryAccessScope = "OnlinePassive"
	// NearlineRecoveryAccessScope shall be consistent with switching access
	// to a different path through a different front-end interconnection
	// infrastructure. Some inconsistency may occur. A restore step may be
	// required before recovery can commence.
	NearlineRecoveryAccessScope RecoveryAccessScope = "Nearline"
	// OfflineRecoveryAccessScope Access to a replica may take a significant
	// amount of time. No direct connection to the replica is assumed. Some
	// inconsistency loss may occur. A restore step is likely to be
	// required.
	OfflineRecoveryAccessScope RecoveryAccessScope = "Offline"
)

// DataProtectionLoSCapabilities is the capabilities to protect data from
// loss by the use of a replica. The requirements shall be met
// collectively by the communication path and the replica. There should
// be one instance associated to a class of service for each replica.
// Each replica independently should have a class of service that
// describes its characteristics.
type DataProtectionLoSCapabilities struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// SupportedLinesOfService collection shall contain known and
	// supported DataProtectionLinesOfService.
	SupportedLinesOfService string
	// SupportedLinesOfService@odata.count is
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// SupportedMinLifetimes each entry shall be an ISO 8601
	// duration that specifies the minimum lifetime required for the replica.
	SupportedMinLifetimes string
	// SupportedRecoveryGeographicObjectives each entry shall
	// specify a supported failure domain.
	SupportedRecoveryGeographicObjectives string
	// SupportedRecoveryPointObjectiveTimes each entry shall
	// specify a supported ISO 8601 time interval defining the maximum source
	// information that may be lost on failure. In the case that IsIsolated =
	// false, failure of the domain is not a consideration.
	SupportedRecoveryPointObjectiveTimes string
	// SupportedRecoveryTimeObjectives each entry shall
	// specify an enumerated value that indicates a supported expectation for
	// the time required to access an alternate replica. In the case that
	// IsIsolated = false, failure of the domain is not a consideration.
	SupportedRecoveryTimeObjectives string
	// SupportedReplicaTypes each entry shall specify a
	// supported replica type
	SupportedReplicaTypes string
	// SupportsIsolated is A value of true shall indicate that allocating a
	// replica in a separate fault domain is supported. The default value for
	// this property is false.
	SupportsIsolated bool
	// SupportedReplicaOptionsCount is the number of supported replica options.
	SupportedReplicaOptionsCount int
	// supportedReplicaOptions shall contain known and
	// supported replica Classes of Service.
	supportedReplicaOptions []string
}

// UnmarshalJSON unmarshals a DataProtectionLoSCapabilities object from the raw JSON.
func (dataprotectionloscapabilities *DataProtectionLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp DataProtectionLoSCapabilities

	// DLinks is the links to related entities for this class.
	type DLinks struct {
		// SupportedReplicaOptions shall contain known and
		// supported replica Classes of Service.
		SupportedReplicaOptions common.Links
		// SupportedReplicaOptionsCount is the number of supported replica options.
		SupportedReplicaOptionsCount int `json:"SupportedReplicaOptions@odata.count"`
	}
	var t struct {
		temp
		Links DLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*dataprotectionloscapabilities = DataProtectionLoSCapabilities(t.temp)

	// Extract the links to other entities for later
	dataprotectionloscapabilities.supportedReplicaOptions = t.Links.SupportedReplicaOptions.ToStrings()

	return nil
}

// GetDataProtectionLoSCapabilities will get a DataProtectionLoSCapabilities instance from the service.
func GetDataProtectionLoSCapabilities(c common.Client, uri string) (*DataProtectionLoSCapabilities, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var dataprotectionloscapabilities DataProtectionLoSCapabilities
	err = json.NewDecoder(resp.Body).Decode(&dataprotectionloscapabilities)
	if err != nil {
		return nil, err
	}

	dataprotectionloscapabilities.SetClient(c)
	return &dataprotectionloscapabilities, nil
}

// ListReferencedDataProtectionLoSCapabilities gets the collection of DataProtectionLoSCapabilities from
// a provided reference.
func ListReferencedDataProtectionLoSCapabilities(c common.Client, link string) ([]*DataProtectionLoSCapabilities, error) {
	var result []*DataProtectionLoSCapabilities
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, dataprotectionloscapabilitiesLink := range links.ItemLinks {
		dataprotectionloscapabilities, err := GetDataProtectionLoSCapabilities(c, dataprotectionloscapabilitiesLink)
		if err != nil {
			return result, err
		}
		result = append(result, dataprotectionloscapabilities)
	}

	return result, nil
}

// SupportedReplicaOptions gets the support replica ClassesOfService.
func (dataprotectionloscapabilities *DataProtectionLoSCapabilities) SupportedReplicaOptions() ([]*ClassesOfService, error) {
	var result []*ClassesOfService

	for _, link := range dataprotectionloscapabilities.supportedReplicaOptions {
		classOfService, err := GetClassesOfService(dataprotectionloscapabilities.Client, link)
		if err != nil {
			return result, err
		}
		result = append(result, classOfService)
	}

	return result, nil
}