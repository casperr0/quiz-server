package handler

import "github.com/gin-gonic/gin"

// GetOfficersHandler handles get requests on route /officers.
func GetOfficersHandler(ctx *gin.Context) {}

// PostOfficersHandler handles post requests on route /officers.
func PostOfficersHandler(ctx *gin.Context) {}

// UpdateOfficerHandler handles update requests on route /officers/<officer_name>.
func UpdateOfficerHandler(ctx *gin.Context) {}

// DeleteOfficerHandler handles delete requests on route /officers/<officer_name>.
func DeleteOfficerHandler(ctx *gin.Context) {}

// GetOfficerRolesHandler handles get requests on route /officers/<officer_name>/roles.
func GetOfficerRolesHandler(ctx *gin.Context) {}

// PostOfficerRolesHandler handles post requests on route /officers/<officer_name>/roles.
func PostOfficerRolesHandler(ctx *gin.Context) {}

// DeleteOfficerRoleHandler handles delete requests on route /officers/<officer_name>/roles/<role_name>.
func DeleteOfficerRoleHandler(ctx *gin.Context) {}
