package main

import (
	"log"
	"os"

	"github.com/MicahParks/go-rsi"
)

func main() {
	// Create a logger.
	logger := log.New(os.Stdout, "", 0)

	// Gather some data.
	//
	// For production systems, it'd be best to gather test data asynchronously.
	avgGains, avgLosses, _ := testData()

	// Determine the number of periods for the initial inputs. Defaults to 14.
	periods := rsi.DefaultPeriods

	// Average the gains and losses over the given period.
	avgGain := avg(avgGains[0:periods])
	avgLoss := avg(avgLosses[0:periods])
	initialInput := rsi.Input{
		AverageGain: avgGain,
		AverageLoss: avgLoss,
	}

	// Create the RSI data structure and get the first result.
	//
	// If the first argument, the initial periods is 0, the default value, 14, will be used.
	r, result := rsi.New(uint(periods), initialInput)
	logger.Printf("Data index: %d\n  Average gain: %.2f\n  Average loss: %.2f\n RSI: %.2f", periods-1, avgGain, avgLoss, result)

	// Use the remaining data to generate the RSI for each period.
	for i := periods; i < len(avgGains); i++ {
		avgGain = avgGains[i]
		avgLoss = avgLosses[i]
		result = r.Calculate(rsi.Input{
			AverageGain: avgGain,
			AverageLoss: avgLoss,
		})
		logger.Printf("Data index: %d\n  Average gain: %.2f\n  Average loss: %.2f\n RSI: %.2f", i, avgGain, avgLoss, result)
	}
}

func avg(s []float64) (avg float64) {
	for _, v := range s {
		avg += v
	}
	avg /= float64(len(s))
	return avg
}

