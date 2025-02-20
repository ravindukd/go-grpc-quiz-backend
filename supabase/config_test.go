package supabase

import (
	"testing"
	"github.com/supabase/supabase-go"
	"os"
)

func TestInitializeSupabase(t *testing.T) {
	os.Setenv("SUPABASE_URL", "your_supabase_url")
	os.Setenv("SUPABASE_KEY", "your_supabase_key")

	client := InitializeSupabase()

	if client == nil {
		t.Errorf("Expected Supabase client to be initialized, but got nil")
	}

	if client.ApiKey != "your_supabase_key" {
		t.Errorf("Expected Supabase client to have the correct API key, but got %s", client.ApiKey)
	}

	if client.BaseUrl != "your_supabase_url" {
		t.Errorf("Expected Supabase client to have the correct base URL, but got %s", client.BaseUrl)
	}
}
