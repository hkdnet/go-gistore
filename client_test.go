package gistore

import "testing"

func TestAuthorize(t *testing.T) {
	c := NewClient("")
	if c.token != "" {
		t.Error("Initial token must be empty")
	}
	c.Authorize("foo")
	if c.token != "foo" {
		t.Errorf("want: %v got: %v", "foo", c.token)
	}
}

func TestSelectAll(t *testing.T) {
	c := newClientWithCache()
	ret := []testUser{}
	err := c.SelectAll("users", &ret)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}
	if len(ret) != 2 {
		t.Errorf("want: %v got: %v", 2, len(ret))
	}
	mami := ret[0]
	if mami.Name != "Mami" {
		t.Errorf("want: %v got: %v", "Mami", mami.Name)
	}
	takane := ret[1]
	if takane.Name != "Takane" {
		t.Errorf("want: %v got: %v", "Takane", takane.Name)
	}
}
func TestSelectAllError(t *testing.T) {
	c := newClientWithCache()
	ret := []testUser{}
	err := c.SelectAll("users", ret)
	if err == nil {
		t.Fatal("Expect to fail but succeeded")
	} else if msg := err.Error(); msg != nonPointerError {
		t.Errorf("want: %v got: %v", nonPointerError, msg)
	}
	var i int32
	err = c.SelectAll("users", &i)
	if err == nil {
		t.Fatal("Expect to fail but succeeded")
	} else if msg := err.Error(); msg != nonPointerOfSliceError {
		t.Errorf("want: %v got: %v", nonPointerOfSliceError, msg)
	}
}

func newClientWithCache() *Client {
	c := NewClient("")
	files := make(map[string]GistFile)
	files["users"] = GistFile{Content: `{"id": 1, "name": "Mami"}
 {"id": 2, "name": "Takane"}`}
	c.cache = &Gist{
		Files: files,
	}
	return c
}

type testUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
