package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// TODO remove
func initdb(c echo.Context) error {
	idParam := c.Param("clear")
	msg := "initDataBase()"
	if idParam == "1" {
		clearDB()
		msg = "cue table cleared\n" + msg
	}
	err := initDB()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("initDataBase error")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.String(http.StatusOK, msg)
}

func getCue(c echo.Context) error {
	var cues []Cue
	cues, err := listCue()
	if err != nil {
		log.Error().Stack().Err(err).Msg("getCue")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, cues)
}

func postCue(c echo.Context) error {
	var cue Cue
	if err := c.Bind(&cue); err != nil {
		log.Error().Stack().Err(err).Msg("postCue")
		return echo.NewHTTPError(http.StatusNotAcceptable)
	}

	err := createCue(cue)
	if err != nil {
		log.Error().Stack().Err(err).Msg("postCue")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func putCue(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not value, id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	var body map[string]map[string]interface{}
	if err := c.Bind(&body); err != nil {
		log.Error().Stack().Err(err).Msg("putCue")
		return echo.NewHTTPError(http.StatusNotAcceptable)
	}

	data, ok := body["data"]
	if !ok {
		log.Error().Msgf("error handling request body at putCue() %v", body)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Debug().Int64("path_param_id", id).Msg("")
	log.Debug().Interface("request_body", data).Msg("")

	if err := updateCue(id, data); err != nil {
		log.Error().Stack().Err(err).Msg("putCue")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func delCue(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not value, id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := removeCue(id); err != nil {
		log.Error().Stack().Err(err).Msg("delCue")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func payCue(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not value, id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := changeCueStatus(id, paid); err != nil {
		log.Error().Stack().Err(err).Msg("delCue")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func checkDues(c echo.Context) error {
	err := checkDueDates()
	time.Sleep(3 * time.Second)
	if err != nil {
		log.Error().Stack().Err(err).Msg("checkDues")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func getHistory(c echo.Context) error {
	history, err := listSchedulerHistory()
	time.Sleep(3 * time.Second)
	if err != nil {
		log.Error().Stack().Err(err).Msg("getHistory")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, history)
}

// func postTenant(c echo.Context) error {
// 	tenant := Tenant{}

// 	if err := c.Bind(&tenant); err != nil {
// 		logger(err)
// 		return err
// 	}

// 	if isValid, msg := isTenantValid(tenant); !isValid {
// 		logger(msg)
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
// 	}

// 	err := saveTenant(tenant)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Propriedade não encontrada!")
// 		} else {
// 			logger(err)
// 			return echo.NewHTTPError(http.StatusInternalServerError)
// 		}
// 	}

// 	return c.JSON(http.StatusCreated, tenant)
// }

// func getTenant(c echo.Context) error {

// 	var tenants []Tenant = []Tenant{}
// 	err := listDocuments(&tenants, primitive.NilObjectID, nil)
// 	if err != nil {
// 		logger(err)
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	}
// 	return c.JSON(http.StatusOK, tenants)
// }

// func putTenant(c echo.Context) error {

// 	// Get Object id from params
// 	id := c.Param("id")

// 	if !primitive.IsValidObjectID(id) {
// 		msg := "Id do objeto inválido"
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
// 	}

// 	tenant := Tenant{}

// 	if err := c.Bind(&tenant); err != nil {
// 		logger(err)
// 		return err
// 	}

// 	// Clear ObjectId if its not null
// 	if !tenant.Id.IsZero() {
// 		tenant.Id = primitive.NilObjectID
// 	}

// 	// TODO validation
// 	// if isValid, msg := isTenantValid(tenant); !isValid {
// 	// 	log.Println(msg)
// 	// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
// 	// }

// 	objectId, _ := primitive.ObjectIDFromHex(id)

// 	result, err := updateDocument(objectId, tenant)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	} else {
// 		if result > 0 {
// 			return c.JSON(http.StatusNoContent, "Registro atualizado")
// 		} else {
// 			return echo.NewHTTPError(http.StatusNotFound, "Registro não encontrado")
// 		}
// 	}
// }

// func delete(c echo.Context, docType interface{}) error {

// 	id := c.Param("id")
// 	if !primitive.IsValidObjectID(id) {
// 		msg := "Id do objeto inválido"
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
// 	}

// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	result, err := removeDocument(objectId, docType)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	} else {
// 		if result > 0 {
// 			return c.JSON(http.StatusNoContent, "Registro atualizado")
// 		} else {
// 			return echo.NewHTTPError(http.StatusNotFound, "Registro não encontrado")
// 		}
// 	}
// }

// func deleteTenant(c echo.Context) error {
// 	return delete(c, Tenant{})
// }

// func postProperty(c echo.Context) error {

// 	property := Property{}

// 	if err := c.Bind(&property); err != nil {
// 		logger(err)
// 		return err
// 	}

// 	// TODO validations

// 	// property.Tenant = Tenant{
// 	// 	Name: "asdasdasd",
// 	// 	Rg:   "324234523432",
// 	// }
// 	err := saveProperty(property)

// 	if err != nil {
// 		echoErr, isEchoErr := err.(*echo.HTTPError)
// 		if isEchoErr {
// 			return echoErr
// 		} else {
// 			logger(err)
// 			return echo.NewHTTPError(http.StatusInternalServerError)
// 		}
// 	}

// 	return c.JSON(http.StatusCreated, property)
// }

// func getProperty(c echo.Context) error {
// 	// time.Sleep(5 * time.Second)
// 	// var properties []Property = []Property{}
// 	// err := listDocuments(&properties, primitive.NilObjectID)
// 	properties, err := listProperties()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, properties)
// }

// func putProperty(c echo.Context) error {

// 	id := c.Param("id")
// 	if !primitive.IsValidObjectID(id) {
// 		msg := "Id do objeto inválido"
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity, msg)
// 	}

// 	property := Property{}

// 	if err := c.Bind(&property); err != nil {
// 		logger(err)
// 		return err
// 	}

// 	// Clear ObjectId if its not null
// 	if !property.Id.IsZero() {
// 		property.Id = primitive.NilObjectID
// 	}

// 	// TODO validation

// 	objectId, _ := primitive.ObjectIDFromHex(id)

// 	result, err := updateDocument(objectId, property)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError)
// 	} else {
// 		if result > 0 {
// 			return c.JSON(http.StatusNoContent, "Registro atualizado")
// 		} else {
// 			return echo.NewHTTPError(http.StatusNotFound, "Registro não encontrado")
// 		}
// 	}
// }

// func deleteProperty(c echo.Context) error {
// 	return delete(c, Property{})
// }

// func postRent(c echo.Context) error {
// 	rent := Rent{}

// 	if err := c.Bind(&rent); err != nil {
// 		logger(err)
// 		return err
// 	}

// 	// err := saveProperty(property)
// 	err := saveRent(rent)

// 	if err != nil {
// 		echoErr, isEchoErr := err.(*echo.HTTPError)
// 		if isEchoErr {
// 			return echoErr
// 		} else {
// 			// logger(err)
// 			return echo.NewHTTPError(http.StatusInternalServerError)
// 		}
// 	}

// 	return c.JSON(http.StatusCreated, rent)
// }
