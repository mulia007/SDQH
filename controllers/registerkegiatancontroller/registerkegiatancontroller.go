package registerkegiatancontroller

import (
	"encoding/json"
	"net/http"
	"sdqh/helper"
	"sdqh/models"
)

func CreateKegiatan(w http.ResponseWriter, r *http.Request) {
	var kegiatan models.Waqaf_Sedeqah
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&kegiatan); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// save data to database
	if err := models.DB.Create(&kegiatan).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
