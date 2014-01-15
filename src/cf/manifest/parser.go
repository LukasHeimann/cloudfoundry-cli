package manifest

import (
	"generic"
	"github.com/cloudfoundry/gamble"
	"io"
	"io/ioutil"
)

func Parse(reader io.Reader) (yamlMap generic.Map, err error) {
	yamlBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	document, err := gamble.Parse(string(yamlBytes))
	if err != nil {
		return
	}

	yamlMap = generic.NewMap(document)
	return
}

func ParseToManifest(reader io.Reader) (m *Manifest, errs ManifestErrors) {
	mapp, err := Parse(reader)
	if err != nil {
		errs = append(errs, err)
		return
	}

	return NewManifest(mapp)
}
