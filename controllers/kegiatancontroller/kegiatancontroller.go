package kegiatancontroller

import (
	"encoding/json"
	"net/http"
	"sdqh/helper"
	"sdqh/models"

	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// ambil data user berdasarkan username
	var sedeqah []models.Waqaf_Sedeqah
	if err := models.DB.Find(&sedeqah).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseJSON(w, http.StatusNotFound, `data not found`)
			return
		default:
			helper.ResponseJSON(w, http.StatusOK, sedeqah)
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, sedeqah)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	// ambil data user berdasarkan username
	var sedeqah models.Waqaf_Sedeqah
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sedeqah); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Where("id = ?", sedeqah.Id).First(&sedeqah).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseJSON(w, http.StatusNotFound, `data not found`)
			return
		default:
			helper.ResponseJSON(w, http.StatusOK, sedeqah)
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, sedeqah)
}
