package model

type LwcComp struct {
	Name     string
	Dir      string
	CssFile  *LwcCss
	HtmlFile *LwcHtml
	JsFile   *LwcJs
	MetaFile *LwcMeta
}

func (lwcComp *LwcComp) SetHtml(htmlFile *LwcHtml) *LwcComp {
	lwcComp.HtmlFile = htmlFile
	htmlFile.Parent = lwcComp
	return lwcComp
}

func (lwcComp *LwcComp) SetJs(jsFile *LwcJs) *LwcComp {
	lwcComp.JsFile = jsFile
	jsFile.Parent = lwcComp
	return lwcComp
}

func (lwcComp *LwcComp) SetMeta(metaFile *LwcMeta) *LwcComp {
	lwcComp.MetaFile = metaFile
	metaFile.Parent = lwcComp
	return lwcComp
}

func (lwcComp *LwcComp) SetCss(cssFile *LwcCss) *LwcComp {
	lwcComp.CssFile = cssFile
	cssFile.Parent = lwcComp
	return lwcComp
}

func (lwcComp *LwcComp) String() string {
	ret := "{" +
		"\n  Name: " + lwcComp.Name +
		"\n  Dir: " + lwcComp.Dir
	if lwcComp.HtmlFile != nil {
		ret += getFileInfo(lwcComp.HtmlFile, "\n  Html", "  ")
	}
	if lwcComp.JsFile != nil {
		ret += getFileInfo(lwcComp.JsFile, "\n  Js", "  ")
	}
	if lwcComp.MetaFile != nil {
		ret += getFileInfo(lwcComp.MetaFile, "\n  Meta", "  ")
	}
	if lwcComp.CssFile != nil {
		ret += getFileInfo(lwcComp.CssFile, "\n  Css", "  ")
	}
	ret += "\n}"
	return ret
}

type LwcFile struct {
	Name   string
	Parent *LwcComp
}

type stringer interface {
	stringInfo(linePrefix string) string
}

type LwcCss struct {
	LwcFile
	Depends []*LwcComp
}

func (lwcCss *LwcCss) stringInfo(linePrefix string) string {
	ret := "{" +
		"\n" + linePrefix + "  Name: " + lwcCss.Name
	if lwcCss.Depends != nil {
		deps := dependenciesList(lwcCss.Depends)
		ret += "\n" + linePrefix + "  Depends: " + deps
	}
	ret += "\n" + linePrefix + "}"
	return ret
}

func getFileInfo(str stringer, key string, linePrefix string) string {
	if str == nil {
		return ""
	}
	return key + ": " + str.stringInfo(linePrefix)
}

type LwcHtml struct {
	LwcFile
	Depends []*LwcComp
}

func (lwcHtml *LwcHtml) stringInfo(linePrefix string) string {
	ret := "{" +
		"\n" + linePrefix + "  Name: " + lwcHtml.Name
	if lwcHtml.Depends != nil {
		deps := dependenciesList(lwcHtml.Depends)
		ret += "\n" + linePrefix + "  Depends: " + deps
	}
	ret += "\n" + linePrefix + "}"
	return ret
}

type LwcJs struct {
	LwcFile
	Depends []*LwcComp
}

func (lwcJs *LwcJs) stringInfo(linePrefix string) string {
	ret := "{" +
		"\n" + linePrefix + "  Name: " + lwcJs.Name
	if lwcJs.Depends != nil {
		deps := dependenciesList(lwcJs.Depends)
		ret += "\n" + linePrefix + "  Depends: " + deps
	}
	ret += "\n" + linePrefix + "}"
	return ret
}

type LwcMeta struct {
	LwcFile
	Exposed string
}

func (lwcMeta *LwcMeta) stringInfo(linePrefix string) string {
	ret := "{" +
		"\n" + linePrefix + "  Name: " + lwcMeta.Name +
		"\n" + linePrefix + "  Exposed: " + lwcMeta.Exposed
	ret += "\n" + linePrefix + "}"
	return ret
}

func dependenciesList(deps []*LwcComp) string {
	if deps == nil {
		return ""
	}
	ret := "["
	for idx, dep := range deps {
		if idx > 0 {
			ret += ", "
		}
		ret += dep.Name
	}
	return ret + "]"
}
