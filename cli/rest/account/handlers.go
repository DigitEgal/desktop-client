package account

import (
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	"github.com/sentinel-official/desktop-client/cli/context"
	"github.com/sentinel-official/desktop-client/cli/utils"
	"github.com/sentinel-official/desktop-client/cli/x/auth"
)

func HandlerGetAccount(ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		address, err := sdk.AccAddressFromHex(vars["address"])
		if err != nil {
			utils.WriteErrorToResponse(w, http.StatusBadRequest, 1, err.Error())
			return
		}

		result, err := ctx.Client().QueryAccount(address)
		if err != nil {
			utils.WriteErrorToResponse(w, http.StatusInternalServerError, 2, err.Error())
			return
		}

		item := auth.NewAccountFromRaw(result)
		utils.WriteResultToResponse(w, http.StatusOK, item)
	}
}
