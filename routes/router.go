package routes

import (
	"github.com/KasonBraley/chat-li/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.GET("/shows/search", controllers.FindShows)

	/* create api group around /rooms endpoint
	   get('/', getAllRooms);
	   post('/', createRoom);
	   delete('/:id', bearerAuth, acl('delete'), deleteRoom);
	*/

	/* create api group around /users endpoint
	    get('/', bearerAuth, acl('delete'), getAllUsers);
	    get('/:id', bearerAuth, acl('delete'), getOneUser);
	    post('/', bearerAuth, acl('delete'), createUser);
	   	delete('/:id', bearerAuth, acl('delete'), deleteUser);
	*/

	/* create api group around auth routes
	authRouter.post('/signup', async (req, res, next) => {
	authRouter.post('/signin', basicAuth, (req, res, next) => {
	authRouter.post('/joinroom', basicAuth, (req, res, next) => {

	*/
	return r
}

func rootHandler(c *gin.Context) {
	c.String(200, "These are not the bugs you are looking for")
}
