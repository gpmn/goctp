package goctp

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/api/v6.7.0_20230209_api_traderapi_se_linux64 -Wl,-rpath,${SRCDIR}/api/v6.7.0_20230209_api_traderapi_se_linux64 -lthostmduserapi_se -lthosttraderapi_se -lLinuxDataCollect -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/api/v6.7.0_20230209_api_traderapi_se_linux64

// windows 不可用，go 部分功能不支持

*/
import "C"
