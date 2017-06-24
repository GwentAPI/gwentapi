# TODO

## Proposing API design change

GwentAPI use the [goa](https://github.com/goadesign/goa) framework to generate most of code of the project.

To modify the API you must first install goagen, modify the design file(s) and then run goa. Goa won't overwrite files in the ``main`` package if they already exists.

If you drastically modify the resources or actions, it's your responsibility to properly implement them in the ``main`` package.

### Installing goagen

Navigate to the ``goagen`` directory in vendor and install the go program:

``cd vendor/github.com/goadesign/goa/goagen``

``go install``

### Modifying the design file(s)

They are located under the ``design`` directory.

| File        | Content           |
| ------------- |:-------------:|
| api_definition.go     | Main definition of the api |
| media_types_definition.go      | Define the different media returned by the API and the different views they have      |
| resources_definition.go | Endpoints, Return types, params, etc.      |
| types_definition.go | A bit like medias except more primitive-like. Medias can be composed of them.     |

Make sure to bump the api version in ``api_definition.go``.

You can find more information about the goa design language [here](https://goa.design/reference/goa/design/apidsl/).


### Running goagen

``goagen bootstrap -d github.com/GwentAPI/gwentapi/design``