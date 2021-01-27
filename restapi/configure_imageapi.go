// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	internal "github.com/jlowellwofford/imageapi/internal/api"
	"github.com/jlowellwofford/imageapi/models"
	"github.com/jlowellwofford/imageapi/restapi/operations"
	"github.com/jlowellwofford/imageapi/restapi/operations/attach"
	"github.com/jlowellwofford/imageapi/restapi/operations/containers"
	"github.com/jlowellwofford/imageapi/restapi/operations/mounts"
)

//go:generate swagger generate server --target ../../imageapi --name Imageapi --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.ImageapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ImageapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	////////////////////////////////
	// Begin: Non-generated block /
	//////////////////////////////

	api.AttachListRbdsHandler = attach.ListRbdsHandlerFunc(func(params attach.ListRbdsParams) middleware.Responder {
		return attach.NewListRbdsOK().WithPayload(internal.Rbds.List())
	})

	api.AttachMapRbdHandler = attach.MapRbdHandlerFunc(func(params attach.MapRbdParams) middleware.Responder {
		var err error
		var r *models.Rbd
		if r, err = internal.Rbds.Map(params.Rbd); err != nil {
			return attach.NewMapRbdDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return attach.NewMapRbdCreated().WithPayload(r)
	})

	api.AttachGetRbdHandler = attach.GetRbdHandlerFunc(func(params attach.GetRbdParams) middleware.Responder {
		var err error
		var r *models.Rbd
		if r, err = internal.Rbds.Get(params.ID); err != nil {
			return attach.NewGetRbdDefault(404).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return attach.NewGetRbdOK().WithPayload(r)
	})

	api.AttachUnmapRbdHandler = attach.UnmapRbdHandlerFunc(func(params attach.UnmapRbdParams) middleware.Responder {
		if err := internal.Rbds.Unmap(params.ID); err != nil {
			return attach.NewUnmapRbdDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return attach.NewUnmapRbdNoContent()
	})

	// MountsRbd

	api.MountsListMountsRbdHandler = mounts.ListMountsRbdHandlerFunc(func(params mounts.ListMountsRbdParams) middleware.Responder {
		return mounts.NewListMountsRbdOK().WithPayload(internal.MountsRbd.List())
	})

	api.MountsMountRbdHandler = mounts.MountRbdHandlerFunc(func(params mounts.MountRbdParams) middleware.Responder {
		var err error
		var r *models.MountRbd
		if r, err = internal.MountsRbd.Mount(params.Mount); err != nil {
			return mounts.NewMountRbdDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return mounts.NewMountRbdCreated().WithPayload(r)
	})

	api.MountsGetMountRbdHandler = mounts.GetMountRbdHandlerFunc(func(params mounts.GetMountRbdParams) middleware.Responder {
		var err error
		var r *models.MountRbd
		if r, err = internal.MountsRbd.Get(params.ID); err != nil {
			return mounts.NewGetMountRbdDefault(404).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewGetMountRbdOK().WithPayload(r)
	})

	api.MountsUnmountRbdHandler = mounts.UnmountRbdHandlerFunc(func(params mounts.UnmountRbdParams) middleware.Responder {
		if err := internal.MountsRbd.Unmount(params.ID); err != nil {
			return mounts.NewUnmountRbdDefault(500).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewUnmountRbdNoContent()
	})

	// MountsOverlay

	api.MountsListMountsOverlayHandler = mounts.ListMountsOverlayHandlerFunc(func(params mounts.ListMountsOverlayParams) middleware.Responder {
		return mounts.NewListMountsOverlayOK().WithPayload(internal.MountsOverlay.List())
	})

	api.MountsMountOverlayHandler = mounts.MountOverlayHandlerFunc(func(params mounts.MountOverlayParams) middleware.Responder {
		var err error
		var r *models.MountOverlay
		if r, err = internal.MountsOverlay.Mount(params.Mount); err != nil {
			return mounts.NewMountOverlayDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return mounts.NewMountOverlayCreated().WithPayload(r)
	})

	api.MountsGetMountOverlayHandler = mounts.GetMountOverlayHandlerFunc(func(params mounts.GetMountOverlayParams) middleware.Responder {
		var err error
		var r *models.MountOverlay
		if r, err = internal.MountsOverlay.Get(params.ID); err != nil {
			return mounts.NewGetMountOverlayDefault(404).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewGetMountOverlayOK().WithPayload(r)
	})

	api.MountsUnmountOverlayHandler = mounts.UnmountOverlayHandlerFunc(func(params mounts.UnmountOverlayParams) middleware.Responder {
		if err := internal.MountsOverlay.Unmount(params.ID); err != nil {
			return mounts.NewUnmountOverlayDefault(500).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewUnmountOverlayNoContent()
	})

	//////////////////////////////
	// End: Non-generated block /
	////////////////////////////

	if api.ContainersCreateContainerHandler == nil {
		api.ContainersCreateContainerHandler = containers.CreateContainerHandlerFunc(func(params containers.CreateContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation containers.CreateContainer has not yet been implemented")
		})
	}
	if api.ContainersDeleteContainerHandler == nil {
		api.ContainersDeleteContainerHandler = containers.DeleteContainerHandlerFunc(func(params containers.DeleteContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation containers.DeleteContainer has not yet been implemented")
		})
	}
	if api.ContainersGetContainerHandler == nil {
		api.ContainersGetContainerHandler = containers.GetContainerHandlerFunc(func(params containers.GetContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation containers.GetContainer has not yet been implemented")
		})
	}
	if api.MountsGetMountOverlayHandler == nil {
		api.MountsGetMountOverlayHandler = mounts.GetMountOverlayHandlerFunc(func(params mounts.GetMountOverlayParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.GetMountOverlay has not yet been implemented")
		})
	}
	if api.MountsGetMountRbdHandler == nil {
		api.MountsGetMountRbdHandler = mounts.GetMountRbdHandlerFunc(func(params mounts.GetMountRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.GetMountRbd has not yet been implemented")
		})
	}
	if api.AttachGetRbdHandler == nil {
		api.AttachGetRbdHandler = attach.GetRbdHandlerFunc(func(params attach.GetRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation attach.GetRbd has not yet been implemented")
		})
	}
	if api.ContainersListContainersHandler == nil {
		api.ContainersListContainersHandler = containers.ListContainersHandlerFunc(func(params containers.ListContainersParams) middleware.Responder {
			return middleware.NotImplemented("operation containers.ListContainers has not yet been implemented")
		})
	}
	if api.MountsListMountsOverlayHandler == nil {
		api.MountsListMountsOverlayHandler = mounts.ListMountsOverlayHandlerFunc(func(params mounts.ListMountsOverlayParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.ListMountsOverlay has not yet been implemented")
		})
	}
	if api.MountsListMountsRbdHandler == nil {
		api.MountsListMountsRbdHandler = mounts.ListMountsRbdHandlerFunc(func(params mounts.ListMountsRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.ListMountsRbd has not yet been implemented")
		})
	}
	if api.AttachListRbdsHandler == nil {
		api.AttachListRbdsHandler = attach.ListRbdsHandlerFunc(func(params attach.ListRbdsParams) middleware.Responder {
			return middleware.NotImplemented("operation attach.ListRbds has not yet been implemented")
		})
	}
	if api.AttachMapRbdHandler == nil {
		api.AttachMapRbdHandler = attach.MapRbdHandlerFunc(func(params attach.MapRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation attach.MapRbd has not yet been implemented")
		})
	}
	if api.MountsMountOverlayHandler == nil {
		api.MountsMountOverlayHandler = mounts.MountOverlayHandlerFunc(func(params mounts.MountOverlayParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.MountOverlay has not yet been implemented")
		})
	}
	if api.MountsMountRbdHandler == nil {
		api.MountsMountRbdHandler = mounts.MountRbdHandlerFunc(func(params mounts.MountRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.MountRbd has not yet been implemented")
		})
	}
	if api.ContainersSetContainerStateHandler == nil {
		api.ContainersSetContainerStateHandler = containers.SetContainerStateHandlerFunc(func(params containers.SetContainerStateParams) middleware.Responder {
			return middleware.NotImplemented("operation containers.SetContainerState has not yet been implemented")
		})
	}
	if api.AttachUnmapRbdHandler == nil {
		api.AttachUnmapRbdHandler = attach.UnmapRbdHandlerFunc(func(params attach.UnmapRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation attach.UnmapRbd has not yet been implemented")
		})
	}
	if api.MountsUnmountOverlayHandler == nil {
		api.MountsUnmountOverlayHandler = mounts.UnmountOverlayHandlerFunc(func(params mounts.UnmountOverlayParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.UnmountOverlay has not yet been implemented")
		})
	}
	if api.MountsUnmountRbdHandler == nil {
		api.MountsUnmountRbdHandler = mounts.UnmountRbdHandlerFunc(func(params mounts.UnmountRbdParams) middleware.Responder {
			return middleware.NotImplemented("operation mounts.UnmountRbd has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
