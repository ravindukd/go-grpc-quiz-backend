package supabase

import (
	"github.com/supabase/supabase-go"
	"os"
)

func InitializeSupabase() *supabase.Client {
	client := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	return client
}
