package matching

import "testing"

func TestMatch(t *testing.T) {
	sample := []struct {
		src   string
		match bool
	}{
		{"{ac[bb]}", true},
		{"[dklf(df(kl))d]{}", true},
		{"{[[[]]]}", true},
		{"{3234[fd", false},
	}

	for _, s := range sample {
		if matchParens(s.src, '{', '}', '[', ']') != s.match {
			t.Errorf("expected %v got %v for %s", s.match, !s.match, s.src)
		}
	}

}
