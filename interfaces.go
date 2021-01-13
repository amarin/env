package env

// StringScanner interface requires implementation provided ScanString method which helps translate string
// value from env into implementor type.
// Usually implemented by basic types successors.
type StringScanner interface {
	// ScanString should implement translation from env string value into target type of implementor.
	// May generate translation or internal validation error.
	ScanString(string) error
}

// Loader interface requires implementation provides LoadEnv method which manage env variables loading.
// Usually implemented by struct types.
// Basic implementation for structures using env tags:
//
//   import "github.com/amarin/env"
//
//   func (i *MyStruct) LoadEnv() error {
//	   return env.Load(i)
//   }
type Loader interface {
	// LoadEnv orchestrates env variables loading.
	LoadEnv() error
}
