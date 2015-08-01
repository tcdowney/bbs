package test_helpers

import (
	"path"

	"github.com/cloudfoundry-incubator/bbs/models"

	. "github.com/onsi/gomega"
)

func (t *TestHelper) RegisterCell(cell models.CellPresence) {
	jsonBytes, err := models.ToJSON(cell)
	Expect(err).NotTo(HaveOccurred())

	err = t.consulSession.AcquireLock(CellSchemaPath(cell.CellID), jsonBytes)
	Expect(err).NotTo(HaveOccurred())
}

const (
	LockSchemaRoot = "v1/locks"
	CellSchemaRoot = LockSchemaRoot + "/cell"
)

func CellSchemaPath(cellID string) string {
	return path.Join(CellSchemaRoot, cellID)
}
