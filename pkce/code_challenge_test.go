package pkce

import "testing"

func TestCodeChallenge(t *testing.T) {
	got := GenerateCodeChallenge("ks02i3jdikdo2k0dkfodf3m39rjfjsdk0wk349rj3jrhf")
	want := "2i0WFA-0AerkjQm4X4oDEhqA17QIAKNjXpagHBXmO_U"

	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}

// 2i0WFA-0AerkjQm4X4oDEhqA17QIAKNjXpagHBXmO_U
// 2i0WFA+0AerkjQm4X4oDEhqA17QIAKNjXpagHBXmO/U=
