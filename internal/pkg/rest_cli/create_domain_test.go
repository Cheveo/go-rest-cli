package rest_cli

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/Cheveo/go-rest-cli/templates"
	"github.com/Cheveo/go-rest-cli/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var path = "path/to/dir"
var domain = "test"
var module = "test/go/mod"

////////////////////////////
// STD RELATED TEST DATA //
///////////////////////////

// map of corresponding std project filepaths --> checksums of files
var stdProjectChecksumMap = map[string]string{
	"path/to/dir/go.mod":                              "5c763f7370eeea4343e25d8d87f20a5cd8ab810a6877ebb31e75c09314528f8d",
	"path/to/dir/cmd/api/main.go":                     "1d6440ccc778cd78e94f4fe4b0cdaeb5e459ee6cf1040b25b6a653b62c06bb72",
	"path/to/dir/server/server.go":                    "4299486e09cb9cb40acbf81a8b60a930a5299739c6b34b26fa2f133f773239a1",
	"path/to/dir/utils/make_http_handler.go":          "c2acab997ce081648f65286da8df845a430c2aa4fb0e4bebbee75e95168cdfae",
	"path/to/dir/utils/write_json.go":                 "784dbab99894a1780e6c44c63d2ffeb70d5ce45986fb1e9279b21fa4a8624328",
	"path/to/dir/test/handler/handler.go":             "436d120abb2d7193db481fdf4e778e592c451d4dd15ff21e6552cd2cedfbdac1",
	"path/to/dir/test/handler/http_error_handler.go":  "30e5b1e17231dc7e8c828133c9bae1c639e945750ad4efa599d160e2bfe9c0fb",
	"path/to/dir/test/handler/test_handler.go":        "6cb07ed06c2cb84affc6442d82114b30e02540ed5d484d1c1544647bd6357a7b",
	"path/to/dir/test/models/test_models.go":          "3b7f8375042b6730d153dc32679038c84f98229d02ac897e5f2675a6a6ed99ca",
	"path/to/dir/test/service/service.go":             "31392e29d4e7f45eda192f2178b2a8ecb0b6329058b270dfd278a6934c070b3f",
	"path/to/dir/test/service/test_service.go":        "50af08e55e1ebff10f1fa92004e263921bc82a0b05f377c9f9811df6979ae33e",
	"path/to/dir/test/storage/storage.go":             "6ea17dfa361d13ce7e53a36c104b57d805174a3300c8f4dc4c1bf3345528df7e",
	"path/to/dir/test/storage/test_sql_statements.go": "31f89a2f74e66d77ee1f55d5365577282873ab3edad5f11d799ade40b2b2b6c7",
	"path/to/dir/test/storage/test_storage.go":        "add52456ec60c355adcdb7f2d74a3b52165a0a93a4527214296c7b85e790911c",
}

// map of corresponding std domain filepaths --> checksums of files
var stdDomainChecksumMap = map[string]string{
	"path/to/dir/test/handler/handler.go":             "436d120abb2d7193db481fdf4e778e592c451d4dd15ff21e6552cd2cedfbdac1",
	"path/to/dir/test/handler/http_error_handler.go":  "30e5b1e17231dc7e8c828133c9bae1c639e945750ad4efa599d160e2bfe9c0fb",
	"path/to/dir/test/handler/test_handler.go":        "6cb07ed06c2cb84affc6442d82114b30e02540ed5d484d1c1544647bd6357a7b",
	"path/to/dir/test/models/test_models.go":          "3b7f8375042b6730d153dc32679038c84f98229d02ac897e5f2675a6a6ed99ca",
	"path/to/dir/test/service/service.go":             "31392e29d4e7f45eda192f2178b2a8ecb0b6329058b270dfd278a6934c070b3f",
	"path/to/dir/test/service/test_service.go":        "50af08e55e1ebff10f1fa92004e263921bc82a0b05f377c9f9811df6979ae33e",
	"path/to/dir/test/storage/storage.go":             "6ea17dfa361d13ce7e53a36c104b57d805174a3300c8f4dc4c1bf3345528df7e",
	"path/to/dir/test/storage/test_sql_statements.go": "31f89a2f74e66d77ee1f55d5365577282873ab3edad5f11d799ade40b2b2b6c7",
	"path/to/dir/test/storage/test_storage.go":        "add52456ec60c355adcdb7f2d74a3b52165a0a93a4527214296c7b85e790911c",
}

