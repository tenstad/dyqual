package matchers

import (
	"bytes"
	"fmt"

	"github.com/gonvenience/ytbx"
	"github.com/homeport/dyff/pkg/dyff"
	yamlv3 "gopkg.in/yaml.v3"
)

func ymlNodes(input interface{}) ([]*yamlv3.Node, error) {
	yml, err := yamlv3.Marshal(input)
	if err != nil {
		return nil, err
	}

	docs, err := ytbx.LoadYAMLDocuments(yml)
	if err != nil {
		return nil, fmt.Errorf("failed to parse as YAML: %s", err)
	}

	return docs[0].Content, nil
}

func compare(expected interface{}, actual interface{}) (string, error) {
	expectedNodes, err := ymlNodes(expected)
	if err != nil {
		return "", err
	}
	actualNodes, err := ymlNodes(actual)
	if err != nil {
		return "", err
	}

	report, err := dyff.CompareInputFiles(
		ytbx.InputFile{Documents: expectedNodes},
		ytbx.InputFile{Documents: actualNodes},
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
