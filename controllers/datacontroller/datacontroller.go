package datacontroller

import (
	"net/http"
	"sdqh/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":            1,
			"nama_kegiatan": "waqaf air",
			"nilai":         10000000,
		},
		{
			"id":            2,
			"nama_kegiatan": "pembebasan tanah masjid",
			"stok":          50000000,
		},
		{
			"id":           1,
			"nama_product": "waqaf alquran",
			"stok":         100000000,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)

}
