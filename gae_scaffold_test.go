package gaescaffold

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"google.golang.org/appengine/datastore"

	"net/http"

	"github.com/favclip/testerator"
	"github.com/k0kubun/pp"
	"github.com/mjibson/goon"
)

func TestMain(m *testing.M) {
	_, _, err := testerator.SpinUp()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	status := m.Run()

	err = testerator.SpinDown()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	os.Exit(status)

}

func TestHandler(t *testing.T) {

	instance, ctx, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed testerator.SpinUp(): %v", err)
	}
	defer testerator.SpinDown()
	_ = ctx
	req, err := instance.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("code want %d insted %d\n", http.StatusOK, rec.Code)
	}

}

type ScafflodSample struct {
	ID   int64 `datastore:"-" goon:"id"`
	Name string
}

func TestGoonPut(t *testing.T) {

	putData := []ScafflodSample{
		// {ID: 1, Name: "name 001"},
		// {ID: 2, Name: "name 002"},
		// {ID: 3, Name: "name 003"},
		{Name: "name 001"},
		{Name: "name 002"},
		{Name: "name 003"},
	}
	instance, ctx, err := testerator.SpinUp()
	if err != nil {
		t.Fatalf("Failed testerator.SpinUp(): %v", err)
	}
	defer testerator.SpinDown()
	_ = instance

	g := goon.FromContext(ctx)
	keys, err := g.PutMulti(putData)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	_ = keys
	dst := []ScafflodSample{}
	_, err = g.GetAll(
		datastore.NewQuery(goon.DefaultKindName(ScafflodSample{})),
		&dst,
	)
	// err = g.GetMulti(&dst)
	if err != nil {
		t.Errorf("%#v", err)
	}

	if len(dst) != len(putData) {
		t.Errorf("length invalid! want %d but insted %d ", len(putData), len(dst))
	}
	pp.Printf("\n%v\n", dst)

}
