//go:generate go run ../../generate/tags/main.go -ListTagsInIDElem=ResourceId -ServiceTagsSlice -TagOp=AddTags -TagInIDElem=ResourceId -UntagOp=RemoveTags -UpdateTags
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package emr
