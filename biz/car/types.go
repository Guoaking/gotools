package car

import "github.com/PuerkitoBio/goquery"

/**
@description
@date: 06/01 15:26
@author Gk
**/

var navMapping = map[string][]string{
	//"基本信息":    {"Basic Info"},
	//"车身":       {"Body"},
	//"发动机":     {"Engine"},
	//"变速箱":     {"Transmission"},
	//"底盘/转向":  {"Chassis", "Steering"},
	//"车轮/制动":  {"Wheels", "Braking"},
	"主动安全":    {},
	"被动安全":    {},
	"辅助/操控配置": {},
	"外部配置":    {},
	"内部配置":    {},
	"舒适/防盗配置": {},
	"座椅配置":    {},
	"智能互联":    {},
	"影音娱乐":    {},
	"灯光配置":    {},
	"玻璃/后视镜":  {},
	"空调/冰箱":   {},
	"智能化配置":   {},
}

type (
	CarHtml struct {
		htmlDoc   *goquery.Document
		car       *CarInfo
		BasicBody *goquery.Selection
	}

	CarContent struct {
		CarInfo
		Tags map[string][]string
	}
)

type (
	RequestParams struct {
		Price          string `url:"price,omitempty"`           // 价格万 "3,5"
		SeriesType     int    `url:"series_type,omitempty"`     //  1 suv 0 轿车
		AgeRange       string `url:"age_range,omitempty"`       // 车龄 "3,5"
		ExpandedOrigin int    `url:"expanded_origin,omitempty"` // 例如 3: 合资自主， 2 自主
		SHCityName     string `url:"sh_city_name,omitempty"`    // 范围 "全国"
		Page           int    `url:"page,omitempty"`            // 分页 页码 1
		Limit          int    `url:"limit,omitempty"`           // 分页大小 20
		MileageRange   string `url:"mileage_range"`             // 公里数 0,6
		FuelForm       string `url:"fuel_form"`                 // 能源类型 1:汽油 2:柴油， 3：混合 5:增城 6:插电
		CapacityL      string `url:"capacity_l"`                // 排量： 1.5,2.0
		GearBoxType    string `url:"gearbox_type"`              // 变速 ： 2 自动
	}

	CarList struct {
		Total               int               `json:"total"`
		HasMore             bool              `json:"has_more"`
		SearchShSkuInfoList []SearchShSkuInfo `json:"search_sh_sku_info_list"`
	}

	SearchShSkuInfo struct {
		SkuId                int         `json:"sku_id"`
		ShopId               string      `json:"shop_id"`
		CarId                int         `json:"car_id"`
		CarName              string      `json:"car_name"`
		SeriesId             int         `json:"series_id"`
		SeriesName           string      `json:"series_name"`
		BrandId              int         `json:"brand_id"`
		BrandName            string      `json:"brand_name"`
		CarYear              int         `json:"car_year"`
		CarSourceCityName    string      `json:"car_source_city_name"`
		BrandSourceCityName  string      `json:"brand_source_city_name"`
		SpuId                int         `json:"spu_id"`
		PlatformType         int         `json:"platform_type"`
		TransferCnt          int         `json:"transfer_cnt"`
		GroupId              int64       `json:"group_id"`
		GroupIdStr           string      `json:"group_id_str"`
		Image                string      `json:"image"`
		RelatedVideoThumb    string      `json:"related_video_thumb"`
		IsVideo              bool        `json:"is_video"`
		ShPrice              string      `json:"sh_price"`
		OfficialPrice        string      `json:"official_price"`
		CarMileage           string      `json:"car_mileage"`
		CarAge               string      `json:"car_age"`
		Title                string      `json:"title"`
		SubTitle             string      `json:"sub_title"`
		CarSourceType        string      `json:"car_source_type"`
		AuthenticationMethod string      `json:"authentication_method"`
		Tags                 []Tag       `json:"tags"`
		TagsV2               []TagV2     `json:"tags_v2"`
		OfficialHintBar      string      `json:"official_hint_bar"`
		SpecialTags          interface{} `json:"special_tags"`
		IsSelfTrade          bool        `json:"is_self_trade"`
	}

	Tag struct {
		Text string `json:"text"`
		//Logo            string `json:"logo"`
		//TextColor       string `json:"text_color"`
		//BackgroundColor string `json:"background_color"`
	}

	TagV2 struct {
		Text string `json:"text"`
		//Height int    `json:"height"`
		//Key    string `json:"key"`
		//Level  int    `json:"level"`
		//Logo            string `json:"logo,omitempty"`
		//Width           int    `json:"width,omitempty"`
		//BackgroundColor string `json:"background_color,omitempty"`
		//BorderColor     string `json:"border_color,omitempty"`
		//TextColor       string `json:"text_color,omitempty"`
	}

	BaseVo[T any] struct {
		Prompts string `json:"prompts"`
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}
)

