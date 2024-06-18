package model

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	html := LwcHtml{
		LwcFile: LwcFile{
			Name: "helloWorld.html",
		},
		Depends: nil,
	}
	js := LwcJs{
		LwcFile: LwcFile{
			Name: "helloWorld.js",
		},
	}
	meta := LwcMeta{
		LwcFile: LwcFile{
			Name: "helloWorld.js-meta.xml",
		},
	}
	comp := LwcComp{
		Name: "helloWorld",
		Dir:  "helloWorld",
	}
	comp.SetHtml(&html).SetJs(&js).SetMeta(&meta)

	fmt.Printf("html = %T\n", html)
	fmt.Printf("html = %v\n", html.Name)
	if "helloWorld.html" != html.Name {
		t.Errorf("Expected helloWorld.html, got %v", html.Name)
	}
	fmt.Printf("comp = %v\n", comp)
	fmt.Printf("html address = %p\n", &html)
	fmt.Printf("js address = %p\n", &js)
	fmt.Printf("meta address = %p\n", &meta)
}
