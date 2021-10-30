package matchers

import (
	"bytes"
	"fmt"

	"github.com/gonvenience/ytbx"
	"github.com/homeport/dyff/pkg/dyff"
	yamlv3 "gopkg.in/yaml.v3"
)

func ymlNodes(input string) ([]*yamlv3.Node, error) {
	docs, err := ytbx.LoadYAMLDocuments([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("failed to parse as YAML: %s", err)
	}

	return docs[0].Content, nil
}

func compare(expected interface{}, actual interface{}) (string, error) {
	expYML, err := yamlv3.Marshal(expected)
	if err != nil {
		return "", err
	}
	actYML, err := yamlv3.Marshal(actual)
	if err != nil {
		return "", err
	}

	expNodes, err := ymlNodes(string(expYML))
	if err != nil {
		return "", err
	}
	actNodes, err := ymlNodes(string(actYML))
	if err != nil {
		return "", err
	}

	report, err := dyff.CompareInputFiles(
		ytbx.InputFile{Documents: expNodes},
		ytbx.InputFile{Documents: actNodes},
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
