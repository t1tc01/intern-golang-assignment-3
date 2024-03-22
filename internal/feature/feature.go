package feature

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	fetch "github.com/yngfoxx/gofiber-fetch"
	"gitlab.com/hedwig-phan/assignment-3/cmd/db"
	"gitlab.com/hedwig-phan/assignment-3/ent/featuretype"
	"gitlab.com/hedwig-phan/assignment-3/ent/source"
	"gitlab.com/hedwig-phan/assignment-3/internal/util"
)

const DATA_URL string = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_week.geojson"

func fetchData(url string) ([]byte, error) {
	res := fetch.Method("GET").FiberFetch(url)
	if res.Error != nil {
		panic(res.Error)
	}

	return res.Data.([]byte), nil
}

func HandlerData(recordNumber int) error {
	var geoJSON GeoJSON

	//Save Json to struct
	data, err := fetchData(DATA_URL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &geoJSON)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	//
	if recordNumber < 1 {
		recordNumber = len(geoJSON.Features)
	}

	//Get types and source
	types_map := make(map[string]int)
	sources_map := make(map[string]int)

	for _, feature := range geoJSON.Features {

		//
		ftypes := feature.Properties.Types
		types := strings.Split(ftypes, ",")
		for i, val := range types {
			types[i] = strings.Trim(val, "\"")

			if types[i] == "" {
				continue
			}

			types_map[types[i]]++
		}

		//

		source_names := feature.Properties.Sources
		names := strings.Split(source_names, ",")
		for i, val := range names {
			names[i] = strings.Trim(val, "\"")

			if names[i] == "" {
				continue
			}

			sources_map[names[i]]++

		}

	}

	for key := range types_map {
		current_time := time.Now()
		_ = db.Client.FeatureType.Create().
			SetFeatType(key).
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(util.Ctx)
	}

	for key := range sources_map {
		current_time := time.Now()
		_ = db.Client.Source.Create().
			SetName(key).
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(util.Ctx)
	}

	//Save to Database

	for _, feature := range geoJSON.Features[:recordNumber] {

		current_time := time.Now()

		//Save location
		place := feature.Properties.Place

		parts := strings.Split(place, ",")

		var address, placeName string

		if len(parts) == 1 {
			// Không có dấu phẩy, lưu vào placeName
			placeName = strings.TrimSpace(parts[0])
		} else {
			// Nhiều hơn 1 dấu phẩy, tách theo dấu phẩy cuối cùng
			placeName = strings.TrimSpace(parts[len(parts)-1])
			address = strings.TrimSpace(strings.Join(parts[:len(parts)-1], ","))
		}

		location := db.Client.Location.Create().
			SetPlace(placeName).
			SetAddress(address).
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(util.Ctx)

		//Save geometry
		location_id := location.ID
		longitude := feature.Geometry.Coordinates[0]
		latitude := feature.Geometry.Coordinates[1]
		depth := feature.Geometry.Coordinates[2]

		geometry := db.Client.Geometry.Create().
			SetLocationID(location_id).
			SetLongitude(longitude).
			SetLatitude(latitude).
			SetDepth(depth).
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(util.Ctx)

		//Save Report

		felt := feature.Properties.Felt
		cdi := feature.Properties.CDI
		mmi := feature.Properties.MMI
		alert := feature.Properties.Alert

		report := db.Client.Report.Create().
			SetFelt(felt).
			SetCdi(cdi).
			SetMmi(mmi).
			SetAlert(alert).
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(util.Ctx)

		//Save earth
		geo_id := geometry.ID
		rp_id := report.ID
		mag := feature.Properties.Mag
		time_int := feature.Properties.Time
		updated_time_int := feature.Properties.Updated
		tz := feature.Properties.Tz
		url := feature.Properties.Detail
		status := feature.Properties.Status
		tsunami := feature.Properties.Tsunami
		sig := feature.Properties.Sig
		net := feature.Properties.Net
		code := feature.Properties.Code
		nst := feature.Properties.NST
		dmin := feature.Properties.Dmin
		rms := feature.Properties.RMS
		gap := feature.Properties.Gap
		magType := feature.Properties.MagType
		eq_type := feature.Properties.Type

		time_cv := time.Unix(time_int, 0)
		updated_cv := time.Unix(updated_time_int, 0)

		equake := db.Client.Earthquake.Create().
			SetGeoID(geo_id).
			SetReportID(rp_id).
			SetMag(mag).
			SetTime(time_cv).
			SetUpdatedTime(updated_cv).
			SetTz(tz).
			SetURL(url).
			SetStatus(status).
			SetTsunami(tsunami).
			SetSig(sig).
			SetNet(net).
			SetCode(code).
			SetNst(nst).
			SetDmin(dmin).
			SetRms(rms).
			SetGap(gap).
			SetMagType(magType).
			SetEqType(eq_type).
			SetCreatedAt(current_time).
			SetUpdatedAt(current_time).
			SaveX(util.Ctx)

		eq_id := equake.ID
		//Save feature type

		ftypes := feature.Properties.Types
		types := strings.Split(ftypes, ",")
		for i, val := range types {
			types[i] = strings.Trim(val, "\"")

			if types[i] == "" {
				continue
			}

			//Get ftype ID
			ftype := db.Client.FeatureType.Query().
				Where(featuretype.FeatType(types[i])).
				OnlyX(util.Ctx)

			//Save to table
			_ = db.Client.FtypeEarthquake.Create().
				SetFtID(ftype.ID).
				SetEqID(eq_id).
				SetCreatedAt(current_time).
				SetUpdatedAt(current_time).
				SaveX(util.Ctx)

		}

		//Save source
		source_names := feature.Properties.Sources
		names := strings.Split(source_names, ",")

		for i, val := range names {
			names[i] = strings.Trim(val, "\"")

			if names[i] == "" {
				continue
			}

			//
			s := db.Client.Source.Query().
				Where(source.Name(names[i])).
				OnlyX(util.Ctx)

			//
			_ = db.Client.SourceEarthquake.Create().
				SetSID(s.ID).
				SetEqID(eq_id).
				SetCreatedAt(current_time).
				SetUpdatedAt(current_time).
				SaveX(util.Ctx)
		}

	}

	return nil
}
