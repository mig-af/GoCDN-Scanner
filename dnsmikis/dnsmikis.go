package dnsmikis

import "gatoscanner/domain"






type Scan interface{
	
	ServiceName()string
	CheckSubdomain() (domain.SubDomains, error)
}