//go:generate go run ../../generate/tags/main.go  -AWSSDKVersion=2 -TagInIDElem=ResourceArn -ListTags -ListTagsInIDElem=ResourceArn -ServiceTagsMap -UpdateTags -KVTValues -SkipTypesImp
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package ssmincidents
