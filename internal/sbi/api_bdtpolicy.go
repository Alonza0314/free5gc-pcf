/*
 * Npcf_BDTPolicyControl Service API
 *
 * The Npcf_BDTPolicyControl Service is used by an NF service consumer to
 * retrieve background data transfer policies from the PCF and to update
 * the PCF with the background data transfer policy selected by the NF
 * service consumer.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/pcf/internal/logger"
	"github.com/free5gc/pcf/internal/util"
)

func (s *Server) getBdtPolicyRoutes() []Route {
	return []Route{
		{
			Name:    "CreateBDTPolicy",
			Method:  http.MethodPost,
			Pattern: "/bdtpolicies",
			APIFunc: s.HTTPCreateBDTPolicy,
		},
		{
			Name:    "GetBDTPolicy",
			Method:  http.MethodGet,
			Pattern: "/bdtpolicies/:bdtPolicyId",
			APIFunc: s.HTTPGetBDTPolicy,
		},
		{
			Name:    "UpdateBDTPolicy",
			Method:  http.MethodPatch,
			Pattern: "/bdtpolicies/:bdtPolicyId",
			APIFunc: s.HTTPUpdateBDTPolicy,
		},
	}
}

func (s *Server) HTTPCreateBDTPolicy(c *gin.Context) {
	var bdtReqData models.BdtReqData
	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.BdtPolicyLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// step 2: convert requestBody to openapi models
	err = openapi.Deserialize(&bdtReqData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.BdtPolicyLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	s.Processor().HandleCreateBDTPolicyContextRequest(c, bdtReqData)
}

func (s *Server) HTTPGetBDTPolicy(c *gin.Context) {
	bdtPolicyId := c.Params.ByName("bdtPolicyId")
	if bdtPolicyId == "" {
		problemDetails := &models.ProblemDetails{
			Title:  util.ERROR_INITIAL_PARAMETERS,
			Status: http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, problemDetails)
		return
	}
	s.Processor().HandleGetBDTPolicyContextRequest(c, bdtPolicyId)
}

// UpdateBDTPolicy - Update an Individual BDT policy
func (s *Server) HTTPUpdateBDTPolicy(c *gin.Context) {
	var bdtPolicyDataPatch models.PcfBdtPolicyControlBdtPolicyDataPatch
	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.BdtPolicyLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// step 2: convert requestBody to openapi models
	err = openapi.Deserialize(&bdtPolicyDataPatch, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.BdtPolicyLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	bdtPolicyId := c.Params.ByName("bdtPolicyId")
	if bdtPolicyId == "" {
		problemDetails := &models.ProblemDetails{
			Title:  util.ERROR_INITIAL_PARAMETERS,
			Status: http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, problemDetails)
		return
	}
	s.Processor().HandleUpdateBDTPolicyContextProcedure(c, bdtPolicyId, bdtPolicyDataPatch)
}
