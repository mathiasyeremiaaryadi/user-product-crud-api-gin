package main

import (
	"os"
)

const ENVIRONMENT string = "development"

var config_env = map[string]map[string]string{
	"development": {
		"API_URL":   "localhost:8080",
		"TOKEN_KEY": "ftr$;C3Uck=2AH/xe(q;}Ak=#%2#@M?BNTrKPP[+zyP.B@G25@%L#AUQ}cvM[ZJ(7}hCNF;qrc$zPz?TB$YT+;BMK6!,SV?PzYXKUvG{:B-XKtL)(awL3ic$AjSzmq9bZ(3WTYrU_V8q*prA._pm;iv_=.FiD+LH+!&U-tpa}/ZzQ:RQ?U?uy75j6v*m[.!t$9UccH+j",
		"DB_HOST":   "127.0.0.1",
		"DB_PORT":   "3306",
		"DB_USER":   "root",
		"DB_PASS":   "",
		"DB_NAME":   "db_user_product",
	},
	"production": {
		"API_URL":   "167.172.73.163:9090",
		"TOKEN_KEY": ".[P}d!T%%Ju2cVCVC;,LJji]e&fZAJq}SB:?:Qn9vdK=2*/?zzJd7htbb#*6aP?==a/K4-*VVy4Heqh@37zG)C86$Eu}a/%;w9,K=wWz))?F?8:c*PN]hEcPa5r37d8yfBX5{X[#_QvgYZ-gTe4kuQ6!C,JMZ7_Z)%nrFZ!CNb7SA?,Zk]ra6n#Zpf!_/FSpqH*GxE$U",
		"DB_HOST":   "167.172.73.163",
		"DB_PORT":   "3306",
		"DB_USER":   "mathias2",
		"DB_PASS":   "DmPUpSN8E394zRTi",
		"DB_NAME":   "mathias2",
	},
}

var CONFIG = config_env[ENVIRONMENT]

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = GetConfig(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}

func GetConfig(key string, config string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return config
}
