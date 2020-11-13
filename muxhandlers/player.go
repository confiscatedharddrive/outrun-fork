package muxhandlers

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db"
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/obj/constobjs"
	"github.com/Mtbcooler/outrun/requests"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
	"github.com/jinzhu/now"
)

func GetPlayerState(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}
	if player.PlayerState.NextNumDailyChallenge <= 0 || int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
		// Initialize daily challenge if it isn't initialized already
		player.PlayerState.NumDailyChallenge = int64(0)
		player.PlayerState.NextNumDailyChallenge = int64(1)
		player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
	}
	if time.Now().UTC().Unix() >= player.PlayerState.DailyMissionEndTime {
		if player.PlayerState.DailyChallengeComplete == 1 && player.PlayerState.DailyChalSetNum < 10 {
			helper.DebugOut("Advancing to next daily mission...")
			player.PlayerState.DailyChalSetNum++
		} else {
			player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
			player.PlayerState.DailyChalSetNum = int64(0)
		}
		if player.PlayerState.DailyChallengeComplete == 0 {
			player.PlayerState.NumDailyChallenge = int64(0)
			player.PlayerState.NextNumDailyChallenge = int64(1)
		} else {
			player.PlayerState.NextNumDailyChallenge++
			if int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1) //restart from beginning
				player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
				player.PlayerState.DailyChalSetNum = int64(0)
			}
		}
		player.PlayerState.DailyChalPosNum = int64(1 + rand.Intn(2))
		player.PlayerState.DailyMissionID = int64((player.PlayerState.DailyChalCatNum * 33) + (player.PlayerState.DailyChalSetNum * 3) + player.PlayerState.DailyChalPosNum)
		player.PlayerState.DailyChallengeValue = int64(0)
		player.PlayerState.DailyChallengeComplete = int64(0)
		player.PlayerState.DailyMissionEndTime = now.EndOfDay().UTC().Unix() + 1
		helper.DebugOut("New daily mission ID: %v", player.PlayerState.DailyMissionID)
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.PlayerState(baseInfo, player.PlayerState)
	helper.SendResponse(response)
}

func GetCharacterState(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	// below is a lazy hack to add event characters to the character state - UPDATE THIS FOR CHRISTMAS IF NECESSARY
	/*charindex := player.IndexOfChara(enums.CTStrGothicAmy)
	if charindex == -1 {
		player.CharacterState = append(player.CharacterState, netobj.DefaultSpecialLockedCharacter(constobjs.CharacterGothicAmy))
	}*/
	/*charindex = player.IndexOfChara(enums.CTStrHalloweenShadow)
	if charindex == -1 {
		player.CharacterState = append(player.CharacterState, netobj.DefaultRouletteOnlyLockedCharacter(constobjs.CharacterHalloweenShadow))
	}
	charindex = player.IndexOfChara(enums.CTStrHalloweenRouge)
	if charindex == -1 {
		player.CharacterState = append(player.CharacterState, netobj.DefaultRouletteOnlyLockedCharacter(constobjs.CharacterHalloweenRouge))
	}
	charindex = player.IndexOfChara(enums.CTStrHalloweenOmega)
	if charindex == -1 {
		player.CharacterState = append(player.CharacterState, netobj.DefaultRouletteOnlyLockedCharacter(constobjs.CharacterHalloweenOmega))
	}*/
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.CharacterState(baseInfo, player.CharacterState)
	helper.SendResponse(response)
}

func GetChaoState(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.ChaoState(baseInfo, player.ChaoState)
	helper.SendResponse(response)
}

func SetUsername(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.SetUsernameRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	// TODO: check if username is already taken
	player.Username = request.Username
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.NewBaseResponse(baseInfo)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
