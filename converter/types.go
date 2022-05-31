package converter

// ConverterType interface defines an interface type constrain
type ConverterType interface {
	HexType | DecType | OctType | BinType
}

// NewConverterType create a new generic ConverterType object
func NewConverterType[T ConverterType]() T {
	return T{
		Base: BaseType{},
	}
}

// BaseType hold base data
type BaseType struct {
	Number   uint64
	Value    string
	LeadZero bool
}

// HexType hold data for hexadecimal types
type HexType struct {
	Base BaseType
}

// DecType hold data for decimal types
type DecType struct {
	Base BaseType
}

// OctType hold data for octal types
type OctType struct {
	Base BaseType
}

// BinType hold data for binary types
type BinType struct {
	Base BaseType
}

//func NewHexType() HexType {
//	return HexType{
//		Base: BaseType{},
//	}
//}
//
//func NewDecType() DecType {
//	return DecType{
//		Base: BaseType{},
//	}
//}
//
//func NewOctType() OctType {
//	return OctType{
//		Base: BaseType{},
//	}
//}
//
//func NewBinType() BinType {
//	return BinType{
//		Base: BaseType{},
//	}
//}
//
