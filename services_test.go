package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSetDroneCIService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/services/drone-ci", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	opt := &SetDroneCIServiceOptions{String("t"), String("u"), Bool(true)}
	_, err := client.Services.SetDroneCIService(1, opt)

	if err != nil {
		t.Fatalf("Services.SetDroneCIService returns an error: %v", err)
	}
}

func TestDeleteDroneCIService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/services/drone-ci", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Services.DeleteDroneCIService(1)

	if err != nil {
		t.Fatalf("Services.DeleteDroneCIService returns an error: %v", err)
	}
}

func TestGetDroneCIService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/services/drone-ci", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &DroneCIService{Service: Service{ID: Int(1)}}

	service, _, err := client.Services.GetDroneCIService(1)

	if err != nil {
		t.Fatalf("Services.GetDroneCIService returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetDroneCIService returned %+v, want %+v", service, want)
	}
}

func TestSetSlackService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/services/slack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	opt := &SetSlackServiceOptions{
		WebHook:  String("webhook_uri"),
		Username: String("username"),
		Channel:  String("#development"),
	}
	_, err := client.Services.SetSlackService(1, opt)

	if err != nil {
		t.Fatalf("Services.SetSlackService returns an error: %v", err)
	}
}

func TestDeleteSlackService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/services/slack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Services.DeleteSlackService(1)

	if err != nil {
		t.Fatalf("Services.DeleteSlackService returns an error: %v", err)
	}
}

func TestGetSlackService(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/1/services/slack", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})
	want := &SetSlackServiceOptions{}

	service, _, err := client.Services.GetSlackService(1)

	if err != nil {
		t.Fatalf("Services.GetSlackService returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, service) {
		t.Errorf("Services.GetSlackService returned %+v, want %+v", service, want)
	}
}
