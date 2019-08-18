package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetDailyChallengeData(helper *helper.Helper) {
    // no player, agnostic
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DailyChallengeData(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}

func GetCostList(helper *helper.Helper) {
    // no player, agonstic
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultCostList(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}

func GetMileageData(helper *helper.Helper) {
    player, err := helper.GetCallingPlayer()
    if err != nil {
        helper.InternalErr("Error getting player", err) // TODO: see if InternalErr is consistent with other usage of this context
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultMileageData(baseInfo, player)
    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}

func GetCampaignList(helper *helper.Helper) {
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultCampaignList(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
