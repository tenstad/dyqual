package dyqual

import (
	"bytes"
	"fmt"

	"github.com/gonvenience/ytbx"
	"github.com/homeport/dyff/pkg/dyff"
	yamlv3 "gopkg.in/yaml.v3"
)

func yml(input string) (*yamlv3.Node, error) {
	docs, err := ytbx.LoadYAMLDocuments([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as YAML: %s", err)
	}
	if len(docs) > 1 {
		return nil, fmt.Errorf("failed to use YAML, because it contains multiple documents")
	}

	return docs[0].Content[0], nil
}

func compare(expected string, actual string) (string, error) {
	expYML, err := yml(expected)
	if err != nil {
		return "", err
	}
	actYML, err := yml(actual)
	if err != nil {
		return "", err
	}

	report, err := dyff.CompareInputFiles(
		ytbx.InputFile{Documents: []*yamlv3.Node{expYML}},
		ytbx.InputFile{Documents: []*yamlv3.Node{actYML}},
	)
	if err != nil {
		return "", err
	}
	humanReport := dyff.HumanReport{
		Report:     report,
		OmitHeader: true,
	}

	buf := bytes.NewBuffer([]byte{})
	if err := humanReport.WriteReport(buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
