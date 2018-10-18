// To parse and unparse this JSON data, add this code to your project and do:
//
//    servers, err := UnmarshalServers(bytes)
//    bytes, err = servers.Marshal()

package models

import "encoding/json"

func UnmarshalServers(data []byte) (Servers, error) {
	var r Servers
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Servers) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Servers struct {
	Servers []ServerClass `json:"servers"`
}

func UnmarshalServer(data []byte) (Server, error) {
	var r Server
	err := json.Unmarshal(data, &r)
	return r, err
}

type Server struct {
	Server ServerClass `json:"server"`
}

type ServerClass struct {
	BackupWindow    *string                `json:"backup_window"`    // Time window (UTC) in which the backup will run, or null if the backups are not enabled
	Created         string                 `json:"created"`          // Point in time when the server was created (in ISO-8601 format)
	Datacenter      Datacenter             `json:"datacenter"`       // Datacenter this server is located at
	ID              float64                `json:"id"`               // ID of server
	Image           *Image                 `json:"image"`            // Image this server was created from.
	IncludedTraffic float64                `json:"included_traffic"` // Free Traffic for the current billing period in bytes
	IngoingTraffic  *float64               `json:"ingoing_traffic"`  // Inbound Traffic for the current billing period in bytes
	ISO             *ISO                   `json:"iso"`              // ISO image that is attached to this server. Null if no ISO is attached.
	Labels          map[string]interface{} `json:"labels"`           // User-defined labels (key-value pairs)
	Locked          bool                   `json:"locked"`           // True if server has been locked and is not available to user.
	Name            string                 `json:"name"`             // Name of the server (must be unique per project and a valid hostname as per RFC 1123)
	OutgoingTraffic *float64               `json:"outgoing_traffic"` // Outbound Traffic for the current billing period in bytes
	Protection      ServerProtection       `json:"protection"`       // Protection configuration for the server
	PublicNet       PublicNet              `json:"public_net"`       // Public network information. The servers ipv4 address can be found in; `public_net->ipv4->ip`
	RescueEnabled   bool                   `json:"rescue_enabled"`   // True if rescue mode is enabled: Server will then boot into rescue system on next reboot.
	ServerType      ServerType             `json:"server_type"`      // Type of server - determines how much ram, disk and cpu a server has
	Status          ServerStatus           `json:"status"`           // Status of the server
	Volumes         []interface{}          `json:"volumes"`          // IDs of Volumes assigned to this server.
}

// Datacenter this server is located at
type Datacenter struct {
	Description string      `json:"description"`  // Description of the datacenter
	ID          float64     `json:"id"`           // ID of the datacenter
	Location    Location    `json:"location"`     // Location where the datacenter resides in
	Name        string      `json:"name"`         // Unique identifier of the datacenter
	ServerTypes ServerTypes `json:"server_types"` // The server types the datacenter can handle
}

// Location where the datacenter resides in
type Location struct {
	City        string  `json:"city"`        // City the location is closest to
	Country     string  `json:"country"`     // ISO 3166-1 alpha-2 code of the country the location resides in
	Description string  `json:"description"` // Description of the location
	ID          float64 `json:"id"`          // ID of the location
	Latitude    float64 `json:"latitude"`    // Latitude of the city closest to the location
	Longitude   float64 `json:"longitude"`   // Longitude of the city closest to the location
	Name        string  `json:"name"`        // Unique identifier of the location
}

// The server types the datacenter can handle
type ServerTypes struct {
	Available []float64 `json:"available"` // IDs of server types that are supported and for which the datacenter has enough resources; left
	Supported []float64 `json:"supported"` // IDs of server types that are supported in the datacenter
}

type ISO struct {
	Deprecated  *string `json:"deprecated"`  // ISO 8601 timestamp of deprecation, null if ISO is still available. After the deprecation; time it will no longer be possible to attach the ISO to servers.
	Description string  `json:"description"` // Description of the ISO
	ID          float64 `json:"id"`          // ID of the ISO
	Name        *string `json:"name"`        // Unique identifier of the ISO. Only set for public ISOs
	Type        ISOType `json:"type"`        // Type of the ISO
}

type Image struct {
	BoundTo     *float64               `json:"bound_to"`               // ID of server the image is bound to. Only set for images of type `backup`.
	Created     string                 `json:"created"`                // Point in time when the image was created (in ISO-8601 format)
	CreatedFrom *CreatedFrom           `json:"created_from"`           // Information about the server the image was created from
	Deprecated  *string                `json:"deprecated"`             // Point in time when the image is considered to be deprecated (in ISO-8601 format)
	Description string                 `json:"description"`            // Description of the image
	DiskSize    float64                `json:"disk_size"`              // Size of the disk contained in the image in GB.
	ID          float64                `json:"id"`                     // ID of the image
	ImageSize   *float64               `json:"image_size"`             // Size of the image file in our storage in GB. For snapshot images this is the value; relevant for calculating costs for the image.
	Labels      map[string]interface{} `json:"labels"`                 // User-defined labels (key-value pairs)
	Name        *string                `json:"name"`                   // Unique identifier of the image. This value is only set for system images.
	OSFlavor    OSFlavor               `json:"os_flavor"`              // Flavor of operating system contained in the image
	OSVersion   *string                `json:"os_version"`             // Operating system version
	Protection  ImageProtection        `json:"protection"`             // Protection configuration for the image
	RapidDeploy *bool                  `json:"rapid_deploy,omitempty"` // Indicates that rapid deploy of the image is available
	Status      ImageStatus            `json:"status"`                 // Whether the image can be used or if it's still being created
	Type        ImageType              `json:"type"`                   // Type of the image
}

