// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	"gocv.io/x/gocv"
)

func main() {
	cascade := gocv.NewCascadeClassifier()
	defer func() { _ = cascade.Close() }()

	if !cascade.Load("lbpcascade_animeface.xml") {
		log.Fatal("Failed to reading cascade file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		img, _, err := image.Decode(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Failed to decode image: %v", err)))
			return
		}

		mat, err := gocv.ImageToMatRGB(img)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Failed to convert image to mat: %v", err)))
			return
		}
		gocv.CvtColor(mat, &mat, gocv.ColorBGRToGray)
		gocv.EqualizeHist(mat, &mat)

		faces := cascade.DetectMultiScaleWithParams(mat, 1.1, 5, 0, image.Pt(24, 24), image.Pt(2000, 2000))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(faces)
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Failed to listen and serve: ", err)
	}
}
