package controllers

import (
	"net/http"
	"strconv"

	"github.com/biseshbhattarai/talwar/inputs"
	"github.com/biseshbhattarai/talwar/models"
	"github.com/gin-gonic/gin"
)

func SearchOrganizationRepos(context *gin.Context) {
	targetId := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetId)
	models.StartSearchRepositories(parseIntTargetId)
	context.JSON(200, gin.H{"message": "Searched"})

}

func RegisterOrganizationRepos(context *gin.Context) {
	var input inputs.OrganizationRepoInput
	targetId := context.Param("target_id")
	parsedTargetId, _ := strconv.ParseInt(targetId, 10, 64)
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	organizationRepo := models.Organization{
		OrgName:          input.OrgName,
		OrganizationType: input.OrganizationType,
		TargetID:         parsedTargetId,
	}
	savedOrganizationRepo, err := organizationRepo.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"organizationRepo": savedOrganizationRepo})
}

func ListOrganizationRepos(context *gin.Context) {
	targetId, _ := strconv.ParseInt(context.Param("target_id"), 10, 64)
	organizationRepos, err := models.ListOrganizationsByTargetID(int(targetId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"organizationRepos": organizationRepos})
}

func SearchOrganizationReposFromGithub(context *gin.Context) {
	targetId := context.Param("target_id")
	parseIntTargetId, _ := strconv.Atoi(targetId)
	models.StartSearchRepositories(parseIntTargetId)
	context.JSON(200, gin.H{"message": "Searched"})

}

func ListRepositoryByOrganizationID(context *gin.Context) {
	organizationId := context.Param("organization_id")
	parseIntOrganizationId, _ := strconv.Atoi(organizationId)
	repositories, err := models.ListRepositoryByOrganizationID(parseIntOrganizationId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"repositories": repositories})
}

func ListOrganizationReposByTargetID(context *gin.Context) {
	targetId, _ := strconv.ParseInt(context.Param("target_id"), 10, 64)
	organizationRepos, err := models.ListOrganizationReposByTargetId(int(targetId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"organizationRepos": organizationRepos})
}
