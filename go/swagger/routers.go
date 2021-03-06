/*
 * Swagger Petstore
 *
 * This is a sample server Petstore server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func NewRouter(impPetApi PetApi, impStoreApi StoreApi, impUserApi UserApi) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/v2/",
			Index,
		},

		Route{
			"AddPet",
			strings.ToUpper("Post"),
			"/v2/pet",
			AddPetHandler(impPetApi),
		},

		Route{
			"DeletePet",
			strings.ToUpper("Delete"),
			"/v2/pet/{petId}",
			DeletePetHandler(impPetApi),
		},

		Route{
			"FindPetsByStatus",
			strings.ToUpper("Get"),
			"/v2/pet/findByStatus",
			FindPetsByStatusHandler(impPetApi),
		},

		Route{
			"FindPetsByTags",
			strings.ToUpper("Get"),
			"/v2/pet/findByTags",
			FindPetsByTagsHandler(impPetApi),
		},

		Route{
			"GetPetById",
			strings.ToUpper("Get"),
			"/v2/pet/{petId}",
			GetPetByIdHandler(impPetApi),
		},

		Route{
			"UpdatePet",
			strings.ToUpper("Put"),
			"/v2/pet",
			UpdatePetHandler(impPetApi),
		},

		Route{
			"UpdatePetWithForm",
			strings.ToUpper("Post"),
			"/v2/pet/{petId}",
			UpdatePetWithFormHandler(impPetApi),
		},

		Route{
			"UploadFile",
			strings.ToUpper("Post"),
			"/v2/pet/{petId}/uploadImage",
			UploadFileHandler(impPetApi),
		},

		Route{
			"DeleteOrder",
			strings.ToUpper("Delete"),
			"/v2/store/order/{orderId}",
			DeleteOrderHandler(impStoreApi),
		},

		Route{
			"GetInventory",
			strings.ToUpper("Get"),
			"/v2/store/inventory",
			GetInventoryHandler(impStoreApi),
		},

		Route{
			"GetOrderById",
			strings.ToUpper("Get"),
			"/v2/store/order/{orderId}",
			GetOrderByIdHandler(impStoreApi),
		},

		Route{
			"PlaceOrder",
			strings.ToUpper("Post"),
			"/v2/store/order",
			PlaceOrderHandler(impStoreApi),
		},

		Route{
			"CreateUser",
			strings.ToUpper("Post"),
			"/v2/user",
			CreateUserHandler(impUserApi),
		},

		Route{
			"CreateUsersWithArrayInput",
			strings.ToUpper("Post"),
			"/v2/user/createWithArray",
			CreateUsersWithArrayInputHandler(impUserApi),
		},

		Route{
			"CreateUsersWithListInput",
			strings.ToUpper("Post"),
			"/v2/user/createWithList",
			CreateUsersWithListInputHandler(impUserApi),
		},

		Route{
			"DeleteUser",
			strings.ToUpper("Delete"),
			"/v2/user/{username}",
			DeleteUserHandler(impUserApi),
		},

		Route{
			"GetUserByName",
			strings.ToUpper("Get"),
			"/v2/user/{username}",
			GetUserByNameHandler(impUserApi),
		},

		Route{
			"LoginUser",
			strings.ToUpper("Get"),
			"/v2/user/login",
			LoginUserHandler(impUserApi),
		},

		Route{
			"LogoutUser",
			strings.ToUpper("Get"),
			"/v2/user/logout",
			LogoutUserHandler(impUserApi),
		},

		Route{
			"UpdateUser",
			strings.ToUpper("Put"),
			"/v2/user/{username}",
			UpdateUserHandler(impUserApi),
		},
	}
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