func testData() (avgGains, avgLosses, results []float64) {
	avgGains = []float64{
		-0.6046602879796196,
		0,
		0,
		0.30091186058528707,
		0.21426387258237492,
		0.4688898449024232,
		-0.6790846759202163,
		-0.360871416856906,
		0,
		0.2065826619136986,
		-0.5238203060500009,
		-0.6072534395455154,
		0,
		0,
		0,
		0,
		-0.7886049150193449,
		0,
		-0.9769168685862624,
		0,
		0,
		0.7302314772948083,
		-0.8969919575618727,
		-0.9222122589217269,
		0.9269868035744142,
		-0.6908388315056789,
		0,
		0.40380328579570035,
		0.8963417453962161,
		-0.6445397825093294,
		-0.6227283173637045,
		0,
		0,
		-0.41032284435628247,
		0,
		0.8305339189948062,
		0,
		0,
		0,
		-0.559392449071014,
		0,
		0.8458327872480417,
		0,
		-0.8143945509670211,
		0.539210105890946,
		0,
		0,
		-0.6278346050000227,
		0,
		0.9296116354490302,
		-0.4117626883450162,
		0,
		0,
		0,
		-0.09838378898573259,
		-0.15184340208190175,
		-0.15965092146489504,
		0.5390745170394794,
		0,
		0,
		0.012825909106361078,
		0,
		0.3443150177263606,
		0,
		0,
		0,
		0.411540363220476,
		-0.09211762444074219,
		0.2507622754291802,
		-0.21256094905031958,
		0,
		0,
		0,
		0.9498832061012532,
		0.4004323171418129,
		0.42868843006993296,
		0,
		0,
		0.07244835679235131,
		0,
		0.24784452318699535,
		0,
		0,
		-0.5838347418625311,
		0.6491884165984236,
		0.33205608571906026,
		0.31051027622482125,
		-0.8332418155281546,
		-0.4173258421903824,
		-0.8958032677441458,
		0.7916723090820832,
		0.43269824007906393,
		-0.0429216421763342,
		0,
		-0.12436351859954159,
		-0.3148047959559753,
		-0.37026753645051064,
		-0.7698890899830834,
		0.34686388037925503,
		0.3511342996652132,
	}
	avgLosses = []float64{
		0,
		-0.4377141871869802,
		0.06563701921747622,
		0,
		0,
		0,
		0,
		0,
		-0.29311424455385804,
		0,
		0,
		0,
		-0.5948085976830626,
		-0.30152268100656,
		-0.544155573000885,
		-0.5305857153507052,
		0,
		-0.2971122606397708,
		0,
		0.6810783123925709,
		-0.932846428518434,
		0,
		0,
		0,
		0,
		0,
		0.6494894605929404,
		0,
		0,
		0,
		0,
		0.5352818906344061,
		0.6280981712183633,
		0,
		-0.5501469205077233,
		0,
		-0.39998376285699544,
		-0.40961827788499267,
		-0.0028430411748625642,
		0,
		0.4584424785756506,
		0,
		0.24746660783662855,
		0,
		0,
		0.2940063127950149,
		-0.35576726540923664,
		0,
		-0.02519395979489504,
		0,
		0,
		-0.9579539135375136,
		0.7830349733960021,
		0.19003276633920804,
		0,
		0,
		0,
		0,
		0.6841751300974551,
		0.654270134424146,
		0,
		0.3691117091643448,
		0,
		0.5550021356347942,
		0.16867966833433606,
		-0.7002878731458151,
		0,
		0,
		0,
		0,
		-0.09297015549992493,
		-0.44846436394045647,
		-0.6718291062346395,
		0,
		0,
		0,
		-0.23632747430436052,
		0.7275772560415229,
		0,
		-0.9777634773577185,
		0,
		0.8351038011543529,
		-0.6320034337887998,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0.5034867900486911,
		0,
		0,
		0,
		0,
		0,
		0,
	}
	results = []float64{
		50.37372814979006,
		42.46506760316381,
		36.455074343328356,
		48.19141056422313,
		44.83183826881235,
		55.75402991428148,
		65.48845086614251,
		52.076954832425926,
		42.07698338873929,
		54.6070966669469,
		63.378465750456634,
		53.692072030084134,
		61.801030904080655,
		75.11905942769799,
		70.9237382460841,
		51.289847633279685,
		68.01512283852966,
		76.43427216113786,
		101.05867776981523,
		170.44976443141798,
		147.5025543564119,
		100.32043640017389,
		100.66763898515933,
		64.4492758322958,
		46.140995751088376,
		46.043230662096285,
		62.76192062871468,
		86.38582878645303,
		45.99829305988126,
		704.430625812193,
		111.6841364461389,
		138.8495600143825,
		-380.17885329730524,
		98.20564701629775,
		99.47091158198144,
		96.52941068543409,
		119.82433081203759,
		90.99612399563593,
		19.597066891868494,
		63.3598555296549,
		152.17608602794462,
		129.2868718557359,
		116.9371564551939,
		111.4636377754057,
		34.54439777560913,
		3.396369035345714,
		1.761056575962911,
		2.7495938316288004,
		2.0959460333819067,
		20.96946935887337,
		15.711724254181803,
		14.520112354584882,
		21.969737016580808,
		41.09598649610068,
		37.3970532079977,
		47.13001378882817,
		38.38541994784097,
		41.62856720991248,
		74.19151021942693,
		-283.21016850672277,
		160.4912184098133,
		139.64945134065763,
		128.37722956290588,
		154.44495974399064,
		92.30336997284398,
		92.62171505950914,
		232.22566829566225,
		193.680855227925,
		94.12119307931962,
		161.98418317624737,
		319.25770441319946,
		154.2936541595806,
		138.38531290124757,
		129.63956633323778,
		186.7752541101841,
		-2083.6051319263383,
		64.48372074758835,
		-457.18970525464147,
		172.87708240856972,
		182.89002322956657,
		66.85383721821556,
		60.12399884284471,
		10.695418395042864,
		256.6490063935015,
		121.85686796662169,
		137.52227459653716,
		271.6187142609139,
	}
	return avgGains, avgLosses, results
}
