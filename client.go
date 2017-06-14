package gistore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Client struct {
	baseURL string
	gistID  string
	token   string

	cache *Gist

	Transport *http.Transport
}

func NewClient(gistID string) *Client {
	return &Client{
		baseURL: "https://api.github.com/gists/",
		gistID:  gistID,
	}
}

// Authorize allows to edit the gist files.
func (c *Client) Authorize(token string) {
	c.token = token
}

func (c *Client) getGistURL() string {
	return c.baseURL + c.gistID
}

func (c *Client) getGist() (ret *Gist, err error) {
	var httpClient http.Client
	if c.Transport != nil {
		httpClient = http.Client{Transport: c.Transport}
	} else {
		httpClient = http.Client{}
	}
	url := c.getGistURL()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ret)
	if err != nil {
		c.cache = ret
	}
	return
}

func (c *Client) SelectAll(name string, output interface{}) (err error) {
	var gist *Gist
	if c.cache == nil {
		gist, err = c.getGist()
		if err != nil {
			return
		}
	} else {
		gist = c.cache
	}
	rv := reflect.ValueOf(output)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("SelectAll: second argument must be a pointer")
	}
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Invalid {
		return fmt.Errorf("SelectAll: nil pointer dereference")
	}
	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("SelectAll: ")
	}
	t := rv.Type().Elem()
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("SelectAll: output must be slice of struct, but %v", rv.Type())
	}
	f, ok := gist.Files[name]
	if !ok {
		return fmt.Errorf("No such name: %s", name)
	}
	lines := strings.Split(f.Content, "\n")
	value, err := linesToSlice(lines, rv.Type())
	if err != nil {
		return err
	}
	rv.Set(value)
	return nil
}

func linesToSlice(lines []string, t reflect.Type) (reflect.Value, error) {
	t = t.Elem()
	slice := reflect.MakeSlice(reflect.SliceOf(t), len(lines), len(lines))
	for i, v := range lines {
		val := reflect.New(t).Interface()
		err := json.Unmarshal([]byte(v), &val)
		if err != nil {
			return reflect.Value{}, err
		}
		slice.Index(i).Set(reflect.Indirect(reflect.ValueOf(val)))
	}
	return slice, nil
}
