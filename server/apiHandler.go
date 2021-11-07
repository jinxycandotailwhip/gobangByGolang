package server

import (
	"fmt"
	"gobang/negative_max"
	"gobang/variable"
	"log"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func GobangHandler(w http.ResponseWriter, r *http.Request) {
	// get index from frontend
	query := r.URL.Query()
	indexReceived := query.Get("index")
	zap.L().Debug("receive index:", zap.String("indexReceived: ", indexReceived))
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
	variable.List2[playerCoor] = true
	variable.List3[playerCoor] = true
	// you win
	if negative_max.GameWin(variable.List2) {
		zap.L().Info("you win")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "youWin")
		return
	}

	// get ai index and response
	aiCoor := negative_max.AI()
	indexReturn := Grid2index(aiCoor)
	variable.List1[aiCoor] = true
	variable.List3[aiCoor] = true
	log.Default().Println("index responsed:", indexReturn)
	// ai win
	fmt.Fprint(w, indexReturn)
	if negative_max.GameWin(variable.List1) {
		zap.L().Info("ai win")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, ",aiWin")
		return
	}
}
