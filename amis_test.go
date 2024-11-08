package amisgo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/zrcoder/amisgo/config"
)

const (
	testFileContent = "test content"
	testFileName    = "test.txt"
	testOKResponse  = "ok"
)

// Helper functions for common test operations
func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status %d, got %d", want, got)
	}
}

func assertBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected body %q, got %q", want, got)
	}
}

func assertContains(t *testing.T, body, substr string) {
	t.Helper()
	if !strings.Contains(body, substr) {
		t.Errorf("expected body to contain %q", substr)
	}
}

func makeRequest(handler http.Handler, method, path string) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, path, nil)
	handler.ServeHTTP(w, r)
	return w, r
}

func createTestFile(t *testing.T, dir string) {
	t.Helper()
	err := os.WriteFile(dir+"/"+testFileName, []byte(testFileContent), 0o644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNew(t *testing.T) {
	e := New()
	if e == nil {
		t.Error("New() should return a non-nil Engine")
	}
}

func TestMount(t *testing.T) {
	e := New()
	page := struct {
		Title string `json:"title"`
		Type  string `json:"type"`
	}{
		Title: "Test Page",
		Type:  "page",
	}
	e.Mount("/page", page)

	t.Run("base path", func(t *testing.T) {
		w, _ := makeRequest(e, "GET", "/page")
		assertStatus(t, w.Code, http.StatusOK)

		body := w.Body.String()
		assertContains(t, body, "<!DOCTYPE html>")

		expectedJSON, _ := json.Marshal(page)
		assertContains(t, body, string(expectedJSON))
	})

	t.Run("sub path", func(t *testing.T) {
		w, _ := makeRequest(e, "GET", "/page/subpath")
		assertStatus(t, w.Code, http.StatusNotFound)
	})
}

func TestRedirect(t *testing.T) {
	e := New()
	e.Redirect("/old", "/new")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/old", nil)
	e.ServeHTTP(w, r)

	if w.Code != http.StatusTemporaryRedirect {
		t.Errorf("expected status 307, got %d", w.Code)
	}

	location := w.Header().Get("Location")
	if location != "/new" {
		t.Errorf("expected redirect to /new, got %s", location)
	}
}

func TestHandle(t *testing.T) {
	e := New()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testOKResponse))
	})
	e.Handle("/api/handle", handler)

	w, _ := makeRequest(e, "GET", "/api/handle")
	assertStatus(t, w.Code, http.StatusOK)
	assertBody(t, w.Body.String(), testOKResponse)
}

func TestHandleFunc(t *testing.T) {
	e := New()
	e.HandleFunc("/api/handlefunc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(testOKResponse))
	})

	w, _ := makeRequest(e, "GET", "/api/handlefunc")
	assertStatus(t, w.Code, http.StatusOK)
	assertBody(t, w.Body.String(), testOKResponse)
}

func TestStaticFiles(t *testing.T) {
	tmpDir := t.TempDir()
	createTestFile(t, tmpDir)

	e := New(config.WithStaticFS("/static/", http.Dir(tmpDir)))

	t.Run("serve static file", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/static/"+testFileName, nil)
		e.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		content, _ := io.ReadAll(w.Body)
		if string(content) != testFileContent {
			t.Error("unexpected file content")
		}
	})

	t.Run("404 for non-existent file", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/static/nonexistent.txt", nil)
		e.ServeHTTP(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", w.Code)
		}
	})
}

func TestAsHandler(t *testing.T) {
	tmpDir := t.TempDir()
	createTestFile(t, tmpDir)

	e := New(config.WithStaticFS("/admin/static/", http.Dir(tmpDir)))
	e.Mount("/admin/page", struct{}{})

	// Create main mux and mount engine
	mux := http.NewServeMux()
	mux.Handle("/admin/", e)

	t.Run("serve mounted component", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/admin/page", nil)
		mux.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}
	})

	t.Run("serve static file", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/admin/static/"+testFileName, nil)
		mux.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		content, _ := io.ReadAll(w.Body)
		if string(content) != testFileContent {
			t.Error("unexpected file content")
		}
	})
}
