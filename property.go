package swag

import (
	"fmt"
	"go/ast"
)

// getPropertyName returns the string value for the given field if it exists, otherwise it panics.
// allowedValues: array, boolean, integer, null, number, object, string
func getPropertyName(field *ast.Field) (name string, fieldType string) {
	return getTypeName(field.Type)

}

func getTypeName(expr ast.Expr) (name string, fieldType string) {
	if astTypeSelectorExpr, ok := expr.(*ast.SelectorExpr); ok {

		// Support for time.Time as a structure field
		if "Time" == astTypeSelectorExpr.Sel.Name {
			return "string", "string"
		}

		// Support bson.ObjectId type
		if "ObjectId" == astTypeSelectorExpr.Sel.Name {
			return "string", "string"
		}

		panic("not supported 'astSelectorExpr' yet.")

	} else if astTypeIdent, ok := expr.(*ast.Ident); ok {
		name = astTypeIdent.Name

		// When its the int type will transfer to integer which is goswagger supported type
		schemeType := TransToValidSchemeType(name)
		return schemeType, schemeType

	} else if astTypeStar, ok := expr.(*ast.StarExpr); ok {
		return getTypeName(astTypeStar.X)
		//panic("not supported astStarExpr yet.")
	} else if _, ok := expr.(*ast.MapType); ok { // if map
		//TODO: support map
		return "object", "object"
	} else if astTypeArray, ok := expr.(*ast.ArrayType); ok { // if array
		str := fmt.Sprintf("%s", astTypeArray.Elt)
		return "array", str
	} else if _, ok := expr.(*ast.StructType); ok { // if struct
		return "object", "object"
	} else if _, ok := expr.(*ast.InterfaceType); ok { // if interface{}
		return "object", "object"
	}

	return name, fieldType
}