// //////////////////////////
// GIN RELATED TEST DATA //
// /////////////////////////

// map of corresponding gin domain filepaths --> checksums of files
var ginProjectChecksumMap = map[string]string{
	"path/to/dir/cmd/api/main.go":              "b45223acf228008c3dac35506f5939a35da04ab184cdd2900b5c2ce2d6b82fd8",
	"path/to/dir/go.mod":                       "5c763f7370eeea4343e25d8d87f20a5cd8ab810a6877ebb31e75c09314528f8d",
	"path/to/dir/errors/http_error.go":         "3aff8bb087494d9b6f6631b8401a3dfb0f163e68a5df3538ab3245fa8ea54e17",
	"path/to/dir/responses/http_response.go":   "401bbb4b1538438cd7784c050c3e30332c1b5d7c9ea4b7b8f643060d612aafeb",
	"path/to/dir/db/db.go":                     "55f71d6028b3eb87bea0c1aa81f8ef5198abe31700aa835a44849f158f5421c8",
	"path/to/dir/db/gorm_db.go":                "683868aac53ef07bb32289a8297d34ecdb49a8b3c65d8f77e4cf409e955be6d3",
	"path/to/dir/middlewares/error_handler.go": "407ded92c80da7fca6e760dc3067b3b2c36584353db61306be2ba2d3ea7e2359",
	"path/to/dir/test/handler/handler.go":      "703c6679e68fcbe35808c1432a2891403300c1eb0823aaaa1e8be4039f21093f",
	"path/to/dir/test/handler/test_handler.go": "21daaf29fdc25ae4f56243cf11d2bdd6bda1d866e283eae5d9753d995229de70",
	"path/to/dir/test/models/test_models.go":   "4cb26cbf263faa2240730b1418a94c2e45f60838a096528d5c4777d0e3dc9fad",
	"path/to/dir/test/service/service.go":      "31392e29d4e7f45eda192f2178b2a8ecb0b6329058b270dfd278a6934c070b3f",
	"path/to/dir/test/service/test_service.go": "50af08e55e1ebff10f1fa92004e263921bc82a0b05f377c9f9811df6979ae33e",
	"path/to/dir/test/storage/storage.go":      "6ea17dfa361d13ce7e53a36c104b57d805174a3300c8f4dc4c1bf3345528df7e",
	"path/to/dir/test/storage/test_storage.go": "5cd33a37fdd2889874d93604fb8e232714ef390e481bb1a4d882aa8ae9284c53",
}
var ginDomainChecksumMap = map[string]string{
	"path/to/dir/test/handler/handler.go":      "703c6679e68fcbe35808c1432a2891403300c1eb0823aaaa1e8be4039f21093f",
	"path/to/dir/test/handler/test_handler.go": "21daaf29fdc25ae4f56243cf11d2bdd6bda1d866e283eae5d9753d995229de70",
	"path/to/dir/test/models/test_models.go":   "4cb26cbf263faa2240730b1418a94c2e45f60838a096528d5c4777d0e3dc9fad",
	"path/to/dir/test/service/service.go":      "31392e29d4e7f45eda192f2178b2a8ecb0b6329058b270dfd278a6934c070b3f",
	"path/to/dir/test/service/test_service.go": "50af08e55e1ebff10f1fa92004e263921bc82a0b05f377c9f9811df6979ae33e",
	"path/to/dir/test/storage/storage.go":      "6ea17dfa361d13ce7e53a36c104b57d805174a3300c8f4dc4c1bf3345528df7e",
	"path/to/dir/test/storage/test_storage.go": "5cd33a37fdd2889874d93604fb8e232714ef390e481bb1a4d882aa8ae9284c53",
}

type CreationTestSuite struct {
	suite.Suite
	path                  string
	stdProjectChecksumMap map[string]string
	stdDomainChecksumMap  map[string]string
	ginProjectChecksumMap map[string]string
	ginDomainChecksumMap  map[string]string
}

