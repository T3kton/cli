/*Package contractor - Automatically generated by cinp-codegen from /api/v1/Directory/ at 2020-02-27T14:18:37.597397
 */
package contractor

import (
	"reflect"
	"time"
	cinp "github.com/cinp/go"
)

//DirectoryZone - Model Zone(/api/v1/Directory/Zone)
/*
Zone(name, parent, ttl, refresh, retry, expire, minimum, updated, created)
 */
type DirectoryZone struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Name string `json:"name"`
	Parent string `json:"parent"`
	TTL int `json:"ttl"`
	Refresh int `json:"refresh"`
	Retry int `json:"retry"`
	Expire int `json:"expire"`
	Minimum int `json:"minimum"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
	Fqdn string `json:"fqdn"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *DirectoryZone) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"name": object.Name,
			"parent": object.Parent,
			"ttl": object.TTL,
			"refresh": object.Refresh,
			"retry": object.Retry,
			"expire": object.Expire,
			"minimum": object.Minimum,
		}
	}
	return &map[string]interface{}{ 
		"parent": object.Parent,
		"ttl": object.TTL,
		"refresh": object.Refresh,
		"retry": object.Retry,
		"expire": object.Expire,
		"minimum": object.Minimum,
	}
}

// DirectoryZoneNew - Make a new object of Model Zone
func (service *Contractor) DirectoryZoneNew() *DirectoryZone {
	return &DirectoryZone{cinp: service.cinp}
}

// DirectoryZoneNewWithID - Make a new object of Model Zone
func (service *Contractor) DirectoryZoneNewWithID(id string) *DirectoryZone {
	result := DirectoryZone{cinp: service.cinp}
	result.SetID("/api/v1/Directory/Zone:"+id+":")
	return &result
}

// DirectoryZoneGet - Get function for Model Zone
func (service *Contractor) DirectoryZoneGet(id string) (*DirectoryZone, error) {
	object, err := service.cinp.Get("/api/v1/Directory/Zone:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DirectoryZone)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Zone
func (object *DirectoryZone) Create() error {
	if err := object.cinp.Create("/api/v1/Directory/Zone", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Zone
func (object *DirectoryZone) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Zone
func (object *DirectoryZone) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// DirectoryZoneList - List function for Model Zone
func (service *Contractor) DirectoryZoneList(filterName string, filterValues map[string]interface{}) <-chan *DirectoryZone {
	in := service.cinp.ListObjects("/api/v1/Directory/Zone", reflect.TypeOf(DirectoryZone{}), filterName, filterValues, 50)
	out := make(chan *DirectoryZone)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*DirectoryZone)
		}
	}()
	return out
}


//DirectoryEntry - Model Entry(/api/v1/Directory/Entry)
/*
Entry(id, zone, type, name, priority, weight, port, target, updated, created)
 */
type DirectoryEntry struct {
	cinp.BaseObject
	cinp *cinp.CInP
	Zone string `json:"zone"`
	Type string `json:"type"`
	Name string `json:"name"`
	Priority int `json:"priority"`
	Weight int `json:"weight"`
	Port int `json:"port"`
	Target string `json:"target"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

// AsMap returns a map[string]interface{} that is required for create and update
func (object *DirectoryEntry) AsMap(isCreate bool) *map[string]interface{} {
	if isCreate {
		return &map[string]interface{}{ 
			"zone": object.Zone,
			"type": object.Type,
			"name": object.Name,
			"priority": object.Priority,
			"weight": object.Weight,
			"port": object.Port,
			"target": object.Target,
		}
	}
	return &map[string]interface{}{ 
		"zone": object.Zone,
		"type": object.Type,
		"name": object.Name,
		"priority": object.Priority,
		"weight": object.Weight,
		"port": object.Port,
		"target": object.Target,
	}
}

// DirectoryEntryNew - Make a new object of Model Entry
func (service *Contractor) DirectoryEntryNew() *DirectoryEntry {
	return &DirectoryEntry{cinp: service.cinp}
}

// DirectoryEntryNewWithID - Make a new object of Model Entry
func (service *Contractor) DirectoryEntryNewWithID(id string) *DirectoryEntry {
	result := DirectoryEntry{cinp: service.cinp}
	result.SetID("/api/v1/Directory/Entry:"+id+":")
	return &result
}

// DirectoryEntryGet - Get function for Model Entry
func (service *Contractor) DirectoryEntryGet(id string) (*DirectoryEntry, error) {
	object, err := service.cinp.Get("/api/v1/Directory/Entry:"+id+":")
	if err != nil {
		return nil, err
	}
	result := (*object).(*DirectoryEntry)
	result.cinp = service.cinp

	return result, nil
}

// Create - Create function for Model Entry
func (object *DirectoryEntry) Create() error {
	if err := object.cinp.Create("/api/v1/Directory/Entry", object); err != nil {
		return err
	}

	return nil
}

// Update - Update function for Model Entry
func (object *DirectoryEntry) Update(fieldList []string) error {
	if err := object.cinp.Update(object, fieldList); err != nil {
		return err
	}

	return nil
}

// Delete - Delete function for Model Entry
func (object *DirectoryEntry) Delete() error {
	if err := object.cinp.Delete(object); err != nil {
		return err
	}

	return nil
}

// DirectoryEntryList - List function for Model Entry
func (service *Contractor) DirectoryEntryList(filterName string, filterValues map[string]interface{}) <-chan *DirectoryEntry {
	in := service.cinp.ListObjects("/api/v1/Directory/Entry", reflect.TypeOf(DirectoryEntry{}), filterName, filterValues, 50)
	out := make(chan *DirectoryEntry)
	go func() {
		defer close(out)
		for v := range in {
			out <- v.(*DirectoryEntry)
		}
	}()
	return out
}

func registerDirectory(cinp *cinp.CInP) { 
	cinp.RegisterType("/api/v1/Directory/Zone", reflect.TypeOf((*DirectoryZone)(nil)).Elem())
	cinp.RegisterType("/api/v1/Directory/Entry", reflect.TypeOf((*DirectoryEntry)(nil)).Elem())
}