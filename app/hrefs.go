// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package app

import (
	"fmt"
	"strings"
)

// ApplicationHref returns the resource href.
func ApplicationHref(projectid, appid interface{}) string {
	paramprojectid := strings.TrimLeftFunc(fmt.Sprintf("%v", projectid), func(r rune) bool { return r == '/' })
	paramappid := strings.TrimLeftFunc(fmt.Sprintf("%v", appid), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/projects/%v/applications/%v", paramprojectid, paramappid)
}

// ClusterHref returns the resource href.
func ClusterHref(projectid, resourceID interface{}) string {
	paramprojectid := strings.TrimLeftFunc(fmt.Sprintf("%v", projectid), func(r rune) bool { return r == '/' })
	paramresourceID := strings.TrimLeftFunc(fmt.Sprintf("%v", resourceID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/projects/%v/cluster/%v", paramprojectid, paramresourceID)
}

// NamespaceHref returns the resource href.
func NamespaceHref(projectid, namespaceid interface{}) string {
	paramprojectid := strings.TrimLeftFunc(fmt.Sprintf("%v", projectid), func(r rune) bool { return r == '/' })
	paramnamespaceid := strings.TrimLeftFunc(fmt.Sprintf("%v", namespaceid), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/projects/%v/namespaces/%v", paramprojectid, paramnamespaceid)
}

// ProjectHref returns the resource href.
func ProjectHref(projectid interface{}) string {
	paramprojectid := strings.TrimLeftFunc(fmt.Sprintf("%v", projectid), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/v1/projects/%v", paramprojectid)
}