// before each test
func (suite *CreationTestSuite) SetupTest() {
	suite.path = path
	suite.stdProjectChecksumMap = stdProjectChecksumMap
	suite.stdDomainChecksumMap = stdDomainChecksumMap
	suite.ginProjectChecksumMap = ginProjectChecksumMap
	suite.ginDomainChecksumMap = ginDomainChecksumMap
	err := templates.LoadTemplates()
	if err != nil {
		suite.T().Error(err.Error())
	}

}

func TestCreationTestSuite(t *testing.T) {
	suite.Run(t, new(CreationTestSuite))
}

func (suite *CreationTestSuite) TestCreateStdProject() {
	// initialization must be done without the new call,
	// since it creates a directory with prefixed homedir
	d := &types.DomainTmpl{
		Domain:            domain,
		Directory:         path,
		CapitalizedDomain: "Test",
		GoMod:             module,
		Type:              types.StandardProject,
	}

	test(suite.T(), types.StandardProject, suite.stdProjectChecksumMap, path, d)
}

func (suite *CreationTestSuite) TestCreateStdDomain() {
	d := &types.DomainTmpl{
		Domain:            domain,
		Directory:         path,
		CapitalizedDomain: "Test",
		GoMod:             module,
		Type:              types.StandardDomain,
	}

	test(suite.T(), types.StandardDomain, suite.stdDomainChecksumMap, path, d)
}

func (suite *CreationTestSuite) TestCreateGinProject() {
	d := &types.DomainTmpl{
		Domain:            domain,
		Directory:         path,
		CapitalizedDomain: "Test",
		GoMod:             module,
		Type:              types.GinProject,
	}

	test(suite.T(), types.GinProject, suite.ginProjectChecksumMap, path, d)
}
func (suite *CreationTestSuite) TestCreateGinDomain() {
	d := &types.DomainTmpl{
		Domain:            domain,
		Directory:         path,
		CapitalizedDomain: "Test",
		GoMod:             module,
		Type:              types.GinDomain,
	}

	test(suite.T(), types.GinDomain, suite.ginDomainChecksumMap, path, d)
}

func test(t *testing.T, projectType types.ProjectType, checksums map[string]string, path string, d *types.DomainTmpl) {
	domain := ProjectTypeFactory(d)
	err := domain.Create()
	if err != nil {
		t.Error(err.Error())
	}

	files, err := getFiles(4, d.Directory)
	if err != nil {
		t.Error(err.Error())
	}

	// extract filenames from the checksum keys and
	// match them with the found files from the directory
	filepaths := getFilePathsFromChecksumMap(checksums)
	ok := assert.ElementsMatch(t, files, filepaths)
	if !ok {
		t.Error("Created files do not match.")
	}

	for _, file := range files {
		compareFileChecksums(t, file, checksums)
	}

	t.Cleanup(func() {
		err := os.RemoveAll("path")
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func getFiles(depth uint8, path string) ([]string, error) {
	type void struct{}
	var v void

	files := make(map[string]void)
	var paths []string
	for i := uint8(0); i < depth; i++ {
		initialPath := ""
		for j := uint8(0); j < i; j++ {
			initialPath += "/*"
		}
		initialPath += "/*.go"

		f, err := filepath.Glob(path + initialPath)
		if err != nil {
			return nil, err
		}

		for _, file := range f {
			files[file] = v
		}
	}

	f, err := filepath.Glob(path + "/go.mod")
	if err != nil {
		return nil, err
	}

	if len(f) > 0 {
		files[f[0]] = v
	}

	for k := range files {
		paths = append(paths, k)
	}

	return paths, nil
}

func getFilePathsFromChecksumMap(checksums map[string]string) []string {
	var filepaths []string
	for p := range checksums {
		filepaths = append(filepaths, p)
	}

	return filepaths
}
func compareFileChecksums(t *testing.T, file string, checksums map[string]string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		t.Errorf("error while creating checksum.\nFile: %s\nFileName: %s", checksums[file], f.Name())
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))

	if hash != checksums[file] {
		t.Errorf("checksums are different.\nShould be: %s\nis actually: %s\nfilename: %s", checksums[file], hash, f.Name())
	}
}
