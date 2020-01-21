package paytm

var (
	paytmMerchantKey     = `xxxxxxxxx`
	mID                  = `xxxxxxxxxxxxxxxxxxxx`
	iNDUSTRY_TYPE_ID     = `Retail`
	cHANNEL_ID           = `WAP`
	wEBSITE              = `APPSTAGING`
	CALLBACK_URL         = `https://securegw-stage.paytm.in/theia/paytmCallback?ORDER_ID=`
	TransactionStatusAPI = `https://securegw-stage.paytm.in/merchant-status/getTxnStatus`
)

//SetPaytmMerchantKey, sets paytm's merchant key
func SetPaytmMerchantKey(key string) {
	paytmMerchantKey = key
}

//SetMID, sets paytm's Merchant ID
func SetMID(mid string) {
	mID = mid
}

//SetIndustryTypeId, sets Industry Type Id
func SetIndustryTypeId(id string) {
	iNDUSTRY_TYPE_ID = id
}

//SetChannerlId, sets channel Id
func SetChannerlId(cid string) {
	cHANNEL_ID = cid
}

//SetWebsite, sets host website
//Use APPSTAGING for development purposes
func SetWebsite(web string) {
	wEBSITE = web
}
