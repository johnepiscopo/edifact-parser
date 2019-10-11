package edisegments

func NewUNBModel() *EdiSegment {
	m := new(EdiSegment)

	m.Name = "UNB"
	m.Elements = []*EdiElement{
		&EdiElement{
			Name: "SyntaxIdentifierElement",
			Tag:  "S001",
			Ref:  "UNB010",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "SyntaxIdentifier",
					Ref:  "UNH010-010",
					Tag:  "0001",
				},
				&EdiElement{
					Name: "SyntaxVersionNumber",
					Ref:  "UNH010-020",
					Tag:  "0002",
				},
				&EdiElement{
					Name: "ServiceCodeList",
					Ref:  "UNH010-030",
					Tag:  "0080",
				},
				&EdiElement{
					Name: "CharacterEncoding",
					Ref:  "UNH010-040",
					Tag:  "0133",
				},
			},
		},
		&EdiElement{
			Name: "InterchangeSender",
			Ref:  "UNB020",
			Tag:  "S002",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "InterchangeSenderIdentifier",
					Ref:  "UNH020-010",
					Tag:  "0004",
				},
				&EdiElement{
					Name: "IdentificationCodeQualifier",
					Ref:  "UNH020-030",
					Tag:  "0007",
				},
				&EdiElement{
					Name: "InterchangeRecipientInternalId",
					Ref:  "UNH020-040",
					Tag:  "0014",
				},
				&EdiElement{
					Name: "InterchangeRecipientInternalSubId",
					Ref:  "UNH020-050",
					Tag:  "0046",
				},
				&EdiElement{
					Name: "AssociationAssignedCode",
					Ref:  "UNH020-050",
					Tag:  "0057",
				},
			},
		},
		&EdiElement{
			Name: "DateTimePreparation",
			Ref:  "UNB040",
			Tag:  "S004",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "Date",
					Ref:  "UNB040-010",
					Tag:  "0017",
				},
				&EdiElement{
					Name: "Time",
					Ref:  "UNB040-020",
					Tag:  "0019",
				},
			},
		},
		&EdiElement{
			Name: "InterchangeControlReference",
			Ref:  "UNB050",
			Tag:  "0020",
		},
		&EdiElement{
			Name: "RecipientReferenceDetails",
			Ref:  "UNB060",
			Tag:  "S004",
			SubElements: []*EdiElement{
				&EdiElement{
					Name: "RecipientReference",
					Ref:  "UNB060-010",
					Tag:  "0022",
				},
				&EdiElement{
					Name: "RecipientReferenceQualifier",
					Ref:  "UNB060-020",
					Tag:  "0025",
				},
			},
		},
		&EdiElement{
			Name: "ApplicationReference",
			Ref:  "UNB070",
			Tag:  "0026",
		},
		&EdiElement{
			Name: "ProcessingPriorityCode",
			Ref:  "UNB080",
			Tag:  "0029",
		},
		&EdiElement{
			Name: "AcknowledgementRequest",
			Ref:  "UNB090",
			Tag:  "0031",
		},
		&EdiElement{
			Name: "InterchangeAgreementId",
			Ref:  "UNB100",
			Tag:  "0032",
		},
		&EdiElement{
			Name: "TestIndicator",
			Ref:  "UNB110",
			Tag:  "0035",
		},
	}

	return m
}