//func (t *Tag) Marshal() ([]byte, error) {
//	data, err := json.Marshal(t)
//	if err != nil {
//		return nil, fmt.Errorf("call Marshal(t)  err: %w", err)
//	}
//	return data, nil
//}

type (
	Top struct {
		Props struct {
			PageProps PageProps `json:"pageProps"`
		} `json:"props"`
		Page  string `json:"page"`
		Query struct {
			Field1 string `json:"0"`
			Field2 string `json:"2"`
		} `json:"query"`
		BuildId      string        `json:"buildId"`
		AssetPrefix  string        `json:"assetPrefix"`
		IsFallback   bool          `json:"isFallback"`
		CustomServer bool          `json:"customServer"`
		Gip          bool          `json:"gip"`
		AppGip       bool          `json:"appGip"`
		ScriptLoader []interface{} `json:"scriptLoader"`
	}

	PageProps struct {
		RawData struct {
			//Properties  []Propertie   `json:"properties"`
			CarInfo     []CarInfo     `json:"car_info"`
			FilterGroup []interface{} `json:"filter_group"`
		} `json:"rawData"`
		PageType       string `json:"_pageType"`
		BackToTopStyle struct {
			ZIndex string `json:"zIndex"`
		} `json:"_backToTopStyle"`
		HasUrlCity            bool     `json:"__hasUrlCity"`
		IsGray                int      `json:"is_gray"`
		HasGray               int      `json:"has_gray"`
		ClientIp              string   `json:"clientIp"`
		SensitiveSeriesIdList []string `json:"sensitiveSeriesIdList"`
	}

	CarInfo struct {
		CarId      string `json:"car_id"`
		CarName    string `json:"car_name"`
		SeriesId   string `json:"series_id"`
		SeriesName string `json:"series_name"`
		BrandId    string `json:"brand_id"`
		BrandName  string `json:"brand_name"`
		CarYear    string `json:"car_year"`

		SeriesType     string `json:"series_type"`
		OfficialPrice  string `json:"official_price"`
		DealerPrice    string `json:"dealer_price"`
		HasDealerPrice bool   `json:"has_dealer_price"`
		SaleStatus     int    `json:"sale_status"`
		CarPageEnable  bool   `json:"car_page_enable"`
		//Info           Info   `json:"info"`
		//DealerText     string `json:"dealer_text"`

	}

	Propertie struct {
		Context          string `json:"context"`
		ConfigureFileUrl string `json:"configure_file_url"`
		Text             string `json:"text"`
		Key              string `json:"key"`
		Type             int    `json:"type"`
		TitleFlag        string `json:"title_flag"`
		TitleFlagList    []struct {
			IconUrl  string `json:"icon_url"`
			Value    string `json:"value"`
			IconType int    `json:"icon_type"`
		} `json:"title_flag_list"`
		SubList []struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"sub_list"`
		WikiInfo  WikiInfo `json:"wiki_info"`
		TypeIndex int      `json:"type_index"`
		TypeText  string   `json:"type_text"`
	}

	Info struct {
		VariableSuspension                  info `json:"variable_suspension"`
		FrontElectricWindow                 info `json:"front_electric_window"`
		RearIndependentAirConditioning      info `json:"rear_independent_air_conditioning"`
		FrontTrack                          info `json:"front_track"`
		SportStyleSeat                      info `json:"sport_style_seat"`
		ChildSeatInterface                  info `json:"child_seat_interface"`
		HighHeadlampType                    info `json:"high_headlamp_type"`
		RearElectricWindow                  info `json:"rear_electric_window"`
		FuelForm                            info `json:"fuel_form"`
		SecondRowSeatDownRatio              info `json:"second_row_seat_down_ratio"`
		HeaderDisplaySystem                 info `json:"header_display_system"`
		CarBodyStructure                    info `json:"car_body_structure"`
		OuterCameraPixel                    info `json:"outer_camera_pixel"`
		RoadTrafficSignRecognition          info `json:"road_traffic_sign_recognition"`
		HeatPumpManagementSystem            info `json:"heat_pump_management_system"`
		ExterMirrorAutoPrevGlare            info `json:"exter_mirror_auto_prev_glare"`
		CylinderVolumeMl                    info `json:"cylinder_volume_ml"`
		EnvironmentalStandards              info `json:"environmental_standards"`
		AccelerationTime                    info `json:"acceleration_time"`
		SteerAssistLight                    info `json:"steer_assist_light"`
		RearBrakeType                       info `json:"rear_brake_type"`
		RemoteControlMove                   info `json:"remote_control_move"`
		CenterScreenSize                    info `json:"center_screen_size"`
		MaintainCost                        info `json:"maintain_cost"`
		FrontParkingRadar                   info `json:"front_parking_radar"`
		FrontFogLight                       info `json:"front_fog_light"`
		TirePressureSystem                  info `json:"tire_pressure_system"`
		CarFragranceDevice                  info `json:"car_fragrance_device"`
		MobileWirelessCharging              info `json:"mobile_wireless_charging"`
		FramelessDesignDoor                 info `json:"frameless_design_door"`
		HeadlampDelayOff                    info `json:"headlamp_delay_off"`
		ElectricBackDoor                    info `json:"electric_back_door"`
		MainDriveBackrestAdjustment         info `json:"main_drive_backrest_adjustment"`
		AutoRoadChange                      info `json:"auto_road_change"`
		ReversingCamera                     info `json:"reversing_camera"`
		PositionService                     info `json:"position_service"`
		SeatBeltPrompted                    info `json:"seat_belt_prompted"`
		AppStore                            info `json:"app_store"`
		Gps                                 info `json:"gps"`
		ViceAirbag                          info `json:"vice_airbag"`
		OverallTurn                         info `json:"overall_turn"`
		ViceDriveWindowSunshadeMirror1      info `json:"vice_drive_window_sunshade_mirror_1"`
		RearCupHolder                       info `json:"rear_cup_holder"`
		MainKneeAirbag                      info `json:"main_knee_airbag"`
		SteerWheelShift                     info `json:"steer_wheel_shift"`
		KeylessStart                        info `json:"keyless_start"`
		RearLcdScreen                       info `json:"rear_lcd_screen"`
		EngineMaxTorqueRevolution           info `json:"engine_max_torque_revolution"`
		DrivingComputerDisplayScreen        info `json:"driving_computer_display_screen"`
		RemoteKey1                          info `json:"remote_key_1"`
		AutoHeadlamp                        info `json:"auto_headlamp"`
		CenterConsoleScreenMaterial         info `json:"center_console_screen_material"`
		OtaUpgrade                          info `json:"ota_upgrade"`
		DragHook                            info `json:"drag_hook"`
		AutoRoadOutIn                       info `json:"auto_road_out_in"`
		CarCall                             info `json:"car_call"`
		ActiveClosedInletGrid               info `json:"active_closed_inlet_grid"`
		ViceDriveBackrestAdjustment         info `json:"vice_drive_backrest_adjustment"`
		PassivePedestrianProtection         info `json:"passive_pedestrian_protection"`
		CarBodyStruct                       info `json:"car_body_struct"`
		SentinelMode                        info `json:"sentinel_mode"`
		ElecSteerWheelAdjustment            info `json:"elec_steer_wheel_adjustment"`
		ForwardCarDepartureReminder         info `json:"forward_car_departure_reminder"`
		AutomaticDriveLevel                 info `json:"automatic_drive_level"`
		LowHeadlampType                     info `json:"low_headlamp_type"`
		SoundBrand                          info `json:"sound_brand"`
		AudioAndVideoSystem4                info `json:"audio_and_video_system_4"`
		InsideMirrorAutoAntiGlare           info `json:"inside_mirror_auto_anti_glare"`
		BackArmrest                         info `json:"back_armrest"`
		EngineMaxPower                      info `json:"engine_max_power"`
		LightProjectionTechnology           info `json:"light_projection_technology"`
		ViceDriveBackAndForthAdjustment     info `json:"vice_drive_back_and_forth_adjustment"`
		EngineMaxHorsepower                 info `json:"engine_max_horsepower"`
		CylinderArrangement                 info `json:"cylinder_arrangement"`
		ViceDriveSeatAdjustment             info `json:"vice_drive_seat_adjustment"`
		MaxSpeed                            info `json:"max_speed"`
		MarketTime                          info `json:"market_time"`
		AdaptiveLight                       info `json:"adaptive_light"`
		RearSlipMethod                      info `json:"rear_slip_method"`
		HighPrecisionMap                    info `json:"high_precision_map"`
		VoiceWakeUpRecognition              info `json:"voice_wake_up_recognition"`
		RearTireSize                        info `json:"rear_tire_size"`
		CarTv                               info `json:"car_tv"`
		CurbWeight                          info `json:"curb_weight"`
		EngineDescription                   info `json:"engine_description"`
		OilTankVolume                       info `json:"oil_tank_volume"`
		SteerWheelAdjustment                info `json:"steer_wheel_adjustment"`
		AirSupply                           info `json:"air_supply"`
		RearTouchControlSystem              info `json:"rear_touch_control_system"`
		EngineRemoteStart                   info `json:"engine_remote_start"`
		WindowOneKeyLift                    info `json:"window_one_key_lift"`
		FrontSuspensionForm                 info `json:"front_suspension_form"`
		GearboxType                         info `json:"gearbox_type"`
		BuiltInTachograph                   info `json:"built_in_tachograph"`
		MultimediaInterface                 info `json:"multimedia_interface"`
		SecondRowSeatBackrestAdjustment     info `json:"second_row_seat_backrest_adjustment"`
		MainDriveSeatAdjustment             info `json:"main_drive_seat_adjustment"`
		BrakeForce                          info `json:"brake_force"`
		BacksidePrivacyGlass                info `json:"backside_privacy_glass"`
		LengthWidthHeight                   info `json:"length_width_height"`
		AirSuspension                       info `json:"air_suspension"`
		DataNetwork                         info `json:"data_network"`
		VoiceprintRecognition               info `json:"voiceprint_recognition"`
		VitalSignsDetection                 info `json:"vital_signs_detection"`
		MainDriveHeightAdjustment           info `json:"main_drive_height_adjustment"`
		HeadlightCleanFunction              info `json:"headlight_clean_function"`
		VoiceSimulate                       info `json:"voice_simulate"`
		GestureControlSystem                info `json:"gesture_control_system"`
		Width                               info `json:"width"`
		MaxEngineNetPower                   info `json:"max_engine_net_power"`
		RearParkingRadar                    info `json:"rear_parking_radar"`
		ElectricSpoiler                     info `json:"electric_spoiler"`
		HiddenDoorHandle                    info `json:"hidden_door_handle"`
		DoorNums                            info `json:"door_nums"`
		MemoryParking                       info `json:"memory_parking"`
		HeatedNozzle                        info `json:"heated_nozzle"`
		EngineUniqueTech                    info `json:"engine_unique_tech"`
		FrontNearCenterAirbag               info `json:"front_near_center_airbag"`
		MainDriveBackAndForthAdjustment     info `json:"main_drive_back_and_forth_adjustment"`
		EmotionRecognition                  info `json:"emotion_recognition"`
		AutoParkEntry                       info `json:"auto_park_entry"`
		MinTurningRadius                    info `json:"min_turning_radius"`
		AutoPark                            info `json:"auto_park"`
		SideAirCurtain                      info `json:"side_air_curtain"`
		ElectromagneticInductSuspension     info `json:"electromagnetic_induct_suspension"`
		FingerprintRecognition              info `json:"fingerprint_recognition"`
		PowerSteeringType                   info `json:"power_steering_type"`
		AqsAirQualityManagementSystem       info `json:"aqs_air_quality_management_system"`
		ActiveNoiseReduction                info `json:"active_noise_reduction"`
		AbsAntiLock                         info `json:"abs_anti_lock"`
		KeylessEntry                        info `json:"keyless_entry"`
		DaytimeLight                        info `json:"daytime_light"`
		EnergyElectMaxTorque                info `json:"energy_elect_max_torque"`
		MultiFingerScreenControl            info `json:"multi_finger_screen_control"`
		ValvesPerCylinderNums               info `json:"valves_per_cylinder_nums"`
		FrontSlipMethod                     info `json:"front_slip_method"`
		EnergyElectMaxPower                 info `json:"energy_elect_max_power"`
		ArRealityNavigation                 info `json:"ar_reality_navigation"`
		FrontAirbag                         info `json:"front_airbag"`
		SecondIndependentSeat               info `json:"second_independent_seat"`
		ElectricDoor                        info `json:"electric_door"`
		HotColdCupHolder                    info `json:"hot_cold_cup_holder"`
		ActiveBrake                         info `json:"active_brake"`
		HighPrecisionPositionSystem         info `json:"high_precision_position_system"`
		RearSuspensionForm                  info `json:"rear_suspension_form"`
		SkylightType                        info `json:"skylight_type"`
		UsbTypecInterfaceCount              info `json:"usb_typec_interface_count"`
		FrontSeatHeating1                   info `json:"front_seat_heating_1"`
		CompressionRatioS                   info `json:"compression_ratio_s"`
		RoofRacks                           info `json:"roof_racks"`
		VisibleToSay                        info `json:"visible_to_say"`
		Cruise                              info `json:"cruise"`
		LineSupport                         info `json:"line_support"`
		CylinderNums                        info `json:"cylinder_nums"`
		GasForm                             info `json:"gas_form"`
		RearMakeupMirror1                   info `json:"rear_makeup_mirror_1"`
		Height                              info `json:"height"`
		SeatCount                           info `json:"seat_count"`
		TractionControl                     info `json:"traction_control"`
		InnerCameraPixel                    info `json:"inner_camera_pixel"`
		CylinderMaterial                    info `json:"cylinder_material"`
		Baggage12VPowerOutlet               info `json:"baggage_12v_power_outlet"`
		Stalls                              info `json:"stalls"`
		FrontWindshieldElectricHeating      info `json:"front_windshield_electric_heating"`
		CentralLockingCar                   info `json:"central_locking_car"`
		CarNetworking                       info `json:"car_networking"`
		FullLoadWeight                      info `json:"full_load_weight"`
		Jb                                  info `json:"jb"`
		HeadlampRainFogMode                 info `json:"headlamp_rain_fog_mode"`
		RearWindowSunshade                  info `json:"rear_window_sunshade"`
		FrontBrakeType                      info `json:"front_brake_type"`
		SignalRecognition                   info `json:"signal_recognition"`
		FuelComprehensive                   info `json:"fuel_comprehensive"`
		SecondRowSeatElectricalAdjustment   info `json:"second_row_seat_electrical_adjustment"`
		NegativeIonGenerator                info `json:"negative_ion_generator"`
		GearboxDescription                  info `json:"gearbox_description"`
		CoPilotRearAdjustableButton         info `json:"co_pilot_rear_adjustable_button"`
		Period                              info `json:"period"`
		SeatMaterial                        info `json:"seat_material"`
		ActiveDmsFatigueDetection           info `json:"active_dms_fatigue_detection"`
		MainDrivePartAdjustment             info `json:"main_drive_part_adjustment"`
		LaneCenter                          info `json:"lane_center"`
		FrontArmrest                        info `json:"front_armrest"`
		Wifi                                info `json:"wifi"`
		ExternalMirrorHeat                  info `json:"external_mirror_heat"`
		RearWiper                           info `json:"rear_wiper"`
		LcdDashboardType                    info `json:"lcd_dashboard_type"`
		SecondRowSmallDesktop               info `json:"second_row_small_desktop"`
		FacialRecognition                   info `json:"facial_recognition"`
		RainInductionWiper                  info `json:"rain_induction_wiper"`
		ParkBrakeType                       info `json:"park_brake_type"`
		FatigueDrivingWarning               info `json:"fatigue_driving_warning"`
		MobileRemoteControl                 info `json:"mobile_remote_control"`
		NightVisionSystem                   info `json:"night_vision_system"`
		TrackReverse                        info `json:"track_reverse"`
		ActiveAmbientLight                  info `json:"active_ambient_light"`
		DriveMode                           info `json:"drive_mode"`
		Length                              info `json:"length"`
		OfficialPrice                       info `json:"official_price"`
		MainDriveWindowSunshadeMirror1      info `json:"main_drive_window_sunshade_mirror_1"`
		LaneWarningSystem                   info `json:"lane_warning_system"`
		ViceDrivePartAdjustment             info `json:"vice_drive_part_adjustment"`
		EngineModel                         info `json:"engine_model"`
		SportsAppearanceKit                 info `json:"sports_appearance_kit"`
		BrakeAssist                         info `json:"brake_assist"`
		InductiveBackDoor                   info `json:"inductive_back_door"`
		SpeechRecognitionSystem             info `json:"speech_recognition_system"`
		Pm25FiltratingEquipment             info `json:"pm25_filtrating_equipment"`
		MultilayerSoundproofGlass           info `json:"multilayer_soundproof_glass"`
		CapacityL                           info `json:"capacity_l"`
		VoiceRecognition                    info `json:"voice_recognition"`
		EngineAntiTheft                     info `json:"engine_anti_theft"`
		EngineSasTech                       info `json:"engine_sas_tech"`
		CylinderHeadMaterial                info `json:"cylinder_head_material"`
		CarFridgeFeature                    info `json:"car_fridge_feature"`
		SteerWheelMaterial2                 info `json:"steer_wheel_material_2"`
		MultifunctionSteerWheel             info `json:"multifunction_steer_wheel"`
		TemperaturePartitionControl3        info `json:"temperature_partition_control_3"`
		LightSensingCanopy                  info `json:"light_sensing_canopy"`
		HeadlightHeightAdjustment           info `json:"headlight_height_adjustment"`
		SpareTireSpecification1             info `json:"spare_tire_specification_1"`
		CentralDifferentialLock             info `json:"central_differential_lock"`
		CarPurifier                         info `json:"car_purifier"`
		RearAirbag                          info `json:"rear_airbag"`
		AirSuspensionV2                     info `json:"air_suspension_v2"`
		RearTrack                           info `json:"rear_track"`
		OilSupply                           info `json:"oil_supply"`
		CarRefrigerator                     info `json:"car_refrigerator"`
		MainAirbag                          info `json:"main_airbag"`
		FuelLabel                           info `json:"fuel_label"`
		PowerOutlet                         info `json:"power_outlet"`
		UsbTypecInterfaceMaxChargingPower   info `json:"usb_typec_interface_max_charging_power"`
		SteepSlope                          info `json:"steep_slope"`
		MaxPowerRevolution                  info `json:"max_power_revolution"`
		RearSeatHeating                     info `json:"rear_seat_heating"`
		ElectricBackDoorMemory              info `json:"electric_back_door_memory"`
		LightSpecialFunction                info `json:"light_special_function"`
		WindowAntiClipFunction              info `json:"window_anti_clip_function"`
		VoiceWakeUpFree                     info `json:"voice_wake_up_free"`
		VoiceWakeUpWord                     info `json:"voice_wake_up_word"`
		LcdDashboardSize                    info `json:"lcd_dashboard_size"`
		OutsideMirrorElectricFolding        info `json:"outside_mirror_electric_folding"`
		Speaker                             info `json:"speaker"`
		InteriorLight                       info `json:"interior_light"`
		AirControlModel                     info `json:"air_control_model"`
		DoorOpenWay                         info `json:"door_open_way"`
		BodyStruct                          info `json:"body_struct"`
		StowableThirdRowSeats               info `json:"stowable_third_row_seats"`
		BluetoothAndCarPhone                info `json:"bluetooth_and_car_phone"`
		Wheelbase                           info `json:"wheelbase"`
		V2XCommunication                    info `json:"v2x_communication"`
		AlloyWheel                          info `json:"alloy_wheel"`
		LaneKeepingAssist                   info `json:"lane_keeping_assist"`
		NavigationSystem                    info `json:"navigation_system"`
		ExplosionTire                       info `json:"explosion_tire"`
		VariableSteerSystem                 info `json:"variable_steer_system"`
		HeadlampFollowUp                    info `json:"headlamp_follow_up"`
		MinClearance                        info `json:"min_clearance"`
		ExterMirrorElecAdjustment           info `json:"exter_mirror_elec_adjustment"`
		BaggageVolume                       info `json:"baggage_volume"`
		SubBrandName                        info `json:"sub_brand_name"`
		UphillSupport                       info `json:"uphill_support"`
		SecondLocalSeat                     info `json:"second_local_seat"`
		SecondRowSeatBackAndForthAdjustment info `json:"second_row_seat_back_and_forth_adjustment"`
		RearWindshieldSunshade              info `json:"rear_windshield_sunshade"`
		FrontTireSize                       info `json:"front_tire_size"`
		RearExhaust                         info `json:"rear_exhaust"`
		BodyStabilitySystem                 info `json:"body_stability_system"`
		DriverForm                          info `json:"driver_form"`
		MobileSystem                        info `json:"mobile_system"`
		NavigationAssistedDriving           info `json:"navigation_assisted_driving"`
		EngineMaxTorque                     info `json:"engine_max_torque"`
	}

	info struct {
		Value       string      `json:"value"`
		IconType    int         `json:"icon_type"`
		IconUrl     string      `json:"icon_url"`
		ConfigPrice string      `json:"config_price"`
		LightConfig interface{} `json:"light_config"`
		WikiInfo    WikiInfo    `json:"wiki_info"`
	}

	WikiInfo struct {
		OpenUrl       string `json:"open_url"`
		WikiType      int    `json:"wiki_type"`
		Abstract      string `json:"abstract"`
		CoverUrl      string `json:"cover_url"`
		VideoUri      string `json:"video_uri"`
		VideoCoverUrl string `json:"video_cover_url"`
		GroupId       string `json:"groupId"`
	}
)