type CreatedFrom struct {
	ID   float64 `json:"id"`   // ID of the server the image was created from
	Name string  `json:"name"` // Server name at the time the image was created
}

// Protection configuration for the image
type ImageProtection struct {
	Delete bool `json:"delete"` // If true, prevents the snapshot from being deleted
}

// Protection configuration for the server
type ServerProtection struct {
	Delete  bool `json:"delete"`  // If true, prevents the server from being deleted
	Rebuild bool `json:"rebuild"` // If true, prevents the server from being rebuilt
}

// Public network information. The servers ipv4 address can be found in
// `public_net->ipv4->ip`
type PublicNet struct {
	FloatingIPS []float64 `json:"floating_ips"` // IDs of floating IPs assigned to this server.
	Ipv4        Ipv4      `json:"ipv4"`         // IP address (v4) and its reverse dns entry of this server.
	Ipv6        Ipv6      `json:"ipv6"`         // IPv6 network assigned to this server and its reverse dns entry.
}

// IP address (v4) and its reverse dns entry of this server.
type Ipv4 struct {
	Blocked bool   `json:"blocked"` // If the IP is blocked by our anti abuse dept
	DNSPtr  string `json:"dns_ptr"` // Reverse DNS PTR entry for the IPv4 addresses of this server.
	IP      string `json:"ip"`      // IP address (v4) of this server.
}

// IPv6 network assigned to this server and its reverse dns entry.
type Ipv6 struct {
	Blocked bool     `json:"blocked"` // If the IP is blocked by our anti abuse dept
	DNSPtr  []DNSPtr `json:"dns_ptr"` // Reverse DNS PTR entries for the IPv6 addresses of this server, `null` by default.
	IP      string   `json:"ip"`      // IP address (v4) of this server.
}

type DNSPtr struct {
	DNSPtr string `json:"dns_ptr"` // DNS pointer for the specific IP address.
	IP     string `json:"ip"`      // Single IPv6 address of this server for which the reverse DNS entry has been set up.
}

// Type of server - determines how much ram, disk and cpu a server has
type ServerType struct {
	Cores       float64     `json:"cores"`        // Number of cpu cores a server of this type will have
	CPUType     CPUType     `json:"cpu_type"`     // Type of cpu.
	Description string      `json:"description"`  // Description of the server type
	Disk        float64     `json:"disk"`         // Disk size a server of this type will have in GB
	ID          float64     `json:"id"`           // ID of the server type
	Memory      float64     `json:"memory"`       // Memory a server of this type will have in GB
	Name        string      `json:"name"`         // Unique identifier of the server type
	Prices      []Price     `json:"prices"`       // Prices in different Locations
	StorageType StorageType `json:"storage_type"` // Type of server boot drive. Local has higher speed. Network has better availability.
}

type Price struct {
	Location     string       `json:"location"`      // Name of the location the price is for
	PriceHourly  PriceHourly  `json:"price_hourly"`  // Hourly costs for a server type in this location
	PriceMonthly PriceMonthly `json:"price_monthly"` // Monthly costs for a server type in this location
}

// Hourly costs for a server type in this location
type PriceHourly struct {
	Gross string `json:"gross"` // Price with VAT added
	Net   string `json:"net"`   // Price without VAT
}

// Monthly costs for a server type in this location
type PriceMonthly struct {
	Gross string `json:"gross"` // Price with VAT added
	Net   string `json:"net"`   // Price without VAT
}

// Type of the ISO
type ISOType string

const (
	Private ISOType = "private"
	Public  ISOType = "public"
)

// Flavor of operating system contained in the image
type OSFlavor string

const (
	Centos          OSFlavor = "centos"
	Debian          OSFlavor = "debian"
	Fedora          OSFlavor = "fedora"
	OSFlavorUnknown OSFlavor = "unknown"
	Ubuntu          OSFlavor = "ubuntu"
	Unknown         OSFlavor = "unknown"
)

// Whether the image can be used or if it's still being created
type ImageStatus string

const (
	Available ImageStatus = "available"
	Creating  ImageStatus = "creating"
)

// Type of the image
type ImageType string

const (
	Backup   ImageType = "backup"
	Snapshot ImageType = "snapshot"
	System   ImageType = "system"
)

// Type of cpu.
type CPUType string

const (
	Dedicated CPUType = "dedicated"
	Shared    CPUType = "shared"
)

// Type of server boot drive. Local has higher speed. Network has better availability.
type StorageType string

const (
	Local   StorageType = "local"
	Network StorageType = "network"
)

// Status of the server
type ServerStatus string

const (
	ServerDeleting      ServerStatus = "deleting"
	ServerInitializing  ServerStatus = "initializing"
	ServerMigrating     ServerStatus = "migrating"
	ServerOff           ServerStatus = "off"
	ServerRebuilding    ServerStatus = "rebuilding"
	ServerRunning       ServerStatus = "running"
	ServerStarting      ServerStatus = "starting"
	ServerStatusUnknown ServerStatus = "unknown"
	ServerStopping      ServerStatus = "stopping"
	ServerFluffyRunning ServerStatus = "running"
)
