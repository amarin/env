= env

Golang env package adds ability to load structure configurable attributes from environment variables.
It uses structure attributes tags to attach env variable name to attribute.

Loadable attributes should be exported (name started from upper case letter).

Example usage:
   
    import "github.com/amarin/env"
    
    type MyStruct struct {
        AttributeA string `env:"MYSTRUCT_A"`
    }
    
    func NewMyStruct() *MyStruct {
        myStruct := &MyStruct{}
        env.Load(myStruct)
        
        return myStruct
    }

NOTE: Currently only string, int and bool fields processed.


