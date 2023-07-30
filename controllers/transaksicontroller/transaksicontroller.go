package transaksicontroller

import (
	"encoding/json"
	"net/http"
	"sdqh/helper"
	"sdqh/models"
)

func CreateTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi models.Transaksi
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaksi); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// save data to database
	if err := models.DB.Create(&transaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
