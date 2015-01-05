goREST
======
Preliminary HTTP REST router.  This currently only supports consecutive paths with a prefix.

Installation:

    go get github.com/euforia/go-rest-router

Example:

    import restrouter "github.com/euforia/go-rest-router"

    
    /api/endpoint/:type/:id

Take a look at the following test function for example usage:

    rest/router_test.go: func Test_NewRESTRouter