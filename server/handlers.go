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
	err := initDb()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("initDataBase error")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.String(http.StatusOK, msg)
}

func postRent(c echo.Context) error {
	var rent Rent
	if err := c.Bind(&rent); err != nil {
		log.Error().Stack().Err(err).Msg("postRent")
		return echo.NewHTTPError(http.StatusNotAcceptable)
	}
	err := createRent(rent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("postRent")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func getRent(c echo.Context) error {
	rents, err := listRent()
	if err != nil {
		log.Error().Stack().Err(err).Msg("getRent")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, rents)
}

func putRent(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not valid, id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	var body map[string]map[string]interface{}
	if err := c.Bind(&body); err != nil {
		log.Error().Stack().Err(err).Msg("putRent")
		return echo.NewHTTPError(http.StatusNotAcceptable)
	}
	data, ok := body["data"]
	if !ok {
		log.Error().Msgf("error handling request body at putRent() %v", body)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Debug().
		Int64("path_param_id", id).
		Interface("request_body", data).
		Msg("")
	if err := updateRent(id, data); err != nil {
		log.Error().Stack().Err(err).Msg("putRent")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func delRent(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not valid, id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := removeRent(id); err != nil {
		log.Error().Stack().Err(err).Msg("delRent")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func getRentHistory(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not valid, id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	// time.Sleep(5 * time.Second)
	logs, err := listRentHystory(id)
	if err != nil {
		log.Error().Stack().Err(err).Msg("getRentHistory")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, logs)
}

func processRent(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		idParam = "0"
	}
	rentId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("rentId not valid, id : %v", rentId))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = createReminder(rentId)
	if err != nil {
		log.Error().Stack().Err(err).Msg("getCue")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func getReminderDetail(c echo.Context) error {
	idParam := c.Param("id")
	reminderId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("reminderId not valid, id : %v", reminderId))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	reminders, err := listReminderDetail(reminderId)
	if err != nil {
		log.Error().Stack().Err(err).Msg("getReminders")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, reminders)
}

func getReminders(c echo.Context) error {
	reminders, err := listReminders(0, time.Now(), time.Now())
	if err != nil {
		log.Error().Stack().Err(err).Msg("getReminders")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, reminders)
}

func payRent(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msg(fmt.Sprintf("id not valid. id : %v", id))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := changeReminderStatus(id, paid); err != nil {
		log.Error().Stack().Err(err).Msg("payRent()")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func processReminders(c echo.Context) error {
	err := processRemindersDates(0)
	// time.Sleep(3 * time.Second)
	if err != nil {
		log.Error().Stack().Err(err).Msg("processReminders")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func getHistory(c echo.Context) error {
	history, err := listSchedulerHistory()
	// time.Sleep(3 * time.Second)
	if err != nil {
		log.Error().Stack().Err(err).Msg("getHistory")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, history)
}

// func sstart(c echo.Context) error {
// 	start()
// 	return c.NoContent(http.StatusOK)
// }
// func sstop(c echo.Context) error {
// 	stop()
// 	return c.NoContent(http.StatusOK)
// }
// func test(c echo.Context) error {
// 	str := sc()
// 	return c.JSON(http.StatusOK, str)
// }

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
