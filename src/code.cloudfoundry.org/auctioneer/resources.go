package auctioneer

// LRPStartRequest and TaskStartRequest have moved to code.cloudfoundry.org/bbs/models.
// These aliases maintain backward compatibility.

import bbsmodels "code.cloudfoundry.org/bbs/models"

type LRPStartRequest  = bbsmodels.LRPStartRequest
type TaskStartRequest = bbsmodels.TaskStartRequest

var NewLRPStartRequest                   = bbsmodels.NewLRPStartRequest
var NewLRPStartRequestFromModel          = bbsmodels.NewLRPStartRequestFromModel
var NewLRPStartRequestFromSchedulingInfo = bbsmodels.NewLRPStartRequestFromSchedulingInfo
var NewTaskStartRequest                  = bbsmodels.NewTaskStartRequest
var NewTaskStartRequestFromModel         = bbsmodels.NewTaskStartRequestFromModel
