goREST
======

Preliminary HTTP REST router.  This currently only supports a single consecutive path with a prefix:

    Valid:

        /api/endpoint/:type/:id

    Invalid:

        /api/endpoint/:type/static/:id    

#### Installation:

    go get github.com/euforia/goREST

#### Usage:

    import gorest "github.com/euforia/goREST"

    type TypeIdHandler struct {
        // Sets all methods not implemented to the default code of 405 //
        DefaultEndpointMethodsHandler
    }

    func (r *TypeIdHandler) GET(r *http.Request, args ...string) (map[string]string, interface{}, int) {
        // positional per the request path //
        _type := args[0]
        _id := args[1]
        
        // do something... //
        
        headers := map[string]string{"Access-Control-Allow-Origin": "*"}
        // some data //
        data := "..."
        // response code //
        respCode := 200
        
        return headers, data, respCode
    }


    router := gorest.NewRESTRouter("/api/prefix", nil)

    router.Register("/", RootHandler)
    router.Reister("/:type", TypeHandler)
    router.Reister("/:type/:_id", TypeIdHandler)

    http.ListenAndServe(":8000", router)

** The handler must implement the EndpointMethodsHandler interface **