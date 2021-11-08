package server

import (
	"fmt"
	"gobang/negative_max"
	"gobang/variable"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func GobangHandler(w http.ResponseWriter, r *http.Request) {
	// get index from frontend
	query := r.URL.Query()
	indexReceived := query.Get("index")
	sessionID := query.Get("sessionID")
	zap.L().Debug("receive index:", zap.String("indexReceived: ", indexReceived))
	zap.L().Debug("sessionID:", zap.String("sessionID: ", indexReceived))
	session, ok := variable.SessionMap[sessionID]
	if !ok {
		// TODO:现在直接创建一个新的session，后续加上推送刷新前端页面
		zap.L().Debug("fail to get session")
		variable.SessionMap[sessionID] = &variable.SessionInfo{
			List1:       make(map[variable.Coordinate]bool),
			List2:       make(map[variable.Coordinate]bool),
			List3:       make(map[variable.Coordinate]bool),
			ListAll:     []int{},
			NextPoint:   [2]int{0, 0},
			CutCount:    0,
			SearchCount: 0,
		}
		session, ok = variable.SessionMap[sessionID]
		if !ok {
			zap.L().Error("fail to creat session")
			return
		}
	}

	indexInt, err := strconv.Atoi(indexReceived)
	if err != nil {
		zap.L().Error("fail to trans string to int", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// restart game
	if indexInt == -1 {
		zap.L().Info("restart game")
	}

	// renew list1, list2, list3
	playerCoor := Index2grid(indexInt)
	session.List2[playerCoor] = true
	session.List3[playerCoor] = true
	// you win
	if negative_max.GameWin(session.List2) {
		zap.L().Info("you win")
		session.Finish = true
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "youWin")
		return
	}

	// get ai index and response
	aiCoor := negative_max.AI(session)
	indexReturn := Grid2index(aiCoor)
	session.List1[aiCoor] = true
	session.List3[aiCoor] = true
	// ai win
	fmt.Fprint(w, indexReturn)
	if negative_max.GameWin(session.List1) {
		zap.L().Info("ai win")
		session.Finish = true
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, ",aiWin")
		return
	}
}
