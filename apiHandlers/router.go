package apiHandlers

import (
	"UIGenerator/api"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/UIGenerator/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")
	DefaultMappings(defaultGroup)
	RouteMappings(group)
}

func RouteMappings(cg fiber.Router) {
cg.Post("/CreateProject", api.CreateProjectApi)
cg.Put("/UpdateProject", api.UpdateProjectApi)
cg.Delete("/DeleteProject", api.DeleteProjectApi)
cg.Get("/FindProject", api.FindProjectApi)
cg.Get("/FindallProject", api.FindallProjectApi)
cg.Post("/UploadProject", api.UploadProjectApi)
cg.Get("/DownloadProject", api.DownloadProjectApi)
cg.Post("/CreateConfig", api.CreateConfigApi)
cg.Put("/UpdateConfig", api.UpdateConfigApi)
cg.Delete("/DeleteConfig", api.DeleteConfigApi)
cg.Get("/FindConfig", api.FindConfigApi)
cg.Get("/FindallConfig", api.FindallConfigApi)
cg.Post("/UploadConfig", api.UploadConfigApi)
cg.Get("/DownloadConfig", api.DownloadConfigApi)
cg.Post("/CreateChangeLog", api.CreateChangeLogApi)
cg.Put("/UpdateChangeLog", api.UpdateChangeLogApi)
cg.Delete("/DeleteChangeLog", api.DeleteChangeLogApi)
cg.Get("/FindChangeLog", api.FindChangeLogApi)
cg.Get("/FindallChangeLog", api.FindallChangeLogApi)
cg.Post("/UploadChangeLog", api.UploadChangeLogApi)
cg.Get("/DownloadChangeLog", api.DownloadChangeLogApi)
cg.Post("/CreateAttachment", api.CreateAttachmentApi)
cg.Put("/UpdateAttachment", api.UpdateAttachmentApi)
cg.Delete("/DeleteAttachment", api.DeleteAttachmentApi)
cg.Get("/FindAttachment", api.FindAttachmentApi)
cg.Get("/FindallAttachment", api.FindallAttachmentApi)
cg.Post("/UploadAttachment", api.UploadAttachmentApi)
cg.Get("/DownloadAttachment", api.DownloadAttachmentApi)
cg.Get("/FindallifConfigByProjectId/ProjectId", api.FindallifConfigByProjectIdApi)
cg.Get("/FindallifChangeLogByProjectId/ProjectId", api.FindallifChangeLogByProjectIdApi)
cg.Get("/FindallifAttachmentByChangeLogId/ChangeLogId", api.FindallifAttachmentByChangeLogIdApi)

cg.Get("/swagger", api.SwaggerHandler)

}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "UIGenerator-APP2084 service is up and running", "version": "1.0"})
	})
}