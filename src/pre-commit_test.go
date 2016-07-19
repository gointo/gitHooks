package gitHooks

import (
  "flag"
  "os"
  "testing"
)

func TestGetheadSuccess(t *testing.T) {
  os.Setenv(`HEAD`, `8b7cf9f5f3d2f465ba5b3a593601fc32ab2b675a`)
  res, err := Gethead()
  t.Logf(`Result: "%s"`, res)
  t.Logf(`Error : "%s"`, err)
  if (res != `8b7cf9f5f3d2f465ba5b3a593601fc32ab2b675a` || err != nil) {
    t.Error(`Error with HEAD`)
  }
}

func TestGetheadFail(t *testing.T) {
  os.Unsetenv(`HEAD`)
  res, err := Gethead()
  t.Logf(`Result: "%s"`, res)
  t.Logf(`Error : "%s"`, err)
  if (res != `4b825dc642cb6eb9a060e54bf8d69288fbee4904` || err == nil) {
    t.Errorf(`Error with HEAD: %s`, res)
  }
}

func TestGetasciisettings(t *testing.T) {
  res := Getasciisettings()
  t.Logf(`Result: "%s"`, res)
}

func TestMain(m *testing.M) {
  flag.Parse()
  os.Exit(m.Run())
}
