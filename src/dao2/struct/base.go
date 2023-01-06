package _struct

type IDb interface {
	// TName should return a unique name that identifies the object,
	// this return value will be used for database creation
	TName() string
	String() string
}
