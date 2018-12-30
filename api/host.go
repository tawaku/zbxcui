package api

type Host struct {
	Hostid            string `json:"hostid"`
	Host              string `json:"host"`
	Available         byte   `json:"available,string"`
	Description       string `json:"description"`
	DisableUntil      int32  `json:"disable_until,string"`
	Error             string `json:"error"`
	ErrorsFrom        int32  `json:"errors_from,string"`
	Flags             byte   `json:"flags,string"`
	InventoryMode     int8   `json:"inventory_mode,string"`
	IpmiAuthtype      int8   `json:"ipmi_authtype,string"`
	IpmiAvailable     byte   `json:"ipmi_available,string"`
	IpmiDisableUntil  int32  `json:"ipmi_disable_until,string"`
	IpmiError         string `json:"ipmi_error"`
	IpmiErrorsFrom    int32  `json:"ipmi_errors_from,string"`
	IpmiPassword      string `json:"ipmi_password"`
	IpmiPrivilege     byte   `json:"ipmi_privilege,string"`
	IpmiUsername      string `json:"ipmi_username"`
	JmxAvailable      byte   `json:"jmx_available,string"`
	JmxDisableUntil   int32  `json:"jmx_disable_until,string"`
	JmxError          string `json:"jmx_error"`
	JmxErrorsFrom     int32  `json:"jmx_errors_from,string"`
	MaintenanceFrom   int32  `json:"maintenance_from,string"`
	MaintenanceStatus byte   `json:"maintenance_status,string"`
	MaintenanceType   byte   `json:"maintenance_type,string"`
	Maintenanceid     string `json:"maintenanceid"`
	Name              string `json:"name"`
	ProxyHostid       string `json:"proxy_hostid"`
	SnmpAvailable     byte   `json:"snmp_available,string"`
	SnmpDisableUntil  int32  `json:"snmp_disable_until,string"`
	SnmpError         string `json:"snmp_error"`
	SnmpErrorsFrom    int32  `json:"snmp_errors_from,string"`
	Status            byte   `json:"status,string"`
	TlsConnect        byte   `json:"tls_connect,string"`
	TlsAccept         byte   `json:"tls_accept,string"`
	TlsIssuer         string `json:"tls_issuer"`
	TlsSubject        string `json:"tls_subject"`
	TlsPskIdentity    string `json:"tls_psk_identity"`
	TlsPsk            string `json:"tls_psk"`
}

func (self *Host) AvailableValue() string {
	switch self.Available {
	case 0:
		return "unknown"
	case 1:
		return "available"
	case 2:
		return "unavailable"
	default:
		return "undefined"
	}
}

func (self *Host) FlagsValue() string {
	switch self.Flags {
	case 0:
		return "a plain host"
	case 1:
		return "a discovered host"
	default:
		return "undefined"
	}
}

func (self *Host) InventoryModeValue() string {
	switch self.InventoryMode {
	case -1:
		return "disabled"
	case 0:
		return "manual"
	case 1:
		return "automatic"
	default:
		return "undefined"
	}
}

func (self *Host) IpmiAuthtypeValue() string {
	switch self.IpmiAuthtype {
	case -1:
		return "disabled"
	case 0:
		return "none"
	case 1:
		return "MD2"
	case 2:
		return "MD5"
	case 4:
		return "straight"
	case 5:
		return "OEM"
	case 6:
		return "RMCP+"
	default:
		return "undefined"
	}
}

func (self *Host) IpmiAvailableValue() string {
	switch self.IpmiAvailable {
	case 0:
		return "unknown"
	case 1:
		return "available"
	case 2:
		return "unavaiable"
	default:
		return "undefined"
	}
}

func (self *Host) IpmiPrivilegeValue() string {
	switch self.IpmiPrivilege {
	case 1:
		return "callback"
	case 2:
		return "user"
	case 3:
		return "operator"
	case 4:
		return "admin"
	case 5:
		return "OEM"
	default:
		return "undefined"
	}
}

func (self *Host) JmxAvailableValue() string {
	switch self.JmxAvailable {
	case 0:
		return "unknown"
	case 1:
		return "available"
	case 2:
		return "unavaiable"
	default:
		return "undefined"
	}
}

func (self *Host) MaintenanceStatusValue() string {
	switch self.MaintenanceStatus {
	case 0:
		return "no maintenance"
	case 1:
		return "maintenance in effect"
	default:
		return "undefined"
	}
}

func (self *Host) MaintenanceTypeValue() string {
	switch self.MaintenanceType {
	case 0:
		return "maintenance with data collection"
	case 1:
		return "maintenance without data collection"
	default:
		return "undefined"
	}
}

func (self *Host) SnmpAvailableValue() string {
	switch self.SnmpAvailable {
	case 0:
		return "unknown"
	case 1:
		return "available"
	case 2:
		return "unavaiable"
	default:
		return "undefined"
	}
}

func (self *Host) StatusValue() string {
	switch self.Status {
	case 0:
		return "monitored host"
	case 1:
		return "unmonitored host"
	default:
		return "undefined"
	}
}

func (self *Host) TlsConnectValue() string {
	switch self.TlsConnect {
	case 1:
		return "No encryption"
	case 2:
		return "PSK"
	case 4:
		return "certificate"
	default:
		return "undefined"
	}
}

func (self *Host) TlsAcceptValue() string {
	switch self.TlsAccept {
	case 1:
		return "No encryption"
	case 2:
		return "PSK"
	case 4:
		return "certificate"
	default:
		return "undefined"
	}
}
