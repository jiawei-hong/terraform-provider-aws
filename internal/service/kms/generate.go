//go:generate go run ../../generate/tags/main.go -ListTags -ListTagsOp=ListResourceTags -ListTagsInIDElem=KeyId -ServiceTagsSlice -TagInIDElem=KeyId -TagTypeKeyElem=TagKey -TagTypeValElem=TagValue -UpdateTags
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package kms
