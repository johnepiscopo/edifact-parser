package edisegments

type TDTModel struct {
	elements   []string
	separators map[string]string
}

func NewTDTModel() *TDTModel {
	m := new(TDTModel)
	m.elements = []string{
		"Carrier",                       //Ref: TDT050 	 Tag: C040
		"CarrierIdentification",         //Ref: TDT050-010 	 Tag: 3127
		"CodeListIdentificationCode",    //Ref: TDT050-020 Tag: 1131
		"CodeListResponsibleAgencyCode", //Ref: TDT050-030 Tag: 3055
		"CarrierName",                   //Ref: TDT050-040 Tag: 3128
	}

	m.separators = map[string]string{
		//This has 5 separators as there are a number of elements
		//missing http://www.unece.org/trade/untdid/d05b/trsd/trsdtdt.htm
		"Carrier":                       "+++++",
		"CarrierIdentification":         ":",
		"CodeListIdentificationCode":    ":",
		"CodeListResponsibleAgencyCode": ":",
		"CarrierName":                   "",
	}

	return m
}

func (u *TDTModel) GetElements() []string {
	return u.elements
}

func (u *TDTModel) GetSeparator(element string) string {
	return u.separators[element]
}
