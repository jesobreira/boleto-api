package app

import (
	"fmt"
	"os"

	"github.com/PMoneda/flow"

	"github.com/mundipagg/boleto-api/api"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/mock"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/robot"
	"github.com/mundipagg/boleto-api/util"
)

//Params this struct contains all execution parameters to run application
type Params struct {
	DevMode    bool
	MockMode   bool
	DisableLog bool
	HTTPOnly   bool
}

//NewParams returns new Empty pointer to ExecutionParameters
func NewParams() *Params {
	return new(Params)
}

//Run starts boleto api Application
func Run(params *Params) {
	configFlags(params.DevMode, params.MockMode, params.DisableLog, params.HTTPOnly)
	installflowConnectors()
	robot.GoRobots()
	if config.Get().MockMode {
		go mock.Run()
	}
	installLog()
	api.InstallRestAPI()

}

func installLog() {
	err := log.Install()
	if err != nil {
		fmt.Println("Log SEQ Fails")
		os.Exit(-1)
	}
}

func installflowConnectors() {
	flow.RegisterConnector("logseq", util.SeqLogConector)
	flow.RegisterConnector("apierro", models.BoletoErrorConector)
	flow.RegisterConnector("tls", util.TlsConector)
}

func configFlags(devMode, mockMode, disableLog, httpOnly bool) {
	if devMode {
		os.Setenv("API_PORT", "3000")
		os.Setenv("API_VERSION", "0.0.1")
		os.Setenv("ENVIROMENT", "Development")
		os.Setenv("SEQ_URL", "http://localhost:5341")   // http://stglog.mundipagg.com/ 192.168.8.119:5341
		os.Setenv("SEQ_API_KEY", "4jZzTybZ9bUHtJiPdh6") //4jZzTybZ9bUHtJiPdh6
		os.Setenv("ENABLE_REQUEST_LOG", "false")
		os.Setenv("ENABLE_PRINT_REQUEST", "true")
		os.Setenv("URL_BB_REGISTER_BOLETO", "https://cobranca.homologa.bb.com.br:7101/registrarBoleto")
		os.Setenv("URL_BB_TOKEN", "https://oauth.hm.bb.com.br:43000/oauth/token")
		os.Setenv("URL_CAIXA", "https://des.barramento.caixa.gov.br/sibar/ManutencaoCobrancaBancaria/Boleto/Externo")
		os.Setenv("URL_CITI", "https://citigroupsoauat.citigroup.com/comercioeletronico/registerboleto/RegisterBoletoSOAP")
		os.Setenv("URL_CITI_BOLETO", "https://ebillpayer.uat.brazil.citigroup.com/ebillpayer/jspInformaDadosConsulta.jsp")
		os.Setenv("APP_URL", "http://localhost:3000/boleto")
		os.Setenv("ELASTIC_URL", "http://localhost:9200")
		os.Setenv("MONGODB_URL", "localhost:27017")
		os.Setenv("BOLETO_JSON_STORE", "/home/philippe/boletodb/upMongo")
		os.Setenv("CERT_BOLETO_CRT", "C:\\cert_boleto_api\\certificate.crt")
		os.Setenv("CERT_BOLETO_KEY", "C:\\cert_boleto_api\\mundi.key")
		os.Setenv("CERT_BOLETO_CA", "C:\\cert_boleto_api\\ca-cert.ca")
		os.Setenv("URL_SANTANDER_TICKET", "https://ymbdlb.isbanbr.dev.corp/dl-ticket-services/TicketEndpointService")
		os.Setenv("URL_SANTANDER_REGISTER", "https://ymbcashhml.santander.com.br:443/ymbsrv/CobrancaEndpointService")
		os.Setenv("SANTANDER_KEY_PATH", "C:\\certificados\\santander\\mundipagg.key")
		os.Setenv("SANTANDER_CA_PATH", "C:\\certificados\\santander\\ca.pem")
		os.Setenv("SANTANDER_CRT_PATH", "C:\\certificados\\santander\\mundipagg.crt")
		if mockMode {
			os.Setenv("URL_BB_REGISTER_BOLETO", "http://localhost:4000/registrarBoleto")
			os.Setenv("URL_BB_TOKEN", "http://localhost:4000/oauth/token")
			os.Setenv("URL_CAIXA", "http://localhost:4000/caixa/registrarBoleto")
			os.Setenv("URL_CITI", "http://localhost:4000/citi/registrarBoleto")
			os.Setenv("URL_SANTANDER_TICKET", "http://localhost:4000/santander/get-ticket")
			os.Setenv("URL_SANTANDER_REGISTER", "http://localhost:4000/santander/register")
			os.Setenv("URL_SANTANDER_TICKET", "http://localhost:4000/santander/get-ticket")
			os.Setenv("URL_SANTANDER_REGISTER", "http://localhost:4000/santander/register")
		}
	}
	config.Install(mockMode, devMode, disableLog, httpOnly)
}
